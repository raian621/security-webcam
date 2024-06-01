import { HLS_SERVER_URL } from "./config";

export async function startStream() : Promise<boolean> {
    try {
        const response = await fetch(`${HLS_SERVER_URL}/start-stream/`, { method: "POST" })
        return response.ok
    } catch (e) {
        console.error(e)
        return false
    }
}

export async function stopStream() : Promise<boolean> {
    try {
        const response = await fetch(`${HLS_SERVER_URL}/stop-stream/`, { method: "POST" })
        return response.ok
    } catch (e) {
        console.error(e)
        return false
    }
}