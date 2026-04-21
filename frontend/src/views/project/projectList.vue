<template>
  <div>
    <n-grid :x-gap="12" :y-gap="12" cols="1 s:1 m:2 l:3 xl:4 2xl:4" responsive="screen">
    <n-grid-item>
      <n-card class="soft-mask" title="待创建" hoverable>

        <div class="mask">
          <n-button type="primary" @click="show=true;title='创建项目';row=null">创建项目</n-button>
        </div>
        <div class="notice">待创建...</div>
        <div class="details">
          <div>
            登录规则:
          </div>
          <n-popover trigger="hover">
            <template #trigger>
              <div>待创建</div>
            </template>
            <span>当前项目应用的是:待创建登录规则</span>
          </n-popover>
        </div>
        <div class="details">
          <div>
            APPKEY:
          </div>
          <n-popover trigger="hover">
            <template #trigger>
              <div>待创建</div>
            </template>
            <span>点击复制</span>
          </n-popover>
        </div>
        <div class="details">
          <div>
            SECRETKEY:
          </div>
          <n-popover trigger="hover">
            <template #trigger>
              <div>待创建</div>
            </template>
            <span>点击复制</span>
          </n-popover>
        </div>
        <div class="details">
          <div>
            运营模式:
          </div>
          <n-popover trigger="hover">
            <template #trigger>
              <div>待创建</div>
            </template>
            <div class="details-popover">
              <ul>
                <li>收费运营：所有接口都必须在用户付费的基础上才可使用</li>
                <li>停止运营：项目的所有接口都将停止</li>
                <li>
                  免费运营：所有的接口都将免费使用，且新注册的用户都会打上免费用户的绑定标签信息，如果将项目调整回收费运营，则所有的免费用户将无法正常使用，必须重新充值才可正常使用！
                </li>
              </ul>
            </div>
          </n-popover>
        </div>
        <div class="details">
          <div>
            加密方式:
          </div>
          <n-popover trigger="hover">
            <template #trigger>
              <div>待创建</div>
            </template>
            <div class="soft-details-item">
              <div class="soft-details-item-title">
                <a-tag style="width: 100%" color="#108ee9">
                  密匙A
                </a-tag>
              </div>
              <div class="view-key">待创建</div>
            </div>
            <div class="soft-details-item">
              <div class="soft-details-item-title">
                <a-tag style="width: 100%" color="#108ee9">
                  密匙B
                </a-tag>
              </div>
              <div class="view-key">待创建</div>
            </div>
            <div class="soft-details-item">
              <div class="soft-details-item-title">
                <a-tag style="width: 100%" color="#108ee9">
                  公钥
                </a-tag>
              </div>
              <div class="view-key">待创建
              </div>
            </div>
            <div class="soft-details-item">
              <div class="soft-details-item-title">
                <a-tag style="width: 100%" color="#108ee9">
                  私钥
                </a-tag>
              </div>
              <div class="view-key">待创建
              </div>
            </div>
          </n-popover>
        </div>
        <div slot="action" class="card-action">
          <n-button block class="card-action">绑定规则</n-button>
          <n-button type="primary" block class="card-action">修改</n-button>
          <n-button type="error" block class="card-action">删除</n-button>
        </div>
      </n-card>
    </n-grid-item>
    <n-grid-item v-for="(item,index) in dataSource" :key="index">
      <n-card :title="item.name" hoverable>
        <div class="notice">{{ withNotice(item.notice) }}</div>
        <div class="details">
          <div>
            登录规则:
          </div>
          <n-popover trigger="hover">
            <template #trigger>
              <div>{{ type[item.type] }}</div>
            </template>
            <span>当前项目应用的是:{{ type[item.type] }}登录规则</span>
          </n-popover>
        </div>
        <div class="details">
          <div>
            APPKEY:
          </div>
          <n-popover trigger="hover">
            <template #trigger>
              <div @click="copy_inline(item.app_key)">{{ encryKey(item.app_key) }}</div>
            </template>
            <span>点击复制</span>
          </n-popover>
        </div>
        <div class="details">
          <div>
            SECRETKEY:
          </div>
          <n-popover trigger="hover">
            <template #trigger>
              <div @click="copy_inline(item.secret_key)">{{ encryKey(item.secret_key) }}</div>
            </template>
            <span>点击复制</span>
          </n-popover>
        </div>
        <div class="details">
          <div>
            运营模式:
          </div>
          <n-popover trigger="hover">
            <template #trigger>
              <div>{{ status[item.status_type] }}</div>
            </template>
            <div class="details-popover">
              <ul>
                <li>收费运营：所有接口都必须在用户付费的基础上才可使用</li>
                <li>停止运营：项目的所有接口都将停止</li>
                <li>
                  免费运营：所有的接口都将免费使用，且新注册的用户都会打上免费用户的绑定标签信息，如果将项目调整回收费运营，则所有的免费用户将无法正常使用，必须重新充值才可正常使用！
                </li>
              </ul>
            </div>
          </n-popover>
        </div>
        <div class="details">
          <div>
            加密方式:
          </div>
          <n-popover trigger="hover">
            <template #trigger>
              <div>{{ encrypt[item.encrypt] }}</div>
            </template>
            <div class="soft-details-item">
              <div class="soft-details-item-title">
                <a-tag style="width: 100%" color="#108ee9">
                  密匙A
                </a-tag>
              </div>
              <div @click="copy_inline(item.key_a)" class="view-key">{{ item.key_a }}</div>
            </div>
            <div class="soft-details-item">
              <div class="soft-details-item-title">
                <a-tag style="width: 100%" color="#108ee9">
                  密匙B
                </a-tag>
              </div>
              <div @click="copy_inline(item.key_b)" class="view-key">{{ item.key_b }}</div>
            </div>
            <div class="soft-details-item">
              <div class="soft-details-item-title">
                <a-tag style="width: 100%" color="#108ee9">
                  公钥
                </a-tag>
              </div>
              <div @click="copy(item.public_key)" class="view-key">{{ item.public_key }}
              </div>
            </div>
            <div class="soft-details-item">
              <div class="soft-details-item-title">
                <a-tag style="width: 100%" color="#108ee9">
                  私钥
                </a-tag>
              </div>
              <div @click="copy(item.private_key)" class="view-key">{{ item.private_key }}
              </div>
            </div>
          </n-popover>
        </div>
        <div slot="action" class="card-action">
          <n-button @click="bindTitle='项目绑定规则';showBind=true;row=item" block class="card-action">绑定规则
          </n-button>
          <n-button @click="show=true;title='修改项目';row=item" type="primary" block
                    class="card-action">修改
          </n-button>
          <n-popconfirm
            @positive-click="handleDel(item.id)"
            @negative-click="handleCancel"
          >
            <template #trigger>
              <n-button type="error" block class="card-action">删除</n-button>
            </template>
            确认删除？(删除后用户数据/激活码数据/在线数据将无法恢复)
          </n-popconfirm>
        </div>
      </n-card>
    </n-grid-item>
  </n-grid>
  <div class="pagination">
    <n-pagination
      v-model:page="param.page"
      :page-count="param.total_pages"
      :page-sizes="[7,20,50]"
      size="medium"
      show-quick-jumper
      show-size-picker
      :on-update:page="fetchList"
      :on-update:page-size="(pageSize)=> {param.limit=pageSize;fetchList(1)}"
    />
  </div>
  <component :is="projectEdit" v-model:show="show" v-model:form="row" v-model:title="title"
             @onUpdate="fetchList"></component>
  <component :is="projectBind" v-model:show="showBind" v-model:form="row" v-model:title="bindTitle"
             @onUpdate="fetchList"></component>
  </div>
