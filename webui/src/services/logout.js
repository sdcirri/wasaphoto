
export default function logout() {
    document.cookie = "WASASESSIONID=;path=/;expires=Thu, 01 Jan 1970 00:00:01 GMT";
}
