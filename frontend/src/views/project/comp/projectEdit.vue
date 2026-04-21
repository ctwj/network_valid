<template>
  <n-modal :title="title" style="width: 600px" preset="card" v-model:show="modalShow">
    <n-form
      ref="formRef"
      :inline="false"
      :label-width="80"
      :model="formData"
      :rules="rules"
    >
      <n-tabs type="line" animated>
        <n-tab-pane name="1" tab="基础设置">
          <n-form-item label="项目名称" path="name">
            <n-input v-model:value="formData.name" placeholder="输入项目名称"/>
          </n-form-item>
          <n-form-item label="登录模式" path="type">
            <n-radio-group v-model:value="formData.type" name="radiogroup2">
              <n-radio v-for="(item,index) in type" :value="index" :key="index">
                {{ item }}
              </n-radio>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="运营模式" path="status_type">
            <n-radio-group v-model:value="formData.status_type" name="radiogroup2">
              <n-radio v-for="(item,index) in status" :value="index" :key="index">
                {{ item }}
              </n-radio>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="加密模式" path="encrypt">
            <n-radio-group v-model:value="formData.encrypt" name="radiogroup2">
              <n-radio v-for="(item,index) in encrypt" :value="index" :key="index">
                {{ item }}
              </n-radio>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="签名算法" path="sign">
            <n-radio-group v-model:value="formData.sign" name="sign">
              <n-radio v-for="(item,index) in hash" :value="index" :key="index">
                {{ item }}
              </n-radio>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="项目公告" path="notice">
            <n-input
              v-model:value="formData.notice"
              type="textarea"
              placeholder="公告内容"
            />
          </n-form-item>
        </n-tab-pane>
        <n-tab-pane v-if="formData.id >0" name="2" tab="密匙相关">
          <n-form-item label="重置RSA" path="type">
            <n-radio-group v-model:value="formData.update_rsa" name="radiogroup2">
              <n-radio v-for="(item,index) in change" :value="index" :key="index">
                {{ item }}
              </n-radio>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="重置AES" path="type">
            <n-radio-group v-model:value="formData.update_key" name="radiogroup2">
              <n-radio v-for="(item,index) in change" :value="index" :key="index">
                {{ item }}
              </n-radio>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="重置appkey" path="type">
            <n-radio-group v-model:value="formData.update_app_key" name="radiogroup2">
              <n-radio v-for="(item,index) in change" :value="index" :key="index">
                {{ item }}
              </n-radio>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="重置secretKey" path="type">
            <n-radio-group v-model:value="formData.update_secret_key" name="radiogroup2">
              <n-radio v-for="(item,index) in change" :value="index" :key="index">
                {{ item }}
              </n-radio>
            </n-radio-group>
          </n-form-item>
        </n-tab-pane>
      </n-tabs>
      <n-form-item>
        <n-button type="primary" attr-type="button" @click="handleOk">
          {{ title }}
        </n-button>
      </n-form-item>
    </n-form>
  </n-modal>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, onMounted, Ref, watch} from 'vue';
import {useAppStore} from "@/store/modules/app";
import {createProject, updateProject} from "@/api";
import {useMessage, MessageReactive} from 'naive-ui'
import setting from '@/settings/componentSetting'
const {project} = setting

const formVal = {
  id: 0,
  name: '',
  type: 0,
  status_type: 0,
  encrypt: 0,
  notice: '',
  api: '',
  sign: 0,
  update_rsa: 0,
  update_key: 0,
  update_app_key: 0,
  update_secret_key: 0
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
    const appStore = useAppStore()
    const {status, encrypt, type,hash} = project
    const change = reactive(["不重置", "重置"])
    const modalShow: Ref<boolean> = ref(false)
    const rules = reactive({});
    const message = useMessage()
    let index: MessageReactive | null = null
    const formData: Ref<Object> = ref({
      id: 0,
      name: '',
      type: 0,
      status_type: 0,
      encrypt: 0,
      notice: '',
      api: '',
      sign: 0,
      update_rsa: 0,
      update_key: 0,
      update_app_key: 0,
      update_secret_key: 0
    })
    const handleOk = async () => {
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      if (formData.value.id > 0) {
        let result = await updateProject(formData.value)
        if (result !== undefined && result > 0) {
          modalShow.value = false
          index.destroy()
          await appStore.fetchCardList(true)
          emit("on-update")
        }else {
          index.destroy()
        }
      } else {
        let result = await createProject(formData.value)
        if (result !== undefined && result > 0) {
          modalShow.value = false
          index.destroy()
          await appStore.fetchCardList(true)
          emit("on-update")
        }else {
          index.destroy()
        }
      }
    }

    onMounted(() => {
      modalShow.value = props.show
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
      status,
      encrypt,
      type,
      formData,
      rules,
      modalShow,
      change,
      hash,
      handleOk
    }
  }
})
</script>

<style scoped>

</style>