</template>

<script lang="ts" setup>
import {ref, onMounted, reactive} from 'vue'
import {deleteProject, getProjectList} from "@/api";
import projectEdit from './comp/projectEdit.vue'
import projectBind from './comp/projectBind.vue'
import setting from '@/settings/componentSetting'
import {MessageReactive, useMessage} from "naive-ui";
import {copy_inline,copy} from "@/utils";
import {useAppStore} from "@/store/modules/app";

const appStore = useAppStore()
const message = useMessage()
let index: MessageReactive | null = null
const {project} = setting
const {status, encrypt, type} = project
const dataSource = ref();
const show = ref(false)
const row = ref(null)
const showBind = ref(false)
const title = ref('')
const bindTitle = ref('')
const param = reactive({
  limit: 7,
  page: 1,
  total_pages: 0,
  count: 0
})

async function fetchList(page = undefined) {
  if (page !== undefined) {
    param.page = page
  }
  let data = await getProjectList(param)
  dataSource.value = data.data
  param.total_pages = data.total_pages
  param.count = data.count
}

function withNotice(data) {
  if (data !== null && data.length > 15) {
    return data.substring(0, 15) + '...'
  } else if (data === null || data.length === 0) {
    return '项目创建者很懒，连公告都没写...'
  } else return data
}

