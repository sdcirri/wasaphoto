<script>
import { ref } from 'vue'

import searchUser from '../services/searchUser';
import getLoginCookie from '../services/getLoginCookie';

export default {
	data: function () {
		return {
			errormsg: null,
			query: ref(),
			results: [],
		}
	},
	methods: {
		async search() {
			try {
				this.results = await searchUser(this.query);
			}
			catch (e) {
				this.errormsg = e.toString();
			}
			this.refresh();
		},
		onProfileError(e) {
			this.errormsg = e.toString();
		},
		refresh() {
			this.errormsg = null;
		}
	},
	mounted() {
		this.refresh();
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Search</h1>
		</div>
		<div class="d-flex flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 centerDiv">
			<input v-model="query" placeholder="type here to search" @input="search()" />
		</div>

		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<ProCard v-for="profile in results" :key="profile.userID" :userID="profile.userID" @profileError="onProfileError" />
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style></style>
