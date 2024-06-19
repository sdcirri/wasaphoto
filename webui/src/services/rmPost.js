/*
      responses:
        '204':
          description: Operation successful
        '401':
          description: Unauthenticated
        '403':
          description: Forbidden, you cannot delete somebody else's posts!
        '404':
          description: Post not found
*/
import { AccessDeniedException, BadAuthException, InternalServerError } from './apiErrors'
import api from './axios'

import { authStatus } from './login'

export default async function rmPost(postID) {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.delete(`/posts/${postID}/delete`,
        { "headers": { "Authorization": `bearer ${authStatus.status}` } });
    switch (resp.status) {
        case 204:
            return;
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
