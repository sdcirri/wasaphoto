<script>
import { ref } from 'vue'

export default {
	data: function() {
		return {
			errormsg: null,
			username: ref(),
			userID: null,
		}
	},
	methods: {
		login() {
			console.log("Logging in as " + this.username);
			try {
				let resp = this.$axios.post("/session", {"name": this.username});
				this.userID = resp.data;
				console.log("Got userID =", this.userID);
			}
			catch(e) {
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
			class="d-flex flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3">
			<input v-model="username" placeholder="username"/>
			<button type="button" class="btn btn-sm btn-outline-secondary" @click="login(username)">
				Login
			</button>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
