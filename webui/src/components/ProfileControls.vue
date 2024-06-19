<script>
import { authStatus } from '../services/login'
import getFollowing from '../services/getFollowing'
import getFollowers from '../services/getFollowers'
import getBlocked from '../services/getBlocked'
import getProfile from '../services/getProfile'
import follow from '../services/follow'
import unfollow from '../services/unfollow'
import rmFollower from '../services/rmFollower'
import block from '../services/block'
import unblock from '../services/unblock'

export default {
	props: {
		userID: {
			type: Number,
			required: true
		}
	},
	data: function () {
		return {
			loading: true,
			authStatus: authStatus,
			following: null,
			follower: null,
			blocked: null,
			profile: null
		}
	},
	methods: {
		async checkFollowing() {
			const followingList = await getFollowing();
			this.following = followingList.some(id => id == this.profile.userID);
		},
		async checkFollower() {
			const followerList = await getFollowers();
			this.follower = followerList.some(id => id == this.profile.userID);
		},
		async checkBlocked() {
			const blockedList = await getBlocked();
			this.blocked = blockedList.some(id => id == this.profile.userID);
		},
		async follow() {
			try {
				await follow(this.profile.userID);
				this.$emit("controlRefresh");
			} catch (e) {
				this.$emit("profileError", e);
			}
		},
		async unfollow() {
			try {
				await unfollow(this.profile.userID);
				this.$emit("unfollowed", this.profile.userID);
				this.$emit("controlRefresh");
			} catch (e) {
				this.$emit("profileError", e);
			}
		},
		async rmFollower() {
			try {
				await rmFollower(this.profile.userID);
				this.$emit("followerRm", this.profile.userID);
				this.$emit("controlRefresh");
			} catch (e) {
				this.$emit("profileError", e);
			}
		},
		async block() {
			try {
				await block(this.profile.userID);
				this.$emit("controlRefresh");
			} catch (e) {
				this.$emit("profileError", e);
			}
		},
		async unblock() {
			try {
				await unblock(this.profile.userID);
				this.$emit("unblock", this.profile.userID);
				this.$emit("controlRefresh");
			} catch (e) {
				this.$emit("profileError", e);
			}
		}
	},
	async mounted() {
		this.loading = true;
		try {
			this.profile = await getProfile(this.userID);
			await this.checkFollowing();
			await this.checkFollower();
			await this.checkBlocked();
		} catch (e) {
			this.$emit("profileError", e);
		}
		this.loading = false;
	}
}
</script>

<template>
	<div class="profileCtrl" v-if="!loading && authStatus.status != null">
		<button class="btn btn-sm btn-outline-primary" v-if="!following" @click="this.follow">Follow</button>
		<button class="btn btn-sm btn-danger" v-else @click="this.unfollow">Unfollow</button>
		<button class="btn btn-sm btn-danger" v-if="follower" @click="this.rmFollower">Remove follower</button>
		<button class="btn btn-sm btn-danger" v-if="!blocked" @click="this.block">Block</button>
		<button class="btn btn-sm btn-outline-primary" v-else @click="this.unblock">Unblock</button>
	</div>
</template>

<style>
.profileCtrl>* {
	margin: 0 .5vh 0 .5vh;
}
</style>
