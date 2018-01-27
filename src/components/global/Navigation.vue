<template>
    <div class="navigation">
        <div class="container">
            <div class="tabs is-boxed is-centered is-medium">
                <ul>
                    <li>
                        <router-link to="/">Home</router-link>
                    </li>
                    <li>
                        <router-link to="/contact" href="contact.html">Contact</router-link>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'navigation',
    }
</script>

<style>

</style>