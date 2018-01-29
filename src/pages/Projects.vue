<template>
    <div class="projects">
        <section class="section">
        <div class="container has-text-centered">
            <h1 class="title">Java Projects</h1>
            <h3 class="subtitle">My current Java projects (Open Source)</h3>
            <hr>
            <div class="columns">
                <div class="column is-one-third">
                    <div class="card">
                        <div class="card-image">
                            <figure class="image is-4by3">
                                <img src="../assets/images/java_signcolors.png" alt="SignColors">
                            </figure>
                        </div>
                        <div class="card-content">
                            <div class="content">
                                SignColors allows players on your <a>#Minecraft</a> server to add colored text to signs.
                                <br>
                                <small>Spigot Plugin - <span id="signcolors">{{ signcolorsText }}</span> for Spigot 1.11.2<br>
                                Updated at <span id="signcolors_updatetime">{{ signcolorsDate }}</span></small>
                            </div>
                        </div>
                        <footer class="card-footer">
                            <a class="card-footer-item" href="https://www.spigotmc.org/resources/signcolors.6135">
                                <span class="icon is-small icon-project">
                                <i class="fa fa-globe"></i>
                            </span> Project Website
                            </a>
                            <a class="card-footer-item" href="https://github.com/Pixelhash/SignColors">
                                <span class="icon is-small icon-project">
                                <i class="fa fa-github"></i>
                            </span> Source
                            </a>
                        </footer>
                    </div>
                </div>
                <div class="column is-one-third">
                    <div class="card">
                        <div class="card-image">
                            <figure class="image is-4by3">
                                <img src="../assets/images/java_playersupport.png" alt="PlayerSupport">
                            </figure>
                        </div>
                        <div class="card-content">
                            <div class="content">
                                PlayerSupport allows players on your <a>#Minecraft</a> server to request help with a simple command.
                                <br>
                                <small>Spigot Plugin - <span id="playersupport">{{ playersupportText }}</span> for Spigot 1.11.2<br>
                                Updated at <span id="playersupport_updatetime">{{ playersupportDate }}</span></small>
                            </div>
                        </div>
                        <footer class="card-footer">
                            <a class="card-footer-item" href="https://www.spigotmc.org/resources/playersupport.11255">
                                <span class="icon is-small icon-project">
                                <i class="fa fa-globe"></i>
                            </span> Project Website
                            </a>
                            <a class="card-footer-item" href="https://github.com/Pixelhash/PlayerSupport">
                                <span class="icon is-small icon-project">
                                <i class="fa fa-github"></i>
                            </span> Source
                            </a>
                        </footer>
                    </div>
                </div>
                <div class="column is-one-third">
                    <div class="card">
                        <div class="card-image">
                            <figure class="image is-4by3">
                                <img src="../assets/images/java_mcwrapper.png" alt="PlayerSupport">
                            </figure>
                        </div>
                        <div class="card-content">
                            <div class="content">
                                A simple <a>#Minecraft</a> Server Wrapper (for Spigot, Bukkit, ...), which provides a console for your browser
                                <br>
                                <small>Wrapper - Not yet ready for production!</small>
                            </div>
                        </div>
                        <footer class="card-footer">
                            <a class="card-footer-item is-disabled" href="https://www.spigotmc.org/resources/playersupport.11255">
                                <span class="icon is-small icon-project">
                                <i class="fa fa-globe"></i>
                            </span>
                                <del>Project Website</del>
                            </a>
                            <a class="card-footer-item" href="https://github.com/Pixelhash/mcwrapper">
                                <span class="icon is-small icon-project">
                                <i class="fa fa-github"></i>
                            </span> Source
                            </a>
                        </footer>
                    </div>
                </div>
            </div>
        </div>
    </section>
    </div>
</template>

<script>
import axios from "axios";
import moment from "moment";
import "moment/locale/de";
import VueCookie from "vue-cookie";

console.log("Current locale: " + VueCookie.get("locale") || "en");
moment.locale(VueCookie.get("locale") || "en");

export default {
  name: "projects",
  data() {
    return {
      signcolorsText: {},
      signcolorsDate: {},
      playersupportText: {},
      playersupportDate: {}
    };
  },
  created() {
    // Set 'SignColors' version and update date
    axios
      .get("https://api.codehat.de/plugin/1")
      .then(response => {
        this.signcolorsText = "v" + response.data.version;
        this.signcolorsDate = moment(response.data.updated_at).format(
          "Do MMM YYYY, HH:mm"
        );
      })
      .catch(error => {
        console.log(error);
      });

    // Set 'PlayerSupport' version and update date
    axios
      .get("https://api.codehat.de/plugin/2")
      .then(response => {
        this.playersupportText = "v" + response.data.version;
        this.playersupportDate = moment(response.data.updated_at).format(
          "Do MMM YYYY, HH:mm"
        );
      })
      .catch(error => {
        console.log(error);
      });
  }
};
</script>
