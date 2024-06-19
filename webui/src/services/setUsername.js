import api from './axios'

import { BadAuthException, UsernameAlreadyTakenException, InternalServerError } from './apiErrors'
import { authStatus } from './login'

export default async function setUsername(username) {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.put(`/setUsername/${authStatus.status}`, username,
        { "headers": { "Authorization": `bearer ${authStatus.status}` } });
    switch (resp.status) {
        case 200:
            return;
        case 401:
            throw BadAuthException;
        case 403:
            throw UsernameAlreadyTakenException;
        default:
            throw InternalServerError;
    }
}
