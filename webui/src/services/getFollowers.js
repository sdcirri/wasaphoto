import api from './axios'

import { BadAuthException, InternalServerError, AccessDeniedException } from './apiErrors';
import getLoginCookie from './getLoginCookie';

export default async function getFollowers() {
    const uid = getLoginCookie();
    if (uid == null) throw BadAuthException;
    let resp = await api.get(`/users/${uid}/followers`, { "headers": { "Authorization": `bearer ${uid}` } }
    );
    switch (resp.status) {
        case 200:
            return resp.data;
        case 400:
        case 401:
            throw BadAuthException;
        case 403:
            throw AccessDeniedException;
        default:
            throw InternalServerError;
    }
}
