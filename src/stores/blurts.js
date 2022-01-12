import { writable } from 'svelte/store';

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

export const blurts = writable([]);

const fetchBlurts = async () => {
	const url = 'http://localhost:3005/blurts.json';
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
