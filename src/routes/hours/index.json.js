import { api } from './_api';

// GET /hours.json
export const get = async (request) => {
	const response = await api(request, 'hours');

	if (response.status === 404) {
		// user hasn't created a todo list.
		// start with an empty array
		return { body: [] };
	}

	return response;
};

// POST /hours.json
export const post = async (request) => {
	const response = await api(request, 'hours', {
		// because index.svelte posts a FormData object,
		// request.body is _also_ a (readonly) FormData
		// object, which allows us to get form data
		// with the `body.get(key)` method
		text: request.body.get('text')
	});

	return response;
};
