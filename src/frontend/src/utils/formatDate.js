export const formatDate = (date) => {
	// change format from yyyy-mm-dd to dd MONTHNAME yyyy
	const monthNames = [
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	];
	const dateObj = new Date(date);
	const day = dateObj.getDate();
	const monthIndex = dateObj.getMonth();
	const year = dateObj.getFullYear();
	return `${day} ${monthNames[monthIndex]} ${year}`;
};