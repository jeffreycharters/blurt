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
			body = await prisma.blurt.findMany({
				include: {
					user: true
				},
				orderBy: {
					created_at: 'desc'
				}
			});
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
			const user = await prisma.user.findUnique({
				where: {
					username: data.username
				}
			});
			newBlurt = await prisma.blurt.create({
				data: {
					userId: user.id,
					blurt: data.blurt
				},
				include: {
					user: true
				}
			});
			body = await prisma.blurt.findUnique({
				where: {
					uid: newBlurt.uid
				},
				include: {
					user: true
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
