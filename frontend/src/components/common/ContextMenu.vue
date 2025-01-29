<script setup>
import { onMounted, ref } from 'vue';
import { useOverlay } from '@/hooks/useOverlay';
// 父组件需要设置好object和options
// 还需要监听select事件
// object是右键菜单的对象
// options是右键菜单的选项
// select事件是选中右键菜单选项时触发的事件
defineEmits(['select']);

const { isActive, showOverlay, hideOverlay, overlayRef, menuPosition } = useOverlay(null);
const object = ref(null);
const options = ref([]);

const handleSelectOption = (option) => {
  emit('select', option, object);
  hideOverlay();
};

onMounted(() => {
  //overlayRef = ref(null);
  console.log(overlayRef);
});

</script>

<template>
  <div 
    v-show="isActive" 
    :style="{ top: `${menuPosition.y}px`, left: `${menuPosition.x}px` }" 
    class="context-menu" 
    ref="overlayRef"
  >
    <ul v-for="(option,index) in options" :key="index">
      <li @click="handleSelectOption('option')">{{option}}</li>
    </ul>
  </div>
</template>

<style scoped>
.context-menu {
  position: absolute;
  background-color: white;
  border: 1px solid #ccc;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  z-index: 1000;
}

.context-menu ul {
  list-style: none;
  margin: 0;
  padding: 0;
}

.context-menu li {
  padding: 8px 12px;
  cursor: pointer;
}

.context-menu li:hover {
  background-color: #f0f0f0;
}
</style>