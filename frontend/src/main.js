import Vue from 'vue';
import Toasted from 'vue-toasted';

import store from './store';
import App from './App.vue';

import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'

Vue.config.productionTip = false

Vue.use(Toasted);

new Vue({
  store,
  render: h => h(App),
}).$mount('#app')
