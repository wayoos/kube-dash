<template lang="pug">
  b-sidebar#sidebar(title='Kubedash' visible)
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