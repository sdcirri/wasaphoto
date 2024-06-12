import api from './axios'

import { BadFollowOperation, BadAuthException, InternalServerError, BlockedException, UserNotFoundException } from './apiErrors';
import getLoginCookie from './getLoginCookie';

export default async function follow(follower, toFollow) {
    let uid = getLoginCookie();
    if (uid == null) throw new BadAuthException();
    let resp = await api.post(`/users/${follower}/follow/${toFollow}`,
        { "headers": { "Authorization": `bearer ${uid}` } }
    );
    switch (resp.status) {
        case 201:
            return;
        case 400:
            throw new BadFollowOperation();
        case 401:
            throw new BadAuthException();
        case 403:
            throw new BlockedException();
        case 404:
            throw new UserNotFoundException();
        default:
            throw new InternalServerError();
    }
}
