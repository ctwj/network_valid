<template>
  <n-modal :title="title" style="width: 800px" preset="card" v-model:show="modalShow">
    <n-form
      ref="formRef"
      :inline="false"
      :label-width="80"
      :model="formData"
      :rules="rules"
    >
      <n-tabs type="line" animated>
        <n-tab-pane name="1" tab="基础设置">
          <n-form-item label="角色名称" path="title">
            <n-input v-model:value="formData.title" placeholder="输入角色名称"/>
          </n-form-item>
          <n-form-item label="角色描述" path="description">
            <n-input v-model:value="formData.description" placeholder="输入角色描述"/>
          </n-form-item>
        </n-tab-pane>
        <n-tab-pane name="2" tab="权限设置">
          <n-transfer
            ref="transfer"
            v-model:value="treeVal"
            :options="options"
            :render-source-list="renderSourceList"
            source-filterable
          />
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
import {defineComponent, reactive, ref, onMounted, Ref, watch, h} from 'vue';
import {updateRoleUser, rolePower, createRoleUser,getUserRole} from "@/api";
import {useMessage, MessageReactive,  NTree, TransferRenderSourceList} from 'naive-ui'
import setting from '@/settings/componentSetting'
import {useAppStore} from "@/store/modules/app";

const {project} = setting
const appStore = useAppStore()
const formVal = {
  "id": 0,
  title: "",
  description: "",
}
type Option = {
  name: string
  value: string
  children?: Option[]
  index: string
  path: string
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
    const valueRef = ref<Array<string | number>>([])
    const {
      keys_type_lock
    } = project
    const modalShow: Ref<boolean> = ref(false)
    const rules = reactive({});
    const message = useMessage()
    let index: MessageReactive | null = null
    const treeData = ref<Array<Option>>([])
    const formData: Ref<Object> = ref({
      "id": 0,
      title: "",
      description: "",
    })
    const options = ref<Array<Option>>([])
    const renderSourceList: TransferRenderSourceList = function ({
                                                                   onCheck,
                                                                   pattern
                                                                 }) {
      return h(NTree, {
        style: 'margin: 0 4px;',
        keyField: 'name',
        labelField: "name",
        checkable: true,
        selectable: false,
        blockLine: true,
        checkOnClick: true,
        data: treeData.value,
        pattern,
        checkedKeys: valueRef.value,
        onUpdateCheckedKeys: (checkedKeys: Array<string | number>) => {
          console.log("checkedKeys", checkedKeys)
          onCheck(checkedKeys)
        }
      })
    }

    function flattenTree(list: null | Option[]): Option[] {
      const result: Option[] = []
      console.log(list)

      function flatten(_list: Option[] = []) {
        _list.forEach((item) => {
          result.push({label: item.name, value: item.name})
          if (item.children !== null) {
            flatten(item.children)
          }
        })
      }

      flatten(list)
      return result
    }

    const handleOk = async () => {
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      let form = JSON.parse(JSON.stringify(formData.value))
      form['name'] = JSON.stringify(valueRef.value)
      if (formData.value.id > 0) {
        let result = await updateRoleUser(form)
        if (result !== undefined && result > 0) {
          modalShow.value = false
          index.destroy()
          emit("on-update")
          await appStore.fetchCardList(true)
        } else {
          index.destroy()
        }

      } else {
        let result = await createRoleUser(form)
        if (result !== undefined && result > 0) {
          modalShow.value = false
          index.destroy()
          emit("on-update")
        } else {
          index.destroy()
        }
      }
    }


    onMounted(async () => {
      modalShow.value = props.show
      await appStore.fetchProjectList(false)
      treeData.value = await rolePower({})
      options.value = flattenTree(JSON.parse(JSON.stringify(treeData.value)))
      console.log(options.value)
    })

    watch(() => props.show, (n) => {
      modalShow.value = n
    })
    watch(() => props.form, async (n) => {
      console.log(n)
      if (n === null) {
        formData.value = formVal
      } else {
        formData.value = n
        if (n.id > 0){
          valueRef.value = await getUserRole({role_id: n.id})
        }
      }

    })
    watch(modalShow, (n) => {
      emit("update:show", n)
    })

    return {
      options,
      keys_type_lock,
      formData,
      rules,
      modalShow,
      handleOk,
      treeVal: valueRef,
      renderSourceList
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
