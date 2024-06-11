<script>
import { ref } from 'vue'

import login from '../services/login'

export default {
	data: function() {
		return {
			errormsg: null,
			username: ref(),
			userID: null,
		}
	},
	methods: {
		login: async function() {
			try {
				this.userID = await login(this.username);
				if (this.$router.options.history.state.back == null)
					this.$router.replace("/");
				else this.$router.back();
			}
			catch (e) {
				this.errormsg = e.toString();
			}
		}
	},
	mounted() {
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Login</h1>
		</div>
		<div
			class="d-flex flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 centerDiv">
			<h5>Login to continue to this site</h5>
		</div>
		<div
			class="d-flex flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 centerDiv">
			<input v-model="username" placeholder="username" @keyup.enter="this.login()"/>
			<button type="button" class="btn btn-sm btn-outline-secondary" @click="this.login()">
				Login
			</button>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
