<script setup>
import { defineProps } from 'vue';
import { NImage } from 'naive-ui';


const {message} = defineProps({message: Object})
</script>

<template>
  <div class="message-content">
    <!--文本消息-->
    <div v-if="message.content_type===0" class="text-message">
      {{ message.content }}
    </div>
    <!--图片消息-->
    <div v-else-if="message.content_type===1">
      <n-image width="100" height="100" :src="message.content" lazy :intersection-observer-options="{
        root: '#chat-messages',
      }">
        <template #placeholder>
          <div class="image-placeholder">加载中...</div>
        </template>
      </n-image>
    </div>
  </div>
</template>

<style scoped>
.message-content {
  
} 
.text-message {
  padding: 0.5em 1em;
  border-radius: 4px;
  background-color: #f0f0f0;
  margin: 0.5em 0;
  max-width: 60%;
}
.image-placeholder {
  width: 100px;
  height: 100px;
  background-color: #f0f0f0;
}
</style>