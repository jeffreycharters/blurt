import { api } from './_api';

// GET /users
export const get = async (event) => {
	const response = await api(event.request, 'users');

	if (response.status === 404) {
		return { body: [] };
	}
	return response;
};
