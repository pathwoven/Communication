import apiClient from "./api"

// 搜索不是好友的用户
export const searchStranger = async (keyword, is_group) =>{
  try {
    const response = await apiClient.post(`/api/v1/contact/search/stranger/${is_group?"group":"single"}`, {keyword});
    if (response.status !== 200) {
      console.log('搜索用户失败:', response.data.message);
      return { success: false, data: null };
    } else {
      return { success: true, data: response.data.data };
    }
  } catch (error) {
    console.error('搜索用户失败:', error);
    return { success: false, data: null };
  }
}