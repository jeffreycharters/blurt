import { browser } from "$app/environment"
import { getContext, setContext } from "svelte"

export class Chaos {
    active = false

    constructor() {
        if (!browser) return
        this.active = localStorage && localStorage.getItem('chaos') === 'true'
    }

    set_chaos(new_active: boolean) {
        if (!browser) return
        this.active = new_active
        localStorage.setItem('chaos', this.active.toString())
    }
}

const CHAOS_KEY = Symbol("chaos")

export function setChaosState() {
    return setContext(CHAOS_KEY, new Chaos())
}

export function getChaosState() {
    return getContext<ReturnType<typeof setChaosState>>(CHAOS_KEY)
}