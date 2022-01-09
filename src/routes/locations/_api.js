import { PrismaClient } from '@prisma/client';
/*
This module is used by the /todos.json and /todos/[uid].json
endpoints to make calls to api.svelte.dev, which stores todos
for each user. The leading underscore indicates that this is
a private module, _not_ an endpoint — visiting /todos/_api
will net you a 404 response.

(The data on the todo app will expire periodically; no
	guarantees are made. Don't use it to organise your life.)
    */

const prisma = new PrismaClient();

export async function api(request, resource, data) {
	let body = {};
	let status = 500;
	switch (request.method.toUpperCase()) {
		case 'DELETE':
			await prisma.location.delete({
				where: {
					uid: resource.split('/').pop()
				}
			});
			status = 204;
			break;
		case 'GET':
			body = await prisma.location.findMany();
			status = 200;
			break;
		case 'PATCH':
			body = await prisma.user.update({
				data: {
					done: data.done,
					text: data.text
				},
				where: {
					uid: resource.split('/').pop()
				}
			});
			status = 200;
			break;
		case 'POST':
			body = await prisma.location.create({
				data: {
					name: data.name
				}
			});
			status = 201;
			break;
	}

	return {
		status,
		body
	};
}
