<template lang="pug">
  b-nav.sidebar(vertical :class='{ open: sidebarOpen }')
    b-nav-text Essential Links
    b-nav-item(href='https://vuejs.org' target='_blank') Core Docs
    b-nav-item(href='https://forum.vuejs.org' target='_blank') Forum
    b-nav-item(href='https://chat.vuejs.org' target='_blank') Community Chat
    b-nav-item(href='https://twitter.com/vuejs' target='_blank') Twitter
    b-nav-item(href='http://vuejs-templates.github.io/webpack/' target='_blank') Docs for This Template

    b-nav-text Essential Links
    b-nav-item(href='http://router.vuejs.org/' target='_blank') vue-router
    b-nav-item(href='http://vuex.vuejs.org/' target='_blank') vuex
    b-nav-item(href='http://vue-loader.vuejs.org/' target='_blank') vue-loader
    b-nav-item(href='https://github.com/vuejs/awesome-vue' target='_blank') awesome-vue
</template>

<script>
export default {
  name: "Sidebar",
  props: {
    sidebarOpen: {
      type: Boolean,
      default: true
    }
  }
};
</script>

<style>
</style>