import { defineStore } from 'pinia';

export const useOverlayStore = defineStore('overlay', {
  state: () => ({
    activeOverlay: null,
  }),
  actions: {
    setActiveOverlay(overlay) {
      this.activeOverlay = overlay;
    },
    clearActiveOverlay() {
      this.activeOverlay = null;
    },
  },
});