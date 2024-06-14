import api from './axios'

import {
    BadUploadException,
    BadPostAuthException,
    BadAuthException,
    InternalServerError
} from './apiErrors';
import getLoginCookie from './getLoginCookie';

export default async function newPost(image, caption) {
    const uid = getLoginCookie();
    if (uid == null) throw BadAuthException;
    let resp = await api.post(`/users/${uid}/newpost`,
        { "image": image, "caption": caption },
        { "headers": { "Authorization": `bearer ${uid}` } }
    );
    switch (resp.status) {
        case 201:
            return resp.data;
        case 400:
            throw BadUploadException;
        case 401:
            throw BadAuthException;
        case 403:
            throw BadPostAuthException;
        default:
            throw InternalServerError;
    }
}
