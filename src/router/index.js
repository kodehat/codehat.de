import Vue from "vue";
import Router from "vue-router";
import AboutMe from "../pages/AboutMe";
import Projects from "../pages/Projects";
import Contact from "../pages/Contact";
import Imprint from "../pages/Imprint";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/",
      name: "AboutMe",
      component: AboutMe
    },
    {
      path: "/projects",
      name: "Projects",
      component: Projects
    },
    {
      path: "/contact",
      name: "Contact",
      component: Contact
    },
    {
      path: "/imprint",
      name: "Imprint",
      component: Imprint
    }
  ]
});
