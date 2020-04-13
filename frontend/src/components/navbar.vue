<template lang="pug">
  b-navbar.border-bottom(toggleable="lg" type="light" variant="light")
    b-button(variant="primary" @click="toggleSidebar")
      b-icon(icon="list" aria-hidden="true")
        span(class="sr-only") Help
    //- b-navbar-brand Blabla
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "Navbar",
  methods: mapActions(["toggleSidebar"])
};
</script>

<style>
</style>