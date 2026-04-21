<template>
  <n-card title="">
    <BasicTable
      :columns="columns"
      :request="loadDataTable"
      :row-key="(row) => row.id"
      ref="actionRef"
      :actionColumn="actionColumn"
      @update:checked-row-keys="handleCheck"
      :scroll-x="1980"
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
              <n-form-item label="账号" path="member">
                <n-input :maxlength="32" v-model:value="formParams.member"
                         placeholder="请输入需查询的账号或者激活码" clearable/>
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
              <n-form-item label="机器码" path="mac">
                <n-input :maxlength="32" v-model:value="formParams.mac"
                         placeholder="请输入需查询的机器码" clearable/>
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
    <component :is="memberEdit" v-model:show="show" v-model:form="row"
               v-model:title="title" @on-update="reloadTable"></component>
  </n-card>
</template>

<script lang="ts" setup>
import {useMessage, useDialog, NTag, DataTableRowKey} from 'naive-ui'
import {computed, h, reactive, ref} from 'vue';
import {BasicTable, TableAction} from '@/components/Table';
import {getMemberList, deleteMember, lockMember, unbindMember} from "@/api";
import {memberColumns} from './columns/list'
import {useAppStore} from "@/store/modules/app";
import memberEdit from "@/views/member/comp/memberEdit.vue";
import Batch from "@/views/member/comp/batch.vue";
const appStore = useAppStore()
const columns = memberColumns
const message = useMessage()
const dialog = useDialog()
const actionRef = ref();
const list = ref([]);
const title = ref('');
const row = ref(null)
const show = ref(false);
for (let i in memberColumns) {
  if (memberColumns[i].title === "状态") {
    memberColumns[i]['render'] = (row) => {
      return h(
        NTag, {
          type: row.is_lock === 0 ? "success" : "error",
          size: "small",
          onClick: async (v) => {
            await lockMember(row.id)
            row.is_lock = row.is_lock === 0 ? 1 : 0
          }
        }, ["正常", "锁定"][row.is_lock])
    }
  }
  if (memberColumns[i].title === "绑定") {
    memberColumns[i]['render'] = (row) => {
      return h(
        NTag, {
          title: row.mac === ""?"已解绑":row.mac,
          type: row.mac === "" ? "success" : "info",
          size: "small",
          onClick: async (v) => {
            if (row.mac !== "") {
              await unbindMember({id: row.id})
              actionRef.value.reload()
            }
          }
        }, ["已解绑", "解绑"][row.mac === "" ? 0 : 1])
    }
  }
}
const checkedRowKeysRef = ref<DataTableRowKey[]>([])
const projectList = computed(() => {
  let data = JSON.parse(JSON.stringify(appStore.getProjectList))
  data.unshift({name: "所有", id: -1})
  return data
})
const cardAllList = computed(() => {
  return appStore.getCardList
})
const cardList = ref([]);
cardList.value = cardAllList.value
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
  project_id: -1,
  is_lock: -1,
  member: "",
  mac: ""
});

const params = ref({
  limit: 10,
  page: 1,
});

function handleEdit(record: Recordable) {
  show.value = true;
  row.value = record
  title.value = "修改会员资料"
  // router.push({ name: 'basic-info', params: { id: record.id } });
}

function handleCheck(rowKeys: DataTableRowKey[]) {
  checkedRowKeysRef.value = rowKeys
}

async function handleDelete(record: Recordable) {
  dialog.warning({
    title: '温馨提示',
    content: '确定删除？(操作不可逆)',
    positiveText: '删除',
    negativeText: '取消',
    onPositiveClick: async () => {
      await deleteMember(record.id)
      reloadTable()
    },
    onNegativeClick: () => {
      message.info('取消操作')
    }
  })
}

const loadDataTable = async (res) => {
  let data = await getMemberList({...formParams, ...params.value, ...res});
  console.log(data)
  return data
};

function addTable() {
  show.value = true;
  row.value = null;
  title.value = "创建新激活码类型"
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
