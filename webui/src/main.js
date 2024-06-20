import { createApp, reactive } from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js'
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import ProCard from './components/ProCard.vue'
import PostCard from './components/PostCard.vue'
import ProfileControls from './components/ProfileControls.vue'
import CommentCard from './components/CommentCard.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App);
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("ProCard", ProCard);
app.component("PostCard", PostCard);
app.component("ProfileControls", ProfileControls);
app.component("LoadingSpinner", LoadingSpinner);
app.component("CommentCard", CommentCard);
app.use(router);
app.mount('#app');
