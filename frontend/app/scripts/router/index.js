import { createRouter, createWebHashHistory } from 'vue-router';
import Dashboard from '@/views/dashboard';
import Users from '@/views/users';

const dashboard = {
    path: '/',
    component: Dashboard
};

const users = {
    path: '/users',
    component: Users
}

const routes = [
    dashboard,
    users
];

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
});

export default router;