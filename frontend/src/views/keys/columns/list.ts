import {h, computed} from 'vue';
import {NTime} from 'naive-ui';
import {useAppStore} from "@/store/modules/app";

const appStore = useAppStore()
appStore.fetchProjectList(false)
appStore.fetchCardList(false)
appStore.fetchAgentList(false)
const projectList = computed(() => {
  let data = JSON.parse(JSON.stringify(appStore.getProjectList))
  let arr = {0: "通用项目"}
  for (let i in data) {
    arr[data[i].id] = data[i].name
  }
  return arr
})

const cardList = computed(() => {
  let data = JSON.parse(JSON.stringify(appStore.getCardList))
  let arr = {}
  for (let i in data) {
    arr[data[i].id] = data[i].title
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


export const typeColumns = [
  {
    title: '#',
    key: 'id',
    width: 100,
  },
  {
    title: '归属软件',
    key: 'project_id',
    width: 120,
    render(row) {
      return h(
        'span', {}, projectList.value[row.project_id])
    }
  },
  {
    title: '激活码名称',
    key: 'title',
    width: 100,
  },
  {
    title: '前缀',
    key: 'key_prefix',
    width: 150,
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
    title: '定价',
    key: 'price',
    width: 100,
  },
  {
    title: '标签',
    key: 'tag',
    width: 100,
  },
  {
    title: '附加属性',
    key: 'key_ext_attr',
    width: 100,
  },
  {
    title: '创建时间',
    key: 'create_time',
    width: 180,
    render(row) {
      return h(
        NTime, {
          time: new Date(row.create_time),
          type: 'datetime'
        })
    }
  },
];
export const keysColumns = [
  {
    type: 'selection',
  },
  {
    title: '#',
    key: 'id',
    width: 100,
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
    title: '归属类型',
    key: 'cards_id',
    width: 110,
    render(row) {
      return h(
        'span', {}, cardList.value[row.cards_id])
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
    title: '激活码',
    key: 'long_keys',
    width: 130,
  },
  {
    title: '充值用户',
    key: 'member',
    width: 130,
    render(row) {
      if (row.member === "") {
        return h("span", "暂未使用")
      } else {
        return h("span", row.member)
      }
    }
  },
  {
    title: '充值时间',
    key: 'use_time',
    width: 160,
    render(row) {
      if (Number(row.use_time) > 0) {
        return h(
          NTime, {
            time: new Date(row.use_time * 1000),
            type: 'datetime'
          })
      } else {
        return h('span', "暂未使用")
      }

    }
  },
  {
    title: '激活码订单',
    key: 'order_id',
    width: 100,
  },
  {
    title: '标签',
    key: 'tag',
    width: 100,
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
  {
    title: '状态',
    key: 'is_lock',
    width: 60,
  },
];
