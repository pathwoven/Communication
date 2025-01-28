import apiClient from './api';

export const login = async (displayId, password) => {
  try {
    const response = await apiClient.post('/auth/login', { displayId, password });
    if (response.status !== 200) {
      console.log('登录失败:', response.data.message);
      return { success: false, data: null };
    }else{
        return { success: true, data: response.data };  
    }
  } catch (error) {
    console.error('登录失败:', error);
    return { success: false, data: null };
  }
};

export const register = async (name, password, email, displayId, sex, birthday, signature) => {
  try{
    const response = await apiClient.post('/auth/register',{name, password,email,displayId,sex,birthday,signature});
    if(response.status !== 200){
      console.log('注册失败:', response.data.message);
      return {success: false, data: null};
    }else{
      return {success: true};
    }
  }catch (error){
    console.error('注册失败:', error);
    return {success: false, data: null};
  }
};