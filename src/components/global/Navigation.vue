<template>
    <div class="navigation">
        <div class="container">
            <div class="tabs is-boxed is-centered is-medium">
                <ul>
                    <li>
                        <router-link to="/">
                            <span class="icon is-small"><i class="fa fa-user"></i></span>
                            <span>{{ $t("tabs.about") }}</span>
                        </router-link>
                    </li>
                    <li>
                        <router-link to="/projects">
                            <span class="icon is-small"><i class="fa fa-briefcase"></i></span>
                            <span>{{ $t("tabs.projects") }}</span>
                        </router-link>
                    </li>
                    <li>
                        <router-link to="/contact">
                            <span class="icon is-small"><i class="fa fa-address-book"></i></span>
                            <span>{{ $t("tabs.contact") }}</span>
                        </router-link>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  name: "navigation"
};
</script>

<style>

</style>
