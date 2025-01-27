import path from 'path';
import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        path: '/',
        name: 'Home',
        component: () => import('@/views/Home.vue'),
    },
    {
        path: '/auth',
        component: () => import('@/views/Auth/Layout.vue'),
        children: [
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
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;
