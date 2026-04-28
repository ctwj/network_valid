<template>
  <n-modal :title="title" style="width: 750px" preset="card" v-model:show="modalShow">
    <n-form
      ref="formRef"
      :inline="false"
      :label-width="80"
      :model="formData"
      :rules="rules"
    >
      <n-tabs type="line" animated>
        <n-tab-pane name="1" tab="基础设置">
          <n-form-item label="项目名称" path="name">
            <n-input v-model:value="formData.name" placeholder="请输入项目名称"/>
          </n-form-item>
          <n-form-item label="项目公告" path="notice">
            <n-input
              v-model:value="formData.notice"
              type="textarea"
              placeholder="公告内容"
            />
          </n-form-item>
        </n-tab-pane>

        <n-tab-pane name="2" tab="运营模式">
          <n-card size="small" :bordered="false" style="background: var(--n-color-modal)">
            <template #header>
              <n-text strong>运营模式</n-text>
            </template>
            <n-radio-group v-model:value="formData.status_type" name="status_type">
              <n-radio-button
                v-for="(item, index) in status"
                :value="index"
                :key="index"
              >
                {{ item }}
              </n-radio-button>
            </n-radio-group>
            <n-alert :type="statusConfig[formData.status_type].type" style="margin-top: 12px">
              {{ statusConfig[formData.status_type].desc }}
            </n-alert>
          </n-card>
        </n-tab-pane>

        <n-tab-pane name="3" tab="加密模式">
          <n-card size="small" :bordered="false" style="background: var(--n-color-modal)">
            <template #header>
              <n-text strong>加密模式</n-text>
            </template>
            <n-radio-group v-model:value="formData.encrypt" name="encrypt">
              <n-radio-button
                v-for="(item, index) in encrypt"
                :value="index"
                :key="index"
              >
                {{ item }}
              </n-radio-button>
            </n-radio-group>
            <n-alert :type="encryptConfig[formData.encrypt].type" style="margin-top: 12px">
              {{ encryptConfig[formData.encrypt].desc }}
            </n-alert>
          </n-card>
        </n-tab-pane>

        <n-tab-pane name="4" tab="签名算法">
          <n-card size="small" :bordered="false" style="background: var(--n-color-modal)">
            <template #header>
              <n-text strong>签名算法</n-text>
            </template>
            <n-radio-group v-model:value="formData.sign" name="sign">
              <n-radio-button
                v-for="(item, index) in hash"
                :value="index"
                :key="index"
              >
                {{ item }}
              </n-radio-button>
            </n-radio-group>
            <n-alert :type="signConfig[formData.sign].type" style="margin-top: 12px">
              {{ signConfig[formData.sign].desc }}
            </n-alert>
          </n-card>
        </n-tab-pane>

        <n-tab-pane name="5" tab="套餐方案">
          <n-space vertical style="width: 100%">
            <n-alert type="info" style="margin-bottom: 12px">
              选择套餐方案后，系统将自动创建对应的套餐类型。您可以在创建后根据需要调整价格和配额。
            </n-alert>

            <n-card size="small" :bordered="false" style="background: var(--n-color-modal); margin-bottom: 12px">
              <template #header>
                <n-text strong>定价基准</n-text>
              </template>
              <n-form-item label="单月费用" path="monthly_price">
                <n-input-number
                  v-model:value="formData.monthly_price"
                  :min="1"
                  :precision="2"
                  :step="5"
                  placeholder="输入单月费用"
                  style="width: 200px"
                >
                  <template #prefix>¥</template>
                </n-input-number>
                <n-text depth="3" style="margin-left: 12px; font-size: 13px">
                  套餐价格将基于此费用自动计算
                </n-text>
              </n-form-item>
            </n-card>

            <n-radio-group v-model:value="formData.scheme" name="scheme" style="width: 100%">
              <n-grid :cols="1" :x-gap="12" :y-gap="12">
                <n-gi v-for="scheme in computedPlanSchemes" :key="scheme.name">
                  <n-card
                    size="small"
                    :style="{
                      border: formData.scheme === scheme.name ? '2px solid #18a058' : '1px solid #e0e0e6',
                      cursor: 'pointer',
                      transition: 'all 0.2s ease',
                      background: formData.scheme === scheme.name ? 'rgba(24, 160, 88, 0.04)' : '#fff'
                    }"
                    @click="formData.scheme = scheme.name">
                    <template #header>
                      <n-space align="center">
                        <n-radio :value="scheme.name">
                          <n-text strong style="font-size: 16px">{{ scheme.name }}</n-text>
                        </n-radio>
                        <n-tag v-if="scheme.name === '标准推荐'" type="success" size="small">
                          <template #icon>
                            <n-icon><StarOutline /></n-icon>
                          </template>
                          推荐
                        </n-tag>
                      </n-space>
                    </template>
                    <n-text depth="3" style="font-size: 13px">{{ scheme.description }}</n-text>

                    <n-divider style="margin: 12px 0"/>

                    <n-space wrap>
                      <n-tag
                        v-for="plan in scheme.plans"
                        :key="plan.name"
                        :type="plan.is_free_tier ? 'default' : (plan.days === 0 ? 'warning' : 'info')"
                        size="medium"
                        round>
                        <n-space :size="4" align="center">
                          <n-text strong>{{ plan.name }}</n-text>
                          <n-text depth="3">|</n-text>
                          <template v-if="plan.computedPrice > 0">
                            <n-text type="error" strong>¥{{ plan.computedPrice.toFixed(2) }}</n-text>
                            <n-text v-if="plan.savings" depth="3" style="font-size: 12px">
                              ({{ plan.savings }})
                            </n-text>
                          </template>
                          <template v-else-if="plan.is_free_tier">
                            <n-text depth="3">免费</n-text>
                          </template>
                          <template v-else>
                            <n-text type="warning">永久</n-text>
                          </template>
                        </n-space>
                      </n-tag>
                    </n-space>

                    <n-space style="margin-top: 12px">
                      <n-text depth="3" style="font-size: 12px">
                        包含 {{ scheme.plans.length }} 个套餐
                      </n-text>
                    </n-space>
                  </n-card>
                </n-gi>
              </n-grid>
            </n-radio-group>
          </n-space>
        </n-tab-pane>

        <n-tab-pane v-if="formData.id > 0" name="6" tab="密钥相关">
          <n-form-item label="重置RSA" path="type">
            <n-radio-group v-model:value="formData.update_rsa" name="radiogroup2">
              <n-radio v-for="(item,index) in change" :value="index" :key="index">
                {{ item }}
              </n-radio>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="重置AES" path="type">
            <n-radio-group v-model:value="formData.update_key" name="radiogroup2">
              <n-radio v-for="(item,index) in change" :value="index" :key="index">
                {{ item }}
              </n-radio>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="重置AppKey" path="type">
            <n-radio-group v-model:value="formData.update_app_key" name="radiogroup2">
              <n-radio v-for="(item,index) in change" :value="index" :key="index">
                {{ item }}
              </n-radio>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="重置SecretKey" path="type">
            <n-radio-group v-model:value="formData.update_secret_key" name="radiogroup2">
              <n-radio v-for="(item,index) in change" :value="index" :key="index">
                {{ item }}
              </n-radio>
            </n-radio-group>
          </n-form-item>
        </n-tab-pane>
      </n-tabs>

      <n-form-item style="margin-top: 16px">
        <n-space>
          <n-button type="primary" attr-type="button" @click="handleOk">
            {{ title }}
          </n-button>
          <n-button attr-type="button" @click="modalShow = false">
            取消
          </n-button>
        </n-space>
      </n-form-item>
    </n-form>
  </n-modal>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, onMounted, Ref, watch, computed} from 'vue';
