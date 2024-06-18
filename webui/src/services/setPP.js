import api from './axios'

import {
    BadAuthException,
    AccessDeniedException,
    InternalServerError,
    ImageTooBigException
} from './apiErrors'
import { authStatus } from './login'

export default async function setPP(imgB64) {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.put(`/setPP/${authStatus.status}`, imgB64,
        { "headers": { "Authorization": `bearer ${authStatus.status}` } });
    switch (resp.status) {
        case 200:
            return;
        case 400:
            throw ImageTooBigException;
        case 401:
            throw BadAuthException;
        case 403:
            throw AccessDeniedException;
        default:
            throw InternalServerError;
    }
}
