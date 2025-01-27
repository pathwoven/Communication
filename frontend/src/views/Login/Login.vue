<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import * as loginApi from '@/api/login';

const username = ref('');
const password = ref('');
const router = useRouter();

const login = () =>{
   // api todo
   // loginApi.login(username.value, password.value);
  console.log("登录");
  router.push('/'); 
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
    <input v-model="username" type="text" placeholder="username" />
    <input v-model="password" type="password" placeholder="password" />
    <button @click="login">Login</button>
    <button @click="goToForget">Forget</button>
    <button @click="goToRegister">Register</button>
  </div>
</template>

<style scoped>

</style>