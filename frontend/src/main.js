import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import naive from 'naive-ui'
import i18n from './i18n'

const app = createApp(App)
const pinia = createPinia()

app.use(router)
app.use(pinia)
app.use(naive)
app.use(i18n)
app.mount('#app') 