import request from '@/request';

const agent = {
    list() {
        return request({
            url: '/api/agents',
            method: 'GET'
        });
    }
};

export default agent;