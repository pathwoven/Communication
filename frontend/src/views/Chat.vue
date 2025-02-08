<script setup>
import { ref } from 'vue'
import ChatList from '@/components/Chat/ChatList.vue'
import ChatBox from '@/components/Chat/ChatBox.vue'
import { NSplit, NLayout, NLayoutContent} from 'naive-ui';
import MainSider from '@/components/common/MainSider.vue';
import { useChatStore } from '@/store/modules/chatStore';

const chatStore = useChatStore()
</script>

<template>
  <div class="chat-container">
    
    <n-layout has-sider>
      <MainSider />
      <n-layout-content>
        <n-split 
          direction="horizontal" 
          style="height: 100vh;" 
          :default-size="0.3"
          :max="0.6" 
          :min="0.2"
        >
          <template #1>
            <ChatList />
          </template>
          <template #2>
            <ChatBox v-if="chatStore.selectedChat"/>
          </template>
        </n-split>
      </n-layout-content>
      
    </n-layout>
  </div>
</template>

<style scoped>
.chat-container {
  display: flex;
  height: 100%;
}

/* .chat-list {
  border-right: 1px solid #ccc;
  height: 100%;
}

.chat-box {
  flex-grow: 1;
    height: 100%;
} */

.resizer {
  width: 5px;
  cursor: col-resize;
  background-color: #ccc;
	height: 100%;
}
</style>
