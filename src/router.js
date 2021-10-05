import Vue from "vue";
import Router from "vue-router";
import Home from "./components/Home.vue";
import Materials from "./components/Materials.vue";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/",
      component: Home
    },
    {
      path: "/materials",
      component: Materials
    }
  ]
})