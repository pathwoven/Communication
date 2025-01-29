<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import * as loginApi from '@/api/login';

const displayId = ref('');
const password = ref('');
const router = useRouter();

const login = async () =>{
   // api todo
  const response = await loginApi.login(displayId.value, password.value);
  if(response.success){
    console.log("登录成功");
    router.push('/'); 
  }else{
    console.log("登录失败");
  }
}
const goToForget = () =>{
    router.push('/login/forget');
}
const goToRegister = () =>{
    router.push('/login/register');
}
</script>

<template>
  <div>
    <input v-model="displayId" type="text" placeholder="账号" />
    <input v-model="password" type="password" placeholder="密码" />
    <button @click="login">Login</button>
    <button @click="goToForget">Forget</button>
    <button @click="goToRegister">Register</button>
  </div>
</template>

<style scoped>

</style>