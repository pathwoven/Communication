<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import * as authApi from '@/api/auth';

const account = ref('');
const nickname = ref('');
const password = ref('');
const confirmPassword = ref('');
const email = ref('');
const emailVerification = ref('');
const sex = ref('2');
const birthday = ref('');
const signature = ref('');
const router = useRouter();

const register = async () => {
  if (password.value !== confirmPassword.value) {
    alert('密码和确认密码不匹配');
    return;
  }

  const response = await authApi.register(nickname.value,password.value, email.value,account.value, parseInt(sex.value,10), birthday.value, signature.value);

  if (response.success) {
    console.log('注册成功');
    router.push('/auth/login');
  } else {
    console.log('注册失败');
  }
};
</script>

<template>
  <div class="register-container">
    <form @submit.prevent="register">
      <div>
        <label for="account">账号（可为空，不可再次修改）</label>
        <input id="account" v-model="account" type="text" placeholder="账号" />
      </div>
      <div>
        <label for="nickname">昵称</label>
        <input id="nickname" v-model="nickname" type="text" placeholder="昵称" required />
      </div>
      <div>
        <label for="password">密码</label>
        <input id="password" v-model="password" type="password" placeholder="密码" required />
      </div>
      <div>
        <label for="confirmPassword">确认密码</label>
        <input id="confirmPassword" v-model="confirmPassword" type="password" placeholder="确认密码" required />
      </div>
      <div>
        <label for="email">邮箱</label>
        <input id="email" v-model="email" type="email" placeholder="邮箱" required />
      </div>
      <div>
        <label for="emailVerification">邮箱验证</label>
        <input id="emailVerification" v-model="emailVerification" type="text" placeholder="邮箱验证" required />
      </div>
      <div>
        <label for="sex">性别</label>
        <select id="sex" v-model="sex">
          <!-- <option value="">请选择</option> -->
          <option value="2">未知</option>
          <option value="0">男</option>
          <option value="1">女</option>
         
        </select>
      </div>
      <div>
        <label for="birthday">生日</label>
        <input id="birthday" v-model="birthday" type="date" />
      </div>
      <div>
        <label for="signature">个性签名</label>
        <textarea id="signature" v-model="signature" placeholder="个性签名"></textarea>
      </div>
      <button type="submit">注册</button>
    </form>
  </div>
</template>

<style scoped>
register-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  background-color: #f5f5f5;
}
form {
  display: flex;
  flex-direction: column;
}

form > div {
  margin-bottom: 1em;
}

label {
  font-size: 14px;
}

input, select, textarea {
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 3px;
  padding-left: 5px;
}

button {
  padding: 0.5em 1em;
  font-size: 1em;
  color: white;
  background-color: #42b983;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #369f6b;
}
</style>