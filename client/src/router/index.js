import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';
import Login from '../views/Login.vue';
import Register from '../views/Register.vue';
import Posts from '../components/Posts.vue';
import Settings from '../components/Settings.vue';
import Profile from '../components/Profile.vue';

Vue.use(VueRouter);

const routes = [
	{
		path: '/',
		name: 'home',
		component: Home,
		meta: { requiresAuth: false },
		children: [
			{
				path: '/profile',
				name: 'profile',
				component: Profile,
				meta: { requiresAuth: false }
			},
			{
				path: '/settings',
				name: 'settings',
				component: Settings,
				meta: { requiresAuth: false }
			}
		]
	},
	{
		path: '/login',
		name: 'login',
		component: Login
	},
	{
		path: '/register',
		name: 'register',
		component: Register
	}
];

const router = new VueRouter({
	mode: 'history',
	base: process.env.BASE_URL,
	routes
});

router.beforeEach((to, from, next) => {
	const token = JSON.parse(localStorage.getItem('token'));
	if (to.meta.requiresAuth) {
		if (!token) {
			next({ name: 'login' });
		} else {
			next();
		}
	} else {
		next();
	}
});

export default router;
