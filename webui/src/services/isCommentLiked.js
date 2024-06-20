import { BadAuthException, InternalServerError } from './apiErrors'
import { authStatus } from './login'
import api from './axios'

export default async function isCommentLiked(pid) {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.get(`/comments/${pid}/liked/${authStatus.status}`,
        { "headers": { "Authorization": `bearer ${authStatus.status}` } });
    switch (resp.status) {
        case 200:
            return resp.data;
        case 401:
            throw BadAuthException;
        case 404:
            throw PostNotFoundError;
        default:
            throw InternalServerError;
    }
}
