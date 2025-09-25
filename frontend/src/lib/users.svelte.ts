import { browser } from "$app/environment"
import { getContext, setContext } from "svelte"


export class User {
    list = $state<string[]>([])

    constructor() {
        this.load()
    }

    get active_user() {
        if (this.list.length == 0) return ""

        return this.list[0]
    }

    get alphabetical_list() {
        return this.list.toSorted()
    }

    set_active_user(name: string) {
        this.list = [name, ...this.list.filter((u) => u != name)]
        this.save()
    }

    load() {
        if (!browser) return

        this.list = JSON.parse(localStorage.getItem("user_list") ?? "[]")
    }

    save() {
        localStorage.setItem("user_list", JSON.stringify(this.list))
    }

    trim_name(name: string) {
        if (name.length > 14) return name.slice(0, 14)
        return name
    }

    add(name: string) {
        if (name.length == 0 || this.list.includes(name)) return

        if (this.list.length >= 14) return

        name = this.trim_name(name)

        this.list = [name, ...this.list]

        this.save()
    }

    remove(name: string) {
        name = this.trim_name(name)
        this.list = this.list.filter((u) => u != name)

        this.save()
    }

    set_random_active() {
        const new_active_index = Math.floor(Math.random() * this.list.length)
        const new_list = [this.list[new_active_index], ...this.list.filter((_, i) => i != new_active_index)]

        this.list = new_list

        this.save()
    }
}

const USER_KEY = Symbol("users")

export function setUserState() {
    return setContext(USER_KEY, new User())
}

export function getUserState() {
    return getContext<ReturnType<typeof setUserState>>(USER_KEY)
}