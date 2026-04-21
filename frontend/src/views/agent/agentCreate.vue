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
        <n-form-item label="归属角色" path="power_id">
          <n-select label-field="title" value-field="id" v-model:value="formData.power_id"
                    :options="roleList" placeholder="请选择归属权限角色"/>
        </n-form-item>
        <n-form-item label="用户名" path="user">
          <n-input :maxlength="32" v-model:value="formData.user"
                   placeholder="输入用户名"/>
        </n-form-item>
        <n-form-item label="密码" path="pwd">
          <n-input :maxlength="32" v-model:value="formData.pwd"
                   placeholder="输入密码"/>
        </n-form-item>
        <n-form-item label="邮箱" path="email">
          <n-input :maxlength="32" v-model:value="formData.email"
                   placeholder="输入邮箱"/>
        </n-form-item>
        <n-form-item label="余额" path="money">
          <n-input-number v-model:value="formData.money" :min="0" :max="9999999"
                          :step="1" clearable/>
        </n-form-item>
      </div>
      <n-form-item>
        <n-button type="primary" attr-type="button" @click="handleOk">
          {{ title }}
        </n-button>
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
import {reactive, onMounted, ref} from "vue";
import {managerAdd,getRole} from '@/api/index'
import setting from '@/settings/componentSetting'
import {copy_inline} from "@/utils";
import {useMessage} from "naive-ui";
import {useAppStore} from "@/store/modules/app";

const appStore = useAppStore()
const message = useMessage()
const {project} = setting
const rules = reactive({})
const roleList = ref([]);
const formData = reactive({
  power_id: null,
  money: 0,
  user: '',
  pwd: '',
  email: ''
})
const title = "创建代理"
const show = ref(false)

async function handleOk() {

  let index = message.create("加载中...", {
    type: "loading",
    duration: 10000
  })
  let result = await managerAdd(formData)
  index.destroy()
  await appStore.fetchAgentList(true)
}


onMounted(async () => {
  roleList.value = await getRole()
})
</script>

<style scoped>

</style>
