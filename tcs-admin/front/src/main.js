import Vue from 'vue'
import App from './App.vue'
import router from './router'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import {store} from './store/store'
import VueMq from 'vue-mq'
import JQuery from 'jquery'
window.$ = JQuery


import moment from 'moment'

Vue.filter('formatDate', function(value) {
  if (value) {
    return moment(String(value)).format('YYYY-MM-DD')
  }
})


Vue.use(BootstrapVue, {
  breakpoints: [`xs`, 'sm', 'md', 'lg', 'xl', 'xxl']
})

Vue.use(BootstrapVue)

Vue.config.productionTip = false

window.axios = require("axios");
window.axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
window.axios.defaults.baseURL = "http://localhost:3000/";
window.axios.defaults.headers.Authorization = "Bearer NOTOKEN";


Vue.use(VueMq, {
  breakpoints: {
    sm: 765,
    md: 990,
    lg: Infinity,
  }
})

new Vue({
  router,
  store: store,
  render: function (h) { return h(App) }
}).$mount('#app')
