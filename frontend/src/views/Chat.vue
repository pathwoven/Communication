<script setup>
import { ref } from 'vue'
import ChatList from '@/components/Chat/ChatList.vue'
import ChatBox from '@/components/Chat/ChatBox.vue'
import { useResizable } from '@/hooks/useResizable'

const {
  chatListWidth,
  chatBoxWidth,
  startDragging,
  stopDragging,
  handleDragging,
} = useResizable();

</script>

<template>
  <div class="chat-container" @mousemove="handleDragging" @mouseup="stopDragging">
    <ChatList :style="{ width: chatListWidth }" />
    <div class="resizer" @mousedown="startDragging"></div>
    <ChatBox :style="{ width: chatBoxWidth }" />
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
