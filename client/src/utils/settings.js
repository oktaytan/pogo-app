import axios from 'axios';

// axios' dan bir örnek oluşturuluyor
export const AXIOS = axios.create({
	baseURL: 'http://localhost:5000/api/'
});

// Axios örneğinden yapılacak isteklerin headers alanına kullanıcının tokenı'ı ekleniyor
if (localStorage.getItem('pogo_user') !== null) {
	AXIOS.defaults.headers.common['Authorization'] =
		'Bearer ' + JSON.parse(localStorage.getItem('pogo_user')).token;
}
