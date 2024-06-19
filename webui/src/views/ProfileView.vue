<script>
import b64AsBlob from '../services/b64AsBlob'
import { authStatus } from '../services/login'
import getProfile from '../services/getProfile'
import { reactive } from 'vue';

export default {
	computed: {
		userID() {
			return this.$route.params.id;
		}
	},
	data: function () {
		return {
			authStatus: authStatus,
			errormsg: null,
			loading: true,
			profile: null,
			following: null,
			blocked: null,
			ownProfile: false
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = "";
			this.profile = await getProfile(this.userID);
			this.ownProfile = (authStatus.status == this.profile.userID);
			const blob = b64AsBlob(this.profile.proPicB64);
			this.blobUrl = URL.createObjectURL(blob);
			this.loading = false;
		}
	},
	mounted() {
		this.refresh();
	},
	beforeUnmount() {
		URL.revokeObjectURL(this.blobUrl);
	},
	watch: {
		"$route.params.id": function (newUID) {
			this.refresh();
		}
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
					<div v-else>
						<h6>
							<RouterLink :to="`/profile/${authStatus.status}/followers`">{{ this.profile.followers }}
								followers</RouterLink>
						</h6>
						<h6>
							<RouterLink :to="`/profile/${authStatus.status}/following`">{{ this.profile.following }}
								following</RouterLink>
						</h6>
						<h6>
							<RouterLink :to="`/profile/${authStatus.status}/blocked`">Manage blocked users</RouterLink>
						</h6>
						<h6>
							<RouterLink :to="`/profile/${authStatus.status}/edit`">Edit profile</RouterLink>
						</h6>
					</div>
				</h4>
			</div>
			<ProfileControls v-if="!ownProfile" :userID="this.profile.userID" @controlRefresh="refresh" />
			<div class="streamContainer">
				<PostCard v-for="post in this.profile.posts" v-bind:key="post.postID" :showControls="!ownProfile"
					:ppostID="post" @postDeleted="refresh" />
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
</style>
