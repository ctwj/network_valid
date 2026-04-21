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
        <n-form-item label="版本号" path="version">
          <n-input-number v-model:value="formData.version" :min="0" :max="999"
                          :step="0.01" clearable/>
        </n-form-item>
        <n-form-item label="启用状态" path="is_active">
          <n-radio-group v-model:value="formData.is_active" name="is_active">
            <n-radio v-for="(item,index) in version_is_active" :value="index" :key="index">
              {{ item }}
            </n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="强制更新" path="is_must_update">
          <n-radio-group v-model:value="formData.is_must_update" name="is_must_update">
            <n-radio v-for="(item,index) in version_is_must_update" :value="index" :key="index">
              {{ item }}
            </n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="热更新地址" path="tag">
          <n-input :maxlength="50" v-model:value="formData.wgt_url" placeholder="输入热更新地址"/>
        </n-form-item>
        <n-form-item label="更新公告" path="notice">
            <n-input
              v-model:value="formData.notice"
              type="textarea"
              :max-length="9000"
              placeholder="公告内容"
            />
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
import {createVersion,updateProjectVersion} from "@/api";
import {useMessage, MessageReactive} from 'naive-ui'
import setting from '@/settings/componentSetting'
import {useAppStore} from "@/store/modules/app";

const {project} = setting
const appStore = useAppStore()
const formVal = {
  "id": 0,
  "manager_id": 0,
  "project_id": null,
  "version": null,
  "is_must_update": 0,
  "is_active": 0,
  "notice": "",
  "wgt_url": "",
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
      version_is_must_update,
      version_is_active
    } = project
    const modalShow: Ref<boolean> = ref(false)
    const rules = reactive({});
    const message = useMessage()
    let index: MessageReactive | null = null
    const projectList = computed(() => {
      let data = JSON.parse(JSON.stringify(appStore.getProjectList))
      return data
    })
    const formData: Ref<Object> = ref({
      "id": 0,
      "manager_id": 0,
      "project_id": null,
      "version": null,
      "is_must_update": 0,
      "is_active": 0,
      "notice": "",
      "wgt_url": "",
    })
    const handleOk = async () => {
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      if (formData.value.id > 0) {
        let result = await updateProjectVersion(formData.value)
        if (result !== undefined && result > 0) {
          modalShow.value = false
          index.destroy()
          emit("on-update")
          await appStore.fetchCardList(true)
        } else {
          index.destroy()
        }
      } else {
        let result = await createVersion(formData.value)
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
      version_is_active,
      version_is_must_update,
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
  max-height: 60vh;
  overflow-y: auto;
}
</style>
