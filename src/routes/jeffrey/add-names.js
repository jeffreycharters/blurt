import { api } from './_api';

// GET /jeffrey/add-names
export const get = async (event) => {
	const response = await api(event.request, 'users');

	if (response.status === 404) {
		return { body: [] };
	}

	return {
		body: {
			users: response.body
		}
	};
};

// POST /jeffrey/add-names
export const post = async (event) => {
	const data = await event.request.json();
	const response = await api(event.request, 'users', data);
	return response;
};

export const patch = async (event) => {
	data = await event.request.json();
	console.log(data);
	const response = await api(event.request, 'users', data);
	return response;
};
