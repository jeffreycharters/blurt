import { prisma } from '$lib/prisma';

export async function api(request, resource, data) {
	let body = {};
	let status = 500;

	if (resource === 'users') {
		switch (request.method.toUpperCase()) {
			case 'GET':
				body = await prisma.user.findMany({
					where: {
						generator: true
					},
					orderBy: {
						id: 'desc'
					},
					take: 25
				});
				status = 200;
				break;
			case 'PATCH':
				body = await prisma.user.update({
					data: {
						username: data.username
					},
					where: {
						id: data.id
					}
				});
				status = 200;
				break;
			case 'POST':
				body = await prisma.user.create({
					data: {
						username: data.username,
						generator: true
					}
				});
				status = 201;
				break;
		}
	}

	return {
		status,
		body
	};
}
