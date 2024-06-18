import { InternalServerError } from './apiErrors'
import api from './axios'

export default async function searchUser(query) {
    if (query == "") return;
    let resp = await api.get(`/searchUser?q=${query}`, {});
    if (resp.status == 200) return resp.data;
    throw InternalServerError;
}
