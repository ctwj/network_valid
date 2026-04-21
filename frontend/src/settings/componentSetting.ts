export default {
  table: {
    apiSetting: {
      // 当前页的字段名
      pageField: 'page',
      // 每页数量字段名
      sizeField: 'limit',
      // 接口返回的数据字段名
      listField: 'data',
      // 接口返回总页数字段名
      totalField: 'total_pages',
    },
    //默认分页数量
    defaultPageSize: 10,
    //可切换每页数量集合
    pageSizes: [10, 20, 30, 40, 50],
  },
  upload: {
    //考虑接口规范不同
    apiSetting: {
      // 集合字段名
      infoField: 'data',
      // 图片地址字段名
      imgField: 'photo',
    },
    //最大上传图片大小
    maxSize: 2,
    //图片上传类型
    fileType: ['image/png', 'image/jpg', 'image/jpeg', 'image/gif', 'image/svg+xml'],
  },
  project: {
    status: ["收费运营", "停止运营", "免费运营"],
    encrypt: ["开放API", "AES"],
    type: ["单码", "用户"],
    hash: ["MD5","SHA1","SHA224","SHA256","SHA384","SHA512"],
    login_mode: ["绑定登录", "普通登录", "点数登录"],
    login_reg_mode: ["带卡注册", "普通注册"],
    login_emial_reg: ["关闭", "开启"],
    login_unbind_mode: ["不允许解绑", "原机解绑", "自动解绑", "任意解绑"],
    login_unbind_weaken_mode: ["不扣时", "解绑就扣时", "超出扣时", "超出扣时"],
    login_number_mode: ["机器码控制", "IP控制"],
    login_unbind_date: ["天", "月"],
    login_number_weaken_time: ["时", "天"],
    keys_type_lock: ["正常", "锁定"],
    keys_create_type: ["纯数字","大写字母数字组合","小写字母数字组合","随机大小写组合"],
    version_is_active: ["启用","关闭"],
    version_is_must_update: ["强制更新","不强制更新"]
  }
};
