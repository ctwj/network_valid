import { RouteRecordRaw } from 'vue-router';
import { Layout } from '@/router/constant';
import { renderIcon } from '@/utils/index';
import {PersonOutline} from '@vicons/ionicons5';
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
    path: '/member',
    name: 'member',
    redirect: '/member/memberList',
    component: Layout,
    meta: {
      title: '会员管理',
      icon: renderIcon(PersonOutline),
      sort: 2,
    },
    children: [
      {
        path: 'memberList',
        name: 'memberList',
        meta: {
          title: '会员列表',
        },
        component: () => import('@/views/member/memberList.vue'),
      },
      {
        path: 'onlineList',
        name: 'onlineList',
        meta: {
          title: '最近在线',
        },
        component: () => import('@/views/member/onlineList.vue'),
      }
    ],
  },
];

export default routes;
