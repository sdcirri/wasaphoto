import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import SearchView from '../views/SearchView.vue'
import NewPostView from '../views/NewPostView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', component: HomeView },
		{ path: '/login', component: LoginView },
		{ path: '/search', component: SearchView },
		{ path: '/newpost', component: NewPostView }
//		{path: '/some/:id/link', component: HomeView},
	]
})

export default router
