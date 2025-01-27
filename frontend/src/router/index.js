import path from 'path';
import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        path: '/',
        name: 'Home',
        component: () => import('@/views/Home.vue'),
    },
    {
        path: '/login',
        component: () => import('@/views/Login/Layout.vue'),
        children: [
            {
                path: '',
                name: 'Login',
                component: () => import('@/views/Login/Login.vue'),
            },
            {
                path: 'register',
                name: 'Register',
                component: () => import('@/views/Login/Register.vue'),
            },
            {
                path: 'forget',
                name: 'Forget',
                component: () => import('@/views/Login/Forget.vue'),
            },
        ]
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;
