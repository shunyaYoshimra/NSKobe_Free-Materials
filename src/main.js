import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from "./router"
import '@mdi/font/css/materialdesignicons.css'
Vue.config.productionTip = false

console.log("1");
new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
