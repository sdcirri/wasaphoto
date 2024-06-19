<script>
import { authStatus } from '../services/login'
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
            const liked = await isCommentLiked(this.comment.commentID);
            if (!liked) {
                await likeComment(this.comment.commentID);
                this.likeCount++;
                this.$refs.likeSvg.classList.add("heartFilled");
            } else {
                await unlikeComment(this.post.postID);
                this.likeCount--;
                this.$refs.likeSvg.classList.remove("heartFilled");
            }
            this.indicatorsRefresh();
        },
        async refresh() {
            this.loading = true;
            this.comment = await getComment(this.commentID);
            this.likeCount = this.comment.likes;
            this.ownPost = (this.comment.author == authStatus.status);
            this.loading = false;
            this.indicatorsRefresh();
        },
        async indicatorsRefresh() {
            const liked = await isCommentLiked(this.comment.commentID);
            if (liked) this.$refs.likeSvg.classList.add("heartFilled");
            else this.$refs.likeSvg.classList.remove("heartFilled");
        },
        async rmComment() {
            await rmComment(this.comment.commentID);
            this.$emit("commentDeleted");
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
                <ProCard :userID="comment.author" :showControls="!ownPost" />
                <button class="delBtn" v-if="ownPost" @click="rmPost">
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
                <!-- Could be nice but I'm undecided
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
.postImg {
    height: 70vh;
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
