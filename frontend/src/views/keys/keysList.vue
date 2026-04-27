<template>
  <n-card title="">
    <BasicTable
      :columns="columns"
      :request="loadDataTable"
      :row-key="(row) => row.id"
      ref="actionRef"
      :actionColumn="actionColumn"
      :scroll-x="1090"
      @update:checked-row-keys="handleCheck"
    >
      <template #form>
        <n-form
          ref="formRef"
          :model="formParams"
          label-placement="left"
          label-width="100"
          require-mark-placement="right-hanging"
        >
          <n-grid cols="1 s:2 m:4 l:4 xl:4 2xl:4" responsive="screen">
            <n-grid-item>
              <n-form-item label="归属项目" path="inputValue">
                <n-select label-field="name" value-field="id" v-model:value="formParams.project_id"
                          :options="projectList" placeholder="请选择归属项目"
                          :on-update:value="updateProjectId"/>
              </n-form-item>
            </n-grid-item>
            <n-grid-item>
              <n-form-item label="套餐类型" path="title">
                <n-select label-field="title" value-field="id" v-model:value="formParams.cards_id"
                          :options="cardList" placeholder="请选择归属套餐类型"/>
              </n-form-item>
            </n-grid-item>
            <n-grid-item>
              <n-form-item label="兑换码" path="long_keys">
                <n-input :maxlength="32" v-model:value="formParams.long_keys"
                         placeholder="请输入需查询的兑换码" clearable/>
              </n-form-item>
            </n-grid-item>
            <n-grid-item>
              <n-form-item label="是否使用" path="is_active">
                <n-select label-field="title" value-field="id" v-model:value="formParams.is_active"
                          :options="[{title:'所有',id:-1},{title:'未使用',id:0},{title:'已使用',id:1}]"
                          placeholder="请选择是否使用"/>
              </n-form-item>
            </n-grid-item>
            <n-grid-item>
              <n-form-item label="是否锁定" path="is_lock">
                <n-select label-field="title" value-field="id" v-model:value="formParams.is_lock"
                          :options="[{title:'所有',id:-1},{title:'未锁定',id:0},{title:'已锁定',id:1}]"
                          placeholder="请选择是否使用"/>
              </n-form-item>
            </n-grid-item>
            <n-grid-item>
              <n-form-item label="关联用户" path="member">
                <n-input :maxlength="32" v-model:value="formParams.member"
                         placeholder="请输入用户账号" clearable/>
              </n-form-item>
            </n-grid-item>
            <n-grid-item>
              <n-form-item label="订单号" path="order_id">
                <n-input-number v-model:value="formParams.order_id" :min="1" :max="999999999"
                                :step="1" clearable placeholder="请输入兑换码订单号"/>
              </n-form-item>
            </n-grid-item>
            <n-grid-item>
              <n-form-item label="操作">
                <component :is="Batch" v-model:rows="checkedRowKeysRef" v-model:form="formParams" @on-update="reloadTable"></component>
              </n-form-item>
            </n-grid-item>
          </n-grid>
        </n-form>
      </template>
      <template #toolbar>
        <n-button type="primary" @click="reloadTable">查询</n-button>
      </template>
    </BasicTable>
  </n-card>
</template>

<script lang="ts" setup>
import {useMessage, useDialog, NTag, DataTableRowKey} from 'naive-ui'
import {computed, h, reactive, ref} from 'vue';
import {BasicTable, TableAction} from '@/components/Table';
import {getKeysList, deleteKeys, lockKey} from "@/api";
import {keysColumns} from './columns/list'
import {useAppStore} from "@/store/modules/app";
import Batch from "@/views/keys/comp/batch.vue";

for (let i in keysColumns) {
  if (keysColumns[i].title === "状态") {
    keysColumns[i]['render'] = (row) => {
      return h(
        NTag, {
          type: row.is_lock === 0 ? "success" : "error",
          size: "small",
          onClick: async (v) => {
            await lockKey(row.id)
            row.is_lock = row.is_lock === 0 ? 1 : 0
          }
        }, ["正常", "锁定"][row.is_lock])
    }
  }
}
const checkedRowKeysRef = ref<DataTableRowKey[]>([])
const appStore = useAppStore()
const columns = keysColumns
const message = useMessage()
const dialog = useDialog()
const actionRef = ref();
const list = ref([]);
const title = ref('');
const row = ref(null)
const show = ref(false);
const projectList = computed(() => {
  let data = JSON.parse(JSON.stringify(appStore.getProjectList))
  data.unshift({name: "通用", id: 0})
  data.unshift({name: "所有", id: -1})
  return data
})

const cardAllList = computed(() => {
  return appStore.getCardList
})
const cardList = ref([]);
cardList.value = cardAllList.value
const actionColumn = reactive({
  width: 80,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
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
  project_id: -1,
  long_keys: null,
  cards_id: null,
  is_active: -1,
  is_lock: -1,
  member: "",
  order_id: null,
  id: "",
  type: "",
  lock: 0
});

const params = ref({
  limit: 10,
  page: 1,
});

function handleEdit(record: Recordable) {
  show.value = true;
  row.value = record
  title.value = "修改套餐类型"
  // router.push({ name: 'basic-info', params: { id: record.id } });
}

async function handleDelete(record: Recordable) {
  dialog.warning({
    title: '温馨提示',
    content: '确定删除？(操作不可逆)',
    positiveText: '删除',
    negativeText: '取消',
    onPositiveClick: async () => {
      await deleteKeys(record.id)
      reloadTable()
    },
    onNegativeClick: () => {
      message.info('取消操作')
    }
  })
}

const loadDataTable = async (res) => {
  let data = await getKeysList({...formParams, ...params.value, ...res});
  console.log(data)
  return data
};

function addTable() {
  show.value = true;
  row.value = null;
  title.value = "创建新套餐类型"
}

function handleCheck(rowKeys: DataTableRowKey[]) {
  checkedRowKeysRef.value = rowKeys
}

function reloadTable() {
  actionRef.value.fetchData()
}

function updateProjectId(val) {
  let list = cardAllList.value
  let arr = []
  if (val !== null) {
    for (let i in list) {
      console.log(list[i].project_id, val)
      if (list[i].project_id === val) {
        arr.push(list[i])
      }
    }
  }
  formParams.project_id = val
  cardList.value = arr

}
</script>

<style scoped>

</style>
