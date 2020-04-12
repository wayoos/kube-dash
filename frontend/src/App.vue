<template lang="pug">
  #app
    #wrapper.d-flex
      Sidebar
      #page-content-wrapper
        nav(class="navbar navbar-expand-lg navbar-light bg-light border-bottom")
          b-button(variant="primary")
            b-icon(icon="list" aria-hidden="true")
            span(class="sr-only") Help

          button#menu-toggle(class="btn btn-primary") Toggle Menu
          button.navbar-toggler(type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation")
            span.navbar-toggler-icon

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

#wrapper.toggled #sidebar-wrapper {
  margin-left: 0;
}

@media (min-width: 768px) {
  #sidebar-wrapper {
    margin-left: 0;
  }

  #page-content-wrapper {
    min-width: 0;
    width: 100%;
  }

  #wrapper.toggled #sidebar-wrapper {
    margin-left: -15rem;
  }
}
</style>
