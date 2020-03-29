import Vue from 'vue';
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.less';
import TextareaAutosize from 'vue-textarea-autosize';
import moment from 'moment';

import App from './App.vue';
import router from './router';
import store from './store';

Vue.config.productionTip = false;
Vue.use(Antd);
Vue.use(TextareaAutosize);

// Filters
Vue.filter('userAvatar', function(value) {
	if (!value) return;
	return value.slice(0, 2).toUpperCase();
});

moment.locale('tr');
Vue.prototype.$moment = moment;

window.vm = new Vue({
	router,
	store,
	render: (h) => h(App)
}).$mount('#app');
