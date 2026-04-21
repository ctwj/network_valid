<template>
  <n-modal :title="title" style="width: 600px" preset="card" v-model:show="modalShow">
    <n-form
      ref="formRef"
      :inline="false"
      :label-width="80"
      :model="formData"
      :rules="rules"
    >
      <div class="form-view">
        <n-form-item label="账号" path="title">
          <n-input :maxlength="32" v-model:value="formData.name" placeholder="输入账号名称"/>
        </n-form-item>
        <n-form-item label="密码" path="password">
          <n-input :maxlength="32" v-model:value="formData.password" placeholder="输入密码"/>
        </n-form-item>
        <n-form-item label="安全密码" path="safe_password">
          <n-input :maxlength="32" v-model:value="formData.safe_password"
                   placeholder="输入安全密码"/>
        </n-form-item>
        <n-form-item label="邮箱" path="email">
          <n-input :maxlength="32" v-model:value="formData.email" placeholder="输入邮箱"/>
        </n-form-item>
        <n-form-item label="机器码" path="mac">
          <n-input :maxlength="32" v-model:value="formData.mac" placeholder="输入机器码"/>
        </n-form-item>
        <n-form-item label="手机" path="phone">
          <n-input :maxlength="11" v-model:value="formData.phone" placeholder="输入手机号码"/>
        </n-form-item>
        <n-form-item label="状态" path="is_lock">
          <n-radio-group v-model:value="formData.is_lock" name="is_lock">
            <n-radio v-for="(item,index) in keys_type_lock" :value="index" :key="index">
              {{ item }}
            </n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="天数" path="days">
          <n-input-number v-model:value="formData.days" :min="0" :max="9999999999"
                          :step="0.01" clearable/>
        </n-form-item>
        <n-form-item label="点数" path="points">
          <n-input-number v-model:value="formData.points" :min="0" :max="99999"
                          :step="1" clearable/>
        </n-form-item>
        <n-form-item label="标签" path="tag">
          <n-input :maxlength="200" v-model:value="formData.tag" placeholder="输入激活码绑定标签"/>
        </n-form-item>
        <n-form-item label="附加属性" path="key_ext_attr">
          <n-input :maxlength="200" v-model:value="formData.key_ext_attr"
                   placeholder="输入激活码绑定附加属性"/>
        </n-form-item>
      </div>
      <n-form-item>
        <n-button type="primary" attr-type="button" @click="handleOk">
          {{ title }}
        </n-button>
      </n-form-item>
    </n-form>
  </n-modal>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, onMounted, Ref, watch, computed} from 'vue';
import {updateMember} from "@/api";
import {useMessage, MessageReactive} from 'naive-ui'
import setting from '@/settings/componentSetting'
import {useAppStore} from "@/store/modules/app";

const {project} = setting
const appStore = useAppStore()
const formVal = {
  "id": 0,
  "manager_id": 0,
  "project_id": 0,
  "email": "",
  "name": "",
  "nick_name": "",
  "password": "",
  "safe_password": "",
  "money": 0,
  "days": 0,
  "points": 0,
  "mac": "",
  "phone": 0,
  "scope": 0,
  "key_ext_attr": "",
  "tag": "",
  "end_time": 0,
  "is_lock": 0,
}
export default defineComponent({
  props: {
    show: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
      default: ''
    },
    form: {
      type: Object,
      default() {
        return formVal
      }
    }
  },
  setup(props, {emit}) {

    const {
      keys_type_lock
    } = project
    const modalShow: Ref<boolean> = ref(false)
    const rules = reactive({});
    const message = useMessage()
    let index: MessageReactive | null = null
    const projectList = computed(() => {
      let data = JSON.parse(JSON.stringify(appStore.getProjectList))
      data.unshift({name: "通用项目-所有项目都可用", id: 0})
      return data
    })
    const formData: Ref<Object> = ref({
      "id": 0,
      "manager_id": 0,
      "project_id": 0,
      "email": "",
      "name": "",
      "nick_name": "",
      "password": "",
      "safe_password": "",
      "money": 0,
      "days": 0,
      "points": 0,
      "mac": "",
      "phone": 0,
      "scope": 0,
      "key_ext_attr": "",
      "tag": "",
      "end_time": 0,
      "is_lock": 0,
    })
    const handleOk = async () => {
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      await updateMember(formData.value)
      modalShow.value = false
        index.destroy()
        emit("on-update")
    }

    onMounted(async () => {
      modalShow.value = props.show
      await appStore.fetchProjectList(false)

    })

    watch(() => props.show, (n) => {
      modalShow.value = n
    })
    watch(() => props.form, (n) => {
      if (n === null) {
        formData.value = formVal
      } else {
        formData.value = n
      }

    })
    watch(modalShow, (n) => {
      emit("update:show", n)
    })

    return {
      keys_type_lock,
      formData,
      rules,
      modalShow,
      handleOk,
      projectList
    }
  }
})
</script>

<style scoped>
.form-view {
  height: 60vh;
  overflow-y: auto;
}
</style>
