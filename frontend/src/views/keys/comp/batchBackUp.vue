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

export default defineComponent({
  props: {
    rows:{
      type: Array,
      default(){
        return []
      }
    },
    form: {
      type: Object,
      default(){
        return {}
      }
    }
  },
  setup(props) {
    const options = [
      {
        label: '操作选中',
        key: '1',
        children: [
          {
            label: '锁定选中',
            key: 'jordan baker'
          },
          {
            label: '删除选中',
            key: 'tom buchanan'
          }
        ]
      },
      {
        label: '操作筛选',
        key: '2',
        children: [
          {
            label: '锁定筛选',
            key: 'jordan baker'
          },
          {
            label: '删除筛选',
            key: 'tom buchanan'
          }
        ]
      }
    ]
    const handleSelect = (key: string | number) => {
      console.log(key)
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