function encryKey(key) {
  if (key === null || key === undefined) return '无';
  return key.substring(0, 3) + '*****' + key.substring(key.length - 4, key.length)
}

async function handleDel(id) {
  console.log(id)
  index = message.create("加载中...", {
    type: "loading",
    duration: 10000
  })
  let result = await deleteProject(id)
  if (result !== undefined && result > 0) {
    index.destroy()
    fetchList()
    await appStore.fetchProjectList(true)
  }
}

function handleCancel() {
  console.log("取消操作")
}

onMounted(() => {
  fetchList()
})
</script>

<style lang="less" scoped>

.mask {
  position: absolute;
  background-color: rgba(255, 255, 255, 0.1);
  right: 0;
  left: 0;
  bottom: 0;
  top: 0;
  overflow: hidden;
  z-index: 1;
  display: flex;
  justify-content: center;
  align-items: center;
}

.soft-mask {
  .details {
    -webkit-filter: blur(2px); /* Chrome, Opera */
    -moz-filter: blur(2px);
    -ms-filter: blur(2px);
    filter: blur(2px);
  }

  .notice {
    -webkit-filter: blur(2px); /* Chrome, Opera */
    -moz-filter: blur(2px);
    -ms-filter: blur(2px);
    filter: blur(2px);
  }

  .card-action {
    -webkit-filter: blur(2px); /* Chrome, Opera */
    -moz-filter: blur(2px);
    -ms-filter: blur(2px);
    filter: blur(2px);
  }
}

.card-action {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;

  .card-action {
    width: calc((100% / 3) - 5px);
    display: flex;
    justify-content: center;
  }
}


.notice {
  line-height: 30px;
  font-size: 14px;
  color: #999;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  margin: 10px 0;
  height: 60px;
  background-color: #eee;
  padding: 0 6px;
}

.details-popover {
  overflow-y: hidden;
  overflow-x: auto;
  max-width: 270px;
}

.details {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
  overflow: hidden;
  text-overflow: ellipsis;
  font-size: 12px;


  .details-action {
    a {
      padding: 0 5px;

      &:hover {
        background-color: #f0f0f0;
      }
    }
  }
}

.soft-details-item {
  min-width: 230px;
  margin: 10px 0;

  &:after {
    content: '';
    clear: both;
    display: block;
  }

  .soft-details-item-title {
    float: left;
    width: 70px;
    margin-right: 6px;
  }

  .view-key {
    max-width: 230px;
    max-height: 60px;
    word-wrap: break-word;
    overflow-y: auto;
    display: block;
  }
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
