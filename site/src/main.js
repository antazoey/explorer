import Vue from "vue"
import App from "./app"
import Vuelidate from "vuelidate"

// sync store and router...
import { sync } from "vuex-router-sync"
import store from "./store/index.js"
import router from "./router/index.js"
sync(store, router)

Vue.use(Vuelidate)
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: "#app",
  router,
  render: h => h(App),
  store
})
