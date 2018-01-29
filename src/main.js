import Vue from "vue";
import VueCookie from "vue-cookie";
import VueMomentJs from "vue-momentjs";
import router from "./router";
import i18n from "./i18n";
import moment from "moment";
import "moment/locale/de";
import App from "./App";
import Buefy from "buefy";
import "buefy/lib/buefy.css";

Vue.use(Buefy);

Vue.config.productionTip = false;

const locale = VueCookie.get("locale") || "en";
console.log("Current locale: " + locale);
moment.locale(locale);
console.log("Loaded locale: " + moment.locale());

Vue.use(VueMomentJs, moment);

new Vue({
  el: "#app",
  router,
  i18n,
  render: h => h(App),
  template: "<App/>",
  components: { App }
});
