<template>
  <div id="app">
    <section class="hero is-primary">
      <top-navbar></top-navbar>
      <div class="hero-body">
        <div class="container has-text-centered">
          <!--<h1 class="title is-1 special-font big">
            CodeHat
          </h1>-->
          <figure class="image is-128x128 img-center animated bounceInDown">
            <img src="./assets/images/codehat_logo.png" alt="CodeHat Avatar">
          </figure>
          <h4 class="title is-4 some-margin-top">CodeHat</h4>
          <h4 class="subtitle is-4">
            {{ $t("messages.job") }} &#9679; {{ $t("messages.education") }}
          </h4>
          <h6 class="subtitle is-6">
            <a href="https://www.fh-luebeck.de/">
              🏫 FH-L&uuml;beck</a> &#9679; {{ $t("messages.origin") }}
          </h6>
        </div>
      </div>
      <div class="hero-foot">
        <my-navigation></my-navigation>
      </div>
    </section>

    <router-view></router-view>

    <my-footer></my-footer>
  </div>
</template>

<script>
import MyNavigation from "./components/global/Navigation";
import TopNavbar from "./components/global/TopNavbar";
import MyFooter from "./components/global/Footer";

export default {
  name: "app",
  components: {
    MyNavigation,
    TopNavbar,
    MyFooter
  }
};
</script>

<style lang="scss">
// Import additional fonts
@import url("https://fonts.googleapis.com/css?family=Supermercado+One");
@font-face {
  font-family: "m5x7";
  src: url("./assets/fonts/m5x7.ttf") format("truetype");
}
$pixel-font: "m5x7", sans-serif;

// Import Bulma's core
@import "../node_modules/bulma/sass/utilities/_all";

// Set your colors
$primary: #d0a465;
$primary-invert: findColorInvert($primary);
$twitter: #4099ff;
$twitter-invert: findColorInvert($twitter);

// Setup $colors to use as bulma classes (e.g. 'is-twitter')
$colors: (
  "white": ($white, $black),
  "black": ($black, $white),
  "light": ($light, $light-invert),
  "dark": ($dark, $dark-invert),
  "primary": ($primary, $primary-invert),
  "info": ($info, $info-invert),
  "success": ($success, $success-invert),
  "warning": ($warning, $warning-invert),
  "danger": ($danger, $danger-invert),
  "twitter": ($twitter, $twitter-invert)
);

// Links
$link: #513000;
$link-invert: $primary-invert;
$link-focus-border: $primary;

// Navbar
$navbar-dropdown-background-color: $primary;

// Import Bulma and Buefy styles
@import "../node_modules/bulma/bulma";
@import "../node_modules/buefy/src/scss/buefy";

.img-center {
  display: block;
  margin: 0 auto;
}

.icon-project {
  margin-right: 0.25rem;
  margin-top: 0.15em;
}

.special-font {
  font-family: "Supermercado One", cursive;
  //font-family: $pixel-font;
}

.special-font.big {
  font-size: 96px !important;
  font-weight: 200 !important;
}

.li-nobullets {
  list-style: none !important;
}

.some-margin-top {
  margin-top: 2rem;
}

.emoji-fix {
  padding-top: 0.2em;
  padding-right: 0.2em;
  padding-left: 0.2em;
}
</style>
