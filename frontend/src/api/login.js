import apiClient from './apiClient';

export const login = async (displayId, password) => {
  try {
    const response = await apiClient.post('/login', { displayId, password });
    if (response.status !== 200) {
      console.log('登录失败:', response.data.message);
      return { success: false, data: null };
    }
  } catch (error) {
    console.error('登录失败:', error);
    return { success: false, data: null };
  }
};
