import adapter from '@sveltejs/adapter-node';
import { Server } from 'socket.io';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter(),
		vite: {
			plugins: [
				{
					name: 'sveltekit-socket-io',
					configureServer(server) {
						const io = new Server(server.httpServer);
						io.on('connection', (socket) => {
							// Receive incoming messages and broadcast them
							socket.on('blurt', (newBlurt) => {
								io.emit('blurt', newBlurt);
							});

							socket.on('lik', (likdBlurt) => {
								io.emit('lik', likdBlurt);
							});
						});
					}
				}
			]
		},
		// hydrate the <div id="svelte"> element in src/app.html
		target: '#svelte'
	}
};

export default config;
