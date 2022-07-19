import Axios from 'axios';

const axios = Axios.create({
    timeout: 1000
});

axios.interceptors.response.use(
response => {
    return response.data;
}, error => {
    console.error(error);
    return Promise.reject(error);
});

export default axios;