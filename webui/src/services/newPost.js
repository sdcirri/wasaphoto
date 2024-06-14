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
    if (uid == null) throw new BadAuthException();
    try {
        let resp = await api.post(`/users/${uid}/newpost`,
            { "image": image, "caption": caption },
            { "headers": { "Authorization": `bearer ${uid}` } }
        );
        switch (resp.status) {
            case 201:
                return resp.data;
            case 400:
                throw new BadUploadException();
            case 401:
                throw new BadAuthException();
            case 403:
                throw new BadPostAuthException();
            default:
                throw new InternalServerError();
        }
    }
    catch (e) {
        throw e;
    }
}
