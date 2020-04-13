<template lang="pug">
  #app
    #wrapper.d-flex
      Sidebar
      #page-content-wrapper
        Navbar
        .container-fluid
          router-view
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import Sidebar from "@/components/Sidebar.vue";

export default {
  name: "App",
  components: {
    Navbar,
    Sidebar
  }
};
</script>
 
<style lang="scss">
body {
  overflow-x: hidden;
}

#page-content-wrapper {
  min-width: 100vw;
}

@media (min-width: 768px) {
  #sidebar-wrapper {
    margin-left: 0;
  }

  #page-content-wrapper {
    min-width: 0;
    width: 100%;
  }
}
</style>
