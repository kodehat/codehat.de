<template>
    <div class="footer">
        <footer class="footer">
            <div class="container">
                <div class="columns">
                    <div class="column is-third has-text-centered">
                        <ul class="ch-hor-list">
                            <li>
                                <router-link to="/imprint">Imprint</router-link>
                            </li>
                            <li>
                                <a href="disclaimer.html">Disclaimer</a>
                            </li>
                        </ul>
                    </div>
                    <div class="column is-third has-text-centered">
                        <p>&copy; CodeHat {{ year }}</p>
                    </div>
                    <div class="column is-third has-text-centered">
                        Source code licensed
                        <a href="https://opensource.org/licenses/mit-license.php">MIT</a>.
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
  name: "my-footer",
  data() {
    return {
      year: new Date().getFullYear()
    };
  }
};
</script>
