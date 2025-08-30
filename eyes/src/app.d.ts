// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export type User = {
	id: Uuid
	username: string
	create_time: Date
}

export type Blurt = {
	id: Uuid
	author: string
	create_time: Date
	content: string
	liks: number
	userLikd: boolean
	usersBlurt: boolean
	edges: {
		author: User
		liks: Lik[] | undefined
	}
}

export type Lik = {
	user_id: Uuid
	blurt_id: Uuid
}
