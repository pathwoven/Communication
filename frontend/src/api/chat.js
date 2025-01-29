import apiClient from './api';

// 聊天列表----------------------------------------------------------------
export const getChatList = async () => {
  try {
    const response = await apiClient.get('/api/v1/chat/list');
    if (response.status !== 200) {
      console.log('获取聊天列表失败:', response.data.message);
      return { success: false, data: null };
    } else {
      return { success: true, data: response.data };
    }
  } catch (error) {
    console.error('获取聊天列表失败:', error);
    return { success: false, data: null };
  }
}
// 创建聊天
export const createChat = async (target_id) => {
  try {
    const response = await apiClient.post('/api/v1/chat/list/create', { target_id });
    if (response.status !== 200) {
      console.log('创建聊天失败:', response.data.message);
      return { success: false, data: null };
    } else {
      return { success: true ,data: response.data};
    }
  } catch (error) {
    console.error('创建聊天失败:', error);
    return { success: false, data: null };
  }
}
// 删除聊天
export const deleteChat = async (chat_id) => {
  try {
    const response = await apiClient.post(`/api/v1/chat/list/delete`, { chat_id });
    if (response.status !== 200) {
      console.log('删除聊天失败:', response.data.message);
      return { success: false };
    } else {
      return { success: true };
    }
  } catch (error) {
    console.error('删除聊天失败:', error);
    return { success: false };
  }
}
// 置顶或取消置顶聊天
export const pinChat = async (chat_id, is_pinned) => {
  try {
    const response = await apiClient.post(`/api/v1/chat/list/pin`, { chat_id, is_pinned });
    if (response.status !== 200) {
      console.log('置顶聊天失败:', response.data.message);
      return { success: false };
    } else {
      return { success: true };
    }
  } catch (error) {
    console.error('置顶聊天失败:', error);
    return { success: false };
  }
}
// 消息免打扰或取消
export const muteChat = async (chat_id, is_muted) => {
  try {
    const response = await apiClient.post(`/api/v1/chat/list/mute`, { chat_id, is_muted });
    if (response.status !== 200) {
      console.log('消息免打扰失败:', response.data.message);
      return { success: false };
    } else {
      return { success: true };
    }
  } catch (error) {
    console.error('消息免打扰失败:', error);
    return { success: false };
  }
}
// 屏蔽或取消屏蔽聊天
export const blockChat = async (chat_id, is_blocked) => {
  try {
    const response = await apiClient.post(`/api/v1/chat/list/block`, { chat_id, is_blocked });
    if (response.status !== 200) {
      console.log('屏蔽聊天失败:', response.data.message);
      return { success: false };
    } else {
      return { success: true };
    }
  } catch (error) {
    console.error('屏蔽聊天失败:', error);
    return { success: false };
  }
}
// 已读或未读消息
export const readChat = async (chat_id, is_read) => {
  try {
    const response = await apiClient.post(`/api/v1/chat/list/read`, { chat_id, is_read });
    if (response.status !== 200) {
      console.log('已读或未读消息失败:', response.data.message);
      return { success: false };
    } else {
      return { success: true };
    }
  } catch (error) {
    console.error('已读或未读消息失败:', error);
    return { success: false };
  }
}
// 为聊天添加标签
export const addTagToChat = async (chat_id, tag_name) => {
  try {
    const response = await apiClient.post(`/api/v1/chat/list/addTag`, { chat_id, tag_name });
    if (response.status !== 200) {
      console.log('为聊天添加标签失败:', response.data.message);
      return { success: false };
    } else {
      return { success: true };
    }
  } catch (error) {
    console.error('为聊天添加标签失败:', error);
    return { success: false };
  }
}
// 为聊天删除标签
export const removeTagFromChat = async (chat_id, tag_name) => {
  try {
    const response = await apiClient.post(`/api/v1/chat/list/removeTag`, { chat_id, tag_name });
    if (response.status !== 200) {
      console.log('为聊天删除标签失败:', response.data.message);
      return { success: false };
    } else {
      return { success: true };
    }
  } catch (error) {
    console.error('为聊天删除标签失败:', error);
    return { success: false };
  }
}

// 标签----------------------------------------------------------------
// 获取标签列表
export const getTags = async () => {
	try {
		const response = await apiClient.get('/api/v1/chat/tag/list');
		if (response.status !== 200) {
		console.log('获取标签列表失败:', response.data.message);
		return { success: false, data: null };
		} else {
		return { success: true, data: response.data };
		}
	} catch (error) {
		console.error('获取标签列表失败:', error);
		return { success: false, data: null };
	}
}
// 创建标签
export const createTag = async (tag_name) => {
  try {
    const response = await apiClient.post('/api/v1/chat/tag/create', { tag_name});
    if (response.status !== 200) {
    console.log('创建标签失败:', response.data.message);
    return { success: false};
    } else {
    return { success: true };
    }
  } catch (error) {
    console.error('创建标签失败:', error);
    return { success: false };
  }
}
// 删除标签
export const deleteTag = async (tag_name) => {
  try {
    const response = await apiClient.post(`/api/v1/chat/tag/delete`, {tag_name});
    if (response.status !== 200) {
    console.log('删除标签失败:', response.data.message);
    return { success: false };
    } else {
    return { success: true };
    }
  } catch (error) {
    console.error('删除标签失败:', error);
    return { success: false };
  }
}
// 重命名标签
export const renameTag = async (old_tag_name, new_tag_name) => {
  try {
    const response = await apiClient.post(`/api/v1/chat/tag/rename`, {old_tag_name, new_tag_name });
    if (response.status !== 200) {
    console.log('重命名标签失败:', response.data.message);
    return { success: false };
    } else {
    return { success: true };
    }
  } catch (error) {
    console.error('重命名标签失败:', error);
    return { success: false };
  }
}