<script setup>
import { ref, computed, onBeforeMount } from 'vue'
import {useChatStore} from '@/store/modules/chatStore'
import ContextMenu from '@/components/common/ContextMenu.vue'
import ChatTag from './ChatTag.vue'
import * as chatApi from '@/api/chat'
import eventBus from '@/utils/eventBus'
import { NIcon, NDropdown, NTag, NAvatar, NBadge} from 'naive-ui'
import {Search, AddCircle} from '@vicons/ionicons5';

const searchQuery = ref('')
const chatStore = useChatStore()

// 标签的逻辑
const tags = ref(['好友', '群聊', '未读', '全部', '屏蔽'])
const selectedTag = ref('全部')
// 选择标签
const selectTag = (tag) => {
	selectedTag.value = tag
  if(tag !== '好友' && tag !== '群聊' && tag !== '未读' && tag !== '全部' && tag !== '屏蔽') {
    // 显示更多标签
    selectedTag.value = '更多'
  }
}
// 更多标签
const moreTags = computed(() => {
  if(!chatStore.moreTags) {
    return []
  }
  return chatStore.moreTags.map(tag => {
    return {label: tag, key: tag}
  })
})

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
		} else if(selectedTag.value === '屏蔽') {
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

const selectChat = async (chat) => {
  // 选择聊天逻辑
  chatStore.selectChat(chat)
  const response = await chatApi.readChat(chat.id, true)
  if(response.success) {
    chat.unread_count = 0
  } else {
    console.log('标记为已读失败')
  }
}

// 右键菜单逻辑
const contextMenu = ref({
  x: 0,
  y: 0,
  show: false,
  chat: null,
  options: []
})
// 显示右键菜单
const showContextMenu = (event, chat) => {
  // 保存聊天对象
  contextMenu.value.chat = chat   // note:ref是浅层的，.chat不需要.value 
  // 设置菜单选项
  const options = [];
  if(chat.is_pinned) {
    options.push({label:'取消置顶', key:'取消置顶'})
  } else {
    options.push({label:'置顶', key:'置顶'})
  }
  if(chat.is_mute) {
    options.push({label: '取消免打扰', key: '取消免打扰'})
  } else {
    options.push({label: '免打扰', key: '免打扰'})
  }
  if(chat.is_blocked) {
    options.push({label: '取消屏蔽', key: '取消屏蔽'})
  } else {
    options.push({label: '屏蔽', key: '屏蔽'})
  }
  if(chat.unread_count > 0) {
    options.push({label: '标记为已读', key: '标记为已读'})
  } else {
    options.push({label: '标记为未读', key: '标记为未读'})
  }
  options.push({label: '删除聊天', key: '删除聊天'})
  options.push({label: '添加标签', key: '添加标签'})
  options.push({label: '移除标签', key: '移除标签'})
  contextMenu.value.options = options
  // 设置菜单位置
  contextMenu.value.x = event.clientX
  contextMenu.value.y = event.clientY
  // 显示菜单
  contextMenu.value.show = true
}
// 处理菜单选项
const handleSelectOption = async (option) => { 
  if(option==='置顶'){
    const response = await chatApi.pinChat(contextMenu.value.chat.id, true)
    if(response.success) {
      contextMenu.value.chat.is_pinned = true
    }else{
      eventBus.emit('notify', {message: '置顶失败，请重试', type:'error'})
    }
  }
  else if(option==='取消置顶'){
    const response = await chatApi.pinChat(contextMenu.value.chat.id, false)
    if(response.success) {
      contextMenu.value.chat.is_pinned = false
    }else{
      eventBus.emit('notify', {message: '取消置顶失败，请重试', type:'error'})
    }
  }
  else if(option === '免打扰') {
    const response = await chatApi.muteChat(contextMenu.value.chat.id, true)
    if(response.success) {
      contextMenu.value.chat.is_mute = true
    } else {
      eventBus.emit('notify', {message: '设置免打扰失败，请重试', type:'error'})
    }
  } 
  else if(option === '取消免打扰') {
    const response = await chatApi.muteChat(contextMenu.value.chat.id, false)
    if(response.success) {
      contextMenu.value.chat.is_mute = false
    } else {
      eventBus.emit('notify', {message: '取消免打扰失败，请重试', type:'error'})
    }
  } 
  else if(option === '屏蔽') {
    const response = await chatApi.blockChat(contextMenu.value.chat.id, true)
    if(response.success) {
      contextMenu.value.chat.is_blocked = true
    } else {
      eventBus.emit('notify', {message: '屏蔽失败，请重试', type:'error'})
    }
  } 
  else if(option === '取消屏蔽') {
    const response = await chatApi.blockChat(contextMenu.value.chat.id, false)
    if(response.success) {
      contextMenu.value.chat.is_blocked = false
    } else {
      eventBus.emit('notify', {message: '取消屏蔽失败，请重试', type:'error'})
    }
  } 
  else if(option === '标记为已读') {
    const response = await chatApi.readChat(contextMenu.value.chat.id, true)
    if(response.success) {
      contextMenu.value.chat.unread_count = 0
    } else {
      eventBus.emit('notify', {message: '标记为已读失败，请重试', type:'error'})
    }
  } else if(option === '标记为未读') {
    const response = await chatApi.readChat(contextMenu.value.chat.id, false)
    if(response.success) {
      contextMenu.value.chat.unread_count = 1
    } else {
      eventBus.emit('notify', {message: '标记为未读失败，请重试', type:'error'})
    }
  }
  contextMenu.value.show = false
}
// 处理菜单点击到外面
const onMenuClickedOutside = () => {
  contextMenu.value.show = false
}
// 添加菜单的选项
const handleAddMenu = (option)=> {
  if(option === '添加好友') {
    // 添加好友逻辑
  } else if(option === '添加群聊') {
    // 添加群聊逻辑
  } else if(option === '创建群聊') {
    // 创建群聊逻辑
  }
}

onBeforeMount(() => {
	// 检查chatStore中是否有聊天数据
	if(!useChatStore().chatList || useChatStore().chatList.length === 0) {
		// 从服务器获取聊天数据
		useChatStore().fetchChatList()
	}
	// 检查chatStore中是否有标签数据
	if(!useChatStore().moreTags || useChatStore().moreTags.length === 0) {
		// 从服务器获取标签数据
		useChatStore().fetchTags()
	}
})
</script>

<template>
  <!--聊天框-->
  <div class="chat-list-container">
    <div class="search-bar">
      <n-input 
        type="text" 
        v-model:value="searchQuery" 
        placeholder="搜索聊天" 
      >
        <template #prefix>
          <n-icon  :component="Search"/>
        </template>
      </n-input>
      <n-dropdown 
        trigger="click" 
        :options="[
          {label: '添加好友', key: '添加好友'},
          {label: '添加群聊', key: '添加群聊'},
          {label: '创建群聊', key: '创建群聊'},
        ]" 
        @select="handleAddMenu"
      >
          <n-icon :component="AddCircle"/>
      </n-dropdown>
    </div>
		<!--标签-->
    <div class="tags">
      <n-tag
        v-for="tag in tags"
        :key="tag"
        :checked="tag === selectedTag"
        @click="selectTag(tag)"
        checkable
        size="small"
      >
        {{ tag }}
      </n-tag>
      <n-dropdown trigger="click" :options="moreTags" @select="selectTag">
        <n-tag 
          @click="selectTag('更多')"
          :checked="selectedTag === '更多'"
          checkable
          size="small"
        > 
          更多
        </n-tag>
      </n-dropdown>
    </div>
		
		<!--聊天列表-->
    <ul class="chat-list" v-for="chat in filteredChats">
      <li 
        :key="chat.id" 
        @click="selectChat(chat)"
        @contextmenu.prevent="showContextMenu($event, chat)"
        :class="{selected: chatStore.selectedChat && chatStore.selectedChat.id === chat.id}"
      >
        <!--头像-->
        <n-badge dot :color="chat.online_status? 'green':'grey'">
          <n-avatar 
            round
            :src="chat.avatar"
            alt="头像"
            size="medium"
          />
        </n-badge>
        <div class="chat-info">
          <div class="chat-header">
            <span class="chat-name">{{ chat.remark? chat.remark : chat.name }}</span>
            <span class="chat-time">{{ chat.last_time }}</span>
          </div>
          <div class="chat-content">
            <span class="chat-message">{{ (chat.last_person?chat.last_person+":":"" )+ chat.last_message }}</span>
            <span class="chat-unread" v-if="chat.unread_count">{{ chat.unread_count }}</span>
          </div>
        </div>
      </li>
    </ul>
    <!--菜单-->
    <n-dropdown 
      placement="bottom-start"
      trigger="manual"
      :x="contextMenu.x"
      :y="contextMenu.y"
      :options="contextMenu.options"
      :show="contextMenu.show"
      @select="handleSelectOption"
      :on-clickoutside="onMenuClickedOutside"
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
.chat-list li.selected{
  background-color: #d9d7d7;
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
