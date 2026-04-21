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


export const orderColumns = [
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
    title: '归属类型',
    key: 'cards_id',
    width: 110,
    render(row) {
      return h(
        'span', {}, cardList.value[row.cards_id])
    }
  },
  {
    title: '操作人',
    key: 'manager_id',
    width: 110,
    render(row) {
      return h(
        'span', {}, agentList.value[row['manager_id']])
    }
  },
  {
    title: '数量',
    key: 'count',
    width: 100,
  },
  {
    title: '金额',
    key: 'price',
    width: 150,
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

