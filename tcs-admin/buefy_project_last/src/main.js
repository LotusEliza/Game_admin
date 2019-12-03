import Vue from 'vue'
import * as VeeValidate from 'vee-validate';
import App from './App.vue'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import VueRouter from 'vue-router'
import Routes from './routes'
import Vuex from 'vuex';
import {store} from './store/store'
import JQuery from 'jquery'
import moment from 'moment'

Vue.use(Vuex);
Vue.use(VueRouter);
Vue.use(Buefy);
Vue.use(VeeValidate, {
  events: '',
    errorBagName: 'vErrors',
});

window.$ = JQuery;

/* eslint-disable no-new */

//*************AXIOS**********************
window.axios = require("axios");
window.axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
window.axios.defaults.baseURL = "http://localhost:3000/";
window.axios.defaults.headers.Authorization = "Bearer NOTOKEN";

const router = new VueRouter({
  routes: Routes
});

Vue.filter('formatDate', function(value) {
  if (value) {
    return moment(String(value)).format('YYYY-MM-DD');
  }
});

Vue.config.productionTip = false;

new Vue({
  render: h => h(App),
  store: store,
  router: router,
}).$mount('#app');
