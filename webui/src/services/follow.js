import api from './axios'

import {
    BadFollowOperation,
    BadAuthException,
    InternalServerError,
    BlockedException,
    UserNotFoundException
} from './apiErrors'
import { authStatus } from './login'

export default async function follow(toFollow) {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.post(`/users/${authStatus.status}/follow/${toFollow}`, {},
        { "headers": { "Authorization": `bearer ${authStatus.status}` } }
    );
    switch (resp.status) {
        case 201:
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
