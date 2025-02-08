import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import {createPinia} from 'pinia'
import * as plugins from './plugins'

const app = createApp(App)
app.use(router)
app.use(createPinia())

plugins.setUpNaiveUI(app)
app.mount('#app')