import {useAppStore} from "@/store/modules/app";
import {createProject, updateProject, getPlanSchemes} from "@/api";
import {useMessage, MessageReactive} from 'naive-ui'
import {StarOutline} from '@vicons/ionicons5'
import setting from '@/settings/componentSetting'
const {project} = setting

// 运营模式配置说明
const statusConfig = [
  { type: 'success', desc: '收费运营：用户需要购买套餐才能使用，适合商业化运营场景。系统将启用完整的付费流程和订单管理功能。SDK 调用需验证用户会员状态，未付费用户将返回错误码。' },
  { type: 'warning', desc: '停止运营：项目暂停服务，所有用户无法登录使用。适合维护期间或暂停业务时使用。SDK 所有接口将返回维护中提示。' },
  { type: 'info', desc: '免费运营：用户无需付费即可使用，适合内部测试、演示或公益项目。SDK 调用无需验证会员状态，新注册用户自动获得免费标签。' }
]

// 加密模式配置说明
const encryptConfig = [
  { type: 'warning', desc: '开放API：不进行数据加密，请求参数明文传输。适合内部系统或已通过 HTTPS 保护的环境，性能最优但安全性较低。SDK 无需额外配置加密参数。' },
  { type: 'success', desc: 'AES加密：使用 AES 对称加密算法保护数据传输安全。SDK 需配置 KeyA 和 KeyB 进行加解密，安全性更高，推荐用于生产环境。' }
]

// 签名算法配置说明
const signConfig = [
  { type: 'warning', desc: 'MD5：生成 128 位哈希值，速度快但安全性较低。SDK 使用 GenerateSign() 函数，参数拼接顺序：appkey + secretkey + version + timestamp + mac。' },
  { type: 'info', desc: 'SHA1：生成 160 位哈希值，比 MD5 更安全，但已发现碰撞漏洞。SDK 需使用对应 SHA1 签名函数。' },
  { type: 'info', desc: 'SHA224：生成 224 位哈希值，安全性较好，适合一般应用场景。SDK 需使用对应 SHA224 签名函数。' },
  { type: 'success', desc: 'SHA256：生成 256 位哈希值，安全性高，推荐使用。适合大多数安全敏感场景，是目前最常用的签名算法。SDK 需使用对应 SHA256 签名函数。' },
  { type: 'success', desc: 'SHA384：生成 384 位哈希值，安全性更高。适合对安全性要求极高的场景，计算开销略大。SDK 需使用对应 SHA384 签名函数。' },
  { type: 'success', desc: 'SHA512：生成 512 位哈希值，安全性最高。适合金融、安全认证等高安全需求场景，计算开销最大。SDK 需使用对应 SHA512 签名函数。' }
]

