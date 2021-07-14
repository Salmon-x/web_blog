import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/antui.js'
import axios from 'axios'
import './assets/css/style.css'

axios.defaults.baseURL = 'http://localhost:8020'

// 请求拦截器
axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})

Vue.prototype.$http = axios
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')