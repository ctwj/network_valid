import { http } from '@/utils/http/axios';

export interface BasicResponseModel<T = any> {
  code: number;
  message: string;
  result: T;
}

export interface BasicPageParams {
  pageNumber: number;
  pageSize: number;
  total: number;
}

/**
 * @description: 获取用户信息
 */
export function getUserInfo() {
  return http.request({
    url: '/user/info',
    method: 'POST',
  });
}

/**
 * @description: 获取统计信息
 */
export function getInfo() {
  return http.request({
    url: '/user/getInfo',
    method: 'POST',
  });
}

export function getSysNotice() {
  return http.request({
    url: '/user/getSysNotice',
    method: 'post',
  });
}

export function getMemberEcharts() {
  return http.request({
    url: '/user/getMemberEcharts',
    method: 'post',
  });
}

export function getKeysEcharts() {
  return http.request({
    url: '/user/getKeysEcharts',
    method: 'post',
  });
}

export function update(params) {
  return http.request({
    url: '/user/update',
    method: 'post',
    params
  }, {
    isShowSuccessMessage: true
  });
}

/**
 * @description: 用户登录
 */
export function login(params) {
  return http.request<BasicResponseModel>(
    {
      url: '/user/login',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 用户修改密码
 */
export function changePassword(params, uid) {
  return http.request(
    {
      url: `/user/u${uid}/changepw`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 用户登出
 */
export function logout(params) {
  return http.request({
    url: '/login/logout',
    method: 'POST',
    params,
  });
}
