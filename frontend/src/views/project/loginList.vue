<template>
  <n-card title="">
  <BasicTable
    :columns="loginListColumns"
    :request="loadDataTable"
    :row-key="(row) => row.id"
    ref="actionRef"
    :actionColumn="actionColumn"
    :scroll-x="1090"
  >
    <template #tableTitle>
      <n-button type="primary" @click="addTable">
        <template #icon>
          <n-icon>
            <PlusOutlined/>
          </n-icon>
        </template>
        新建
      </n-button>
    </template>

    <template #toolbar>
      <n-button type="primary" @click="reloadTable">刷新数据</n-button>
    </template>
  </BasicTable>
  <component :is="loginEdit" v-model:show="show" v-model:form="row"
             v-model:title="title" @on-update="reloadTable"></component>
  </n-card>
</template>

<script lang="ts" setup>
import {useMessage, useDialog} from 'naive-ui'
import {h, reactive, ref, Ref} from 'vue';
import {BasicTable, TableAction} from '@/components/Table';
import {PlusOutlined} from '@vicons/antd';
import {deleteLoginRule, getLoginRuleList} from "@/api";
import {loginListColumns} from './columns/list'
import loginEdit from "@/views/project/comp/loginEdit.vue";

const message = useMessage()
const dialog = useDialog()
const actionRef = ref();
const list = ref([]);
const title = ref('');
const row = ref(null)
const show = ref(false);
const actionColumn = reactive({
  width: 140,
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
  title.value = "修改登录规则"
  // router.push({ name: 'basic-info', params: { id: record.id } });
}

async function handleDelete(record: Recordable) {
  dialog.warning({
    title: '温馨提示',
    content: '确定删除？(操作不可逆)',
    positiveText: '删除',
    negativeText: '取消',
    onPositiveClick: async () => {
      await deleteLoginRule(record.id)
      reloadTable()
    },
    onNegativeClick: () => {
      message.info('取消操作')
    }
  })
}

const loadDataTable = async (res) => {
  let data = await getLoginRuleList({...formParams, ...params.value, ...res});
  console.log(data)
  return data
};

function addTable() {
  show.value = true;
  row.value = null;
  title.value = "创建新登录规则"
}

function reloadTable() {
  actionRef.value.fetchData()
}
</script>

<style scoped>

</style>
