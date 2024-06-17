
export default function b64AsBlob (imgB64) {
    const bin = window.atob(imgB64);
    const arrayBuffer = new ArrayBuffer(bin.length);
    const bytes = new Uint8Array(arrayBuffer);
    for (let i = 0; i < bin.length; i++)
        bytes[i] = bin.charCodeAt(i);
    return new Blob([arrayBuffer], { type: "image/jpg" });
}
