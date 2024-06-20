import {
  AccessDeniedException,
  BadAuthException,
  InternalServerError,
  CommentNotFoundException
} from './apiErrors'
import api from './axios'

import { authStatus } from './login'

export default async function rmComment(cid) {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.delete(`/comments/${cid}/delete/${authStatus.status}`,
        { "headers": { "Authorization": `bearer ${authStatus.status}` } });
    switch (resp.status) {
        case 204:
            return;
        case 401:
            throw BadAuthException;
        case 403:
            throw AccessDeniedException;
        case 404:
            throw CommentNotFoundException;
        default:
            throw InternalServerError;
    }
}
