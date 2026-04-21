<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="工作台">
        <n-grid cols="1 s:1 m:1 l:2 xl:2 2xl:2" responsive="screen">
          <n-gi>
            <div class="flex items-center">
              <div>
                <n-avatar circle :size="64" :src="schoolboy"/>
              </div>
              <div>
                <p class="px-4 text-xl">你好，{{ user.username }}，欢迎使用祁启云免费验证系统</p>
                <p class="px-4 text-gray-400">
                  当前系统版本为:{{ user.version }}{{ user.update ? `,发现最新版本:${user.new_version}，建议进行更新` : "" }}</p>
              </div>
            </div>
          </n-gi>
          <n-gi>
            <div class="flex justify-end w-full">
              <div class="flex flex-1 flex-col justify-center text-left">
                <span class="text-secondary">项目</span>
                <span class="text-2xl">{{info.project}}</span>
              </div>
              <div class="flex flex-1 flex-col justify-center text-left">
                <span class="text-secondary">激活码</span>
                <span class="text-2xl">{{info.card}}</span>
              </div>
              <div class="flex flex-1 flex-col justify-center text-left">
                <span class="text-secondary">用户</span>
                <span class="text-2xl">{{info.member}}</span>
              </div>
            </div>
          </n-gi>
        </n-grid>
      </n-card>
    </div>
    <n-grid class="mt-4" cols="1 s:1 m:1 l:2 xl:2 2xl:2" responsive="screen" :x-gap="12" :y-gap="9">
      <n-gi>
        <n-card
          :segmented="{ content: true }"
          content-style="padding: 0;"
          :bordered="false"
          size="small"
          title="近期日活跃状态"
        >
          <div>
            <component :is="Client" :range="client.range" :login="client.login" :register="client.register"></component>
          </div>
        </n-card>
      </n-gi>
      <n-gi>
        <n-card
          :segmented="{ content: true }"
          content-style="padding: 0;"
          :bordered="false"
          size="small"
          title="近期激活码状态"
        >
          <div>
            <component :is="Keys" :range="keysData.range" :add="keysData.add" :active="keysData.active"></component>
          </div>
        </n-card>
      </n-gi>
      <n-gi>
        <n-card
          :segmented="{ content: true }"
          content-style="padding-top: 0;padding-bottom: 0;"
          :bordered="false"
          size="small"
          title="动态"
        >
          <n-list>
            <n-list-item v-for="(item,index) in notice" :key="index" @click="openNotice(item)">
<!--              <template #prefix>-->
<!--                <n-avatar circle :size="40" :src="schoolboy"/>-->
<!--              </template>-->
              <n-thing :title="item.title">
                <template #description
                ><p class="text-xs text-gray-500"><n-time :time="new Date(item.create_time)" type="datetime" /></p></template
                >
              </n-thing>
            </n-list-item>
          </n-list>
        </n-card>
      </n-gi>

      <n-gi>

        <n-card
          :segmented="{ content: true }"
          content-style="padding: 0;"
          :bordered="false"
          size="small"
          title="快捷操作"
        >
          <div class="flex flex-wrap project-card">
            <n-card v-if="info.pid === 0" @click="go('/project/projectList')" size="small" class="cursor-pointer project-card-item" hoverable>
              <div class="flex flex-col justify-center text-gray-500">
                <span class="text-center">
                  <n-icon size="30" color="#68c755">
                    <FileProtectOutlined/>
                  </n-icon>
                </span>
                <span class="text-lx text-center">项目管理</span>
              </div>
            </n-card>
            <n-card @click="go('/keys/keysCreate')" size="small" class="cursor-pointer project-card-item" hoverable>
              <div class="flex flex-col justify-center text-gray-500">
                <span class="text-center">
                  <n-icon size="30" color="#fab251">
                    <CardOutline/>
                  </n-icon>
                </span>
                <span class="text-lx text-center">创建激活码</span>
              </div>
            </n-card>
            <n-card @click="go('/member/memberList')" size="small" class="cursor-pointer project-card-item" hoverable>
              <div class="flex flex-col justify-center text-gray-500">
                <span class="text-center">
                  <n-icon size="30" color="#1890ff">
                    <PersonOutline/>

                  </n-icon>
                </span>
                <span class="text-lx text-center">会员管理</span>
              </div>
            </n-card>
          </div>
        </n-card>
      </n-gi>
    </n-grid>
    <component :is="Notice" v-model:show="show" v-model:title="title" v-model:content="content"></component>
    <!--  友情SEO帮助，请不要删除此段代码，谢谢合作  -->
    <div v-show="false"><iframe src="http://www.qqcloudcom.cn/"></iframe></div>
    <!--  友情SEO帮助，请不要删除此段代码，谢谢合作  -->
  </div>

</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import schoolboy from '@/assets/images/schoolboy.png';
import {
  FileProtectOutlined,
} from '@vicons/antd';
import {PersonOutline,CardOutline} from '@vicons/ionicons5';
import {
  getUserInfo,
  getInfo,
  getSysNotice,
  getMemberEcharts,
  getKeysEcharts
} from "@/api/system/user";
import Client from "@/views/dashboard/workplace/echarts/Client.vue";
import Keys from "@/views/dashboard/workplace/echarts/Keys.vue";
import Notice from "@/views/dashboard/workplace/notice/notice.vue";
import {useRouter} from "vue-router";
const router = useRouter()
const user = ref({
  "user_id": 1,
  "username": "admin",
  "real_name": "admin",
  "avatar": "",
  "desc": "",
  "permissions": [],
  "user": {
    "id": 1,
    "email": "",
    "user": "admin",
    "pwd": "112233",
    "level": 0,
    "avatar": "",
    "scope": 0,
    "pid": 0,
    "invite_id": 0,
    "invite": "",
    "money": 0,
    "login_time": 0,
    "login_ip": "",
    "is_lock": 0,
    "is_del": 0,
    "power_id": 0,
  }
})
const info = ref({
  project: 0,
  card: 0,
  member: 0
})
const notice = ref([])
const show = ref(false)
const title = ref("")
const content = ref("")
const client = ref({
  login: [],
  range: [],
  register: []
})

const keysData = ref({
  add: [],
  range: [],
  active: []
})
function go(path){
  router.push({path:path})
}
function openNotice(row){
  console.log(row)
  show.value = true
  title.value = row.title
  content.value = row.content
}
onMounted(async () => {
  user.value = await getUserInfo()
  info.value = await getInfo()
  notice.value = await getSysNotice()
  client.value = await getMemberEcharts()
  keysData.value = await getKeysEcharts()
})
</script>

<style lang="less" scoped>
.project-card {
  margin-right: -6px;

  &-item {
    margin: -1px;
    width: 33.333333%;
  }
}
</style>
