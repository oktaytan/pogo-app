import axios from 'axios';

export const AXIOS = axios.create({
	baseURL: 'http://localhost:5000/api/'
});

if (localStorage.getItem('pogo_user') !== null) {
	AXIOS.defaults.headers.common['Authorization'] =
		'Bearer ' + JSON.parse(localStorage.getItem('pogo_user')).token;
}

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
