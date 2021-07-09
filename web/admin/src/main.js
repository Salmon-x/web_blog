import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/antui.js'
import axios from 'axios'
import './assets/css/style.css'

axios.defaults.baseURL = 'http://localhost:8020'

Vue.prototype.$http = axios
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')