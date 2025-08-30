import { env } from '$env/dynamic/public';

export const prerender = true;
export const trailingSlash = 'always';

export async function load() {
	return {
		PUBLIC_API_ADDRESS: env.PUBLIC_API_ADDRESS ?? 'http://localhost:3320'
	};
}
