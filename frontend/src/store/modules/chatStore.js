import { defineStore } from "pinia";
import * as chatApi from "@/api/chat";

export const useChatStore = defineStore("chat", { 
	state: () => ({
		selectedChat: null,
		chatList: [],
		moreTags: [],
	}), 
	actions: {
		async fetchChatList() {
			const response = await chatApi.getChatList();
			if (response.success) {
				this.chatList = response.data.data;
			}else{
				console.error('获取聊天列表失败');
			}
		},
		selectChat(chat) {
			this.selectedChat = chat;
		},
		addChat(newChat) {
			this.chatList.push(newChat);
		},
		removeChat(chatId) {
			this.chatList = this.chatList.filter((chat) => chat.id !== chatId);
		},
		updateChat(chatId, newChat) {
			const index = this.chatList.findIndex((chat) => chat.id === chatId);
			if (index !== -1) {
				this.chatList.splice(index, 1, newChat);
			}
		},


		async fetchTags() {
			const response = await chatApi.getTags();
			if (response.success) {
				this.moreTags = response.data.data.tags;
			}else{
				console.error('获取标签列表失败');
			}
		},
		addMoreTag(newTag) {
			this.moreTags.push(newTag);
		},
		deleteMoreTag(oldTag) {
			this.moreTags = this.moreTags.filter(tag => tag !== oldTag);
		},
	},
});