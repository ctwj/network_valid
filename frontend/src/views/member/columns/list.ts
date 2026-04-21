import {h, computed} from 'vue';
import {NTime} from 'naive-ui';
import {useAppStore} from "@/store/modules/app";

const appStore = useAppStore()
appStore.fetchProjectList(false)
appStore.fetchAgentList(false)
const projectList = computed(() => {
  let data = JSON.parse(JSON.stringify(appStore.getProjectList))
  let arr = {0: "通用项目"}
  for (let i in data) {
    arr[data[i].id] = data[i].name
  }
  return arr
})

const agentList = computed(() => {
  let data = JSON.parse(JSON.stringify(appStore.getAgentList))
  let arr = {}
  for (let i in data) {
    arr[data[i].id] = data[i].user
  }
  return arr
})

export const clientColumns = [

  {
    title: '归属软件',
    key: 'project_id',
    width: 110,
    render(row) {
      return h(
        'span', {}, projectList.value[row.project_id])
    }
  },
  {
    title: '账号',
    key: 'name',
    width: 130,
  },
  {
    title: '机器码',
    key: 'mac',
    width: 130,
  },
  {
    title: 'IP地址',
    key: 'ip',
    width: 140,
  },
  {
    title: '心跳时间',
    key: 'client_time',
    width: 180,
    render(row) {
      if (Number(row.client_time) > 0) {
        return h(
          NTime, {
            time: new Date(row.client_time * 1000),
            type: 'datetime'
          })
      } else {
        return h('span', "暂未使用")
      }

    }
  }
];

export const memberColumns = [
  {
    type: 'selection',
  },
  {
    title: '#',
    key: 'id',
    width: 100,
  },
  {
    title: '状态',
    key: 'is_lock',
    width: 70,
  },
  {
    title: '绑定',
    key: 'bind',
    width: 90,
  },
  {
    title: '归属软件',
    key: 'project_id',
    width: 110,
    render(row) {
      return h(
        'span', {}, projectList.value[row.project_id])
    }
  },
  {
    title: '归属代理',
    key: 'manager_id',
    width: 110,
    render(row) {
      return h(
        'span', {}, agentList.value[row['manager_id']])
    }
  },
  {
    title: '过期时间',
    key: 'end_time',
    width: 180,
    render(row) {
      if (Number(row.end_time) > 0) {
        return h(
          NTime, {
            time: new Date(row.end_time * 1000),
            type: 'datetime'
          })
      } else {
        return h('span', "暂未使用")
      }

    }
  },
  {
    title: '账号',
    key: 'name',
    width: 130,
  },
  {
    title: '机器码',
    key: 'mac',
    width: 130,
  },
  {
    title: '天数',
    key: 'days',
    width: 100,
  },
  {
    title: '点数',
    key: 'points',
    width: 100,
  },
  {
    title: '标签',
    key: 'tag',
    width: 100,
  },
  {
    title: '附属性',
    key: 'key_ext_attr',
    width: 140,
  },
  {
    title: 'IP地址',
    key: 'last_login_ip',
    width: 140,
  },
  {
    title: '登录时间',
    key: 'last_login_time',
    width: 180,
    render(row) {
      if (Number(row.last_login_time) > 0) {
        return h(
          NTime, {
            time: new Date(row.last_login_time * 1000),
            type: 'datetime'
          })
      } else {
        return h('span', "暂未使用")
      }

    }
  },
  {
    title: '创建时间',
    key: 'create_time',
    width: 180,
    render(row) {
      return h(
        NTime, {
          time: new Date(row.create_time),
          type: 'datetime',
        })
    }
  },

];
export const onlineColumns = [
  {
    title: '归属软件',
    key: 'project_id',
    width: 110,
    render(row) {
      return h(
        'span', {}, projectList.value[row.project_id])
    }
  },
  {
    title: '归属代理',
    key: 'manager_id',
    width: 110,
    render(row) {
      return h(
        'span', {}, agentList.value[row['manager_id']])
    }
  },
  {
    title: '账号',
    key: 'name',
    width: 130,
  },
  {
    title: '机器码',
    key: 'mac',
    width: 180,
  },
  {
    title: '识别码',
    key: 'client',
    width: 180,
  },
  {
    title: 'IP地址',
    key: 'ip',
    width: 140,
  },
  {
    title: '心跳时间',
    key: 'clienttime',
    width: 180,
    render(row) {
      if (Number(row.clienttime) > 0) {
        return h(
          NTime, {
            time: new Date(row.clienttime * 1000),
            type: 'datetime'
          })
      } else {
        return h('span', "暂未使用")
      }

    }
  },
  {
    title: '登录时间',
    key: 'addtime',
    width: 180,
    render(row) {
      if (row['addtime'] != undefined && Number(row.addtime) > 0) {
        return h(
          NTime, {
            time: new Date(row.addtime * 1000),
            type: 'datetime'
          })
      } else {
        return h('span', "")
      }

    }
  }

];

