import { createRouter, createWebHashHistory } from 'vue-router';
import Dashboard from '@/views/dashboard';
import Users from '@/views/users';

const dashboard = {
    path: '/',
    component: Dashboard,
    icon: 'dashboard',
    name: '仪表盘'
};

const users = {
    path: '/users',
    component: Users,
    icon: 'person',
    name: '用户列表'
};

const devices = {
    path: '/devices',
    component: {template: 'devices'},
    icon: 'computer',
    name: '资产列表'
};

const tags = {
    path: '/tags',
    component: {template: 'tags'},
    icon: 'tag',
    name: '标签管理'
};

const groups = [
    {
        name: '概览',
        pages: [dashboard]
    }, {
        name: '用户管理',
        pages: [users]
    }, {
        name: '资产管理',
        pages: [devices, tags]
    }
]

const routes = [
    dashboard,
    users,
    devices,
    tags
];

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
});

export {
    groups,
    router
};