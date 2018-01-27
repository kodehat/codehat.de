<template>
  <div id="app">
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container has-text-centered">
          <h1 class="title is-1 special-font">
            CodeHat
          </h1>
          <figure class="image is-128x128 img-center animated bounceInDown">
            <img src="./assets/images/avatar_128x128.png" alt="CodeHat Avatar">
          </figure>
          <h3 class="subtitle is-3 some-margin-top">
            Software Developer &#8722; Computer Science Student
          </h3>
          <h5 class="subtitle is-5">
            <a href="https://www.fh-luebeck.de/">FH-L&uuml;beck</a> - Germany
          </h5>
        </div>
      </div>
      <div class="hero-foot">
        <Navigation></Navigation>
      </div>
    </section>

    <router-view></router-view>

    <Footer></Footer>
  </div>
</template>

<script>
  import Navigation from './components/global/Navigation';
  import Footer from './components/global/Footer';

  export default {
      name: 'app',
      components: {
          Navigation,
          Footer,
      }
  }
</script>

<style lang="scss">
  // Import additional font
  @import url('https://fonts.googleapis.com/css?family=Supermercado+One');

  // Import Bulma's core
  @import "../node_modules/bulma/sass/utilities/_all";

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
  $link: #0f5f3e;
  $link-invert: $primary-invert;
  $link-focus-border: $primary;

  // Import Bulma and Buefy styles
  @import "../node_modules/bulma/bulma";
  @import "../node_modules/buefy/src/scss/buefy";

  .img-center {
      display: block;
      margin: 0 auto;
  }

  .icon-project {
    margin-right: .25rem;
    margin-top: .15em;
  }

  .special-font {
    font-family: 'Supermercado One', cursive;
    font-size: 96px !important;
    font-weight: 200 !important;
  }

  .li-nobullets {
    list-style: none !important;
  }

  .some-margin-top {
    margin-top: 2rem;
  }
</style>