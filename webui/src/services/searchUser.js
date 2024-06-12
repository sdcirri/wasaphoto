import api from './axios'

import getProfile from './getProfile';

export default async function searchUser(query) {
    if (query == "") return;
    let results = [];
    try {
        let resp = await api.get(`/searchUser?q=${query}`, {});
        for (let i = 0; i < resp.data.length; i++) {
            let p = await getProfile(resp.data[i]);
            results.push(p);
        }
        return results;
    }
    catch (e) {
        throw e;
    }
}
