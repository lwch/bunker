import { createRouter, createWebHashHistory } from 'vue-router';
import Dashboard from '@/views/dashboard';
import Users from '@/views/users';

const dashboard = {
    path: '/',
    component: Dashboard,
    icon: 'dashboard'
};

const users = {
    path: '/users',
    component: Users,
    icon: 'person'
};

const devices = {
    path: '/devices',
    component: '<template></template>',
    icon: 'computer'
};

const tags = {
    path: '/tags',
    component: '<template></template>',
    icon: 'tag'
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

export default {
    groups,
    router
};