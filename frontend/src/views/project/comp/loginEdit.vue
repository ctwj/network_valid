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
        <n-form-item label="规则名称" path="title">
          <n-input v-model:value="formData.title" placeholder="输入规则名称"/>
        </n-form-item>
        <n-form-item label="绑定模式" path="mode">
          <n-radio-group v-model:value="formData.mode" name="mode">
            <n-radio v-for="(item,index) in login_mode" :value="index" :key="index">
              {{ item }}
            </n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="注册模式" path="reg_mode">
          <n-radio-group v-model:value="formData.reg_mode" name="reg_mode">
            <n-radio v-for="(item,index) in login_reg_mode" :value="index" :key="index">
              {{ item }}
            </n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="解绑模式" path="unbind_mode">
          <n-radio-group v-model:value="formData.unbind_mode" name="unbind_mode">
            <n-radio v-for="(item,index) in login_unbind_mode" :value="index" :key="index">
              {{ item }}
            </n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="解绑扣时" path="unbind_weaken">
          <n-input-number v-model:value="formData.unbind_weaken" :min="0" :max="99999" :step="0.01"
                          clearable/>
        </n-form-item>
        <n-form-item label="解绑扣点" path="unbind_weaken_points">
          <n-input-number v-model:value="formData.unbind_weaken_points" :min="0" :max="99999"
                          :step="1" clearable/>
        </n-form-item>
        <n-form-item label="解绑上限" path="unbind_times">
          <n-input-number v-model:value="formData.unbind_times" :min="0" :max="99999" :step="1"
                          clearable/>
        </n-form-item>
        <n-form-item label="解绑周期" path="unbind_date">
          <n-radio-group v-model:value="formData.unbind_date" name="unbind_mode">
            <n-radio v-for="(item,index) in login_unbind_date" :value="index" :key="index">
              {{ item }}
            </n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="前置解绑时间" path="unbind_before">
          <n-input-number v-model:value="formData.unbind_before" :min="0" :max="99999" :step="1"
                          clearable/>
        </n-form-item>
        <n-form-item label="点数登录模式" path="number_more">
          <n-radio-group v-model:value="formData.number_more" name="unbind_mode">
            <n-radio v-for="(item,index) in login_number_mode" :value="index" :key="index">
              {{ item }}
            </n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="用户/激活码对应最大的可分发数量" path="number_more">
          <n-input-number v-model:value="formData.number_more" :min="0" :max="99999" :step="1"
                          clearable/>
        </n-form-item>
        <n-form-item label="登录扣除点数" path="number_weaken">
          <n-input-number v-model:value="formData.number_weaken" :min="0" :max="99999" :step="1"
                          clearable/>
        </n-form-item>
        <n-form-item label="点数登录扣点周期间隔" path="number_weaken_time">
          <n-radio-group v-model:value="formData.number_weaken_time" name="unbind_mode">
            <n-radio v-for="(item,index) in login_number_weaken_time" :value="index" :key="index">
              {{ item }}
            </n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="单机多开最大可在线量" path="pc_more">
          <n-input-number v-model:value="formData.pc_more" :min="0" :max="99999" :step="1"
                          clearable/>
        </n-form-item>
        <n-form-item label="多机多开最大可在线量" path="pc_code_more">
          <n-input-number v-model:value="formData.pc_code_more" :min="0" :max="99999" :step="1"
                          clearable/>
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
import {defineComponent, reactive, ref, onMounted, Ref, watch} from 'vue';
import {createLoginRule, updateLoginRule} from "@/api";
import {useMessage, MessageReactive} from 'naive-ui'
import setting from '@/settings/componentSetting'

const {project} = setting

const formVal = {
  "id": 0,
  "manager_id": 0,
  "title": "",
  "mode": 0,
  "reg_mode": 0,
  "email_reg": 0,
  "unbind_mode": 0,
  "unbind_weaken_mode": 0,
  "unbind_weaken": 0,
  "unbind_weaken_points": 0,
  "unbind_times": 0,
  "unbind_date": 0,
  "unbind_before": 0,
  "number_mode": 0,
  "number_more": 0,
  "number_weaken": 0,
  "number_weaken_time": 0,
  "pc_more": 0,
  "pc_code_more": 0,
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
      login_mode,
      login_number_mode,
      login_unbind_mode,
      login_unbind_weaken_mode,
      login_emial_reg,
      login_reg_mode,
      login_unbind_date,
      login_number_weaken_time
    } = project
    const change = reactive(["不重置", "重置"])
    const modalShow: Ref<boolean> = ref(false)
    const rules = reactive({});
    const message = useMessage()
    let index: MessageReactive | null = null
    const formData: Ref<Object> = ref({
      "id": 0,
      "manager_id": 0,
      "title": "",
      "mode": 0,
      "reg_mode": 0,
      "email_reg": 0,
      "unbind_mode": 0,
      "unbind_weaken_mode": 0,
      "unbind_weaken": 0,
      "unbind_weaken_points": 0,
      "unbind_times": 0,
      "unbind_date": 0,
      "unbind_before": 0,
      "number_mode": 0,
      "number_more": 0,
      "number_weaken": 0,
      "number_weaken_time": 0,
      "pc_more": 0,
      "pc_code_more": 0,
    })
    const handleOk = async () => {
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      if (formData.value.id > 0) {
        let result = await updateLoginRule(formData.value)
        if (result !== undefined && result > 0) {
          modalShow.value = false
          index.destroy()
          emit("on-update")
        } else {
          index.destroy()
        }
      } else {
        let result = await createLoginRule(formData.value)
        if (result !== undefined && result > 0) {
          modalShow.value = false
          index.destroy()
          emit("on-update")
        } else {
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
      login_mode,
      login_number_mode,
      login_unbind_mode,
      login_unbind_weaken_mode,
      login_emial_reg,
      login_reg_mode,
      login_unbind_date,
      login_number_weaken_time,
      formData,
      rules,
      modalShow,
      change,
      handleOk
    }
  }
})
</script>

<style scoped>
.form-view{
  height: 60vh;
  overflow-y: auto;
}
</style>
