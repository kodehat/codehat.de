<template>
    <div class="contact">
        <section class="section">
            <div class="container">
                <div class="content is-medium">
                    <h2 class="title is-2">{{ $t("tabs.contact") }}</h2>
                    <hr>
                    <p><span v-html="$t('contact.note')"></span></p>
                    <ul class="li-nobullets">
                        <li>
                            <span class="icon icon-contact">
                                <i class="fa fa-envelope"></i>
                            </span> Email: &#100;&#101;&#118;&#064;&#099;&#111;&#100;&#101;&#104;&#097;&#116;&#046;&#100;&#101;
                        </li>
                        <li>
                            <span class="icon icon-contact telegram">
                                <i class="fa fa-telegram"></i>
                            </span> Telegram: <a href="http://telegram.me/CodeHat" target="_blank">@CodeHat</a>
                        </li>
                        <li>
                            <span class="icon icon-contact twitter">
                                <i class="fa fa-twitter"></i>
                            </span> Twitter: <a href="https://twitter.com/CodeHat" target="_blank">@CodeHat</a>
                        </li>
                        <!--<li>
                            <span class="icon icon-contact">
                                <i class="fa fa-github"></i>
                            </span> GitHub: <a href="https://github.com/Pixelhash">PixelHash</a>
                        </li>-->
                        <li>
                            <span class="icon icon-contact keybase">
                                <i class="fa fa-user-secret"></i>
                            </span> Keybase: <a href="https://keybase.io/codehat" target="_blank">CodeHat</a>
                        </li>
                    </ul>
                </div>
            </div>
        </section>
    </div>
</template>

<script>
export default {
  name: "contact"
};
</script>

<style>
.icon-contact {
  margin-top: 0.19em;
  margin-right: 0.5em;
}

.icon-contact.telegram {
  color: #279eda;
}

.icon-contact.twitter {
  color: #55acee;
}

.icon-contact.keybase {
  color: #ff6f21;
}
</style>
