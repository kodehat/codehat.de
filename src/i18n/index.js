import Vue from "vue";
import VueI18n from "vue-i18n";
import VueCookie from "vue-cookie";

// Import locale files
import en from "../assets/locales/en";
import de from "../assets/locales/de";

Vue.use(VueI18n);

// Ready translated locale messages
const messages = {
  en: en,
  de: de
};

export default new VueI18n({
  locale: VueCookie.get("locale") || "en",
  messages
});
