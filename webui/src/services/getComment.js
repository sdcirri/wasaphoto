import api from './axios'

import { authStatus } from './login'
import {
    BadIdsException,
    BlockedException,
    UserNotFoundException,
    InternalServerError,
    BadAuthException
} from './apiErrors'

export default async function getComment(cid) {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.get(`/comments/${cid}`, { "headers": { "Authorization": `bearer ${authStatus.status}` } });
    switch (resp.status) {
        case 200:
            return resp.data;
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
