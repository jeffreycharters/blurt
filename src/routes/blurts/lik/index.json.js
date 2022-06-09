import { api } from './_api';

export const get = async (event) => {
	return api(event, '');
};
