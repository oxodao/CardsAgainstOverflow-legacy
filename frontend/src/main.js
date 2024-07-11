import Vue from 'vue';
import VTooltip from 'v-tooltip';
import App from './App.vue';
import store from './store';

Vue.use(VTooltip);

Vue.config.productionTip = false;

new Vue({
    store,
    render: h => h(App)
}).$mount('#app');
