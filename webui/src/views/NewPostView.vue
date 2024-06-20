<script>
import { ref } from 'vue'

import newPost from '../services/newPost'

export default {
	data: function () {
		return {
			errormsg: null,
			uploadB64: null,
			caption: ref(""),
			imgLoading: false
		}
	},
	methods: {
		onUpload() {
			const file = document.getElementById("upForm").files[0];
			const reader = new FileReader();

			reader.addEventListener(
				"load", () => { this.uploadB64 = reader.result; },
				false
			);

			if (file) reader.readAsDataURL(file);
		},
		deleteImg() {
			this.uploadB64 = null;
			this.errormsg = null;
		},
		async publish() {
			this.errormsg = "";
			const b64split = this.uploadB64.split("base64,");
			if (!(new RegExp("^(data\:image\/)(jpeg|png)")).exec(b64split[0])) {
				this.errormsg = "invalid image type!";
				return
			}
			try {
				await newPost(b64split[1], this.caption);
				this.$router.replace("/");
			} catch (e) {
				this.errormsg = e.toString();
			}
		}
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">New post</h1>
		</div>
		<div class="justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<input id="upForm" v-if="!this.uploadB64" type="file" @change="onUpload()" accept="image/*" capture />
			<img id="preview" class="uploadPreview" v-if="this.uploadB64" :src="this.uploadB64" />
			<br />
			<button type="button" class="btn btn-danger" v-if="this.uploadB64" @click="deleteImg">Delete</button>
			<br />
			<textarea id="inputCaption" name="inputCaption" rows="10" cols="100" placeholder="post caption (optional)"
				v-model="caption"></textarea>
			<br />
			<button type="button" class="btn btn-sm btn-outline-primary" @click="publish">Publish</button>
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		</div>
	</div>
</template>

<style>
.uploadPreview {
	width: 300px;
}
</style>
