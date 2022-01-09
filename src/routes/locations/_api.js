import { prisma } from '$lib/prisma';

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
