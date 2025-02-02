<script setup>
import { onMounted, onUnmounted, ref } from 'vue';
import Notification from '@/components/common/Notification.vue';
import eventBus from '@/utils/eventBus';


// 通知
const notification = ref(null)
const notify = (message, type='info', duration=2000) => {
	if (notification.value) {
		notification.value.type = type
		notification.value.message = message
		notification.value.duration = duration
		notification.value.show()
	}
}

onMounted(() => {
	eventBus.on('notify', (data) => {
		notify(data.message, data.type, data.duration);
	})
})

onUnmounted(() => {
	eventBus.off('notify');
})
</script>

<template>
	<div id="app">
		<Notification ref="notification" />
		<router-view></router-view>
	</div>
</template>

<style src="@/assets/styles/main.css">
#app {
	height: 100%;
	width: 100%;
	background-color: #f0f2f5;
}
</style>
