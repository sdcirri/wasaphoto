import api from './axios'

import getLoginCookie from './getLoginCookie'
import { BadIdsException, UserNotFoundException, InternalServerError } from './apiErrors'

export default async function getProfile(uid) {
    let auth = getLoginCookie();
    let headers = (auth != null) ? { "authorization": `bearer ${uid}` } : {};
    try {
        let resp = await api.get(`/users/${uid}`, { "headers": headers });
        switch (resp.status) {
            case 200:
                return resp.data;
                break;
            case 400:
                throw new BadIdsException();
            case 403:
                throw new BlockedException();
            case 404:
                throw new UserNotFoundException();
            default:
                throw new InternalServerError();
        }
    }
    catch (e) {
        throw e;
    }
}
