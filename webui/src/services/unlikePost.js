import {
    AccessDeniedException,
    BadAuthException,
    InternalServerError,
    LikeImpersonationException
} from './apiErrors';
import getLoginCookie from './getLoginCookie';
import api from './axios'

export default async function unlikePost(pid) {
    const uid = getLoginCookie();
    if (uid == null) throw BadAuthException;
    let resp = await api.delete(`/posts/${pid}/unlike/${uid}`, { "headers": { "Authorization": `bearer ${uid}` } });
    switch (resp.status) {
        case 201:
            return;
        case 401:
            throw BadAuthException;
        case 403:
            throw AccessDeniedException;
        case 404:
            throw LikeImpersonationException;
        default:
            throw InternalServerError;
    }
}
