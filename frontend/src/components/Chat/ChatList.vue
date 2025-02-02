<script setup>
import { ref, computed, onBeforeMount } from 'vue'
import {useChatStore} from '@/store/modules/chatStore'
import ContextMenu from '@/components/common/ContextMenu.vue'
import ChatTag from './ChatTag.vue'
import * as chatApi from '@/api/chat'
import eventBus from '@/utils/eventBus'

const searchQuery = ref('')
const chatStore = useChatStore()

// 标签的逻辑
const tags = ref(['好友', '群聊', '未读', '全部', '屏蔽'])
const selectedTag = ref('全部')
// 选择标签
const selectTag = (tag) => {
	selectedTag.value = tag
}

// const chats = ref([
//   // 示例聊天
// 	{
//     id: 1,
//     target_id: 101,
//     is_group: false,
//     avatar: 'avatar1.png',
//     display_id: 'user101',
//     name: '张三',
//     remark: '好友',
//     unread_count: 2,
//     last_message: '你好',
//     last_time: '10:00',
//     online_status: true,
//     status: 'active',
//     is_pinned: false,
//     is_mute: false,
//     is_blocked: false,
//     tag1: '好友',
//     tag2: '常联系',
//     tag3: null,
//   },
//   {
//     id: 2,
//     target_id: 202,
//     is_group: true,
//     avatar: 'avatar2.png',
//     display_id: 'group202',
//     name: '群聊202',
//     remark: '工作群',
//     unread_count: 0,
//     last_message: '在吗',
//     last_time: '11:00',
//     online_status: false,
//     status: 'inactive',
//     is_pinned: true,
//     is_mute: true,
//     is_blocked: false,
//     tag1: '群聊',
//     tag2: '工作',
//     tag3: null,
//   },
//   // ...更多聊天数据
// ])

const filteredChats = computed(() => {
  if(!chatStore.chatList) {
    return []
  }
	// 根据标签筛选
	if(searchQuery.value === '') {
		if(selectedTag.value === '全部') {
			return chatStore.chatList.filter(chat => !chat.is_blocked)
		} else if(selectedTag.value === '未读') {
			return chatStore.chatList.filter(chat => chat.unread_count > 0 && !chat.is_blocked)
		} else if(selectTag.value === '屏蔽') {
			return chatStore.chatList.filter(chat => chat.is_blocked)
		}else if(selectedTag.value === '好友') {
			return chatStore.chatList.filter(chat => !chat.is_group && !chat.is_blocked)
		}else if(selectedTag.value === '群聊') {
			return chatStore.chatList.filter(chat => chat.is_group && !chat.is_blocked)
		}else {
			return chatStore.chatList.filter(chat => chat.tag1===selectedTag.value||chat.tag2===selectedTag.value||chat.tag3===selectedTag.value)
		}
	}
	// 根据搜索关键字筛选
  return chats.value.filter(chat => {
    return chat.name.includes(searchQuery.value) || chat.content.includes(searchQuery.value)
  })
})

const createContact = () => {
  // 添加聊天逻辑
}

const selectChat = (chat) => {
  // 选择聊天逻辑
}

// 右键菜单逻辑
const contextMenu = ref(null)
// 显示右键菜单
const showContextMenu = (event, chat) => {
  // 保存聊天对象
  contextMenu.value.object = chat   // note:ref是浅层的，.object不需要.value 
  // 设置菜单选项
  const options = [];
  if(chat.is_pinned) {
    options.push('取消置顶')
  } else {
    options.push('置顶')
  }
  if(chat.is_mute) {
    options.push('取消免打扰')
  } else {
    options.push('免打扰')
  }
  if(chat.is_blocked) {
    options.push('取消屏蔽')
  } else {
    options.push('屏蔽')
  }
  if(chat.unread_count > 0) {
    options.push('标记为已读')
  }else{
    options.push('标记为未读')
  }
  options.push('删除聊天')
  options.push('添加标签')
  options.push('移除标签')
  contextMenu.value.options = options
  // 显示菜单
  contextMenu.value.showOverlay(event)
}
// 处理菜单选项
const handleSelectOption = async (option, object) => {
  const chat = object;   
  if(option==='置顶'){
    const response = await chatApi.pinChat(chat.id, true)
    if(response.success) {
      chat.is_pinned = true
    }else{
      eventBus.emit('notify', {message: '置顶失败，请重试', type:'error'})
    }
  }
}

onBeforeMount(() => {
	// 检查chatStore中是否有聊天数据
	if(useChatStore().chatList.length === 0) {
		// 从服务器获取聊天数据
		useChatStore().fetchChatList()
	}
	// 检查chatStore中是否有标签数据
	if(useChatStore().moreTags.length === 0) {
		// 从服务器获取标签数据
		useChatStore().fetchTags()
	}
})
</script>

<template>
  <div class="chat-list-container">
		<!--聊天框-->
    <div class="search-bar">
      <input type="text" v-model="searchQuery" placeholder="搜索聊天" />
      <button @click="createContact">+</button>
    </div>
		<!--标签-->
		<ChatTag 
			:tags="tags" 
			:selectedTag="selectedTag" 
			@update:selectedTag="selectTag" 
		/>
		<!--聊天列表-->
    <ul class="chat-list" v-for="chat in filteredChats">
      <li 
        :key="chat.id" 
        @click="selectChat(chat)"
        @contextmenu.prevent="showContextMenu($event, chat)"
      >
        <img :src="chat.avatar" alt="头像" class="avatar" />
        <div class="chat-info">
          <div class="chat-header">
            <span class="chat-name">{{ chat.remark? chat.remark : chat.name }}</span>
            <span class="chat-time">{{ chat.last_time }}</span>
          </div>
          <div class="chat-content">
            <span class="chat-message">{{ chat.last_message }}</span>
            <span class="chat-unread" v-if="chat.unread_count">{{ chat.unread_count }}</span>
          </div>
        </div>
        <div class="chat-status" :class="{ online: chat.online_status }"></div>
      </li>
    </ul>
    <!--菜单-->
    <ContextMenu 
      ref="contextMenu"
      @select="handleSelectOption" 
    />
  </div>
</template>

<style scoped>
.chat-list-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.search-bar {
  display: flex;
  justify-content: space-between;
  padding: 10px;
  border-bottom: 1px solid #ccc;
}

.search-bar input {
  flex-grow: 1;
  padding: 5px;
}

.search-bar button {
  margin-left: 10px;
  padding: 5px 10px;
}

.tags {
  display: flex;
  justify-content: space-around;
  padding: 10px;
  border-bottom: 1px solid #ccc;
}

.chat-list {
  flex-grow: 1;
  overflow-y: auto;
  list-style: none;
  padding: 0;
  margin: 0;
}

.chat-list li {
  display: flex;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #ccc;
  cursor: pointer;
}

.chat-list li:hover {
  background-color: #f5f5f5;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 10px;
}

.chat-info {
  flex-grow: 1;
}

.chat-header {
  display: flex;
  justify-content: space-between;
}

.chat-name {
  font-weight: bold;
}

.chat-time {
  color: #999;
}

.chat-content {
  display: flex;
  justify-content: space-between;
}

.chat-message {
  color: #666;
}

.chat-unread {
  background-color: red;
  color: white;
  border-radius: 50%;
  padding: 2px 5px;
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
</style>
