import Vue from "vue";
import router from "./router";
import i18n from "./i18n";
import App from "./App";
import Buefy from "buefy";
import "buefy/lib/buefy.css";

Vue.use(Buefy);

Vue.config.productionTip = false;
Vue.config.lang = "en";

new Vue({
  el: "#app",
  router,
  i18n,
  render: h => h(App),
  template: "<App/>",
  components: { App }
});
