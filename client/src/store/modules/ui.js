import * as _ from '../keys';

// Sağ barın açık / kapalı olma durumu localStorage'den alınıp state' e ekleniyor
let rightSideBarShow = localStorage.getItem('pogo_settings')
	? JSON.parse(localStorage.getItem('pogo_settings')).rightSideBarShow
	: true;

// Üst barın açık / kapalı olma durumu localStorage'den alınıp state' e ekleniyor
let topBarShow = localStorage.getItem('pogo_settings')
	? JSON.parse(localStorage.getItem('pogo_settings')).topBarShow
	: true;

// Temanın durumu localStorage'den alınıp state' e ekleniyor
let darkTheme = localStorage.getItem('pogo_settings')
	? JSON.parse(localStorage.getItem('pogo_settings')).darkTheme
	: true;

const state = {
	settings: {
		rightSideBarShow,
		topBarShow,
		darkTheme
	},
	colors: {
		dark: {
			background: '#102a43',
			backgroundLight: '#1a446c',
			borderColor: '#1a446c',
			textPrimary: '#ffffff',
			textSecondary: '#829ab1'
		},
		light: {
			background: '#f0f4f8',
			backgroundLight: '#d9e2ec',
			borderColor: '#bcccdc',
			textPrimary: '#243B53',
			textSecondary: '#486581'
		}
	}
};

const getters = {
	[_.GET_RIGHT_BAR_SHOW]: (state) => state.settings.rightSideBarShow,
	[_.GET_TOP_BAR_SHOW]: (state) => state.settings.topBarShow,
	[_.GET_THEME]: (state) => state.settings.darkTheme,
	[_.GET_COLORS]: (state) => state.colors
};

const actions = {
	// Sağ barın açık / kapalı olma durumu ayarı
	[_.CHANGE_RIGHT_BAR_SHOW]: ({ commit }) => commit(_.SET_RIGHT_BAR_SHOW),
	// Üst barın açık / kapalı olma durumu ayarı
	[_.CHANGE_TOP_BAR_SHOW]: ({ commit }) => commit(_.SET_TOP_BAR_SHOW),
	// Temanın durum ayarı
	[_.CHANGE_THEME]: ({ commit }) => commit(_.SET_THEME)
};

const mutations = {
	[_.SET_RIGHT_BAR_SHOW]: (state) => {
		// Sağ barın açık / kapalı olma durumu göre state güncelleniyor
		state.settings.rightSideBarShow = !state.settings.rightSideBarShow;
		// Bu bilgi kaybolmamması için localStorage' e kaydediliyor
		localStorage.setItem('pogo_settings', JSON.stringify(state.settings));
	},
	[_.SET_TOP_BAR_SHOW]: (state) => {
		// Üst barın açık / kapalı olma durumu göre state güncelleniyor
		state.settings.topBarShow = !state.settings.topBarShow;
		// Bu bilgi kaybolmamması için localStorage' e kaydediliyor
		localStorage.setItem('pogo_settings', JSON.stringify(state.settings));
	},
	[_.SET_THEME]: (state) => {
		// Üst barın açık / kapalı olma durumu göre state güncelleniyor
		state.settings.darkTheme = !state.settings.darkTheme;
		// Bu bilgi kaybolmamması için localStorage' e kaydediliyor
		localStorage.setItem('pogo_settings', JSON.stringify(state.settings));
	}
};

export default {
	state,
	getters,
	actions,
	mutations
};
