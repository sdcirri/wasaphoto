<script>
import { authStatus } from '../services/login'
import getFollowers from '../services/getFollowers'
import { BlockedException } from '../services/apiErrors';

export default {
	data: function () {
		return {
			loading: true,
			errormsg: null,
			userList: []
		}
	},
	methods: {
		async refresh() {
			this.userList = [];
			if (authStatus.status == null)
				this.$router.push("/login");
			else {
				this.loading = true;
				this.errormsg = null;
				this.userList = await getFollowers();
				this.loading = false;
			}
		},
		onProfileError(e) {
			if (e.error === BlockedException.toString())
				this.removeFollower(e.userID);
			else this.errormsg = e.toString();
		},
		removeFollower(uid) {
			let i = this.userList.indexOf(uid);
			if (i !== -1) this.userList.splice(i, 1);
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
			<h1 class="h2">Followers</h1>
		</div>
		<div class="proCardList">
			<LoadingSpinner v-if="loading" />
			<ProCard v-else v-for="uid in userList" v-bind:key="uid" :userID="uid" :showControls="true"
				@profileError="this.onProfileError" @followerRm="this.removeFollower" />
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style></style>
