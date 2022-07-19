import request from '@/request';

const shell = {
    new(id) {
        return request({
            url: '/api/shell/'+id,
            method: 'POST'
        });
    }
};

export default shell;