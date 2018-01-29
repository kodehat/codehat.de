<template>
  <div class="locale-switcher">
    <!-- {{ initLocale() }} -->
    <b-dropdown hoverable>
        <button class="button is-info" slot="trigger">
            <span>Select language</span>
            <b-icon icon="menu-down"></b-icon>
        </button>

        <b-dropdown-item
        class="locale-link"
        v-for="locale in locales"
        :key="locale.id"
        @click="setLocale(locale)"
        href="#">
        {{ getLanguageString(locale) }}
        </b-dropdown-item>
    </b-dropdown>
  </div>
</template>

<script>
import Vue from "vue";
// Restore locale from cookie, if it was set
import VueCookie from "vue-cookie";
Vue.use(VueCookie);

const localeStrings = {
  en: "English",
  de: "Deutsch"
};

Vue.config.lang = VueCookie.get("locale") || "en";

console.log(
  "Locale from cookie = " +
    Vue.config.lang +
    ": language = " +
    localeStrings[Vue.config.lang]
);

export default {
  name: "language-switcher",
  props: {
    locales: {
      type: Array,
      default: ["en"]
    },
    showFull: Boolean,
    dropup: Boolean,
    locLabel: {
      type: String,
      default: ""
    }
  },
  data: function() {
    return {
      activeLocale: Vue.config.lang
    };
  },
  computed: {
    dropdownLbl: function() {
      return this.locLabel ? this.locLabel : this.activeLocale;
    }
  },
  methods: {
    setLocale: function(locale) {
      Vue.config.lang = locale;
      this.activeLocale = locale;
      this.$cookie.set("locale", locale);
      this.$i18n.locale = Vue.config.lang;
      console.log(
        "New locale = " +
          Vue.config.lang +
          ": language = " +
          localeStrings[Vue.config.lang]
      );
    },
    getLanguageString: function(locale) {
      return this.showFull ? localeStrings[locale] : locale;
    }
  }
};
</script>
