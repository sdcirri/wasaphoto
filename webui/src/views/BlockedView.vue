<script>
import { authStatus } from '../services/login'
import getBlocked from '../services/getBlocked'

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
				this.userList = await getBlocked();
				this.loading = false;
			}
		},
		onProfileError(e) {
			this.errormsg = e.toString();
		},
		unblock(uid) {
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
			<h1 class="h2">Blocked users</h1>
		</div>
		<div class="proCardList">
			<LoadingSpinner v-if="loading" />
			<ProCard v-else v-for="uid in userList" v-bind:key="uid" :userID="uid" :showControls="true"
				@profileError="onProfileError" @unblock="unblock" />
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style></style>
