import {http} from '@/utils/http/axios';

export function getProjectList(params) {
  return http.request({
    url: '/project/getProjectList',
    method: 'post',
    params,
  });
}

export function projectList() {
  return http.request({
    url: '/project/projectList',
    method: 'post',
    params: {},
  });
}

export function createProject(params) {
  return http.request({
    url: '/project/createProject',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function getPlanSchemes() {
  return http.request({
    url: '/project/getPlanSchemes',
    method: 'post',
  });
}

export function updateProject(params) {
  return http.request({
    url: '/project/updateProject',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function deleteProject(id) {
  return http.request({
    url: '/project/deleteProject',
    method: 'post',
    params: {id},
  }, {
    isShowSuccessMessage: true
  });
}

export function bindProjectLogin(params) {
  return http.request({
    url: '/project/bindProjectLogin',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function loginRuleList() {
  return http.request({
    url: '/project/loginRuleList',
    method: 'post',
  });
}

export function getLoginRuleList(params) {
  return http.request({
    url: '/project/getLoginRuleList',
    method: 'post',
    params,
  });
}

export function createLoginRule(params) {
  return http.request({
    url: '/project/createLoginRule',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function updateLoginRule(params) {
  return http.request({
    url: '/project/updateLoginRule',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function deleteLoginRule(id) {
  return http.request({
    url: '/project/deleteLoginRule',
    method: 'post',
    params: {id},
  }, {
    isShowSuccessMessage: true
  });
}

export function getVersionList(params) {
  return http.request({
    url: '/project/getVersionList',
    method: 'post',
    params,
  });
}

export function createVersion(params) {
  return http.request({
    url: '/project/createVersion',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function updateProjectVersion(params) {
  return http.request({
    url: '/project/updateProjectVersion',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function deleteProjectVersion(id) {
  return http.request({
    url: '/project/deleteProjectVersion',
    method: 'post',
    params: {id},
  }, {
    isShowSuccessMessage: true
  });
}

export function getCardList(params) {
  return http.request({
    url: '/project/getCardList',
    method: 'post',
    params,
  });
}

export function cardList() {
  return http.request({
    url: '/project/cardList',
    method: 'post',
  });
}

export function deleteCard(id) {
  return http.request({
    url: '/project/deleteCard',
    method: 'post',
    params: {id},
  }, {
    isShowSuccessMessage: true
  });
}

export function createCard(params) {
  return http.request({
    url: '/project/createCard',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function updateCard(params) {
  return http.request({
    url: '/project/updateCard',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function queryOrderRmb(params) {
  return http.request({
    url: '/project/queryOrderRmb',
    method: 'post',
    params,
  });
}

export function createKeys(params) {
  return http.request({
    url: '/project/createKeys',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function getKeysList(params) {
  return http.request({
    url: '/project/getKeysList',
    method: 'post',
    params,
  });
}

export function batchKeys(params) {
  return http.request({
    url: '/project/batchKeys',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function deleteKeys(id) {
  return http.request({
    url: '/project/deleteKeys',
    method: 'post',
    params: {id},
  }, {
    isShowSuccessMessage: true
  });
}


export function lockKey(id) {
  return http.request({
    url: '/project/lockKey',
    method: 'post',
    params: {id},
  }, {
    isShowSuccessMessage: true
  });
}



export function getMemberList(params) {
  return http.request({
    url: '/project/getMemberList',
    method: 'post',
    params,
  });
}

export function batchMember(params) {
  return http.request({
    url: '/project/batchMember',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function deleteMember(id) {
  return http.request({
    url: '/project/deleteMember',
    method: 'post',
    params: {id},
  }, {
    isShowSuccessMessage: true
  });
}

export function lockMember(id) {
  return http.request({
    url: '/project/lockMember',
    method: 'post',
    params: {id},
  }, {
    isShowSuccessMessage: true
  });
}

export function updateMember(params) {
  return http.request({
    url: '/project/updateMember',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function unbindMember(params) {
  return http.request({
    url: '/project/unbindMember',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function getOnlineList(params) {
  return http.request({
    url: '/project/getOnlineList',
    method: 'post',
    params,
  });
}

export function memberLogout(params) {
  return http.request({
    url: '/project/memberLogout',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}


export function getRoleUser(params) {
  return http.request({
    url: '/project/getRoleUser',
    method: 'post',
    params,
  });
}

export function rolePower(params) {
  return http.request({
    url: '/project/rolePower',
    method: 'post',
    params,
  });
}

export function createRoleUser(params) {
  return http.request({
    url: '/project/createRoleUser',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function updateRoleUser(params) {
  return http.request({
    url: '/project/updateRoleUser',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function getUserRole(params) {
  return http.request({
    url: '/project/getUserRole',
    method: 'post',
    params,
  });
}

export function deleteRole(params) {
  return http.request({
    url: '/project/deleteRole',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function getRole() {
  return http.request({
    url: '/project/getRole',
    method: 'post',
  });
}

export function managerAdd(params) {
  return http.request({
    url: '/project/managerAdd',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}


export function getManagerList(params) {
  return http.request({
    url: '/project/getManagerList',
    method: 'post',
    params
  });
}

export function managerUpdate(params) {
  return http.request({
    url: '/project/managerUpdate',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function managerDelete(params) {
  return http.request({
    url: '/project/managerDelete',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}


export function getManagerCards(params) {
  return http.request({
    url: '/project/getManagerCards',
    method: 'post',
    params,
  });
}

export function addManagerCards(params) {
  return http.request({
    url: '/project/addManagerCards',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}

export function updateManagerCards(params) {
  return http.request({
    url: '/project/updateManagerCards',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}


export function deleteManagerCards(params) {
  return http.request({
    url: '/project/deleteManagerCards',
    method: 'post',
    params,
  }, {
    isShowSuccessMessage: true
  });
}


export function updateCache() {
  return http.request({
    url: '/project/updateCache',
    method: 'post',
  }, {
    isShowSuccessMessage: true
  });
}

export function getOrder(params) {
  return http.request({
    url: '/project/getOrder',
    method: 'post',
    params,
  });
}

export function getAgent() {
  return http.request({
    url: '/project/getAgent',
    method: 'post',
  });
}
