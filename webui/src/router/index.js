import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import SearchView from '../views/SearchView.vue'
import NewPostView from '../views/NewPostView.vue'
import ProfileView from '../views/ProfileView.vue'
import FollowersView from '../views/FollowersView.vue'
import FollowingView from '../views/FollowingView.vue'
import BlockedView from '../views/BlockedView.vue'
import EditProfileView from '../views/EditProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', component: HomeView },
		{ path: '/login', component: LoginView },
		{ path: '/search', component: SearchView },
		{ path: '/newpost', component: NewPostView },
		{ path: '/profile/:id', component: ProfileView },
		{ path: '/profile/:id/followers', component: FollowersView },
		{ path: '/profile/:id/following', component: FollowingView },
		{ path: '/profile/:id/blocked', component: BlockedView },
		{ path: '/profile/:id/edit', component: EditProfileView }
	]
})

export default router
