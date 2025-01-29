<template>
  <div class="chat-box-container">
    <div class="chat-header">
      <img :src="chat.avatar" alt="头像" class="avatar" />
      <div class="chat-info">
        <span class="chat-name">{{ chat.name }}</span>
        <span class="chat-status" :class="{ online: chat.online }"></span>
      </div>
    </div>
    <div class="chat-messages" @scroll="loadMoreMessages">
      <div v-for="message in messages" :key="message.id" class="message">
        <img :src="message.avatar" alt="头像" class="message-avatar" />
        <div class="message-content">
          <div class="message-header">
            <span class="message-name">{{ message.name }}</span>
            <span class="message-time">{{ message.time }}</span>
          </div>
          <div class="message-body">
            <p>{{ message.content }}</p>
          </div>
        </div>
      </div>
    </div>
    <div class="chat-input">
      <input v-model="newMessage" type="text" placeholder="输入消息..." @keyup.enter="sendMessage" />
      <button @click="sendMessage">发送</button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const chat = ref({
  avatar: 'avatar.png',
  name: '张三',
  online: true,
})

const messages = ref([
  // 示例消息数据
  { id: 1, avatar: 'avatar1.png', name: '张三', time: '10:00', content: '你好' },
  { id: 2, avatar: 'avatar2.png', name: '李四', time: '10:01', content: '你好' },
  // ...更多消息数据
])

const newMessage = ref('')

const loadMoreMessages = () => {
  // 加载更多消息逻辑
}

const sendMessage = () => {
  if (newMessage.value.trim() !== '') {
    messages.value.push({
      id: messages.value.length + 1,
      avatar: 'avatar.png',
      name: '我',
      time: new Date().toLocaleTimeString(),
      content: newMessage.value,
    })
    newMessage.value = ''
    scrollToBottom()
  }
}

const scrollToBottom = () => {
  const chatMessages = document.querySelector('.chat-messages')
  chatMessages.scrollTop = chatMessages.scrollHeight
}

onMounted(() => {
  scrollToBottom()
})
</script>

<style scoped>
.chat-box-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.chat-header {
  display: flex;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #ccc;
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
  flex-grow: 1;
  overflow-y: auto;
  padding: 10px;
}

.message {
  display: flex;
  margin-bottom: 10px;
}

.message-avatar {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  margin-right: 10px;
}

.message-content {
  flex-grow: 1;
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
