<template>
    <div class="hero-head">
        <nav class="navbar is-primary">
            <div class="container">
                <div class="navbar-brand">
                    <router-link to="/" class="navbar-item">
                        <h3 class="title is-3 special-font small">CodeHat</h3>
                    </router-link>
                    <a class="navbar-item is-hidden-desktop" href="https://github.com/Pixelhash" target="_blank">
                        <span class="icon" style="color: lightgrey;">
                            <i class="fa fa-lg fa-github"></i>
                        </span>
                    </a>
                    <a class="navbar-item is-hidden-desktop" href="https://twitter.com/CodeHat" target="_blank">
                        <span class="icon" style="color: #55acee;">
                            <i class="fa fa-lg fa-twitter"></i>
                        </span>
                    </a>
                    <span class="navbar-burger burger" data-target="navbarTop">
                        <span></span>
                        <span></span>
                        <span></span>
                    </span>
                </div>
                <div id="navbarTop" class="navbar-menu">
                    <div class="navbar-end">
                        <a class="navbar-item is-hidden-touch" href="https://github.com/Pixelhash" target="_blank">
                            <span class="icon" style="color: lightgrey;">
                                <i class="fa fa-lg fa-github"></i>
                            </span>
                        </a>
                        <a class="navbar-item is-hidden-touch" href="https://twitter.com/CodeHat" target="_blank">
                            <span class="icon" style="color: #55acee;">
                                <i class="fa fa-lg fa-twitter"></i>
                            </span>
                        </a>
                        <locale-switcher :locales="['en', 'de']" :show-full="true" ></locale-switcher>
                    </div>
                </div>
            </div>
        </nav>
  </div>
</template>

<script>
document.addEventListener("DOMContentLoaded", function() {
  // Get all "navbar-burger" elements
  var $navbarBurgers = Array.prototype.slice.call(
    document.querySelectorAll(".navbar-burger"),
    0
  );

  // Check if there are any navbar burgers
  if ($navbarBurgers.length > 0) {
    // Add a click event on each of them
    $navbarBurgers.forEach(function($el) {
      $el.addEventListener("click", function() {
        // Get the target from the "data-target" attribute
        var target = $el.dataset.target;
        var $target = document.getElementById(target);

        // Toggle the class on both the "navbar-burger" and the "navbar-menu"
        $el.classList.toggle("is-active");
        $target.classList.toggle("is-active");
      });
    });
  }
});

import LocaleSwitcher from "./LocaleSwitcher";

export default {
  name: "top-navbar",
  components: {
    LocaleSwitcher
  }
};
</script>

<style scoped>
.special-font.small {
  font-weight: 100;
}
</style>
