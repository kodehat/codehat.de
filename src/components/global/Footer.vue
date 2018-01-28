<template>
  <div class="footer">
    <footer class="footer">
        <div class="container">
            <div class="columns">
                <div class="column is-third has-text-centered">
                    <ul class="ch-hor-list">
                        <!--<li><a href=".">Home</a></li>-->
                        <li><a href="imprint.html">Imprint</a></li>
                        <li><a href="disclaimer.html">Disclaimer</a></li>
                        <!--<li><a href="contact.html">Contact</a></li>-->
                    </ul>
                </div>
                <div class="column is-third has-text-centered">
                    <p>&copy; CodeHat {{ year }}</p>
                </div>
                <div class="column is-third has-text-centered">
                    This website is licensed <a href="https://choosealicense.com/licenses/mit/">MIT</a>.
                    <br><br>
                    <a class="icon" href="https://github.com/PixelHash/website">
                        <i class="fa fa-github"></i>
                    </a>
                </div>
            </div>
        </div>
    </footer>
  </div>
</template>

<script>
export default {
  name: "footer",
  data() {
    return {
      year: new Date().getFullYear()
    };
  }
};
</script>
