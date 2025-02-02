<script setup>
import { ref } from 'vue';

const type = ref('info');
const message = ref('');
const duration = ref(2000);

const visible = ref(false)
const show = () => {
  visible.value = true;
  setTimeout(() => {
    visible.value = false;
  }, duration.value);
};

defineExpose({show, type, message, duration})
</script>

<template>
  <div v-if="visible" class="notification" :class="type">
    {{ message }}
  </div>
</template>

<style scoped>
.notification {
  position: fixed;
  top: 20px;
  left: 50%;
  transform: translateX(-50%);
  padding: 10px 20px;
  background-color: #333;
  color: #fff;
  border-radius: 5px;
  z-index: 1000;
  opacity: 1;
  transition: opacity 0.5s ease-in-out;
}
.notification.info {
  background-color: #2196f3;
}
.notification.success {
  background-color: #4caf50;
}
.notification.error {
  background-color: #f44336;
}
.notification.fade-out {
  opacity: 0;
}
</style>