import api from './axios'

import getLoginCookie from './getLoginCookie'
import {
    BadIdsException,
    BlockedException,
    UserNotFoundException,
    InternalServerError,
    BadAuthException
} from './apiErrors'

export default async function getPost(pid) {
    const uid = getLoginCookie();
    if (uid == null) throw BadAuthException;
    let resp = await api.get(`/posts/${pid}`, { "headers": { "Authorization": `bearer ${uid}` } });
    switch (resp.status) {
        case 200:
            return resp.data;
            break;
        case 400:
            throw BadIdsException;
        case 401:
            throw BadAuthException;
        case 403:
            throw BlockedException;
        case 404:
            throw UserNotFoundException;
        default:
            throw InternalServerError;
    }
}
