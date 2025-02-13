import axios from 'axios';

export const baseURL = 'http://localhost:8080';
const apiClient = axios.create({
  baseURL: baseURL,
  withCredentials: true,
});
// 请求拦截器
// apiClient.interceptors.request.use(
//   (config) => {
//     // 从 sessionStorage 中获取会话信息
//     const sessionToken = sessionStorage.getItem('sessionToken');
//     console.log(sessionToken)
//     if (sessionToken) {
//       // 将会话信息附加到请求头中
//       config.headers.Authorization = `Bearer ${sessionToken}`;
//     }
//     return config;
//   },
//   (error) => {
//     return Promise.reject(error);
//   }
// );
export default apiClient;