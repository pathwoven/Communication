
<script setup>
import { NAvatar } from 'naive-ui';
</script>

<template>
  <div class="add-friend">
    <n-input placeholder="请输入" />
    <ul>
      <li >
        <n-avatar round size="small" ></n-avatar>
        <div class="info">
          <span>用户名</span>
          <span>用户id</span>
        </div>
        <n-button type="primary">添加</n-button>
      </li>
    </ul>
  </div>
</template>

<style scoped>

</style>