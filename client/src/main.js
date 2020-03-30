import Vue from 'vue';

import Antd from 'ant-design-vue'; // Ant Desing Vue ui kütüphanesi kullanıldı
import 'ant-design-vue/dist/antd.less'; // Ant design style
import TextareaAutosize from 'vue-textarea-autosize'; // Projede kullanılacak custom textarea
import moment from 'moment'; // Tarih formatlama için moment.js kullanıldı

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

// Moment için türkçe dil seçeneği ayarı yapılıyor
moment.locale('tr');
Vue.prototype.$moment = moment;

window.vm = new Vue({
	router,
	store,
	render: (h) => h(App)
}).$mount('#app');
