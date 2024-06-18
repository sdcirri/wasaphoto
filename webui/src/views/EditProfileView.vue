<script>
import { ref } from 'vue'

import { UsernameAlreadyTakenException } from '../services/apiErrors'
import { authStatus } from '../services/login'
import setUsername from '../services/setUsername'
import getProfile from '../services/getProfile'
import setPP from '../services/setPP'

export default {
    data: function () {
        return {
            loading: true,
            errormsg: null,
            profile: null,
            usernamebuf: ref(),
            usernamevalid: false,
            uploadB64: null,
            uploadNotOG: false,
            authStatus: authStatus
        }
    },
    methods: {
        async refresh() {
            this.loading = true;
            this.errormsg = "";
            this.profile = await getProfile(authStatus.status);
            this.usernamebuf = ref(this.profile.username);
            this.uploadB64 = "data:image/jpg;base64," + this.profile.proPicB64;
            this.uploadNotOG = false;
            this.loading = false;
        },
        async setUsername() {
            this.validateUsername();
            if (this.usernamevalid)
                try {
                    await setUsername(this.usernamebuf);
                    this.refresh();
                } catch (e) {
                    this.errormsg = e.toString();
                    if (e.toString() === UsernameAlreadyTakenException.toString)
                        this.usernamevalid = false;
                }
        },
        async setPP() {
            if (this.uploadNotOG) {
                const b64split = this.uploadB64.split("base64,");
                if (!(new RegExp("^(data\:image\/)(jpeg|png)")).exec(b64split[0])) {
                    this.errormsg = "invalid image type!";
                    return;
                }
                try {
                    await setPP(b64split[1]);
                    this.refresh();
                } catch (e) {
                    this.errormsg = e.toString();
                }
            }
        },
        validateUsername() {
            this.usernamevalid = true;
            if (this.usernamebuf < 3) {
                this.errormsg = "username too short! At least 3 characters are required";
                this.usernamevalid = false;
            }
            if (this.usernamebuf > 40) {
                this.errormsg = "username too long! Max 40 characters are allowed";
                this.usernamevalid = false;
            }
        },
        onUpload() {
            const file = document.getElementById("upForm").files[0];
            const reader = new FileReader();

            reader.addEventListener(
                "load", () => { this.uploadB64 = reader.result; },
                false
            );

            if (file) {
                reader.readAsDataURL(file);
                this.uploadNotOG = true;
            }
        },
        deleteImg() {
            this.uploadB64 = this.profile.proPicB64;
            this.errormsg = null;
        },
        async submitAll() {
            await this.setUsername();
            await this.setPP();
            this.refresh();
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
            <h1 class="h2">Edit profile</h1>
        </div>
        <div class="editContainer" v-if="!loading">
            <img class="propicPreview" :src="this.uploadB64" alt="Your profile picture" />
            <div class="editPanel">
                <span>
                    Upload a new profile picture
                    <input id="upForm" type="file" @change="onUpload()" accept="image/*" capture />
                    <button type="button" class="btn btn-danger" v-if="uploadNotOG" @click="deleteImg">Delete</button>
                </span>
                <input v-model="usernamebuf" @keyup.enter="setUsername" placeholder="pick a new username" />
                <button type="button" class="btn btn-sm btn-outline-primary" @click="submitAll">Submit</button> <br />
            </div>
        </div>
        <LoadingSpinner v-else />
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    </div>
</template>

<style>
.propicPreview {
    width: 50vh;
    height: 50vh;
}

.editContainer {
    display: flex;
    flex-direction: row;
}

.editPanel {
    display: flex;
    flex-direction: column;
    justify-content: center;
    height: 100%;
}

.editPanel>* {
    margin: 1vh;
}
</style>
