import { RouteRecordRaw } from 'vue-router';
import { Layout } from '@/router/constant';
import { renderIcon } from '@/utils/index';
import {Users} from '@vicons/tabler'
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
    path: '/agent',
    name: 'agent',
    redirect: '/agent/agentList',
    component: Layout,
    meta: {
      title: '代理管理',
      icon: renderIcon(Users),
      sort: 5,
      permissions: ['developer'],
    },
    children: [
      {
        path: 'agentCreate',
        name: 'agentCreate',
        meta: {
          title: '创建代理',
          permissions: ['developer'],
        },
        component: () => import('@/views/agent/agentCreate.vue'),
      },
      {
        path: 'agentList',
        name: 'agentList',
        meta: {
          title: '代理列表',
          permissions: ['developer'],
        },
        component: () => import('@/views/agent/agentList.vue'),
      }
    ],
  },
];

export default routes;
