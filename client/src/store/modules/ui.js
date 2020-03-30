import * as _ from '../keys';

// Sağ barın açık / kapalı olma durumu localStorage'den alınıp state' e ekleniyor
let rightSideBarShow = localStorage.getItem('pogo_settings')
	? JSON.parse(localStorage.getItem('pogo_settings')).rightSideBarShow
	: true;

// Üst barın açık / kapalı olma durumu localStorage'den alınıp state' e ekleniyor
let topBarShow = localStorage.getItem('pogo_settings')
	? JSON.parse(localStorage.getItem('pogo_settings')).topBarShow
	: true;

const state = {
	settings: {
		rightSideBarShow,
		topBarShow
	}
};

const getters = {
	[_.GET_RIGHT_BAR_SHOW]: (state) => state.settings.rightSideBarShow,
	[_.GET_TOP_BAR_SHOW]: (state) => state.settings.topBarShow
};

const actions = {
	// Sağ barın açık / kapalı olma durumu ayarı
	[_.CHANGE_RIGHT_BAR_SHOW]: ({ commit }) => commit('SET_RIGHT_BAR_SHOW'),
	// Üst barın açık / kapalı olma durumu ayarı
	[_.CHANGE_TOP_BAR_SHOW]: ({ commit }) => commit('SET_TOP_BAR_SHOW')
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
	}
};

export default {
	state,
	getters,
	actions,
	mutations
};
