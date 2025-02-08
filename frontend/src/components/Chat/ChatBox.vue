<script setup>
import { ref, onMounted, watch } from 'vue'
import { NLayout, NLayoutSider, NLayoutContent, NLayoutHeader, NLayoutFooter } from 'naive-ui'
import {useSettingStore} from '@/store/modules/setting';
import {useChatStore} from '@/store/modules/chatStore';
import * as chatApi from '@/api/chat';

import MessageContent from './MessageContent.vue';
import MessageInput from './MessageInput.vue';

const settingStore = useSettingStore();
const chatStore = useChatStore();


const messages = ref(null)
const fetchMessages = async () => {
  const response = await chatApi.getMessages(chatStore.selectedChat.target_id, 0, 50)
  if (response.success) {
    messages.value = response.data
  } else {
    console.log('获取消息列表失败')
    // 提醒 todo
  }
}
// 监听聊天人
watch(() => chatStore.selectedChat, (newVal) => {
  if(newVal){
    fetchMessages()
  }
})

const operation = ref(null)
const op_message_id = ref(null)
const receiver_id = ref('')

const sendMessage = (data) => {
  const message = {
    id: data.id,
    avatar: settingStore.user.avatar,
    name: settingStore.user.nickname,
    time: data.create_time,
    content: data.content,
    content_type: data.content_type,
    sender_id: settingStore.user.id,
    receiver_id: receiver_id.value,
    operation: operation.value,
    op_message_id: op_message_id.value,
  }
  messages.value.push(message)
  scrollToBottom()
}

const loadMoreMessages = () => {
  // 加载更多消息逻辑
}

const scrollToBottom = () => {
  const chatMessages = document.querySelector('.chat-messages')
  chatMessages.scrollTop = chatMessages.scrollHeight
}

onMounted(() => {
  scrollToBottom()
  // 获取消息列表
  if(chatStore.selectedChat === null){
    return
  }
  fetchMessages()
})
</script>

<template>
  <div class="chat-box-container">
    <n-layout  has-sider>
      <n-layout-content >
        <div style="display: flex; flex-direction: column;height: 100vh;">
          <div class="chat-header">
            <img :src="chatStore.selectedChat.avatar" alt="头像" class="avatar" />
            <div class="chat-info">
              <span class="chat-name">{{ chatStore.selectedChat.name }}</span>
              <span class="chat-status" :class="{ online: chatStore.selectedChat.online }"></span>
            </div>
          </div>
          <!-- 展示消息列表-->
          <div class="chat-messages" @scroll="loadMoreMessages" id="chat-messages">
            <!--消息列表-->
            <div v-for="message in messages" :key="message.id" class="message-item">
              <!--我发出的消息-->
              <div v-if="message.sender_id===settingStore.user.id" class="my-message">
                <MessageContent :message="message" />
                <img :src="message.avatar" alt="头像" class="message-avatar" />
              </div>
              <!--别人发出的消息-->
              <div v-else class="other-message">
                <img :src="message.avatar" alt="头像" class="message-avatar" />
                <MessageContent :message="message" />
              </div>
            </div>
          </div>
          <!--输入区域-->
          <MessageInput
            :isGroup="chatStore.selectedChat.isGroup"
            :operation="operation"
            :op_message_id="op_message_id"
            :receiver_id="chatStore.selectedChat.target_id"
            @send-message="sendMessage"
          />
        </div>
        
      </n-layout-content>
      
      <!--侧边栏-->
      <n-layout-sider
        collapsed=false
      >
      </n-layout-sider>
    </n-layout>
    
  </div>
</template>


<style scoped>
.chat-box-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.chat-header {
  display: flex;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #ccc;
  flex-grow: 1;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 10px;
}

.chat-info {
  display: flex;
  flex-direction: column;
}
.chat-name {
  font-weight: bold;
}
.chat-status {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background-color: gray;
}
.chat-status.online {
  background-color: green;
}
.chat-messages {
  flex-grow: 19;
  overflow-y: auto;
  padding: 10px;
  display: flex;
  flex-direction: column;
}
.message-item {
  width: 100%;
  margin-bottom: 10px;
  
}
.message-avatar {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  margin-right: 10px;
}
.my-message {
  display: flex;
  justify-content: flex-end;
}
.other-message {
  display: flex;
  justify-content: flex-start;
}

.message-header {
  display: flex;
  justify-content: space-between;
}

.message-name {
  font-weight: bold;
}

.message-time {
  color: #999;
}

.message-body {
  margin-top: 5px;
}

.chat-input {
  display: flex;
  padding: 10px;
  border-top: 1px solid #ccc;
  flex-grow: 2;
}

.chat-input input {
  flex-grow: 1;
  padding: 5px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.chat-input button {
  margin-left: 10px;
  padding: 5px 10px;
}
</style>
