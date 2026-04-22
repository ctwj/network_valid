import type { PluginOption } from 'vite';
import vue from '@vitejs/plugin-vue';
import vueJsx from '@vitejs/plugin-vue-jsx';
import progress from 'vite-plugin-progress';
import EslintPlugin from 'vite-plugin-eslint';
import { VitePWA } from 'vite-plugin-pwa';
import { visualizer } from 'rollup-plugin-visualizer';
import { createMockServer } from 'vite-plugin-mock';
import PurgeIcons from 'vite-plugin-purge-icons';
import { createHtmlPlugin } from 'vite-plugin-html';
import Compression from 'vite-plugin-compression';

export function createVitePlugins(viteEnv: ViteEnv, isBuild: boolean, prodMock: boolean) {
  const { VITE_GLOB_APP_TITLE, VITE_PUBLIC_PATH } = viteEnv;

  const vitePlugins: (PluginOption | PluginOption[])[] = [
    // Vue3 插件
    vue(),
    // Vue3 JSX 插件
    vueJsx(),
    // 进度条
    progress(),
    // 自动导入图标
    PurgeIcons(),
    // Eslint 插件
    EslintPlugin({
      cache: false,
      include: ['src/**/*.vue', 'src/**/*.ts', 'src/**/*.tsx'],
      exclude: ['node_modules', 'dist'],
    }),
  ];

  // PWA 插件
  vitePlugins.push(
    VitePWA({
      manifest: {
        name: VITE_GLOB_APP_TITLE,
        short_name: VITE_GLOB_APP_TITLE,
        icons: [
          {
            src: './resource/img/pwa-192x192.png',
            sizes: '192x192',
            type: 'image/png',
          },
          {
            src: './resource/img/pwa-512x512.png',
            sizes: '512x512',
            type: 'image/png',
          },
        ],
      },
      workbox: {
        cleanupOutdatedCaches: true,
        globPatterns: ['**/*.{js,css,html,ico,png,svg}'],
      },
    })
  );

  // HTML 插件
  vitePlugins.push(
    createHtmlPlugin({
      inject: {
        data: {
          title: VITE_GLOB_APP_TITLE,
        },
      },
      minify: isBuild,
    })
  );

  // Mock 插件
  vitePlugins.push(
    createMockServer({
      ignore: /^_/,
      mockFiles: ['mock/**/*.ts'],
      enableProd: prodMock,
      watchFiles: true,
    })
  );

  // 打包分析
  if (isBuild) {
    vitePlugins.push(
      visualizer({
        filename: './stats.html',
        open: false,
        gzipSize: true,
        brotliSize: true,
      }) as PluginOption
    );
  }

  // Gzip 压缩
  if (isBuild) {
    vitePlugins.push(
      Compression({
        verbose: true,
        disable: false,
        threshold: 10240,
        algorithm: 'gzip',
        ext: '.gz',
      })
    );
  }

  return vitePlugins;
}