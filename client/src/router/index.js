import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from '../views/Home.vue';
import Login from '../views/Login.vue';
import Logout from '../views/Logout.vue';
import Register from '../views/Register.vue';
import NotFound from '../views/NotFound.vue';
import Settings from '../components/Settings.vue';
import Profile from '../components/Profile.vue';
import PostDetail from '../components/PostDetail.vue';

Vue.use(VueRouter);

const routes = [
	{
		path: '/',
		name: 'Anasayfa',
		component: Home,
		meta: { requiresAuth: true },
		children: [
			{
				path: '/posts/:id',
				name: 'Detaylar',
				component: PostDetail,
				meta: { requiresAuth: true }
			},
			{
				path: '/profile',
				name: 'Profilim',
				component: Profile,
				meta: { requiresAuth: true }
			},
			{
				path: '/settings',
				name: 'Ayarlar',
				component: Settings,
				meta: { requiresAuth: true }
			}
		]
	},
	{
		path: '/login',
		name: 'login',
		component: Login,
		meta: { noLogin: true }
	},
	{
		path: '/logout',
		name: 'logout',
		component: Logout,
		meta: { requiresAuth: true }
	},
	{
		path: '/register',
		name: 'register',
		component: Register,
		meta: { noLogin: true }
	},
	{
		path: '*',
		name: 'notFound',
		component: NotFound,
		meta: { noLogin: true, requiresAuth: false }
	}
];

const router = new VueRouter({
	mode: 'history',
	base: process.env.BASE_URL,
	routes
});

// Route' lar için gerekli authentication kontrolleri yapılıyor
router.beforeEach((to, from, next) => {
	const token = localStorage.getItem('pogo_user')
		? JSON.parse(localStorage.getItem('pogo_user')).token
		: null;
	if (to.meta.requiresAuth) {
		if (!token) {
			next({ name: 'login' });
		} else {
			next();
		}
	} else if (to.meta.noLogin) {
		if (token) {
			if (to.name === 'login' || to.name === 'register') {
				next({ name: 'Anasayfa' });
			} else {
				next();
			}
		} else {
			next();
		}
	}
});

export default router;
