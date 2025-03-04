<script setup>
import { ref, useTemplateRef } from 'vue'
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

const sendMessage = async (formData, content,content_type) =>{
  if(!operation){
    formData.append('operation', '')
  }else{
    formData.append('operation', operation)
  }
  if(!op_message_id){
    formData.append('op_message_id', '')
  } else {
    formData.append('op_message_id', op_message_id)
  }
  let response;
  if(isGroup){
    response = await chatApi.sendGroupMessage(formData)
  }else{
    response = await chatApi.sendSingleMessage(formData)
  }
  if(!response.success){
    console.log('发送消息失败')   // 提醒 todo
  }else{
    emit('send-message', {
      id: response.data.id,
      create_time: response.data.create_time,
      content: content,
      content_type: content_type,
    })
  }
}

// 发送文本消息，0为文本类型
const sendTextMessage = async () => {
  let response;
  var formData = new FormData()
  formData.append('receiver_id', receiver_id)
  formData.append('content', newMessage.value)
  formData.append('content_type', 0)
  await sendMessage(formData, newMessage.value, 0)
  newMessage.value = ''
}

// 文件处理逻辑
const fileInputRef = useTemplateRef('fileInput')
const sendFile = () => {
  fileInputRef.value.click()
}
const handleFile = (e) => {
  const file = e.target.files[0]
  if (!file) return
  sendFileMessage(file)
}
// 发送文件消息，2为文件类型
const sendFileMessage = async (file) => {
  var formData = new FormData()
  formData.append('receiver_id', receiver_id)
  formData.append('content', file)
  formData.append('content_type', 2)
  await sendMessage(formData, '[文件]', 2)
}
</script>

<template>
  <div class="message-input">
    <!-- 操作 -->
    <div class = "operation">
      <n-button text>图片</n-button>
      <n-button text @click="sendFile">文件</n-button>
      <input ref="fileInput" type="file" style="display:none" @change="handleFile"/>
      <n-button text>群投票</n-button>
    </div>
    <n-input v-model:value="newMessage" placeholder="请输入消息" />
    <n-button type="primary" @click="sendTextMessage">发送</n-button>
  </div>
</template>

<style scoped>

</style>