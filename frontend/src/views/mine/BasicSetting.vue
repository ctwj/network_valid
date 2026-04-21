<template>
  <n-grid cols="2 s:2 m:2 l:3 xl:3 2xl:3" responsive="screen">
    <n-grid-item>
      <n-form :label-width="80" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="账号" path="user">
          <n-input v-model:value="formValue.user" placeholder="请输入账号" />
        </n-form-item>
        <n-form-item label="密码" path="pwd">
          <n-input placeholder="请输入新密码" v-model:value="formValue.pwd" />
        </n-form-item>
        <n-form-item label="邮箱" path="email">
          <n-input placeholder="请输入邮箱" v-model:value="formValue.email" />
        </n-form-item>
        <div>
          <n-space>
            <n-button type="primary" @click="formSubmit">更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-grid-item>
  </n-grid>
</template>

<script lang="ts" setup>
  import { reactive, ref, onMounted} from 'vue';
  import { useMessage } from 'naive-ui';
  import {getUserInfo, update} from "@/api/system/user";
  const rules = {
    user: {
      required: true,
      message: '请输入昵称',
      trigger: 'blur',
    },
    email: {
      required: true,
      message: '请输入邮箱',
      trigger: 'blur',
    }
  };
  const formRef: any = ref(null);
  const message = useMessage();

  const formValue = reactive( {
    "id": 0,
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
  });

  function formSubmit() {
    formRef.value.validate(async (errors) => {
      if (!errors) {
        await update(formValue)
      } else {
        message.error('验证失败，请填写完整信息');
      }
    });
  }
  onMounted(async () =>{
    const {user} = await getUserInfo()
    formValue.user = user.user
    formValue.pwd = user.pwd
    formValue.email = user.email
  })
</script>
