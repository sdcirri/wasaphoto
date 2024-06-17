import api from './axios'
import {
    BadUploadException,
    BadPostAuthException,
    BadAuthException,
    InternalServerError
} from './apiErrors';
import { authStatus } from './login'

export default async function newPost(image, caption) {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.post(`/users/${authStatus.status}/newpost`,
        { "image": image, "caption": caption },
        { "headers": { "Authorization": `bearer ${authStatus.status}` } }
    );
    switch (resp.status) {
        case 201:
            return resp.data;
        case 400:
            throw BadUploadException;
        case 401:
            throw BadAuthException;
        case 403:
            throw BadPostAuthException;
        default:
            throw InternalServerError;
    }
}
