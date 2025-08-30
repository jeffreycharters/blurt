// place files you want to import through the `$lib` alias in this folder.

export const displayDate = (blurtDate: Date, date: Date) => {
	const timeGap = date.valueOf() - blurtDate.valueOf()

	const fewSeconds = 1000 * 60
	const fewMinutes = 1000 * 60 * 15
	// const aboutHour = 1000 * 60 * 45;
	const oneHour = 1000 * 60 * 60
	const hourish = 1000 * 60 * 90
	const coupleHours = 1000 * 60 * 180
	const oneDay = 1000 * 60 * 60 * 24

	if (timeGap < fewSeconds) {
		return "A few seconds ago."
	} else if (timeGap < fewMinutes) {
		return "A few minutes ago."
	} else if (timeGap < fewMinutes) {
		return "Less than an hour ago."
	} else if (timeGap < hourish) {
		return "About an hour ago."
	} else if (timeGap < coupleHours) {
		return "A couple of hours ago."
	} else if (timeGap < oneDay) {
		return `About ${Math.ceil(timeGap / oneHour)} hours ago.`
	} else {
		return blurtDate.toDateString()
	}
}
