<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import * as authApi from '@/api/auth';

const displayId = ref('');
const password = ref('');
const router = useRouter();

const login = async () =>{
  // api todo
  const response = await authApi.login(displayId.value, password.value);
  if(response.success){
   console.log("登录成功");
   //router.push('/'); 
  }else{
   console.log("登录失败");
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
   <input v-model="displayId" type="text" placeholder="账号" class="input-field" />
   <input v-model="password" type="password" placeholder="密码" class="input-field" />
   <button @click="login" class="button">Login</button>
   <button @click="goToForget" class="button">Forget</button>
   <button @click="goToRegister" class="button">Register</button>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  background-color: #f5f5f5;
}

.input-field {
  margin: 10px 0;
  padding: 10px;
  width: 200px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.button {
  margin: 10px 0;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  cursor: pointer;
}

.button:hover {
  background-color: #0056b3;
}
</style>