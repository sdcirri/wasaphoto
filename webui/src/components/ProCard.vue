<script>

import follow from '../services/follow'

export default {
    props: {
        pprofile: {
            type: Object,
            required: true
        },
        pauth: {
            type: String,
            required: true
        },
    },
    data: function () {
        return {
            blobUrl: null,
            // Deep copy props for operations
            profile: JSON.parse(JSON.stringify(this.pprofile)),
            auth: this.pauth
        }
    },
    methods: {
        proPic() {
            const bin = window.atob(this.profile.proPicB64);
            const arrayBuffer = new ArrayBuffer(bin.length);
            const bytes = new Uint8Array(arrayBuffer);
            for (let i = 0; i < bin.length; i++)
                bytes[i] = bin.charCodeAt(i);
            return new Blob([arrayBuffer], { type: "image/jpg" });
        },
        async doFollow() {
            await follow(this.auth, this.profile.userID);
        }
    },
    mounted() {
        const blob = this.proPic();
        this.blobUrl = URL.createObjectURL(blob);
        const btn = document.getElementById("followButton");
        if (btn != null) btn.onclick = this.doFollow;
    },
    beforeDestroy() {
        URL.revokeObjectURL(this.blobUrl);
    }
}
</script>

<template>
    <div v-if="profile.userID != auth" class="proBox" id="container">
        <img class="propic" :src="blobUrl" :alt="`${profile.username}'s profile picture`" />
        <p class="spaced">{{ profile.username }} ({{ profile.userID }})</p>
        <button id="followButton" v-if="auth != null">Follow</button>
        <br />
    </div>
</template>

<style>
.proBox {
    display: flex;
    align-items: center;
    width: 100%;
    height: 36px;
}

.propic {
    width: 36px;
    height: 36px;
}

.spaced {
    margin-right: 12pt;
    margin-left: 12pt;
}
</style>
