<script>
import getLoginCookie from '../services/getLoginCookie'

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			userID: null,
			postList: [],
		}
	},
	methods: {
		async refresh() {
			if (this.userID == null) {
				this.userID = getLoginCookie();
				if (this.userID == null)
					this.$router.push("/login");
			}
			if (this.userID != null) {
				this.loading = true;
				this.errormsg = null;
				try {
					let response = await this.$axios.get(
						"/feed/" + this.userID,
						{
							headers: { "authorization": "bearer " + this.userID },
						});
					this.postList = response.data;
				} catch (e) {
					this.errormsg = e.toString();
				}
				this.loading = false;
			}
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
		</div>
		<div>
			<p v-if="this.postList.length == 0">So empty! Add some new friends to view their photos!</p>
			<ul>
				<li v-for="postID in this.postList">postID</li>
			</ul>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
