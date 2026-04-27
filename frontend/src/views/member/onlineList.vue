<template>
  <n-card title="">
    <BasicTable
      :columns="columns"
      :request="loadDataTable"
      :row-key="(row) => row.uid"
      ref="actionRef"
      children-key="client_list"
      :actionColumn="actionColumn"
      :scroll-x="1200"

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
                         placeholder="请输入需查询的账号或者兑换码" clearable/>
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
import {useMessage, useDialog} from 'naive-ui'
import {computed, h, reactive, ref} from 'vue';
import {BasicTable, TableAction} from '@/components/Table';
import {getOnlineList,  memberLogout} from "@/api";
import {onlineColumns} from './columns/list'
import {useAppStore} from "@/store/modules/app";

const appStore = useAppStore()
const columns = onlineColumns
const message = useMessage()
const dialog = useDialog()
const actionRef = ref();
const list = ref([]);
const title = ref('');
const row = ref(null)
const show = ref(false);

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
  width: 100,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: record['client_list'] === null?"未上线": record['client_list'] !== undefined && record['client_list'] !== null ? "全部下线" : "下线",
          onClick: handleLogout.bind(null, record),
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
  member: "",
});

const params = ref({
  limit: 10,
  page: 1,
});


async function handleLogout(record: Recordable) {
  dialog.info({
    title: '温馨提示',
    content: '确定下线？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: async () => {
      await memberLogout({project_id: record.project_id, id: record.member_id, client: record.client})
      reloadTable()
    },
    onNegativeClick: () => {
      message.info('取消操作')
    }
  })
}

const loadDataTable = async (res) => {
  let data = await getOnlineList({...formParams, ...params.value, ...res});
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
