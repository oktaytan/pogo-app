import {
	GET_RIGHT_BAR_SHOW,
	CHANGE_RIGHT_BAR_SHOW,
	SET_RIGHT_BAR_SHOW,
	GET_TOP_BAR_SHOW,
	CHANGE_TOP_BAR_SHOW,
	SET_TOP_BAR_SHOW
} from '../keys';

let rightSideBarShow = localStorage.getItem('pogo_settings')
	? JSON.parse(localStorage.getItem('pogo_settings')).rightSideBarShow
	: true;

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
	[GET_RIGHT_BAR_SHOW]: (state) => state.settings.rightSideBarShow,
	[GET_TOP_BAR_SHOW]: (state) => state.settings.topBarShow
};

const actions = {
	[CHANGE_RIGHT_BAR_SHOW]: ({ commit }) => commit('SET_RIGHT_BAR_SHOW'),
	[CHANGE_TOP_BAR_SHOW]: ({ commit }) => commit('SET_TOP_BAR_SHOW')
};

const mutations = {
	[SET_RIGHT_BAR_SHOW]: (state) => {
		state.settings.rightSideBarShow = !state.settings.rightSideBarShow;
		localStorage.setItem('pogo_settings', JSON.stringify(state.settings));
	},
	[SET_TOP_BAR_SHOW]: (state) => {
		state.settings.topBarShow = !state.settings.topBarShow;
		localStorage.setItem('pogo_settings', JSON.stringify(state.settings));
	}
};

export default {
	state,
	getters,
	actions,
	mutations
};
