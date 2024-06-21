
export default function timeAgo(date) {
    const rtf = new Intl.RelativeTimeFormat("en", { numeric: "auto" });
    const units = [
        { name: 'year', seconds: 31536000 },
        { name: 'month', seconds: 2592000 },
        { name: 'week', seconds: 604800 },
        { name: 'day', seconds: 86400 },
        { name: 'hour', seconds: 3600 },
        { name: 'minute', seconds: 60 },
        { name: 'second', seconds: 1 }
    ];
    const delta = ((new Date()) - date) / 1000;     // s
    for (const unit of units) {
        if (delta >= unit.seconds || unit.name === "second") {
            const value = Math.floor(delta / unit.seconds);
            return rtf.format(-value, unit.name);
        }
    }
}
