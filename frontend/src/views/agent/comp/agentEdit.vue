<template>
  <n-modal id="drawer-target" :title="title" style="width: 800px" preset="card"
           v-model:show="modalShow">
    <n-form
      ref="formRef"
      :inline="false"
      :label-width="80"
      :model="formData"
      :rules="rules"
    >
      <n-tabs type="line" animated>
        <n-tab-pane name="1" tab="基础设置">
          <div class="form-view">
            <n-form-item label="归属角色" path="power_id">
              <n-select label-field="title" value-field="id" v-model:value="formData.power_id"
                        :options="roleList" placeholder="请选择归属权限角色"/>
            </n-form-item>
            <n-form-item label="用户名" path="user">
              <n-input :maxlength="32" v-model:value="formData.user"
                       placeholder="输入用户名"/>
            </n-form-item>
            <n-form-item label="密码" path="pwd">
              <n-input :maxlength="32" v-model:value="formData.pwd"
                       placeholder="输入密码"/>
            </n-form-item>
            <n-form-item label="邮箱" path="email">
              <n-input :maxlength="32" v-model:value="formData.email"
                       placeholder="输入邮箱"/>
            </n-form-item>
            <n-form-item label="余额" path="money">
              <n-input-number v-model:value="formData.money" :min="0" :max="9999999"
                              :step="1" clearable/>
            </n-form-item>
            <n-form-item>
              <n-button type="primary" attr-type="button" @click="handleOk">
                {{ title }}
              </n-button>
            </n-form-item>
          </div>
        </n-tab-pane>
        <n-tab-pane name="2" tab="激活码授权">
          <div class="form-view">
            <n-button style="margin-left: 20px;margin-bottom: 16px" type="primary"
                      attr-type="button"
                      @click="active=true">
              添加授权
            </n-button>
            <n-list hoverable clickable>
              <n-list-item v-for="(item,index) in managerCardsList" :key="index">

                <n-thing :title-extra="getProjectName(item.project_id)" :title="item.title"
                         content-style="margin-top: 10px;">
                  <template #description>
                    <n-space size="small" style="margin-top: 4px">
                      <n-tag :bordered="false" type="info" size="small">
                        天数:{{ getCardsData(item.cards_id, "days") }}
                      </n-tag>
                      <n-tag :bordered="false" type="info" size="small">
                        点数:{{ getCardsData(item.cards_id, "points") }}
                      </n-tag>
                    </n-space>
                    <n-input-number style="margin-top: 6px" size="small"
                                    v-model:value="item.price">
                      <template #prefix>
                        ￥
                      </template>
                    </n-input-number>
                  </template>
                </n-thing>
                <template #suffix>
                  <n-button @click="handleUpdate(item)">
                    修改
                  </n-button>
                  <n-button @click="handleDel(item)" style="margin-top: 4px">
                    删除
                  </n-button>
                </template>
              </n-list-item>
            </n-list>
            <n-drawer to="#drawer-target" :block-scroll="false" :trap-focus="false"
                      v-model:show="active" width="100%" placement="right">
              <n-drawer-content closable title="激活码列表">
                <n-list hoverable clickable>
                  <n-list-item v-for="(item,index) in cardAllList" :key="index">

                    <n-thing :title-extra="getProjectName(item.project_id)" :title="item.title"
                             content-style="margin-top: 10px;">
                      <template #description>
                        <n-space size="small" style="margin-top: 4px">
                          <n-tag :bordered="false" type="info" size="small">
                            天数:{{ item.days }}
                          </n-tag>
                          <n-tag :bordered="false" type="info" size="small">
                            点数:{{ item.points }}
                          </n-tag>
                        </n-space>
                        <n-input-number style="margin-top: 6px" size="small"
                                        v-model:value="item.price">
                          <template #prefix>
                            ￥
                          </template>
                        </n-input-number>
                      </template>
                    </n-thing>
                    <template #suffix>
                      <n-button @click="handleAdd(item)">
                        添加
                      </n-button>
                    </template>
                  </n-list-item>
                </n-list>
              </n-drawer-content>
            </n-drawer>
          </div>
        </n-tab-pane>
      </n-tabs>

    </n-form>
  </n-modal>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, onMounted, Ref, watch, computed} from 'vue';
import {managerUpdate, getRole, getManagerCards, addManagerCards,deleteManagerCards,updateManagerCards} from "@/api";
import {useMessage, MessageReactive} from 'naive-ui'
import setting from '@/settings/componentSetting'
import {useAppStore} from "@/store/modules/app";

const {project} = setting
const appStore = useAppStore()
const formVal = {
  id: 0,
  power_id: null,
  money: 0,
  user: '',
  pwd: '',
  email: ''
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
    const active = ref(false);
    const roleList = ref([]);
    const managerCardsList = ref([]);
    const {
      keys_type_lock
    } = project
    const modalShow: Ref<boolean> = ref(false)
    const rules = reactive({});
    const message = useMessage()
    let index: MessageReactive | null = null
    const cardAllList = computed(() => {
      return appStore.getCardList
    })
    const projectList = computed(() => {
      let data = JSON.parse(JSON.stringify(appStore.getProjectList))
      data.unshift({name: "通用项目-所有项目都可用", id: 0})
      return data
    })

    const formData: Ref<Object> = ref({
      "id": 0,
      title: "",
      description: "",
    })

    const getProjectName = (id) => {
      let list = projectList.value
      for (let i in list) {
        if (list[i].id == id) {
          return list[i].name
        }
      }
    }

    const getCardsData = (id, field) => {
      let list = cardAllList.value
      for (let i in list) {
        if (list[i].id == id) {
          return list[i][field]
        }
      }
    }


    const handleOk = async () => {
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      let result = await managerUpdate(formData.value)
      if (result !== undefined && result > 0) {
        modalShow.value = false
        index.destroy()
        emit("on-update")
        await appStore.fetchCardList(true)
      } else {
        index.destroy()
      }
    }

    const handleAdd = async (row) => {
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      let data = JSON.parse(JSON.stringify(row))
      let form = {
        manager_id: formData.value.id,
        cards_id: data.id,
        price: data.price
      }
      await addManagerCards(form)
      index.destroy()
      const {data: res} = await getManagerCards({id: formData.value.id})
      managerCardsList.value = res
    }

    async function handleDel(item) {
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      await deleteManagerCards({id:item.id,manager_id: item.manager_id})
      index.destroy()
      const {data: res} = await getManagerCards({id: formData.value.id})
      managerCardsList.value = res
    }
    async function handleUpdate(item){
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      await updateManagerCards(item)
      index.destroy()
      const {data: res} = await getManagerCards({id: formData.value.id})
      managerCardsList.value = res
    }


    onMounted(async () => {
      modalShow.value = props.show
      let list = await getRole()
      list.unshift({title: "无权限", id: 0})
      roleList.value = list
      await appStore.fetchProjectList(false)
      await appStore.fetchCardList(false)

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
        const {data} = await getManagerCards({id: n.id})
        managerCardsList.value = data
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
      roleList,
      cardAllList,
      projectList,
      getProjectName,
      active,
      handleAdd,
      managerCardsList,
      getCardsData,
      handleDel,
      handleUpdate
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
