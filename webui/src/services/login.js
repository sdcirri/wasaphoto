import { InternalServerError } from './apiErrors';
import api from './axios'

export default async function login(username) {
    let resp = await api.post("/session", {
        "headers": { "content-type": "application/json" },
        "name": username
    });
    if (resp.status == 201) {
        document.cookie = `WASASESSIONID=${resp.data}; path=/`;
        return resp.data;
    } else throw InternalServerError;
}
