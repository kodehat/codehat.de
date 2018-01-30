<template>
    <div class="about-me">
        <section class="section">
            <div class="container">
                <div class="content is-medium">
                    <h2 class="title is-2">{{ $t("tabs.about") }}</h2>
                    <hr>

                    <p>
                        <strong>HayHo!</strong> I'm CodeHat and you're currently viewing my website.<br>
                        I like learning new technologies like <a href="https://vuejs.org/" target="_blank">Vue.js</a> or
                        <a href="https://kotlinlang.org/" target="_blank">Kotlin</a>, but primary I focus on <strong>Java</strong>.
                        That's the reason why I'm maintaining some plugins on <a href="https://www.spigotmc.org/resources/signcolors.6135/">SpigotMC.org</a>.<br><br>

                        You can take a look at my other <strong>projects</strong> by visiting the <strong>"Projects"</strong> tab above.<br><br>

                        To have something special, which changes every time you visit my website, here is a random <strong>programming quote</strong>:
                        <blockquote>
                            {{ quote.message }}<br>
                            <small>{{ quote.author }} &middot; <a :href="quote.url">Source</a></small>
                        </blockquote>
                    </p>
                </div>
            </div>
        </section>
    </div>
</template>

<script>
import axios from "axios";

const quote = {
  message: "...",
  author: "...",
  url: "..."
};

export default {
  name: "about-me",
  data() {
    return {
      quote
    };
  },
  created() {
    axios
      .get("http://quotes.stormconsultancy.co.uk/random.json")
      .then(response => {
        this.quote.message = response.data.quote;
        this.quote.author = response.data.author;
        this.quote.url = response.data.permalink;
      })
      .catch(error => {
        console.log(error);
      });
  }
};
</script>
