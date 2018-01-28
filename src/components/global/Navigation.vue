<template>
    <div class="navigation">
        <div class="container">
            <div class="tabs is-boxed is-centered is-medium">
                <ul>
                    <li>
                        <router-link to="/">
                            <span class="icon is-small"><i class="fa fa-home"></i></span>
                            <span>Home</span>
                        </router-link>
                    </li>
                    <li>
                        <router-link to="/projects">
                            <span class="icon is-small"><i class="fa fa-briefcase"></i></span>
                            <span>Projects</span>
                        </router-link>
                    </li>
                    <li>
                        <router-link to="/contact">
                            <span class="icon is-small"><i class="fa fa-address-book"></i></span>
                            <span>Contact</span>
                        </router-link>
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