import { prisma } from '$lib/prisma';

export async function api(request, resource, data) {
	let body = {};
	let status = 500;
	switch (request.method.toUpperCase()) {
		case 'GET':
			generatedUserCount = await prisma.user.count({
				where: {
					generator: true
				}
			});
			const skipCount = Math.floor(Math.random() * generatedUserCount);
			let randomUser = await prisma.user.findMany({
				where: {
					generator: true
				},
				skip: skipCount,
				take: 1
			});
			body = randomUser[0] ?? {};
			status = 200;
			break;
	}

	return {
		status,
		body
	};
}
