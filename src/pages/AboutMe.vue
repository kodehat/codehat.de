<template>
  <section class="section">
      <div class="container">
          <div class="content is-medium">
              <h2 class="title is-2">{{ $t("tabs.about") }}</h2>
              <hr>
              <p>
                <span v-html="$t('aboutMe.firstParagraph')"></span>
                <span v-html="$t('aboutMe.secondParagraph')"></span>
                <span v-html="$t('aboutMe.thirdParagraph')"></span>
                <span v-html="$t('aboutMe.fourthParagraph')"></span>
              </p>
              <blockquote>
                  {{ quote.message }}<br>
                  <small>- {{ quote.author }} &middot; <a :href="quote.url">{{ $t("aboutMe.quoteSource") }}</a></small>
              </blockquote>
          </div>
      </div>
  </section>
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
      .get(
        "https://cors-anywhere.herokuapp.com/http://quotes.stormconsultancy.co.uk/random.json"
      )
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
