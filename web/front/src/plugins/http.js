import Vue from 'vue'
import axios from 'axios'

// axios请求地址

axios.defaults.baseURL = 'http://localhost:8020'
Vue.prototype.$http = axios
