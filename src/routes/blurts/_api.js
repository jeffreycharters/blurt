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
			const params = request.url.searchParams;
			const takeNumber = Number(params.get('take')) || 25;
			const cursorUid = params.get('cursor');
			const afterTime = params.get('after');
			if (!cursorUid && !afterTime) {
				body = await prisma.blurt.findMany({
					include: {
						user: true,
						liks: {
							include: {
								user: true
							}
						}
					},
					orderBy: {
						created_at: 'desc'
					},
					take: takeNumber
				});
			} else if (cursorUid) {
				body = await prisma.blurt.findMany({
					cursor: {
						uid: cursorUid
					},
					skip: 1,
					include: {
						user: true,
						liks: {
							include: {
								user: true
							}
						}
					},
					orderBy: {
						created_at: 'desc'
					},
					take: takeNumber
				});
			} else if (afterTime) {
				body = await prisma.blurt.findMany({
					include: {
						user: true,
						liks: {
							include: {
								user: true
							}
						}
					},
					orderBy: {
						created_at: 'desc'
					},
					where: {
						created_at: {
							gt: afterTime
						}
					}
				});
			}
			status = 200;
			if (body.length === 0) {
				status = 404;
			}
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
			const newBlurt = await prisma.blurt.create({
				data: {
					userId: user.id,
					blurt: data.blurt
				},
				include: {
					user: true
				}
			});
			body = newBlurt;
			status = 201;
			break;
	}

	return {
		status,
		body
	};
}
