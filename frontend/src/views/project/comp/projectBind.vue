<template>
  <n-modal :title="title" style="width: 720px" preset="card" v-model:show="modalShow">
    <n-list hoverable clickable>
      <n-list-item v-for="(item,index) in loginData" :key="index">
        <n-thing :title="item.title" content-style="margin-top: 10px;">
          <template #description>
            <n-space size="small" style="margin-top: 4px">
              <n-tag :bordered="false" type="info" size="small">
                {{ ["绑定登录", "普通登录", "点数登录"][item.mode] }}
              </n-tag>
              <n-tag :bordered="false" type="info" size="small">
                {{ ["不允许解绑", "原机解绑", "自动解绑", "任意解绑"][item.unbind_mode] }}
              </n-tag>
              <n-tag :bordered="false" type="info" size="small">
                {{ ["不扣时", "解绑就扣时", "超出扣时", "超出不扣时"][item.unbind_weaken_mode] }}
              </n-tag>
            </n-space>
          </template>

        </n-thing>
        <template #suffix>
            <n-button :disabled="formData.login_type === item.id" @click="handleBind(item)">{{formData.login_type === item.id ? "已绑定" : "绑定"}}</n-button>
          </template>
      </n-list-item>
    </n-list>
    <!--    <n-data-table-->
    <!--      v-if="modalShow"-->
    <!--      :columns="loginColumns"-->
    <!--      :row-key="(row)=>{return row.id}"-->
    <!--      :data="loginData"-->
    <!--      :pagination="false"-->
    <!--      :bordered="false"-->
    <!--    />-->
    <!--    <n-tabs type="line" animated>-->
    <!--      <n-tab-pane name="登录规则" tab="登录规则">-->
    <!--        -->
    <!--      </n-tab-pane>-->
    <!--    </n-tabs>-->
  </n-modal>
</template>

<script lang="ts">
import {defineComponent, ref, Ref, onMounted, watch, reactive, h} from "vue";
import {bindProjectLogin, loginRuleList} from "@/api";
import {loginColumns} from "@/views/project/columns/list";
import {TableAction} from '@/components/Table';
import {MessageReactive, useMessage} from "naive-ui";

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
        return {
          id: 0,
          login_type: 0
        }

      }
    }
  },
  setup(props, {emit}) {
    const actionColumn = reactive({
      width: 80,
      title: '操作',
      key: 'action',
      fixed: 'right',
      align: 'center',
      render(record) {
        return h(TableAction as any, {
          style: 'text',
          actions: createActions(record),
        });
      },
    });
    const modalShow: Ref<boolean> = ref(false)
    const loginData = ref([]);
    const formData = reactive({id: 0, login_type: 0})
    const message = useMessage()
    let index: MessageReactive | null = null

    function createActions(record) {
      return [
        {
          label: formData.login_type === record.id ? "已绑定" : "绑定",
          type: 'primary',
          disabled: formData.login_type === record.id,
          onClick: handleBind.bind(null, record),
          ifShow: () => {
            return true;
          },
          // auth: ['basic_list'],
        }
      ];
    }

    onMounted(async () => {
      modalShow.value = props.show
      await fetchLogin()
    })
    watch(() => props.show, (n) => {
      modalShow.value = n
    })
    watch(modalShow, (n) => {
      emit("update:show", n)
    })
    watch(() => props.form, (n) => {
      if (n === null) {
        formData.id = 0
        formData.login_type = 0
      } else {
        formData.id = n.id
        formData.login_type = n.login_type
      }

    })

    async function handleBind(row) {
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      const result = await bindProjectLogin({id: formData.id, project_login_id: row.id})
      if (result !== undefined && result > 0) {
        index.destroy()
        emit("on-update")
      } else {
        index.destroy()
      }
    }

    async function fetchLogin() {
      const {data} = await loginRuleList()
      console.log(data)
      loginData.value = data
    }

    loginColumns.push(actionColumn)
    return {
      modalShow,
      loginColumns,
      loginData,
      formData,
      handleBind
    }
  }
})
</script>

<style scoped>

</style>
