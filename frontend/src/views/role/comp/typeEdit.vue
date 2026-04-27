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
        <n-form-item v-if="formData.id === 0" label="归属项目" path="project_id">
          <n-select label-field="name" value-field="id" v-model:value="formData.project_id"
                    :options="projectList" placeholder="请选择归属项目"/>
        </n-form-item>
        <n-form-item label="套餐类型" path="title">
          <n-input :maxlength="28" v-model:value="formData.title" placeholder="输入套餐类型名称"/>
        </n-form-item>
        <n-form-item label="启用状态" path="is_lock">
          <n-radio-group v-model:value="formData.is_lock" name="is_lock">
            <n-radio v-for="(item,index) in keys_type_lock" :value="index" :key="index">
              {{ item }}
            </n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="前缀" path="key_prefix">
          <n-input :maxlength="4" v-model:value="formData.key_prefix"
                   placeholder="输入4位数前缀（字母或者数字组合）可空"/>
        </n-form-item>
        <n-form-item label="定价" path="price">
          <n-input-number v-model:value="formData.price" :min="0" :max="999999"
                          :step="0.01" clearable/>
        </n-form-item>
        <n-form-item label="天数" path="days">
          <n-input-number v-model:value="formData.days" :min="0" :max="999999"
                          :step="0.01" clearable/>
        </n-form-item>
        <n-form-item label="点数" path="points">
          <n-input-number v-model:value="formData.points" :min="0" :max="99999"
                          :step="1" clearable/>
        </n-form-item>
        <n-form-item label="标签" path="tag">
          <n-input :maxlength="200" v-model:value="formData.tag" placeholder="输入兑换码绑定标签"/>
        </n-form-item>
        <n-form-item label="附加属性" path="key_ext_attr">
          <n-input :maxlength="200" v-model:value="formData.key_ext_attr"
                   placeholder="输入兑换码绑定附加属性"/>
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
import {createCard, updateCard} from "@/api";
import {useMessage, MessageReactive, useDialog} from 'naive-ui'
import setting from '@/settings/componentSetting'
import {useAppStore} from "@/store/modules/app";

const {project} = setting
const appStore = useAppStore()
const formVal = {
  "id": 0,
  "project_id": null,
  "manager_id": 0,
  "title": "",
  "price": 0,
  "key_prefix": "",
  "level_id": 0,
  "days": 0,
  "points": 0,
  "key_ext_attr": "",
  "tag": "",
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
    const dialog = useDialog()
    let index: MessageReactive | null = null
    const projectList = computed(() => {
      let data = JSON.parse(JSON.stringify(appStore.getProjectList))
      data.unshift({name: "通用项目-所有项目都可用", id: 0})
      return data
    })
    const formData: Ref<Object> = ref({
      "id": 0,
      "project_id": null,
      "manager_id": 0,
      "title": "",
      "price": 0,
      "key_prefix": "",
      "level_id": 0,
      "days": 0,
      "points": 0,
      "key_ext_attr": "",
      "tag": "",
      "is_lock": 0,
    })
    const handleOk = async () => {
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      if (formData.value.id > 0) {
        dialog.warning({
          title: '温馨提示',
          content: '修改已经创建的套餐类型，将会导致已经创建且未使用的兑换码属性跟随变换，如非必要请勿修改天数点数属性',
          positiveText: '确定',
          negativeText: '取消',
          onPositiveClick: async() => {
            let result = await updateCard(formData.value)
            if (result !== undefined && result > 0) {
              modalShow.value = false
              index.destroy()
              emit("on-update")
              await appStore.fetchCardList(true)
            } else {
              index.destroy()
            }
          },
          onNegativeClick: () => {
            message.error('取消操作')
            index.destroy()
          }
        })

      } else {
        let result = await createCard(formData.value)
        if (result !== undefined && result > 0) {
          modalShow.value = false
          index.destroy()
          emit("on-update")
          await appStore.fetchCardList(true)
        } else {
          index.destroy()
        }
      }
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
