<template>
  <n-card title="">
    <n-form
      ref="formRef"
      :inline="false"
      :label-width="80"
      :model="formData"
      :rules="rules"
    >
      <div class="form-view">
        <n-form-item label="归属项目" path="project_id">
          <n-select label-field="name" value-field="id" v-model:value="formData.project_id"
                    :options="projectList" placeholder="请选择归属项目"
                    :on-update:value="updateProjectId"/>
        </n-form-item>
        <n-form-item label="激活码类型" path="title">
          <n-select label-field="title" value-field="id" v-model:value="formData.cards_id"
                    :options="cardList" placeholder="请选择归属激活码类型"
                    :on-update:value="updateCards"/>
        </n-form-item>
        <n-form-item label="激活码长度" path="length">
          <n-input-number v-model:value="formData.length" :min="8" :max="32"
                          :step="1" clearable/>
        </n-form-item>
        <n-form-item label="激活码组合格式" path="is_lock">
          <n-radio-group v-model:value="formData.create_type" name="is_lock">
            <n-radio v-for="(item,index) in project.keys_create_type" :value="index" :key="index">
              {{ item }}
            </n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="激活码数量" path="days">
          <n-input-number v-model:value="formData.count" :min="1" :max="500"
                          :step="1" clearable :on-update:value="updateCount"/>
        </n-form-item>
        <n-form-item label="绑定标签" path="tag">
          <n-input :maxlength="200" v-model:value="formData.tag"
                   placeholder="输入激活码绑定标签，不填则默认绑定激活码内置标签"/>
        </n-form-item>
      </div>
      <div>
        <n-statistic label="订单金额" tabular-nums>
          <n-number-animation
            ref="numberAnimationInstRef"
            :from="0.0"
            :to="rmb"
            :active="active"
            :precision="2"
          />
        </n-statistic>
        <n-tag size="small" v-if="queryMsg!== ''" :type="queryStatus?'success':'error'">
          {{ queryMsg }}
        </n-tag>
      </div>
      <n-form-item>
        <n-button type="primary" attr-type="button" @click="handleOk">
          {{ title }}
        </n-button>
        <span style="margin-left: 6px" v-if="userInfo.pid > 0">账号余额：{{ userInfo.money }}</span>
      </n-form-item>
      <n-form-item v-if="show" label="请复制下面的激活码">
        <n-input
          v-model:value="keys_str"
          type="textarea"
          placeholder="生成的激活码"
        />
        <n-button style="margin-left: 10px" @click="copy_inline(keys_str)">复制</n-button>
      </n-form-item>
    </n-form>
  </n-card>
</template>

<script lang="ts" setup>
import {reactive, onMounted, computed, ref} from "vue";
import {createKeys, queryOrderRmb} from '@/api/index'
import {useAppStore} from "@/store/modules/app";
import setting from '@/settings/componentSetting'
import {copy_inline} from "@/utils";
import {useMessage, NumberAnimationInst} from "naive-ui";
import {getUserInfo} from "@/api/system/user";

const numberAnimationInstRef = ref<NumberAnimationInst | null>(null)
const message = useMessage()
const {project} = setting
const appStore = useAppStore()
const projectList = computed(() => {
  let data = JSON.parse(JSON.stringify(appStore.getProjectList))
  data.unshift({name: "通用项目-所有项目都可用", id: 0})
  return data
})

const cardAllList = computed(() => {
  return appStore.getCardList
})
const rules = reactive({})
const cardList = ref([]);
const active = ref(false)
const formData = reactive({
  project_id: null,
  cards_id: null,
  length: 12,
  create_type: 0,
  count: 1,
  tag: ""
})
const keys_str = ref('')
const keys_list = ref([]);
const title = "创建激活码"
const show = ref(false)
const rmb = ref(0)
const queryMsg = ref("")
const queryStatus = ref(false)
const userInfo = reactive({
  "id": 0,
  "email": "",
  "user": "",
  "pwd": "",
  "level": 0,
  "avatar": "",
  "scope": 0,
  "pid": 0,
  "invite_id": 0,
  "invite": "",
  "money": 0,
  "login_time": 0,
  "login_ip": "",
  "is_lock": 0,
  "is_del": 0,
  "power_id": 0,
});

async function handleOk() {

  let index = message.create("加载中...", {
    type: "loading",
    duration: 10000
  })
  let {manager_money, price, keys, count} = await createKeys(formData)
  keys_list.value = keys
  let arr = []
  for (let i in keys) {
    arr.push(keys[i].long_keys)
  }
  keys_str.value = arr.join('\n')
  show.value = true
  await queryInfo()
  index.destroy()
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
  formData.project_id = val
  cardList.value = arr
  formData.cards_id = null
}

async function updateCards(val) {
  formData.cards_id = val
  const {cost, status, msg} = await queryOrderRmb(formData)
  rmb.value = cost
  queryMsg.value = msg
  queryStatus.value = status
  numberAnimationInstRef.value?.play()
  active.value = true
}

async function updateCount(val) {
  formData.count = val
  const {cost, status, msg} = await queryOrderRmb(formData)
  rmb.value = cost
  queryMsg.value = msg
  queryStatus.value = status
  numberAnimationInstRef.value?.play()
  active.value = true
}
async function queryInfo(){
  const {user} = await getUserInfo()
  userInfo.user = user.user
  userInfo.money = user.money
  userInfo.pid = user.pid
}

onMounted(async () => {
  await appStore.fetchProjectList(false)
  await appStore.fetchCardList(false)
  await queryInfo()
})
</script>

<style scoped>

</style>
