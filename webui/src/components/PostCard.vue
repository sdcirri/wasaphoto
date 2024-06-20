<script>
import { authStatus } from '../services/login'
import getPost from '../services/getPost'
import isLiked from '../services/isLiked'
import likePost from '../services/likePost'
import unlikePost from '../services/unlikePost'
import rmPost from '../services/rmPost'

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
            ownPost: null,
            likeCount: 0,
            loading: true
        }
    },
    methods: {
        async toggleLike() {
            try {
                const liked = await isLiked(this.post.postID);
                if (!liked) {
                    await likePost(this.post.postID);
                    this.likeCount++;
                    this.$refs.likeSvg.classList.add("heartFilled");
                } else {
                    await unlikePost(this.post.postID);
                    this.likeCount--;
                    this.$refs.likeSvg.classList.remove("heartFilled");
                }
                this.indicatorsRefresh();
            } catch (e) {
                this.propagateToView(e);
            }
        },
        goToComments() {
            this.$router.push(`/posts/${this.post.postID}/comments`);
        },
        async refresh() {
            this.loading = true;
            try {
                this.post = await getPost(this.ppostID);
                this.likeCount = this.post.likeCount;
                this.ownPost = (this.post.author == authStatus.status);
                this.loading = false;
                this.indicatorsRefresh();
            } catch (e) {
                this.propagateToView(e);
            }
        },
        async indicatorsRefresh() {
            try {
                const liked = await isLiked(this.post.postID);
                if (liked) this.$refs.likeSvg.classList.add("heartFilled");
                else this.$refs.likeSvg.classList.remove("heartFilled");
            } catch (e) {
                this.propagateToView(e);
            }
        },
        async rmPost() {
            try {
                await rmPost(this.post.postID);
                this.$emit("postDeleted");
            } catch (e) {
                this.propagateToView(e);
            }
        },
        propagateProCardError(e) {
            this.propagateToView(e.error);
        },
        propagateToView(e) {
            this.$emit("renderError", e);
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
            <span class="flex d-flex align-items-center">
                <ProCard :userID="this.post.author" :showControls="!ownPost" @profileError="propagateProCardError" />
                <button class="delBtn" v-if="ownPost" @click="rmPost">
                    <svg class="feather featherBtn">
                        <use href="/feather-sprite-v4.29.0.svg#trash-2" />
                    </svg>
                </button>
            </span>
            <p class="date">on {{ post.pubTime }}</p>
            <div class="imgContainer">
                <img class="postImg" :src="'data:image/jpg;base64,' + post.imageB64" />
            </div>
            <br />
            <p class="caption">{{ post.caption }}</p> <br />
            <div class="flex d-flex justify-center postCtrl">
                <button @click="toggleLike()">
                    <div class="flex d-flex align-items-center">
                        <svg ref="likeSvg" class="feather featherBtn">
                            <use href="/feather-sprite-v4.29.0.svg#heart" />
                        </svg>
                        {{ likeCount }}
                    </div>
                </button>
                <RouterLink v-if="ownPost" :to="`/posts/${post.postID}/likes`">
                    <svg class="feather featherBtn">
                        <use href="/feather-sprite-v4.29.0.svg#eye" />
                    </svg>
                </RouterLink>
                <button @click="goToComments()">
                    <div class="flex d-flex align-items-center">
                        <svg class="feather featherBtn">
                            <use href="/feather-sprite-v4.29.0.svg#message-circle" />
                        </svg>
                        {{ post.comments.length }}
                    </div>
                </button>
            </div>
        </div>
    </div>
</template>

<style>
.postImg {
    max-width: 40vw;
    max-height: 70vh;
}

.postCtrl {
    gap: 3vh;
}

.postCtrl button {
    display: contents;
    font-size: 36px;
}

.postCtrl svg {
    margin-right: 1vh;
}

.heartFilled use {
    fill: red;
}

.postContainer {
    margin: 16px;
    padding: 12px;
    border: 1px solid black;
    inline-size: min-content;
}

.imgContainer {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}

.caption {
    font-size: 16pt;
    margin: 16px 0 0 0;
}

.date {
    margin: 0 0 0 16px;
}

.delBtn {
    display: contents;
}

.delBtn>* {
    border: 1px solid black;
    background-color: red;
    color: white;
}
</style>
<style scoped>
.featherBtn {
    width: 8vh;
    height: 8vh;
}
</style>
