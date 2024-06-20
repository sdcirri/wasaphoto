import api from './axios'
import {
    BadAuthException,
    BadCommentException,
    InternalServerError,
    PostNotFoundException
} from './apiErrors'
import { authStatus } from './login'

export default async function commentPost(pid, text) {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.post(`/posts/${pid}/comment/${authStatus.status}`,
        text, { "headers": { "Authorization": `bearer ${authStatus.status}` } }
    );
    switch (resp.status) {
        case 201:
            return resp.data;
        case 400:
            throw BadCommentException;
        case 401:
            throw BadAuthException;
        case 404:
            throw PostNotFoundException;
        default:
            throw InternalServerError;
    }
}
