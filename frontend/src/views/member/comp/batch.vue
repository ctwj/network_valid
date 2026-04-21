<template>
  <div>
    <n-dropdown
      :options="options"
      placement="bottom-start"
      trigger="click"
      @select="handleSelect"
    >
      <n-button>批量操作</n-button>
    </n-dropdown>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import {useMessage, MessageReactive, useDialog} from 'naive-ui'
import {batchMember} from "@/api";

export default defineComponent({
  props: {
    rows: {
      type: Array,
      default() {
        return []
      }
    },
    form: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  setup(props,{emit}) {
    const dialog = useDialog()
    let index: MessageReactive | null = null
    const message = useMessage()
    const options = [
      {
        label: '操作选中',
        key: '1',
        children: [
          {
            label: '锁定选中',
            key: 'lock_1'
          },
          {
            label: '解锁选中',
            key: 'unlock_1'
          },
          {
            label: '解绑选中',
            key: 'unbind_1'
          },
          {
            label: '删除选中',
            key: 'delete_1'
          }
        ]
      },
      {
        label: '操作筛选',
        key: '2',
        children: [
          {
            label: '锁定筛选',
            key: 'lock_2'
          },
          {
            label: '解锁筛选',
            key: 'unlock_2'
          },
          {
            label: '解绑筛选',
            key: 'unbind_2'
          },
          {
            label: '删除筛选',
            key: 'delete_2'
          }
        ]
      }
    ]
    let formData = props.form

    const formStr = JSON.stringify(formData)
    const check = () => {
      return props.rows.length !== 0;
    }
    const handleSelect = async (key: string) => {
      let form = JSON.parse(JSON.stringify(props.form))
      let list = JSON.parse(JSON.stringify(props.rows))
      const keyArr = key.split("_")
      if (Number(keyArr[1]) === 2) {
        if (formStr == JSON.stringify(formData)) return message.warning('请至少设置一个筛选条件')
      } else {
        if (!check()) return message.warning('请至少选中一项数据')
        form.id = JSON.stringify(list)
      }
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      if (keyArr[0] === "delete") {
        dialog.warning({
          title: '温馨提示',
          content: '删除操作将不可逆!(请谨慎操作)',
          positiveText: '确定',
          negativeText: '取消',
          onPositiveClick: async () => {
            form.type = key
            await batchMember(form)
            index.destroy()
            emit("on-update")
          },
          onNegativeClick: () => {
            message.error('取消操作')
            index.destroy()
          }
        })
      } else {
        form.type = key
        await batchMember(form)
        index.destroy()
        emit("on-update")
      }
    }
    return {
      options,
      handleSelect
    }
  }
})
</script>

<style scoped>

</style>
