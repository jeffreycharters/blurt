import { prisma } from '$lib/prisma';

export async function api(request, resource, data) {
	let body = {};
	let status = 500;
	switch (request.method.toUpperCase()) {
		case 'DELETE':
			await prisma.user.delete({
				where: {
					uid: resource.split('/').pop()
				}
			});
			status = 200;
			break;
		case 'GET':
			body = await prisma.user.findMany();
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
			body = await prisma.user.create({
				data: {
					created_at: new Date(),
					username: data.username,
					email: data.email
				}
			});
			body = 'thumbs up';
			status = 201;
			break;
	}

	return {
		status,
		body
	};
}