const formVal = {
  id: 0,
  name: '',
  type: 1,  // 默认用户模式（后端会强制设置为用户模式）
  status_type: 0,
  encrypt: 0,
  notice: '',
  api: '',
  sign: 0,
  scheme: '标准推荐',  // 默认标准推荐方案
  monthly_price: 30,  // 默认单月费用
  update_rsa: 0,
  update_key: 0,
  update_app_key: 0,
  update_secret_key: 0
}
export default defineComponent({
  components: {
    StarOutline
  },
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
    const appStore = useAppStore()
    const {status, encrypt, hash} = project
    const change = reactive(["不重置", "重置"])
    const modalShow: Ref<boolean> = ref(false)
    const rules = reactive({});
    const message = useMessage()
    let index: MessageReactive | null = null
    const planSchemes: Ref<any[]> = ref([])
    const formData: Ref<Object> = ref({
      id: 0,
      name: '',
      type: 1,
      status_type: 0,
      encrypt: 0,
      notice: '',
      api: '',
      sign: 0,
      scheme: '标准推荐',
      monthly_price: 30,
      update_rsa: 0,
      update_key: 0,
      update_app_key: 0,
      update_secret_key: 0
    })

    // 计算套餐价格
    const calculatePlanPrice = (plan: any, monthlyPrice: number, schemeName: string): number => {
      if (plan.is_free_tier) return 0
      if (plan.days === 0) {
        // 永久套餐
        if (schemeName === '高级专业') return Math.round(monthlyPrice * 66.6 * 100) / 100
        return Math.round(monthlyPrice * 66.6 * 100) / 100
      }
      switch (schemeName) {
        case '入门引导':
          if (plan.days === 30) return Math.round(monthlyPrice * 100) / 100
          if (plan.days === 90) return Math.round(monthlyPrice * 2.67 * 100) / 100
          break
        case '标准推荐':
          if (plan.days === 30) return Math.round(monthlyPrice * 100) / 100
          if (plan.days === 90) return Math.round(monthlyPrice * 2.63 * 100) / 100
          if (plan.days === 365) return Math.round(monthlyPrice * 8.14 * 100) / 100
          break
        case '高级专业':
          if (plan.days === 30) return Math.round(monthlyPrice * 3.3 * 100) / 100
          if (plan.days === 90) return Math.round(monthlyPrice * 8.3 * 100) / 100
          if (plan.days === 365) return Math.round(monthlyPrice * 23.3 * 100) / 100
          break
      }
      return Math.round(monthlyPrice * plan.days / 30 * 100) / 100
    }

    // 计算后的套餐方案
    const computedPlanSchemes = computed(() => {
      const monthlyPrice = (formData.value as any).monthly_price || 30
      return planSchemes.value.map(scheme => ({
        ...scheme,
        plans: scheme.plans.map((plan: any) => ({
          ...plan,
          computedPrice: calculatePlanPrice(plan, monthlyPrice, scheme.name)
        }))
      }))
    })

    // 获取预设方案
    const fetchPlanSchemes = async () => {
      try {
        const result = await getPlanSchemes()
        if (result) {
          planSchemes.value = result
        }
      } catch (e) {
        console.error('获取预设方案失败', e)
      }
    }

    const handleOk = async () => {
      if (!formData.value.name) {
        message.error('请输入项目名称')
        return
      }
      index = message.create("加载中...", {
        type: "loading",
        duration: 10000
      })
      if (formData.value.id > 0) {
        let result = await updateProject(formData.value)
        if (result !== undefined && result > 0) {
          modalShow.value = false
          index.destroy()
          await appStore.fetchCardList(true)
          emit("on-update")
        }else {
          index.destroy()
        }
      } else {
        let result = await createProject(formData.value)
        if (result !== undefined && result > 0) {
          modalShow.value = false
          index.destroy()
          await appStore.fetchCardList(true)
          emit("on-update")
        }else {
          index.destroy()
        }
      }
    }

    onMounted(async () => {
      modalShow.value = props.show
      await fetchPlanSchemes()
    })

    watch(() => props.show, async (n) => {
      modalShow.value = n
      if (n) {
        await fetchPlanSchemes()
      }
    })
    watch(() => props.form, (n) => {
      if (n === null) {
        formData.value = {...formVal}
      } else {
        formData.value = {...n}
      }
    })
    watch(modalShow, (n) => {
      emit("update:show", n)
    })

    return {
      status,
      encrypt,
      formData,
      rules,
      modalShow,
      change,
      hash,
      planSchemes,
      computedPlanSchemes,
      handleOk,
      statusConfig,
      encryptConfig,
      signConfig
    }
  }
})
</script>

<style scoped>
</style>
