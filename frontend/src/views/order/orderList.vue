<template>
  <n-card title="">
    <BasicTable
      :columns="columns"
      :request="loadDataTable"
      :row-key="(row) => row.id"
      ref="actionRef"
      :scroll-x="1090"
    >
      <template #toolbar>
        <n-button type="primary" @click="reloadTable">刷新数据</n-button>
      </template>
    </BasicTable>
  </n-card>
</template>

<script lang="ts" setup>
import {useMessage, useDialog} from 'naive-ui'
import {h, reactive, ref, Ref} from 'vue';
import {BasicTable, TableAction} from '@/components/Table';
import {getOrder,deleteRole} from "@/api";
import {orderColumns as columns} from './columns/list'

const message = useMessage()
const dialog = useDialog()
const actionRef = ref();
const list = ref([]);
const title = ref('');
const row = ref(null)
const show = ref(false);
// const actionColumn = reactive({
//   width: 70,
//   title: '操作',
//   key: 'action',
//   fixed: 'right',
//   render(record) {
//     return h(TableAction as any, {
//       style: 'button',
//       actions: [
//         {
//           label: '删除',
//           onClick: handleDelete.bind(null, record),
//           // 根据业务控制是否显示 isShow 和 auth 是并且关系
//           ifShow: () => {
//             return true;
//           },
//         },
//       ]
//     });
//   },
// });
const formParams = reactive({
  name: '',
  address: '',
  date: null,
});

const params = ref({
  limit: 10,
  page: 1,
});


async function handleDelete(record: Recordable) {
  dialog.warning({
    title: '温馨提示',
    content: '删除角色将会影响所有已给予权限的代理（操作不可逆，请谨慎操作）',
    positiveText: '删除',
    negativeText: '取消',
    onPositiveClick: async () => {
      await deleteRole({id: record.id})
      reloadTable()
    },
    onNegativeClick: () => {
      message.info('取消操作')
    }
  })
}

const loadDataTable = async (res) => {
  let data = await getOrder({...formParams, ...params.value, ...res});
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
