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
    <component :is="typeEdit" v-model:show="show" v-model:form="row"
               v-model:title="title" @on-update="reloadTable"></component>
  </n-card>
</template>

<script lang="ts" setup>
import {useMessage, useDialog} from 'naive-ui'
import {h, reactive, ref, Ref} from 'vue';
import {BasicTable, TableAction} from '@/components/Table';
import {PlusOutlined} from '@vicons/antd';
import {getCardList, deleteCard} from "@/api";
import {typeColumns as columns} from './columns/list'
import typeEdit from "@/views/keys/comp/typeEdit.vue";

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
  title.value = "修改套餐类型"
  // router.push({ name: 'basic-info', params: { id: record.id } });
}

async function handleDelete(record: Recordable) {
  dialog.warning({
    title: '温馨提示',
    content: '删除套餐类型，并且删除所有已经创建好的兑换码！（操作不可逆，请谨慎操作）',
    positiveText: '删除',
    negativeText: '取消',
    onPositiveClick: async () => {
      await deleteCard(record.id)
      reloadTable()
    },
    onNegativeClick: () => {
      message.info('取消操作')
    }
  })
}

const loadDataTable = async (res) => {
  let data = await getCardList({...formParams, ...params.value, ...res});
  console.log(data)
  return data
};

function addTable() {
  show.value = true;
  row.value = null;
  title.value = "创建新套餐类型"
}

function reloadTable() {
  actionRef.value.fetchData()
}
</script>

<style scoped>

</style>
