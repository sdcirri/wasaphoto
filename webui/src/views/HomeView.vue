<script>
import { authStatus } from '../services/login'
import getFeed from '../services/getFeed'

export default {
	data: function () {
		return {
			errormsg: null,
			loading: true,
			postList: []
		}
	},
	methods: {
		async refresh() {
			this.postList = [];
			if (authStatus.status == null)
				this.$router.push("/login");
			else {
				this.loading = true;
				this.errormsg = null;
				this.postList = await getFeed();
				this.loading = false;
			}
		},
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
			<h1 class="h2">Home page</h1>
		</div>
		<div class="streamContainer">
			<LoadingSpinner v-if="loading" />
			<div v-else>
				<p v-if="this.postList.length == 0">So empty! Add some new friends to view their photos!</p>
				<PostCard v-for="postID in this.postList" v-bind:key="postID" :ppostID="postID"
					@postDeleted="refresh" />
			</div>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style></style>
