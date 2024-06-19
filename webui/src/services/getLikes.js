import api from './axios'

import { BadAuthException, InternalServerError, AccessDeniedException, PostNotFoundException } from './apiErrors'
import { authStatus } from './login'

export default async function getLikes(postID) {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.get(`/posts/${postID}/likes`,
        { "headers": { "Authorization": `bearer ${authStatus.status}` } }
    );
    switch (resp.status) {
        case 200:
            return resp.data;
        case 400:
        case 401:
            throw BadAuthException;
        case 403:
            throw AccessDeniedException;
        case 404:
            throw PostNotFoundException;
        default:
            throw InternalServerError;
    }
}
