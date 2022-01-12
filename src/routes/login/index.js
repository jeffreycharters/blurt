import { prisma } from '$lib/prisma';

export const post = async (request) => {
	const { url, body } = request;

	const payload = Object.fromEntries(body);
	const username = payload.username;

	const forward = new URL(`http://${url.host}/blurts`);

	const user = await prisma.user.findUnique({
		where: {
			username: username
		}
	});

	if (!user) {
		await prisma.user.create({
			data: {
				username: username
			}
		});
	}

	return {
		headers: { Location: forward.toString() },
		status: 302,
		body: {
			user: user
		}
	};
};
