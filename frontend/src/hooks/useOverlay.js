import { ref, onMounted, onBeforeUnmount } from 'vue';
import { useOverlayStore } from '@/store/modules/overlayStore';
export function useOverlay() {
  const overlayStore = useOverlayStore();
  const isActive = ref(false);
  const overlayRef = ref(null);
  const menuPosition = ref({ x: 0, y: 0 }); // 用于记录菜单的位置

  const showOverlay = (event) => {
    // overlayStore.setActiveOverlay(overlayName);
    isActive.value = true;
    // 设置菜单的位置
    const { clientX, clientY } = event;
    const { innerWidth, innerHeight } = window;
    const { offsetWidth, offsetHeight } = overlayRef.value;
    const x = clientX + offsetWidth > innerWidth ? innerWidth - offsetWidth : clientX;
    const y = clientY + offsetHeight > innerHeight ? innerHeight - offsetHeight : clientY;
    menuPosition.value = { x, y };
    console.log(menuPosition.value);
  };

  const hideOverlay = () => {
    // overlayStore.clearActiveOverlay();
    isActive.value = false;
  };

  const handleClick = (event) => {
    if(isActive.value && !overlayRef.value.contains(event.target)) {
      hideOverlay();
    }
  };

  onMounted(() => {
    document.addEventListener('click', handleClick);
  });
  onBeforeUnmount(() => {
    document.removeEventListener('click', handleClick);
  });

  return {
    isActive,
    showOverlay,
    hideOverlay,
    overlayRef,
    menuPosition,
  };
}