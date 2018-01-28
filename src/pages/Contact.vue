<template>
    <div class="contact">
        <section class="section">
            <div class="container">
            <div class="content is-medium">
                <h2 class="title is-2">Contact</h2>
                <h4 class="subtitle is-4">You can contact me in many different ways</h4>
                <hr>
                <p>
                    <ul class="li-nobullets">
                        <li>
                            <span class="icon icon-contact twitter">
                                <i class="fa fa-twitter"></i>
                            </span> Twitter: <a href="https://twitter.com/CodeHat">@CodeHat</a>
                        </li>
                        <li>
                            <span class="icon icon-contact telegram">
                                <i class="fa fa-telegram"></i>
                            </span> Telegram: <a href="http://telegram.me/CodeHat">@CodeHat</a>
                        </li>
                        <li>
                            <span class="icon icon-contact">
                                <i class="fa fa-github"></i>
                            </span> GitHub: <a href="https://github.com/Pixelhash">PixelHash</a>
                        </li>
                        <li>
                            <span class="icon icon-contact">
                                <i class="fa fa-envelope"></i>
                            </span> Email: <a href="mailto:dev@codehat.de">dev [ett] codehat.de</a>
                        </li>
                        <li>
                            <span class="icon icon-contact keybase">
                                <i class="fa fa-user-secret"></i>
                            </span> Keybase: <a href="https://keybase.io/codehat">CodeHat</a>
                        </li>
                    </ul>
                </p>
            </div>
            </div>
        </section>
    </div>
</template>

<script>
    export default {
        name: 'contact',
    }
</script>

<style>
    .icon-contact {
        margin-top: .19em;
        margin-right: .5em;
    }

    .icon-contact.telegram {
        color: #279EDA;
    }

    .icon-contact.twitter {
        color: #55ACEE;
    }

    .icon-contact.keybase {
        color: #FF6F21;
    }
</style>
