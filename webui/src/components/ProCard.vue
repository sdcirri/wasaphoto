<script>
import { authStatus } from '../services/login'
import getProfile from '../services/getProfile'

export default {
    props: {
        userID: {
            type: Number,
            required: true
        },
        showControls: {
            type: Boolean,
            required: true
        }
    },
    data: function () {
        return {
            loading: true,
            profile: null,
            following: null,
            ownProfile: null
        }
    },
    methods: {
        emitError(e) {
            this.$emit("profileError", { "error": e, "userID": this.userID });
        },
        propagateFRM(uid) {
            this.$emit("followerRm", uid);
        },
        propagateUnfollowed(uid) {
            this.$emit("unfollowed", uid);
        },
        propagateUnblock(uid) {
            this.$emit("unblock", uid);
        },
        async refresh() {
            try {
                this.loading = true;
                this.profile = await getProfile(this.userID);
                this.loading = false;
            } catch (e) {
                this.emitError(e);
            }
        }
    },
    async mounted() {
        await this.refresh();
        if (this.profile != null)
            this.ownProfile = (this.profile.userID == authStatus.status);
    }
}
</script>

<template>
    <div class="proBox" id="container" v-if="!loading">
        <img class="propic" :src="`data:image/jpg;base64,${this.profile.proPicB64}`"
            :alt="`${this.profile.username}'s profile picture`" />
        <RouterLink :to="`/profile/${this.profile.userID}`" class="spaced">
            <h3>{{ this.profile.username }}</h3>
        </RouterLink>
        <ProfileControls v-if="!ownProfile && showControls" :userID="this.profile.userID" @controlRefresh="refresh"
            @profileError="emitError" @followerRm="propagateFRM" @unfollowed="propagateUnfollowed"
            @unblock="propagateUnblock" />
        <br />
    </div>
    <LoadingSpinner v-else />
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
