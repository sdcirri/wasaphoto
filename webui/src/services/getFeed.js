import api from './axios'
import { BadAuthException, BadFeedException } from './apiErrors';
import getLoginCookie from './getLoginCookie'

export default async function getFeed() {
    const uid = getLoginCookie();
    if (uid == null) throw BadAuthException;
    let resp = await api.get(`/feed/${uid}`, { headers: { "authorization": `bearer ${uid}`} });
    switch (resp.status) {
        case 200:
            return resp.data;
        case 401:
            throw BadAuthException;
        case 403:
            throw BadFeedException;
    }
}
