import { getContext, setContext } from "svelte"

export type Blurt = {
    id: string
    content: string
    author: string
    created: string
    likkers?: string[]
}


export class Blurts {
    list = $state<Blurt[]>([])

    constructor(blurts: Blurt[]) {
        this.list = blurts
    }

    add(blurt: Blurt) {
        this.list = [blurt, ...this.list]
    }

    add_bulk(blurts: Blurt[]) {
        const new_list = [...this.list, ...blurts]
        this.list = [...new Set(new_list)]
    }

}

const BLURT_KEY = Symbol("blurts")

export function setBlurtState(blurts: Blurt[]) {
    return setContext(BLURT_KEY, new Blurts(blurts))
}

export function getBlurtState() {
    return getContext<ReturnType<typeof setBlurtState>>(BLURT_KEY)
}