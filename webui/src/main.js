import { createApp, reactive } from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js'
import ErrorMsg from './components/ErrorMsg.vue'
import ProCard from './components/ProCard.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App);
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("ProCard", ProCard);
app.component("LoadingSpinner", LoadingSpinner);
app.use(router);
app.mount('#app');
