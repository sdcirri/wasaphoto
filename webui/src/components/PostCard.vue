<script>
import getPost from '../services/getPost'
import getLoginCookie from '../services/getLoginCookie'
import isLiked from '../services/isLiked'
import likePost from '../services/likePost'
import unlikePost from '../services/unlikePost'

export default {
    props: {
        ppostID: {
            type: Number,
            required: true
        }
    },
    data: function () {
        return {
            post: null,
            auth: null,
            likeIndicator: "🩶",
            likeCount: 0,
            loading: true
        }
    },
    methods: {
        async toggleLike() {
            const liked = await isLiked(this.post.postID);
            if (!liked) {
                await likePost(this.post.postID);
                this.likeCount++;
                this.likeIndicator = "❤️";
            } else {
                await unlikePost(this.post.postID);
                this.likeCount--;
                this.likeIndicator = "🩶";
            }
            this.indicatorsRefresh();
        },
        goToComments() {
            // STUB! This function will push the comment view onto the router stack
            this.indicatorsRefresh();
        },
        async refresh() {
            this.loading = true;
            this.post = await getPost(this.ppostID);
            this.auth = getLoginCookie();
            this.likeCount = this.post.likeCount;
            this.indicatorsRefresh();
            this.loading = false;
        },
        async indicatorsRefresh() {
            const liked = await isLiked(this.post.postID);
            this.likeIndicator = liked ? "❤️" : "🩶";
        }
    },
    mounted() {
        this.refresh();
    }
}
</script>

<template>
    <div>
        <LoadingSpinner v-if="loading" />
        <div v-if="!loading" class="postContainer">
            <ProCard :userID="this.post.author"/>
            <p class="date">on {{ post.pubTime }}</p>
            <img class="postImg" :src=" 'data:image/jpg;base64,' + post.imageB64" /> <br />
            <p class="caption">{{ post.caption }}</p> <br />
            <div class="postCtrl">
                <button @click="toggleLike()">{{ likeIndicator }} {{ likeCount }}</button> <button @click="goToComments()">💬 {{ post.comments.length }}</button>
            </div>
        </div>
    </div>
</template>

<style>
  .postImg {
    height: 60vh;
  }
  .postCtrl button {
    display: contents;
    margin: 0 32px 0 0;
    font-size: 36px;
  }
  .postContainer {
    margin: 16px;
    padding: 12px;
    border: 1px solid black;
    inline-size: min-content;
  }
  .caption {
    margin: 16px 0 0 0;
  }
  .date {
    margin: 0 0 0 16px;
  }
</style>
