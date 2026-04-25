package main

import (
	"embed"
	"io/fs"
	"mime"
	"path/filepath"
	"runtime"
	"strings"
	"verification/controllers/common"
	"verification/models"
	_ "verification/routers"

	"github.com/astaxie/beego/validation"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"
)

//go:embed static
var staticFiles embed.FS

// staticFileHandler serves static files from embedded FS with SPA fallback
func staticFileHandler(ctx *context.Context) {
	path := ctx.Request.URL.Path

	// API routes should not be handled here (handled by Beego routers)
	if strings.HasPrefix(path, "/admin/") || strings.HasPrefix(path, "/api/") {
		return // Continue to Beego routers
	}

	// Create a sub-filesystem from the "static" directory
	subFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		ctx.Output.SetStatus(500)
		return
	}

	// Remove leading slash for file path
	filePath := strings.TrimPrefix(path, "/")

	// Try to read the file from the sub-filesystem
	fileContent, err := fs.ReadFile(subFS, filePath)
	if err != nil {
		// File not found, serve index.html for SPA routing
		indexFile, indexErr := fs.ReadFile(subFS, "index.html")
		if indexErr != nil {
			ctx.Output.SetStatus(404)
			return
		}
		ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
		ctx.Output.Body(indexFile)
		return
	}

	// Determine MIME type based on file extension
	ext := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		// Default MIME type for unknown extensions
		mimeType = "application/octet-stream"
	}
	// Ensure charset for text types
	if strings.HasPrefix(mimeType, "text/") || strings.Contains(mimeType, "javascript") || strings.Contains(mimeType, "json") {
		if !strings.Contains(mimeType, "charset") {
			mimeType += "; charset=utf-8"
		}
	}

	ctx.Output.Header("Content-Type", mimeType)
	ctx.Output.Body(fileContent)
}

func init() {
	status, Conf := common.ReadIni()
	if status == false {
		logs.Error("配置文件读取失败")
	}

	// 从配置文件设置端口
	if Conf.HttpPort > 0 {
		beego.BConfig.Listen.HTTPPort = Conf.HttpPort
	}

	cpuNum := runtime.NumCPU()
	maxIdle := 4
	maxConn := int(((0.2 + 0.4) / 0.2) * cpuNum)
	logs.Info("线程数量:", maxConn)
	if Conf.Sql == "sqlite" {
		file, _ := beego.AppConfig.String("sqlFile")
		_ = orm.RegisterDriver("sqlite", orm.DRSqlite)
		_ = orm.RegisterDataBase("default", "sqlite", file, orm.MaxIdleConnections(maxIdle), orm.MaxOpenConnections(maxConn))
	} else {
		_ = orm.RegisterDataBase("default", "mysql", Conf.SqlUser+":"+Conf.SqlPwd+"@tcp("+Conf.SqlIp+":"+Conf.SqlPort+")/"+Conf.SqlDb+"?charset=utf8", orm.MaxIdleConnections(maxIdle), orm.MaxOpenConnections(maxConn))
	}
	orm.RegisterModel(
		new(models.Manager),
		new(models.Keys),
		new(models.Level),
		new(models.Member),
		new(models.MemberLogin),
		new(models.Project),
		new(models.ProjectLogin),
		new(models.ProjectVersion),
		new(models.Cards),
		new(models.Order),
		new(models.Role),
		new(models.RoleItem),
		new(models.ManagerCards),
		new(models.UnbindLog),
	)
	_ = orm.RunSyncdb("default", Conf.SqlRebuild, true)
	var m *models.Manager
	m.InitManager()
}

func main() {
	var messages = map[string]string{
		"Required": "不能为空",
		"MinSize":  "最短长度为 %d",
		"MaxSize":  "最长长度为 %d",
		"Length":   "长度必须为 %d",
		"Numeric":  "必须是有效的数字",
		"Email":    "必须是有效的电子邮件地址",
		"Mobile":   "必须是有效的手机号码",
	}
	validation.SetDefaultMessage(messages)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Register static file handler with SPA fallback using Filter
	beego.InsertFilter("*", beego.BeforeRouter, staticFileHandler)

	beego.Run()
}
