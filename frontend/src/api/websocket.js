import { baseURL } from "./api";
import {useChatStore} from "@/store/modules/chatStore";
import { useSettingStore } from "@/store/modules/setting";
import eventBus from "@/utils/eventBus";  

const chatStore = useChatStore();
const settingStore = useSettingStore();

export const createWebSocketConnection = () => {
  const ws = new WebSocket(`${baseURL}/api/v1/ws`);

  // 打开连接
  ws.onopen = () => {
    console.log('连接成功');
  };
  // 接收消息
  ws.onmessage = async (event) => {
    console.log('收到消息:', event.data);
    // 反序列化消息
    const reader = new FileReader();
    reader.onload = () => {
      try{
        const data = JSON.parse(reader.result);
        handleData(data);
      }
      catch(e){
        console.error('解析消息失败:', e);
      }
    }
    reader.onerror = (e) => {
      console.error('读取消息失败:', e);  
    }
    reader.readAsText(event.data);
  };
  // 关闭连接
  ws.onclose = () => {
    console.log('连接关闭');
  };
  // 连接错误
  ws.onerror = (error) => {
    console.error('连接错误:', error);
  };

  return ws;
};

const handleData = (data) =>{
  console.log('解析消息:', data);
  // 新消息提醒
  if(data.needNotice){
    // 提醒 todo
  }
  // message
  if(data.type === 'message'){
    if(chatStore.selectedChat && (chatStore.selectedChat.target_id === data.data.chat.target_id)){
      eventBus.emit('new-message', data.data.message);
    }
    chatStore.updataChatMessage(data.data.chat.target_id, data.data.chat.last_message, data.data.chat.last_person, data.data.message.create_time);
  }
}