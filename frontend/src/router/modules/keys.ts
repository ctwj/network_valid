import { RouteRecordRaw } from 'vue-router';
import { Layout } from '@/router/constant';
import { renderIcon } from '@/utils/index';
import {CardOutline} from '@vicons/ionicons5';
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
    path: '/keys',
    name: 'Keys',
    redirect: '/keys/typeList',
    component: Layout,
    meta: {
      title: '激活码管理',
      icon: renderIcon(CardOutline),
      sort: 2,
    },
    children: [
      {
        path: 'typeList',
        name: 'typeList',
        meta: {
          title: '激活码类型',
          permissions: ['developer'],
        },
        component: () => import('@/views/keys/typeList.vue'),
      },
      {
        path: 'keysCreate',
        name: 'keysCreate',
        meta: {
          title: '创建激活码',
        },
        component: () => import('@/views/keys/keysCreate.vue'),
      },
      {
        path: 'keysList',
        name: 'keysList',
        meta: {
          title: '激活码列表',
        },
        component: () => import('@/views/keys/keysList.vue'),
      }
    ],
  },
];

export default routes;
