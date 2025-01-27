import { defineStore } from 'pinia';

export const useSettingStore = defineStore('setting', {
  state: () => ({
    setting:{
        birthday: null,
        sex: null,
        font_size: null,
        font_style: null,
        theme: null,
        background: null,
        receive_notice: null,
        notice_sound: null,
        password: null,
        email: null,
        create_time: null,
        last_login_time: null,
        is_deleted: null,
        found_by_account_id: null,
        found_by_nickname: null,
        found_by_group: null,
        found_by_friend: null,
        allow_stranger_message: null,
        is_invisible: null,

    }
  }),
  actions: {
    updateSettings(newSettings) {
        // todo
        this.setting = { ...this.setting, ...newSettings };
    },
  },
});