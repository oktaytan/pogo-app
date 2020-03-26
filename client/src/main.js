import Vue from 'vue';
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.less';
import infiniteScroll from 'vue-infinite-scroll';
import TextareaAutosize from 'vue-textarea-autosize';
import moment from 'moment';

import App from './App.vue';
import router from './router';
import store from './store';

Vue.config.productionTip = false;
Vue.use(Antd);
Vue.use(infiniteScroll);
Vue.use(TextareaAutosize);

Vue.prototype.$moment = moment;

window.vm = new Vue({
	router,
	store,
	render: (h) => h(App)
}).$mount('#app');
