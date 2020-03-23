import axios from 'axios';

export const ax = axios.create({
	baseURL: 'http://localhost:5000/api'
});

if (localStorage.getItem('token') !== null) {
	ax.defaults.headers.common['Authorization'] =
		'Bearer ' + JSON.parse(localStorage.getItem('token'));
}

window.ax = ax;

export function capitalize(value) {
	let newValue;
	if (value.split(' ').length > 0) {
		newValue = value
			.split(' ')
			.map((val) => {
				return val.charAt(0).toUpperCase() + val.substr(1);
			})
			.join(' ');
	} else {
		newValue = value.charAt(0).toUpperCase() + value.substr(1);
	}

	return newValue;
}
