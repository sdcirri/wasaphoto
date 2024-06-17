import api from './axios'

import { BadAuthException, InternalServerError, AccessDeniedException } from './apiErrors'
import { authStatus } from './login'

export default async function getFollowers() {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.get(`/users/${authStatus.status}/followers`, { "headers": { "Authorization": `bearer ${authStatus.status}` } }
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
