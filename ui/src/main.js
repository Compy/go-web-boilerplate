import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import Welcome from './views/Welcome'
import Login from './views/Login'
import store from './store'
import vuetify from './plugins/vuetify'

Vue.use(VueRouter)
const routes = [
  {
    path: "/",
    name: "Welcome",
    component: Welcome,
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
  }
]

Vue.config.productionTip = false

const router = new VueRouter({
  mode: "history",
  base: "/",
  routes,
})

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App)
}).$mount("#app")