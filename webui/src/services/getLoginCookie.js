export default function getLoginCookie() {
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
