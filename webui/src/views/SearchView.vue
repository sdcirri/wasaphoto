<script>
import { ref } from 'vue'

import searchUser from '../services/searchUser';

export default {
	data: function() {
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
			<input v-model="query" placeholder="type here to search" @input="search()"/>
			<button type="button" class="btn btn-sm btn-outline-secondary" @click="search()">
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
