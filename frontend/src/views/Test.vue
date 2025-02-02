<script setup>
import { getCurrentInstance, ref } from 'vue';
import eventBus from '@/utils/eventBus';
import * as chatAPI from '@/api/chat';

const testApi = async () => {
  testChatApi();
};

const testChatApi = async () => {
  //const response = await chatAPI.getChatList();
  await chatAPI.createChat(2);
};

const testNotification = ()=>
{
  eventBus.emit('notify', {message: '测试消息', type: 'info', duration: 2000});
}
</script>

<template>
  <div>
    <h1>API 测试页面</h1>
    <button @click="testNotification">测试接口</button>
  </div>
</template>
