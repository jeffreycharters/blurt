import { prisma } from '$lib/prisma';

const base = '/';

export async function api(request, resource, data) {
	let body = {};
	let status = 500;
	switch (request.method.toUpperCase()) {
		case 'DELETE':
			await prisma.todo.delete({
				where: {
					uid: resource.split('/').pop()
				}
			});
			status = 200;
			break;
		case 'GET':
			body = await prisma.hour.findMany();
			status = 200;
			break;
		case 'PATCH':
			body = await prisma.todo.update({
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
			body = await prisma.todo.create({
				data: {
					created_at: new Date(),
					done: false,
					text: data.text
				}
			});
			status = 201;
			break;
	}

	// if the request came from a <form> submission, the browser's default
	// behaviour is to show the URL corresponding to the form's "action"
	// attribute. in those cases, we want to redirect them back to the
	// /todos page, rather than showing the response
	if (request.method !== 'GET' && request.headers.accept !== 'application/json') {
		return {
			status: 303,
			headers: {
				location: '/hours'
			}
		};
	}

	return {
		status,
		body
	};
}
