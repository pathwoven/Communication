<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import {useMessage, NForm, NFormItem} from 'naive-ui'
import { useSettingStore } from '@/store/modules/setting';
import { createWebSocketConnection } from '@/api/websocket';
import * as authApi from '@/api/auth';

const account = ref({
  displayId: '',
  password: '',
});
const router = useRouter();
const message = useMessage();

const settingStore = useSettingStore();

const login = async () =>{
  // api todo
  const response = await authApi.login(account.value.displayId, account.value.password);
  if(response.success){
    message.success('登录成功');
    // 写入个人信息
    settingStore.updateSettings(response.data.setting);
    settingStore.setUser({
      id: response.data.id,
      nickname: response.data.name,
      avatar: response.data.avatar,
      display_id: response.data.display_id,
    });
    // 创建ws连接
    createWebSocketConnection();
    router.push('/chat');
  }else{
    message.error('登录失败');
  }

}
const goToForget = () =>{
   router.push('/auth/forget');
}
const goToRegister = () =>{
   router.push('/auth/register');
}
</script>

<template>
  <div class="login-container">
    <n-form
      class="input-field"
      :model="account"
      :rules="{
        displayId: [{ required: true, message: '请输入账号', trigger: 'blur' }],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
      }"
    >
      <n-form-item path="displayId" >
        <n-input v-model:value="account.displayId" placeholder="账号" class="item"/>
      </n-form-item>
      <n-form-item path="password">
        <n-input v-model:value="account.password" placeholder="密码" class="item"/>
      </n-form-item>
      <n-form-item style="display: flex;flex-direction: column;align-items: center;">
        <n-button attr-type="submit" @click="login" class="item">Login</n-button>
      </n-form-item>
    </n-form> 
    <div class="action">
      <n-button @click="goToForget" class="button" text>Forget</n-button>
      <n-button @click="goToRegister" class="button" text>Register</n-button>
    </div>
  </div>
</template>

<style scoped>
.login-container {
  height: 300px;
  display: flex;
  justify-content: center;
  flex-direction: column;
}
.input-field {
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.item {
  margin: 10px 0;
  padding: 10px 20px;
}
.action {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
}
</style>