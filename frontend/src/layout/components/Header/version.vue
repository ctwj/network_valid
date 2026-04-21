<template>
  <n-modal :title="title" style="width: 1200px" preset="card" v-model:show="modalShow">
    <iframe class="iframe-view" src="http://www.qqcloudcom.cn/home/free/sys.html"></iframe>
  </n-modal>
</template>

<script lang="ts">
import {defineComponent,  ref,  Ref, watch} from 'vue';

const formVal = {
  "id": 0,
  title: "",
  description: "",
}

export default defineComponent({
  props: {
    show: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
      default: '祁启云免费系统最新发布'
    },
    form: {
      type: Object,
      default() {
        return formVal
      }
    }
  },
  setup(props, {emit}) {
    const modalShow: Ref<boolean> = ref(false)

    watch(() => props.show, (n) => {
      modalShow.value = n
    })
    watch(modalShow, (n) => {
      emit("update:show", n)
    })

    return {
      modalShow,
    }
  }
})
</script>

<style scoped>
.iframe-view{
  border: 0;
  width: 100%;
  height: calc(100vh - 200px);
}
</style>
