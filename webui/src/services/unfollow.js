import api from './axios'

import {
    BadFollowOperation,
    BadAuthException,
    InternalServerError,
    BlockedException,
    UserNotFoundException
} from './apiErrors'
import { authStatus } from './login'

export default async function unfollow(toUnfollow) {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.delete(`/users/${authStatus.status}/unfollow/${toUnfollow}`,
        { "headers": { "Authorization": `bearer ${authStatus.status}` } }
    );
    switch (resp.status) {
        case 204:
            return;
        case 400:
            throw BadFollowOperation;
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
