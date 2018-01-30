import Vue from "vue";
import VueCookie from "vue-cookie";
import VueMomentJs from "vue-momentjs";
import router from "./router";
import i18n from "./i18n";
// Import German language with 'src' identifier to make it work!
import "moment/src/locale/de";
import moment from "moment";
import App from "./App";
import Buefy from "buefy";
import "buefy/lib/buefy.css";

const locale = VueCookie.get("locale") || "en";
moment.locale(locale);

Vue.use(Buefy);
Vue.use(VueMomentJs, moment);

Vue.config.productionTip = false;

new Vue({
  el: "#app",
  router,
  i18n,
  render: h => h(App),
  template: "<App/>",
  components: { App }
});
