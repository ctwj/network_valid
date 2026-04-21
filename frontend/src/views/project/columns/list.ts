import {computed, h} from 'vue';
import {NTag, NTime} from 'naive-ui';
import {useAppStore} from "@/store/modules/app";
const appStore = useAppStore()
const projectList = computed(() => {
  let data = JSON.parse(JSON.stringify(appStore.getProjectList))
  let arr = {}
  for (let i in data){
    arr[data[i].id] = data[i].name
  }
  return arr
})

export const columns = [
  {
    title: '#',
    key: 'id',
    width: 100,
  },
  {
    title: '项目名称',
    key: 'name',
    width: 100,
  },
  {
    title: '地址',
    key: 'address',
    auth: ['basic_list'], // 同时根据权限控制是否显示
    ifShow: (_column) => {
      return true; // 根据业务控制是否显示
    },
    width: 150,
  },
  {
    title: '开始日期',
    key: 'beginTime',
    width: 160,
  },
  {
    title: '结束日期',
    key: 'endTime',
    width: 160,
  },
  {
    title: '创建时间',
    key: 'create_time',
    width: 100,
  },
];


export const loginColumns = [
  {
    title: '#',
    key: 'id',
    width: 100,
  },
  {
    title: '规则名称',
    key: 'title',
  },
  {
    title: '规则模式',
    key: 'mode',
    width: 100,
    render(row) {
      return h(
        NTag, {
          type: row.mode === 0 ? "success" : row.mode === 1 ? "info" : "warning",
        }, ["绑定登录", "普通登录", "点数登录"][row.mode])
    }
  },
  {
    title: '解绑规则',
    key: 'UnbindMode',
    width: 120,
    render(row) {
      return h(
        'span', {}, ["不允许解绑", "原机解绑", "自动解绑", "任意解绑"][row.unbind_mode])
    }
  },
  {
    title: '解绑扣时',
    key: 'UnbindWeakenMode',
    width: 120,
    render(row) {
      return h(
        'span', {}, ["不扣时", "解绑就扣时", "超出扣时", "超出不扣时"][row.unbind_weaken_mode])
    }
  }
];


export const loginListColumns = [
  {
    title: '#',
    key: 'id',
    width: 100,
  },
  {
    title: '规则名称',
    key: 'title',
  },
  {
    title: '规则模式',
    key: 'mode',
    width: 100,
    render(row) {
      return h(
        NTag, {
          type: row.mode === 0 ? "success" : row.mode === 1 ? "info" : "warning",
        }, ["绑定登录", "普通登录", "点数登录"][row.mode])
    }
  },
  {
    title: '解绑规则',
    key: 'UnbindMode',
    width: 120,
    render(row) {
      return h(
        'span', {}, ["不允许解绑", "原机解绑", "自动解绑", "任意解绑"][row.unbind_mode])
    }
  },
  {
    title: '解绑扣时',
    key: 'UnbindWeakenMode',
    width: 120,
    render(row) {
      return h(
        'span', {}, ["不扣时", "解绑就扣时", "超出扣时", "超出不扣时"][row.unbind_weaken_mode])
    }
  }
];

export const versionColumns = [
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
    title: '版本号',
    key: 'version',
  },
  {
    title: '强制更新',
    key: 'is_must_update',
    width: 120,
    render(row) {
      return h(
        'span', {}, ["强制", "不强制"][row.is_must_update])
    }
  },
  {
    title: '是否启用',
    key: 'is_active',
    width: 120,
    render(row) {
      return h(
        'span', {}, ["启用", "不启用"][row.is_active])
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
          type: 'datetime'
        })
    }
  },
];
