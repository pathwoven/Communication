import path from 'path';
import {createRouter, createWebHistory} from 'vue-router'

const routes = [
	{
		path: '/',
		redirect: { name: 'Chat' },
	},
	{
		path: '/auth',
		component: () => import('@/views/Auth/Layout.vue'),
		children: [
			{
				path: '',
				redirect: { name: 'Login' },
			},
			{
				path: 'login',
				name: 'Login',
				component: () => import('@/views/Auth/Login.vue'),
			},
			{
				path: 'register',
				name: 'Register',
				component: () => import('@/views/Auth/Register.vue'),
			},
			{
				path: 'forget',
				name: 'Forget',
				component: () => import('@/views/Auth/Forget.vue'),
			},
		]
	},
	{
		path: '/chat',
		component: () => import('@/views/Chat.vue'),
	},
	{
		path: '/test',   // debug todo
		component: () => import('@/views/Test.vue'),
	}
];

const router = createRouter({
	history: createWebHistory(),
	routes
});

export default router;
