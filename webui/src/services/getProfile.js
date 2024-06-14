import api from './axios'

import getLoginCookie from './getLoginCookie'
import {
    BadIdsException,
    BlockedException,
    UserNotFoundException,
    InternalServerError
} from './apiErrors'

export default async function getProfile(uid) {
    let auth = getLoginCookie();
    let headers = (auth != null) ? { "Authorization": `bearer ${uid}` } : {};
    let resp = await api.get(`/users/${uid}`, { "headers": headers });
    switch (resp.status) {
        case 200:
            return resp.data;
            break;
        case 400:
            throw BadIdsException;
        case 403:
            throw BlockedException;
        case 404:
            throw UserNotFoundException;
        default:
            throw InternalServerError;
    }
}
