<script setup>
import { ref } from 'vue'
import * as chatApi from '@/api/chat'
import { defineProps, defineEmits } from 'vue'

const props = defineProps({
  isGroup: Boolean,
  operation: String,
  op_message_id: Number,
  receiver_id: Number
})

const { isGroup, operation, op_message_id, receiver_id } = props

const emit = defineEmits(['send-message'])

const newMessage = ref('')

// 发送文本消息，0为文本类型
const sendTextMessage = async () => {
  let response;
  if(isGroup){
    response = await chatApi.sendGroupMessage(receiver_id,newMessage.value, 0, operation,op_message_id)
  }else{
    response = await chatApi.sendSingleMessage(receiver_id,newMessage.value, 0,operation,op_message_id)
  }
  if(!response.success){
    console.log('发送消息失败')   // 提醒 todo
  }else{
    emit('send-message', {
      id: response.data.id,
      create_time: response.data.create_time,
      content: newMessage.value,
      content_type: 0,
    })
  }
  newMessage.value = ''
}
</script>

<template>
  <div class="message-input">
    <!-- 操作 -->
    <div class = "operation">
      <n-button text>图片</n-button>
      <n-button text>文件</n-button>
      <n-button text>群投票</n-button>
    </div>
    <n-input v-model:value="newMessage" placeholder="请输入消息" />
    <n-button type="primary" @click="sendTextMessage">发送</n-button>
  </div>
</template>

<style scoped>

</style>