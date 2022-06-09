import { prisma } from '$lib/prisma';

export async function api(event, resource, data) {
	let body = {};
	let user;
	let status = 500;
	const request = event.request;
	const url = event.url;
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
			const params = url.searchParams;
			const takeNumber = Number(params.get('take')) || 25;
			const cursorUid = params.get('cursor');
			const afterTime = params.get('after');
			const username = decodeURIComponent(params.get('username'));
			user = await prisma.user.findUnique({
				where: {
					username
				}
			});
			let blurts;
			if (!cursorUid && !afterTime) {
				blurts = await prisma.blurt.findMany({
					include: {
						user: true,
						_count: {
							select: {
								liks: true
							}
						}
					},
					orderBy: {
						created_at: 'desc'
					},
					take: takeNumber
				});
			} else if (cursorUid) {
				blurts = await prisma.blurt.findMany({
					cursor: {
						uid: cursorUid
					},
					skip: 1,
					include: {
						user: true,
						_count: {
							select: {
								liks: true
							}
						}
					},
					orderBy: {
						created_at: 'desc'
					},
					take: takeNumber
				});
			} else if (afterTime) {
				blurts = await prisma.blurt.findMany({
					include: {
						user: true,
						_count: {
							select: {
								liks: true
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

			let blurtIds = blurts.map((b) => b.uid);
			let liks = await prisma.lik.findMany({
				where: {
					blurtId: { in: blurtIds },
					userId: user.id
				}
			});
			body = blurts.map((b) => {
				const foundLik = liks.find((l) => l.blurtId === b.uid);
				return {
					...b,
					userLikd: !!foundLik
				};
			});
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
			user = await prisma.user.findUnique({
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
