<template>
  <n-modal :title="title" style="width: 600px" preset="card" v-model:show="modalShow">
    <div v-html="content"></div>
  </n-modal>
</template>

<script lang="ts">
import {defineComponent,  ref, onMounted, Ref, watch} from 'vue';


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
    content: {
      type: String,
      default: ''
    }
  },
  setup(props, {emit}) {

    const modalShow: Ref<boolean> = ref(false)


    onMounted(() => {
      modalShow.value = props.show
    })

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
.form-view{
  height: 60vh;
  overflow-y: auto;
}
</style>
