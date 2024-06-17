<script>
import { RouterLink, RouterView } from 'vue-router'
import { authStatus } from './services/login'
import logout from './services/logout'

export default {
	data: function () {
		return {
			authStatus: authStatus,
			loading: true
		}
	},
	methods: {
		logout_and_reset() {
			logout();
			this.refresh();
		},
		refresh() {
			this.loading = true;
			console.log("[DBG] refresh triggered")
			this.loading = false;
		}
	},
	mounted() {
		this.refresh();
	}
}
</script>

<template>
	<div>
		<div v-if="!loading">
			<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
				<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAPhoto</a>
				<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse"
					data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false"
					aria-label="Toggle navigation">
					<span class="navbar-toggler-icon"></span>
				</button>
			</header>

			<div class="container-fluid">
				<div class="row">
					<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
						<div class="position-sticky pt-3 sidebar-sticky">
							<ul class="nav flex-column">
								<li class="nav-item">
									<RouterLink to="/" class="nav-link">
										<svg class="feather">
											<use href="/feather-sprite-v4.29.0.svg#home" />
										</svg>
										Home
									</RouterLink>
								</li>
								<li class="nav-item">
									<RouterLink to="/newpost" class="nav-link">
										<svg class="feather">
											<use href="/feather-sprite-v4.29.0.svg#plus" />
										</svg>
										New post
									</RouterLink>
								</li>
								<li class="nav-item">
									<RouterLink :to="`/profile/${authStatus.status}`" class="nav-link" v-if="authStatus.status">
										<svg class="feather">
											<use href="/feather-sprite-v4.29.0.svg#user" />
										</svg>
										My profile
									</RouterLink>
								</li>
								<li class="nav-item">
									<RouterLink to="/search" class="nav-link">
										<svg class="feather">
											<use href="/feather-sprite-v4.29.0.svg#search" />
										</svg>
										Search
									</RouterLink>
								</li>
								<li class="nav-item">
									<RouterLink to="/login" class="nav-link" @click="logout_and_reset">
										<svg class="feather">
											<use href="/feather-sprite-v4.29.0.svg#key" />
										</svg>
										Logout
									</RouterLink>
								</li>
							</ul>
						</div>
					</nav>

					<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
						<RouterView />
					</main>
				</div>
			</div>
		</div>
		<LoadingSpinner v-else />
	</div>
</template>

<style></style>
