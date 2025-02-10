<script setup>
import { ref, onMounted, watch, onBeforeUnmount, onBeforeMount, useTemplateRef, nextTick } from 'vue'
import { NLayout, NLayoutSider, NLayoutContent, NLayoutHeader, NLayoutFooter, NAffix, NTag } from 'naive-ui'
import {useSettingStore} from '@/store/modules/setting';
import {useChatStore} from '@/store/modules/chatStore';
import eventBus from '@/utils/eventBus';
import * as chatApi from '@/api/chat';

import MessageContent from './MessageContent.vue';
import MessageInput from './MessageInput.vue';

const settingStore = useSettingStore();
const chatStore = useChatStore();


const messages = ref(null)
// 获取消息列表
const fetchMessages = async () => {
  const response = await chatApi.getMessages(chatStore.selectedChat.target_id, 0, 50)
  if (response.success) {
    messages.value = response.data

  } else {
    console.log('获取消息列表失败')
    // 提醒 todo
  }
  nextTick(() => {
    scrollToBottom(true)
  })
}
// 接收新消息
eventBus.on('new-message', (data) => {
  if(chatStore.selectedChat &&
      (data.sender_id === chatStore.selectedChat.target_id||data.receiver_id === chatStore.selectedChat.target_id))
  {
    messages.value.push(data)
    scrollToBottom()
  }
})
// 监听聊天人
watch(() => chatStore.selectedChat, (newVal) => {
  if(newVal){
    fetchMessages()
  }
})

const operation = ref(null)
const op_message_id = ref(null)

const sendMessage = (data) => {
  const message = {
    id: data.id,
    avatar: settingStore.user.avatar,
    name: settingStore.user.nickname,
    create_time: data.create_time,
    content: data.content,
    content_type: data.content_type,
    sender_id: settingStore.user.id,
    receiver_id: chatStore.selectedChat.target_id,
    operation: operation.value,
    op_message_id: op_message_id.value,
  }
  messages.value.push(message)
  // 更新chatList
  let lastMessage;
  if(message.content_type === 1){
    lastMessage = '[图片]'
  }else if(message.content_type === 2){
    lastMessage = '[文件]'
  }else if(message.content_type === 3){
    lastMessage = '[群投票]'
  }else if(message.content_type === 4){
    lastMessage = '[聊天记录]'
  }else{
    lastMessage = message.content
  }
  chatStore.updataChatMessage(message.receiver_id, lastMessage, '', message.create_time)
  console.log(message)
  scrollToBottom()
}





// 布局
const collapsed = ref(true)
const contentRef = useTemplateRef('contentRef')
const scrollToBottom = (immediate=false) => {
  if(contentRef.value && contentRef.value.scrollableElRef){
    const scrollableEl = contentRef.value.scrollableElRef
    scrollableEl.scrollTo({
      top: scrollableEl.scrollHeight,
      behavior: immediate? 'instant':'smooth',
    })
  }
}
const loadMoreMessages = () => {
  // 加载更多消息逻辑
}

onMounted(() => {
  // 获取消息列表
  if(chatStore.selectedChat === null){
    return
  }
  fetchMessages()
})
onBeforeUnmount(() => {
  eventBus.off('new-message')
})
</script>

<template>
  <div class="chat-box-container">
    <n-layout  has-sider sider-placement="right">
      <n-layout >
        <!--头部-->
        <n-layout-header class="chat-header">
          <div >
            <img :src="chatStore.selectedChat.avatar" alt="头像" class="avatar" />
            <div class="chat-info">
              <span class="chat-name">{{ chatStore.selectedChat.name }}</span>
              <span class="chat-status" :class="{ online: chatStore.selectedChat.online_status }"></span>
            </div>
            <n-button
              size="small"
              type="text"
              @click="collapsed = !collapsed"
            >
              ···
            </n-button>
          </div>
        </n-layout-header> 
        <!-- 展示消息列表-->
        <n-layout-content 
          position="absolute" 
          native-scrollbar="false" 
          class="chat-content" 
          @scroll="loadMoreMessages"
          ref="contentRef"
        >
          <div class="chat-messages">
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
          <!--滚动至底部-->
          <div class="scroll-to-bottom" @click="scrollToBottom">
            <n-tag type="success">滚动至底部</n-tag>
          </div>
        </n-layout-content>
        <!--底部-->
        <n-layout-footer class="chat-footer" position="absolute">
          <!--输入区域-->
          <MessageInput
            :isGroup="chatStore.selectedChat.isGroup"
            :operation="operation"
            :op_message_id="op_message_id"
            :receiver_id="chatStore.selectedChat.target_id"
            @send-message="sendMessage"
          />
        </n-layout-footer>
      </n-layout>
      <!--侧边栏-->
      <n-layout-sider
        collapse-mode="width"
        :width="200"
        collapsed-width="0"
        :collapsed="collapsed"
      >
        <div >
          侧边栏
        </div>
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
  height: 100px;
}
.chat-footer{
  height: 100px;
  bottom: 0;
}
.chat-content{
  top:100px;
  bottom:100px;
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
.scroll-to-bottom {
  position: absolute;
  bottom: 20px;
  right: 20px;
}
.chat-messages {
 
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
}

.chat-input input {
  padding: 5px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.chat-input button {
  margin-left: 10px;
  padding: 5px 10px;
}
</style>
