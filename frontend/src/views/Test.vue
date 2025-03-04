<script setup>
import { getCurrentInstance, ref, useTemplateRef } from 'vue';
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

const linkRef = useTemplateRef("link")
const testFile = async() =>{
  const response = await chatAPI.downloadMessageFile(61);
  let blob = new Blob([response.data.file], {type: response.data.contentType})
  console.log(blob)
  linkRef.value.href = URL.createObjectURL(blob)
  linkRef.value.download = response.data.name
  linkRef.value.click()
}
</script>

<template>
  <div>
    <h1>API 测试页面</h1>
    <button @click="testFile">测试接口</button>
    <a ref="link" style="visibility: hidden;" download>下载</a>
  </div>
</template>
