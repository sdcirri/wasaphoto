import api from './axios'

import { BadAuthException, InternalServerError, AccessDeniedException } from './apiErrors';
import getLoginCookie from './getLoginCookie';

export default async function getFollowers() {
    let uid = getLoginCookie();
    if (uid == null) throw new BadAuthException();
    let resp = await api.get(`/users/${uid}/followers`, {},
        { "headers": { "Authorization": `bearer ${uid}` } }
    );
    switch (resp.status) {
        case 200:
            return resp.data;
        case 400:
        case 401:
            throw new BadAuthException();
        case 403:
            throw new AccessDeniedException();
        default:
            throw new InternalServerError();
    }
}
