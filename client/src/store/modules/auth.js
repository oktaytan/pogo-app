import { AXIOS } from '../../utils/settings';
import {
	USER_LOGIN,
	USER_LOGOUT,
	USER_REFRESH,
	USER_REGISTER,
	SET_USER,
	GET_USER
} from '../keys';

const state = {
	user: null
};

const getters = {
	GET_USER: (state) => state.user
};

const actions = {
	// Kullanıcı login oluyor
	USER_LOGIN: ({ commit }, user) => {
		return new Promise((resolve, reject) => {
			AXIOS.post('/login', JSON.stringify(user))
				.then((res) => {
					if (!res.data.error) {
						// Hata yoksa kullanıcı biilgileri
						// token ile birlikte localStorage ekleniyor
						localStorage.setItem('pogo_user', JSON.stringify(res.data));
						commit('SET_USER', res.data);
						resolve(res.data);
					} else {
						reject(res.data);
					}
				})
				.catch((err) => {
					reject(err);
				});
		});
	},

	// Kullanıcı logout oluyor
	USER_LOGOUT: ({ commit }) => {
		return new Promise((resolve, reject) => {
			if (localStorage.getItem('pogo_user')) {
				AXIOS.get('/logout')
					.then((res) => {
						commit('SET_USER', res.data);
						resolve(res.data);
					})
					.catch((err) => {
						reject(err);
					});
			} else {
				reject();
			}
		});
	},

	// Yeni kullanıcı kayıt oluyor
	USER_REGISTER: ({ commit }, user) => {
		return new Promise((resolve, reject) => {
			AXIOS.post('/register', JSON.stringify(user))
				.then((res) => {
					if (!res.data.error && res.data.user_added) {
						resolve(res.data);
					} else {
						reject(res.data);
					}
				})
				.catch((err) => {
					reject(err);
				});
		});
	},

	// Sayfa yenilenmesi durumunda localStorage' den
	// kullanıcı bilgileri alınıp state' e ekleniyor
	USER_REFRESH: ({ commit }) => {
		const userData =
			localStorage.getItem('pogo_user') &&
			JSON.parse(localStorage.getItem('pogo_user'));

		commit('SET_USER', userData);
	}
};

const mutations = {
	SET_USER: (state, userData) => {
		state.user = userData.user;
	}
};

export default {
	state,
	getters,
	actions,
	mutations
};
