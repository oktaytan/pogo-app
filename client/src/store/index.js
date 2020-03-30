import Vue from 'vue';
import Vuex from 'vuex';

import auth from './modules/auth';
import posts from './modules/posts';
import ui from './modules/ui';

Vue.use(Vuex);

export default new Vuex.Store({
	modules: {
		auth,
		posts,
		ui
	}
});
