import api from './axios'
import { BadAuthException, BadFeedException } from './apiErrors';
import { authStatus } from './login'

export default async function getFeed() {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.get(`/feed/${authStatus.status}`, { headers: { "authorization": `bearer ${authStatus.status}`} });
    switch (resp.status) {
        case 200:
            return resp.data;
        case 401:
            throw BadAuthException;
        case 403:
            throw BadFeedException;
    }
}
