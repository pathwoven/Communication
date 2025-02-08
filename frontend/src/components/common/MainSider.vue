<script setup>
import { ref, h } from 'vue'
import { NLayout, NLayoutSider, NMenu, NIcon } from 'naive-ui';
import { PersonOutline } from '@vicons/ionicons5';

// 边栏
const collapsed = ref(false)

function renderIcon(icon){
  return () => h(NIcon, null, { default: () => h(icon) });
}
const menuOptions = ref([{
  key: 'chat',
  label: '聊天',
  icon: renderIcon(PersonOutline),
}])
</script>

<template>
 
    <n-layout-sider
      bordered
      collapse-mode="width"
      :collapsed-width="48"
      :width="150"
      :collapsed="collapsed"
      show-trigger
      @collapse="collapsed = true"
      @expand="collapsed = false"
    >
      <img src="" alt="logo" style="width: 100%; height: 100px;"/>
      <n-menu 
        
        :collapsed="collapsed"
        :collapsed-width="48"
        :collapsed-icon-size="20"
        :options="menuOptions"
        
      />
    </n-layout-sider>
  
</template>