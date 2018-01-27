<template>
  <div id="app">
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container has-text-centered">
          <h1 class="title is-1 ch-pixelfont">
            CodeHat
          </h1>
          <figure class="image is-128x128 ch-logo animated bounceInDown">
            <img src="./assets/images/avatar_128x128.png" alt="CodeHat Avatar">
          </figure>
          <h3 class="subtitle is-3 ch-margin-top">
            💻 Science Student &#8722; Developer
          </h3>
          <h5 class="subtitle is-5">
            Living in the middle of nowhere in Germany
          </h5>
        </div>
      </div>
      <div class="hero-foot">
        <navigation></navigation>
      </div>
    </section>

    <router-view></router-view>
  </div>
</template>

<script>
  import Navigation from './components/global/Navigation';

  export default {
      name: 'app',
      components: {
          Navigation
      }
  }
</script>

<style lang="scss">
  // Import Bulma's core
  @import "~bulma/sass/utilities/_all";

  // Set your colors
  $primary: #2a231a;
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
  $link: $primary;
  $link-invert: $primary-invert;
  $link-focus-border: $primary;

  // Import Bulma and Buefy styles
  @import "~bulma";
  @import "~buefy/src/scss/buefy";

  .ch-logo {
      display: block;
      margin: 0 auto;
  }
</style>