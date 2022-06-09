import { api } from './_api';

// PATCH /todos/:uid.json
// export const patch: RequestHandler<Locals, FormData> = async (request) => {
//   return api(request, `todos/${request.locals.userid}/${request.params.uid}`, {
//     text: request.body.get("text"),
//     done: request.body.has("done") ? !!request.body.get("done") : undefined,
//   })
// }

export const post = async (event) => {
	const body = await event.request.json();
	return api(event, '', {
		username: body.username,
		uid: body.uid
	});
};
