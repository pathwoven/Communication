<script setup>
import { NAvatar } from 'naive-ui';
import * as contactApi from '@/api/contact';
import { ref } from 'vue';

// 搜索
const keyword = ref('')
const search = async () => {
  const response = await contactApi.searchStranger(keyword.value, false);
  if(response.success){
    users.value = response.data
  }else{
    // todo
  }
}

// 用户
const users = ref([])
</script>

<template>
  <div class="add-friend">
    <n-input placeholder="请输入" v-model:value="keyword" @keyup.enter="search"/>
    <ul>
      <li v-for="user in users" :key="user.user_id">
        <n-avatar round size="small" ></n-avatar>
        <div class="info">
          <span>用户名</span>
          <span>用户id</span>
        </div>
        <n-button type="primary">添加</n-button>
      </li>
    </ul>
  </div>
</template>

<style scoped>

</style>