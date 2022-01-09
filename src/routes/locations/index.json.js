import { api } from './_api';

// GET /locations.json
export const get = async (request) => {
	const response = await api(request, 'locations');

	if (response.status === 404) {
		// user hasn't created a todo list.
		// start with an empty array
		return { body: [] };
	}

	return response;
};

// POST /locations.json
export const post = async (request) => {
	const response = await api(request, 'locations', {
		name: request.body.name
	});
	return response;
};
