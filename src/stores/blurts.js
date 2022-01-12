import { writable } from 'svelte/store';
export const blurts = writable([]);

const { VITE_WEB_URL } = import.meta.env;

const months = [
	'January',
	'February',
	'March',
	'April',
	'May',
	'June',
	'July',
	'August',
	'September',
	'October',
	'November',
	'December'
];

const fetchBlurts = async () => {
	const url = `${VITE_WEB_URL}/blurts.json`;
	const res = await fetch(url);
	const data = await res.json();
	const loadedBlurts = data.map((blurt) => {
		const date = new Date(blurt.created_at);
		blurt.date = `${months[date.getMonth()]} ${date.getDate()}, ${date.getFullYear()}`;
		return blurt;
	});
	blurts.set(loadedBlurts);
};

fetchBlurts();
setInterval(() => fetchBlurts(), 2000);
