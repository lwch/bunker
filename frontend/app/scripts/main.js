import { createApp } from 'vue';
import { router } from '@/router';
import App from '@/views/layouts/app';
import BalmUI from 'balm-ui'; // Official Google Material Components
import BalmUIPlus from 'balm-ui/dist/balm-ui-plus'; // BalmJS Team Material Components

const app = createApp(App);

app.use(BalmUI); // Mandatory
app.use(BalmUIPlus); // Optional
app.use(router);

app.mount('#app');
