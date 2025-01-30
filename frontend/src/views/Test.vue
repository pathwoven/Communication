<script setup>
import { ref } from 'vue';
import * as chatAPI from '@/api/chat';

const testApi = async () => {
  testChatApi();
};

const testChatApi = async () => {
  //const response = await chatAPI.getChatList();
  await chatAPI.createChat(2);
};
</script>

<template>
  <div>
    <h1>API 测试页面</h1>
    <button @click="testApi">测试接口</button>
  </div>
</template>
