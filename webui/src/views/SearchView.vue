<script>
import { ref } from 'vue'

export default {
	data: function() {
		return {
			errormsg: null,
			query: ref(),
			results: [],
		}
	},
	methods: {
		async getUsername(uid) {
			try {
				let resp = await this.$axios.get("/users/" + uid, {});
				return resp.data["username"];
			}
			catch (e) {
				this.errormsg = e.toString();
				return null;
			}
		},
		async search(query) {
			if (query == "") return;
			try {
				this.results = [];
				let resp = await this.$axios.get("/searchUser?q=" + query, {});
				for (let i = 0; i < resp.data.length; i++) {
					let u = await this.getUsername(resp.data[i]);
					this.results.push(u);
				}
			}
			catch (e) {
				this.errormsg = e.toString();
			}
		},
	},
	mounted() {
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Search</h1>
		</div>
		<div
			class="d-flex flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 centerDiv">
			<input v-model="query" placeholder="type here to search" @input="search(query)"/>
			<button type="button" class="btn btn-sm btn-outline-secondary" @click="search(query)">
				Search
			</button>
		</div>

		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<ul>
				<li v-for="user in results">
					{{ user }}
				</li>
			</ul>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
