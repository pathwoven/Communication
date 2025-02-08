<script setup>
import { ref,computed } from 'vue';
import { useRouter } from 'vue-router';
import * as authApi from '@/api/auth';
import { NForm,NFormItem, NAutoComplete, NDatePicker, NRadio, NRadioGroup} from 'naive-ui';

const router = useRouter();

const model= ref({
  account: '',
  nickname: '',
  password: '',
  confirmPassword: '',
  email: '',
  emailVerification: '',
  sex: '2',
  birthday: null,
  signature: '',
});
const register = async () => {
  if (model.value.password !== model.value.confirmPassword) {
    alert('密码和确认密码不匹配');
    return;
  }
  const response = await authApi.register(model.value.nickname, model.value.password, model.value.email,model.value.account, parseInt(model.value.sex,10), model.value.birthday, model.value.signature);
  if (response.success) {
    console.log('注册成功');
    router.push('/auth/login');
  } else {
    console.log('注册失败');
  }
};

// 表单
// 规则
const rules = {
  account: [
    {
      validator(rule, value){
        if(value){
          if(!/^[a-zA-Z0-9_]*$/.test(value)){
            return new Error('账号需由字母、数字、下划线组成');
          }else if(value.length < 4 || value.length > 16){
            return new Error('账号长度在 4 到 16 个字符');
          }
        }
        return true;
      }, 
      trigger: 'blur' 
    },
  ],
  nickname: [
    {
      validator(rule, value){
        if(value){
          if(value.length < 2 || value.length > 20){
            return new Error('昵称长度在 2 到 20 个字符');
          }
        }else{
          return new Error('请输入昵称');
        }
      }, 
      trigger: 'blur'
    }
  ],
  password: [
    {
      validator(rule, value){
        if(value){
          if(value.length < 6 || value.length > 20){
            return new Error('密码长度在 6 到 20 个字符');
          }
        }else{
          return new Error('请输入密码');
        }
      }, 
      trigger: 'blur'
    }
  ],
  confirmPassword: [
    {
      validator(rule, value){
        if(value !== model.value.password){
          return new Error('两次输入密码不一致');
        }
        else{
          return new Error('请输入确认密码');
        }
      }, 
      trigger: 'blur'
    }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱', trigger: 'blur' },
  ],
  
}
// 邮箱自动补全
const emailOptions = computed(()=>{
  return ["@gmail.com", "@163.com", "@qq.com"].map((suffix)=>{
    const prefix = model.value.email.split('@')[0];
    return{
      label: prefix + suffix,
      value: prefix + suffix,
    }
  });
})
</script>

<template>
  <div class="register-container">
    <n-form
      :model="model"
      :rules="rules"
      label-placement="left"
      label-width="100px"
      require-mark-placement="right-hanging"
      class="form"
    >
      <n-form-item path="account" label="账号" class="item">
        <n-input v-model:value="model.account" placeholder="账号" />
      </n-form-item>
      <n-form-item show-require-mark path="nickname" label="昵称" class="item">
        <n-input v-model:value="model.nickname" placeholder="昵称" />
      </n-form-item>
      <n-form-item show-require-mark path="password" label="密码" class="item">
        <n-input v-model:value="model.password" type="password" placeholder="密码" show-password-on="click"/>
      </n-form-item>
      <n-form-item show-require-mark path="confirmPassword" label="确认密码" class="item">
        <n-input v-model:value="model.confirmPassword" type="password" placeholder="确认密码" show-password-on="click"/>
      </n-form-item>
      <n-form-item show-require-mark first path="email" label="邮箱" class="item">
        <n-auto-complete
          v-model:value="model.email"
          :input-props="{
            autocomplete: 'disabled',
          }"
          :options="emailOptions"
          placeholder="邮箱"
          clearable
        />
      </n-form-item>
      <n-form-item show-require-mark path="emailVerification" label="邮箱验证" class="item">
        <n-input v-model:value="model.emailVerification" placeholder="邮箱验证" />
      </n-form-item>
      <n-form-item show-require-mark path="sex" label="性别" class="item">
        <n-radio-group v-model:value="model.sex">
          <n-radio value="2">未知</n-radio>
          <n-radio value="0">男</n-radio>
          <n-radio value="1">女</n-radio>
        </n-radio-group>
      </n-form-item>
      <n-form-item path="birthday" label="生日" class="item">
        <n-date-picker v-model:value="model.birthday" type="date" />
      </n-form-item>
      <n-form-item path="signature" label="个性签名" class="item">
        <n-input v-model:value="model.signature" placeholder="个性签名" />
      </n-form-item>
      <n-button attr-type="submit" @click="register" class="item">注册</n-button>
    </n-form>
  </div>
</template>

<style scoped>
.register-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 500px;
}
.form {
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.item{
  
}
</style>