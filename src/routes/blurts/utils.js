export const humanizeDates = (blurts) => {
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
	const datedBlurts = blurts.map((blurt) => {
		const date = new Date(blurt.created_at);
		blurt.date = `${months[date.getMonth()]} ${date.getDate()}, ${date.getFullYear()}`;
		return blurt;
	});
	return datedBlurts;
};
