import { BadAuthException, InternalServerError } from './apiErrors';
import getLoginCookie from './getLoginCookie';
import api from './axios'

export default async function isLiked(pid) {
    const uid = getLoginCookie();
    if (uid == null) throw BadAuthException;
    let resp = await api.get(`/posts/${pid}/liked/${uid}`, { "headers": { "Authorization": `bearer ${uid}` } });
    switch (resp.status) {
        case 200:
            return resp.data;
        case 401:
            throw BadAuthException;
        case 404:
            throw PostNotFoundError;
        default:
            throw InternalServerError;
    }
}
