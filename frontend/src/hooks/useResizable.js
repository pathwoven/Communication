import { ref } from 'vue';

export function useResizable(minWidth = 20, maxWidth = 50) {
  const leftWidth = ref('30%');
  const rightWidth = ref('70%');
  const isDragging = ref(false);

  const startDragging = () => {
    isDragging.value = true;
  };

  const stopDragging = () => {
    isDragging.value = false;
  };

  const handleDragging = (event) => {
    if (isDragging.value) {
      const totalWidth = event.target.parentElement.clientWidth;
      const newChatListWidth = (event.clientX / totalWidth) * 100;
      if (newChatListWidth < minWidth || newChatListWidth > maxWidth) {
        stopDragging();
        return;
      }
      leftWidth.value = `${newChatListWidth}%`;
      rightWidth.value = `${100 - newChatListWidth}%`;
    }
  };

  return {
    leftWidth,
    rightWidth,
    startDragging,
    stopDragging,
    handleDragging,
  };
}