import { env } from "$env/dynamic/public"

export const API_ADDRESS = env.PUBLIC_API_ADDRESS ?? "http://localhost:3320"
