import { InternalServerError } from './apiErrors'
import api from './axios'
import { reactive } from 'vue'

export function getLoginCookie() {
    let cookies = decodeURIComponent(document.cookie).split(";");
    for (let i = 0; i < cookies.length; i++) {
        let c = cookies[i];
        while (c.charAt(0) == ' ')
            c = c.substring(1);
        if (c.indexOf("WASASESSIONID=") == 0)
            return c.substring(14, c.length);
    }
    return null;
}

export const authStatus = reactive({ status: getLoginCookie() });

export default async function login(username) {
    let resp = await api.post("/session", {
        "headers": { "content-type": "application/json" },
        "name": username
    });
    if (resp.status == 201) {
        document.cookie = `WASASESSIONID=${resp.data}; path=/; SameSite=Strict;`;
        authStatus.status = resp.data;
        return resp.data;
    } else throw InternalServerError;
}
