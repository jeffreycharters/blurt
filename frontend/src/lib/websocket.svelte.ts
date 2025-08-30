import { WS_ADDRESS } from "$lib/index"
import type { Blurt, Blurts } from "./blurts.svelte"

export enum ContentType {
    Blurt = "blurt",
    Lik = "lik"
}

export type Message = {
    content_type: ContentType
    payload: LikPayload | BlurtPayload | Blurt
}

type LikPayload = {
    likker: string
    blurt_id: string
}

type BlurtPayload = {
    username: string
    blurt: string
}

let websocket: WebSocket | null = $state(null)

export function get_websocket(blurts: Blurts) {
    if (websocket != null) return websocket

    websocket = new WebSocket(WS_ADDRESS)

    websocket.onopen = function () {
        console.log("Connected")
    }

    websocket.onmessage = function (evt) {
        const message = JSON.parse(evt.data) as Message

        if (message.content_type == ContentType.Blurt) {
            blurts.add(message.payload as Blurt)
            return
        }

        if (message.content_type == ContentType.Lik) {
            const lik = message.payload as LikPayload
            const blurt_index = blurts.list.findIndex((blurt) => blurt.id === (lik.blurt_id))

            blurts.list[blurt_index].likkers = [...blurts.list[blurt_index].likkers ?? [], lik.likker]
            return
        }

        console.log("unknown message", message)
    }

    websocket.onerror = function (evt) {
        console.error("websocket error", evt)
        if (websocket?.CLOSED) reconnect(blurts)
    }

    websocket.onclose = function (evt) {
        console.log("websocket closed", evt)
        reconnect(blurts)
    }

    return websocket
}

function reconnect(blurts: Blurts) {
    websocket = null
    const gummies = setInterval(() => {
        websocket = get_websocket(blurts)
        if (websocket != null) clearInterval(gummies)
    }, 1000)
}