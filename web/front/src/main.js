import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import './plugins/http'
import moment from 'moment'

Vue.filter('dataformat', function(indate,outdate){
	return moment(indate).format(outdate)
})

Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
