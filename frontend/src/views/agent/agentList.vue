<template>
  <n-card title="">
    <BasicTable
      :columns="columns"
      :request="loadDataTable"
      :row-key="(row) => row.id"
      ref="actionRef"
      :actionColumn="actionColumn"
      :scroll-x="1090"
    >
      <template #toolbar>
        <n-button type="primary" @click="reloadTable">刷新数据</n-button>
      </template>
    </BasicTable>
    <component :is="agentEdit" v-model:show="show" v-model:form="row"
               v-model:title="title" @on-update="reloadTable"></component>
  </n-card>
</template>

<script lang="ts" setup>
import {useMessage, useDialog} from 'naive-ui'
import {h, reactive, ref} from 'vue';
import {BasicTable, TableAction} from '@/components/Table';
import {getManagerList,managerDelete} from "@/api";
import { agentColumns as columns} from './columns/list'
import agentEdit from "@/views/agent/comp/agentEdit.vue";
import {useAppStore} from "@/store/modules/app";

const appStore = useAppStore()
const message = useMessage()
const dialog = useDialog()
const actionRef = ref();
const list = ref([]);
const title = ref('');
const row = ref(null)
const show = ref(false);
const actionColumn = reactive({
  width: 100,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '编辑',
          onClick: handleEdit.bind(null, record),
          ifShow: () => {
            return true;
          }
        },
        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          // 根据业务控制是否显示 isShow 和 auth 是并且关系
          ifShow: () => {
            return true;
          },
        },
      ]
    });
  },
});
const formParams = reactive({
  name: '',
  address: '',
  date: null,
});

const params = ref({
  limit: 10,
  page: 1,
});

function handleEdit(record: Recordable) {
  console.log('点击了编辑', record);
  show.value = true;
  row.value = record
  title.value = "修改代理"
  // router.push({ name: 'basic-info', params: { id: record.id } });
}

async function handleDelete(record: Recordable) {
  dialog.warning({
    title: '温馨提示',
    content: '删除代理，将会清空代理名下激活码以及用户(请谨慎操作！)',
    positiveText: '删除',
    negativeText: '取消',
    onPositiveClick: async () => {
      await managerDelete({id: record.id})
      reloadTable()
      await appStore.fetchAgentList(true)
    },
    onNegativeClick: () => {
      message.info('取消操作')
    }
  })
}

const loadDataTable = async (res) => {
  let data = await getManagerList({...formParams, ...params.value, ...res});
  console.log(data)
  return data
};

function addTable() {
  show.value = true;
  row.value = null;
  title.value = "创建新角色"
}

function reloadTable() {
  actionRef.value.fetchData()
}
</script>

<style scoped>

</style>
