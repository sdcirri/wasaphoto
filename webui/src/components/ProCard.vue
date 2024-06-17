<script>

import follow from '../services/follow'
import getFollowing from '../services/getFollowing'
import getLoginCookie from '../services/getLoginCookie'
import b64AsBlob from '../services/b64AsBlob'

export default {
    props: {
        profile: {
            type: Object,
            required: true
        }
    },
    data: function () {
        return {
            blobUrl: null,
            uid: null,
            username: null,
            proPicB64: null,
            following: null,
            auth: null
        }
    },
    methods: {
        async follow() {
            await follow(this.uid);
            await this.checkFollowing();
        },
        async checkFollowing() {
            const followingList = await getFollowing();
            this.following = followingList.some(id => id == this.uid);
        }
    },
    async mounted() {
        // Deep copy props for operations
        this.uid = new Number(this.profile.userID);
        this.username = this.profile.username;
        this.proPicB64 = this.profile.proPicB64;
        this.auth = getLoginCookie();
        await this.checkFollowing();
        const blob = b64AsBlob(this.proPicB64);
        this.blobUrl = URL.createObjectURL(blob);
    },
    beforeUnmount() {
        URL.revokeObjectURL(this.blobUrl);
    }
}
</script>

<template>
    <div v-if="uid != auth" class="proBox" id="container">
        <img class="propic" :src="blobUrl" :alt="`${username}'s profile picture`" />
        <RouterLink :to="`/profile/${ uid }`" class="spaced">{{ username }}</RouterLink>
        <button id="followButton" class="btn btn-sm btn-outline-primary" v-if="auth != null && !following" @click="this.follow">Follow</button>
        <br />
    </div>
</template>

<style>
.proBox {
    display: flex;
    align-items: center;
    width: 100%;
    height: 10vh;
}

.propic {
    width: 5vh;
    height: 5vh;
}

.spaced {
    margin-right: 2vh;
    margin-left: 2vh;
    font-size: 2vh;
    font-style: bold;
}
</style>
