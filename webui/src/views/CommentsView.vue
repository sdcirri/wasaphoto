<script>
import { ref } from 'vue'

import { authStatus } from '../services/login'
import getPost from '../services/getPost'
import commentPost from '../services/commentPost'

export default {
	computed: {
		postID() {
			return this.$route.params.id;
		}
	},
	data: function () {
		return {
			errormsg: null,
			loading: true,
			commentDraft: ref(),
			commentList: []
		}
	},
	methods: {
		async refresh() {
			this.errormsg = "";
			this.commentList = [];
			if (authStatus.status == null)
				this.$router.push("/login");
			else {
				this.loading = true;
				try {
					let post = await getPost(this.postID);
					this.commentList = post.comments;
				} catch (e) {
					this.errormsg = e.toString();
				}
				this.loading = false;
			}
		},
		async submitComment() {
			try {
				let newID = await commentPost(this.postID, this.commentDraft);
				this.commentList.push(new Number(newID));
				this.refresh();
			} catch (e) {
				this.errormsg = e.toString();
			}
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
			<h1 class="h2">Comments</h1>
		</div>
		<div class="streamContainer">
			<LoadingSpinner v-if="loading" />
			<div v-else>
				<div class="commentForm">
					<textarea rows="10" cols="100" placeholder="leave a comment" v-model="commentDraft" />
					<button type="button" class="btn btn-sm btn-outline-primary" @click="submitComment">Submit</button>
				</div>
				<p v-if="this.commentList.length == 0">So empty! Start the discussion!</p>
				<CommentCard v-for="commentID in this.commentList" v-bind:key="commentID" :commentID="commentID"
					@commentDeleted="refresh" />
			</div>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
.commentForm > * {
	vertical-align: middle;
	flex-direction: row;
	margin: 2vh;
}
</style>
