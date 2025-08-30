import { browser } from "$app/environment"
import { getContext, setContext } from "svelte"

export class PersonaDrawer {
    open = $state(false)

    constructor() {
        if (!browser) return
        this.open = localStorage.getItem('personas_open') === 'true'
    }

    toggle() {
        this.open = !this.open
        localStorage.setItem('personas_open', this.open.toString())
    }
}

const PERSONA_CTX = Symbol("personas")

export function setPersonaDrawerState() {
    return setContext(PERSONA_CTX, new PersonaDrawer())
}

export function getPersonaDrawerState() {
    return getContext<ReturnType<typeof setPersonaDrawerState>>(PERSONA_CTX)
}