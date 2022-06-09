import adapter from '@sveltejs/adapter-node';
import { Server } from 'socket.io';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter()
	}
};

export default config;
