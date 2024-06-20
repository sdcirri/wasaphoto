<script>
import { authStatus } from '../services/login'
import isCommentLiked from '../services/isCommentLiked'
import likeComment from '../services/likeComment'
import unlikeComment from '../services/unlikeComment'
import getComment from '../services/getComment'
import rmComment from '../services/rmComment'
import getPost from '../services/getPost'

export default {
    props: {
        commentID: {
            type: Number,
            required: true
        }
    },
    data: function () {
        return {
            comment: null,
            ownPost: null,
            ownComment: null,
            likeCount: 0,
            loading: true
        }
    },
    methods: {
        async toggleLike() {
            try {
                const liked = await isCommentLiked(this.comment.commentID);
                if (!liked) {
                    await likeComment(this.comment.commentID);
                    this.likeCount++;
                    this.$refs.likeSvg.classList.add("heartFilled");
                } else {
                    await unlikeComment(this.comment.commentID);
                    this.likeCount--;
                    this.$refs.likeSvg.classList.remove("heartFilled");
                }
                this.indicatorsRefresh();
            } catch (e) {
                this.propagateToView(e);
            }
        },
        async refresh() {
            this.loading = true;
            try {
                this.comment = await getComment(this.commentID);
                let post = await getPost(this.comment.postID);
                this.likeCount = this.comment.likes;
                this.ownPost = (post.author == authStatus.status);
                this.ownComment = (this.comment.author == authStatus.status);
                this.loading = false;
                this.indicatorsRefresh();
            } catch (e) {
                this.propagateToView(e);
            }
        },
        async indicatorsRefresh() {
            try {
                const liked = await isCommentLiked(this.comment.commentID);
                if (liked) this.$refs.likeSvg.classList.add("heartFilled");
                else this.$refs.likeSvg.classList.remove("heartFilled");
            } catch (e) {
                this.propagateToView(e);
            }
        },
        async rmComment() {
            try {
                await rmComment(this.comment.commentID);
                this.$emit("commentDeleted");
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
    <div class="cardRoot">
        <LoadingSpinner v-if="loading" />
        <div v-else class="postContainer">
            <span class="flex d-flex align-items-center">
                <ProCard :userID="comment.author" :showControls="!ownComment" @profileError="propagateProCardError" />
                <button class="delBtn" v-if="ownPost || ownComment" @click="rmComment">
                    <svg class="feather featherBtn">
                        <use href="/feather-sprite-v4.29.0.svg#trash-2" />
                    </svg>
                </button>
            </span>
            <p class="date">on {{ comment.time }}</p>
            <p class="caption">{{ comment.content }}</p> <br />
            <div class="flex d-flex justify-center postCtrl">
                <button @click="toggleLike()">
                    <div class="flex d-flex align-items-center">
                        <svg ref="likeSvg" class="feather featherBtn">
                            <use href="/feather-sprite-v4.29.0.svg#heart" />
                        </svg>
                        {{ likeCount }}
                    </div>
                </button>
                <!-- Would be nice but I'm undecided
                <RouterLink v-if="ownPost" :to="`/comments/${comment.commentID}/likes`">
                    <svg class="feather featherBtn">
                        <use href="/feather-sprite-v4.29.0.svg#eye" />
                    </svg>
                </RouterLink>
                -->
            </div>
        </div>
    </div>
</template>

<style>
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
    width: 42vw;
    margin: 16px;
    padding: 12px;
    border: 1px solid black;
}

.caption {
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
