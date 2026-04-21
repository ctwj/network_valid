import {RouteRecordRaw} from 'vue-router';
import {Layout} from '@/router/constant';
import {FileProtectOutlined} from '@vicons/antd';
import {renderIcon} from '@/utils/index';

/**
 * @param name 路由名称, 必须设置,且不能重名
 * @param meta 路由元信息（路由附带扩展信息）
 * @param redirect 重定向地址, 访问这个路由时,自定进行重定向
 * @param meta.disabled 禁用整个菜单
 * @param meta.title 菜单名称
 * @param meta.icon 菜单图标
 * @param meta.keepAlive 缓存该路由
 * @param meta.sort 排序越小越排前
 *
 * */
const routes: Array<RouteRecordRaw> = [
  {
    path: '/project',
    name: 'Project',
    redirect: '/project/projectList',
    component: Layout,
    meta: {
      title: '项目管理',
      icon: renderIcon(FileProtectOutlined),
      sort: 1,
      permissions: ['developer'],
    },
    children: [
      {
        path: 'projectList',
        name: 'projectList',
        meta: {
          title: '项目列表',
          permissions: ['developer'],
        },
        component: () => import('@/views/project/projectList.vue'),
      },
      {
        path: 'loginList',
        name: 'loginList',
        meta: {
          title: '登录规则',
          permissions: ['developer'],
        },
        component: () => import('@/views/project/loginList.vue'),
      },
      {
        path: 'versionList',
        name: 'versionList',
        meta: {
          title: '版本号管理',
          permissions: ['developer'],
        },
        component: () => import('@/views/project/versionList.vue'),
      }
    ],
  },
];

export default routes;
