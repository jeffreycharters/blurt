import { api } from './_api';

// GET /locations.json
export const get = async (event) => {
	const response = await api(event, 'blurts');
	if (response.status === 404) {
		// no blurts!
		// start with an empty array
		return { body: [] };
	}
	return response;
};

// POST /locations.json
export const post = async (event) => {
	const body = await event.request.json();
	if (body.blurt === '') {
		return {
			status: 200,
			body: {}
		};
	}
	const response = await api(event, '/blurts', {
		username: body.username,
		blurt: body.blurt
	});
	return response;
};
