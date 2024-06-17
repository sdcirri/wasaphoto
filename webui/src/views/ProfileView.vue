<script>
import b64AsBlob from '../services/b64AsBlob'
import getBlocked from '../services/getBlocked'
import getFollowing from '../services/getFollowing'
import getLoginCookie from '../services/getLoginCookie'
import getProfile from '../services/getProfile'
import follow from '../services/follow'
import unfollow from '../services/unfollow'
import block from '../services/block'
import unblock from '../services/unblock'

export default {
	computed: {
		userID() {
			return this.$route.params.id;
		}
  	},
	data: function () {
		return {
			errormsg: null,
			loading: true,
			auth: null,
			profile: null,
			following: null,
			blocked: null,
			ownProfile: false
		}
	},
	methods: {
		async checkFollowing() {
			try {
				const followingList = await getFollowing();
				this.following = followingList.some(id => id == this.profile.userID);
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		async checkBlocked() {
			try {
				const blockedList = await getBlocked();
				this.blocked = blockedList.some(id => id == this.profile.userID);
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		async follow() {
			try {
				await follow(this.profile.userID);
			} catch (e) {
				this.errormsg = e.toString();
			}
            this.refresh();		// to update follower count
		},
		async unfollow() {
            try {
				await unfollow(this.profile.userID);
			} catch (e) {
				this.errormsg = e.toString();
			}
            this.refresh();		// to update follower count
		},
		async block() {
            try {
				await block(this.profile.userID);
			} catch (e) {
				this.errormsg = e.toString();
			}
            await this.checkBlocked();
		},
		async unblock() {
            try {
				await unblock(this.profile.userID);
			} catch (e) {
				this.errormsg = e.toString();
			}
            await this.checkBlocked();
		},
		async refresh() {
			this.loading = true;
			this.errormsg = "";
			this.auth = getLoginCookie();
			this.profile = await getProfile(this.userID);
			this.ownProfile = (this.auth == this.profile.userID);
			const blob = b64AsBlob(this.profile.proPicB64);
			this.blobUrl = URL.createObjectURL(blob);
			this.following = this.checkFollowing();
			this.blocked = this.checkBlocked();
			this.loading = false;
		}
	},
	mounted() {
		this.refresh();
	},
	beforeUnmount() {
		URL.revokeObjectURL(this.blobUrl);
	}
}
</script>

<template>
	<div>
		<LoadingSpinner v-if="loading" />
		<div v-if="!loading" class="columnFlex pt-3 pb-2 mb-3 border-bottom">
			<div class="proHeading">
				<img class="propicTop" :src="'data:image/jpg;base64,' + this.profile.proPicB64" />
				<h1>{{ this.profile.username }}</h1>
				<h4 class="counters">
					<div v-if="!ownProfile">
						<h6>{{ this.profile.followers }} followers</h6>
						<h6>{{ this.profile.following }} following</h6>
					</div>
					<div v-if="ownProfile">
						<h6><RouterLink to="/">{{ this.profile.followers }} followers</RouterLink></h6>
						<h6><RouterLink to="/">{{ this.profile.following }} following</RouterLink></h6>
					</div>
				</h4>
			</div>
			<div class="profileCtrl" v-if="!ownProfile">
				<button class="btn btn-sm btn-outline-primary" v-if="!loading && auth != null && !following" @click="this.follow">Follow</button>
				<button class="btn btn-sm btn-danger" v-if="!loading && auth != null && following" @click="this.unfollow">Unfollow</button>
				<button class="btn btn-sm btn-danger" v-if="!loading && auth != null && !blocked" @click="this.block">Block</button>
				<button class="btn btn-sm btn-outline-primary" v-if="!loading && auth != null && blocked" @click="this.unblock">Unblock</button>
			</div>
			<div class="streamContainer">
				<PostCard v-for="post in this.profile.posts" v-bind:key="post.postID" :ppostID="post" />
			</div>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
.propicTop {
	width: 20vh;
	height: 20vh;
	margin: 0 16px 16px 0;
}

.columnFlex {
	display: flex;
	flex-direction: column;
}

.proHeading {
	display: flex;
	flex-direction: row;
	vertical-align: middle;
	align-items: center;
}

.counters h6 {
	display: flex;
	flex-direction: column;
	margin: auto 32px auto 32px;
}

.profileCtrl > * {
	margin: 0 16px 0 16px;
}
</style>
