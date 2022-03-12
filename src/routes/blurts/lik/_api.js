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
					user: true,
					liks: true
				},
				orderBy: {
					created_at: 'desc'
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
			await prisma.lik.create({
				data: {
					userId: user.id,
					blurtId: data.uid
				}
			});
			status = 201;
			let likdBlurt = await prisma.blurt.findUnique({
				where: {
					uid: data.uid
				},
				include: {
					user: true,
					liks: true
				}
			});
			body = likdBlurt;
			break;
	}

	return {
		status,
		body
	};
}
