<script>

import follow from '../services/follow';
import getFollowing from '../services/getFollowing';
import getLoginCookie from '../services/getLoginCookie';

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
        proPic() {
            const bin = window.atob(this.proPicB64);
            const arrayBuffer = new ArrayBuffer(bin.length);
            const bytes = new Uint8Array(arrayBuffer);
            for (let i = 0; i < bin.length; i++)
                bytes[i] = bin.charCodeAt(i);
            return new Blob([arrayBuffer], { type: "image/jpg" });
        },
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
        const blob = this.proPic();
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
        <p class="spaced">{{ username }}</p>
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
}
</style>
