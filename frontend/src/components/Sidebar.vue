<template lang="pug">
  #sidebar-wrapper(class="bg-light border-right" v-bind:class="[ sidebarOpen ? '' : 'toggled' ]")
    .sidebar-heading(class="border-bottom") Kubedash
    div(class="list-group list-group-flush")
      router-link(class="list-group-item list-group-item-action bg-light" to="/") Dashboard
      router-link(class="list-group-item list-group-item-action bg-light" to="/namespaces") Namespaces
      router-link(class="list-group-item list-group-item-action bg-light" to="/deprecated") Deprecated
      router-link(class="list-group-item list-group-item-action bg-light" to="/about") About

      .sidebar-heading Externals
      a(href="https://kubernetes.io" target="_blank" class="list-group-item list-group-item-action bg-light") K8s Dashboard&nbsp;
        b-icon(icon="box-arrow-up-right" aria-hidden="true")
      a(href="https://prometheus.io" target="_blank" class="list-group-item list-group-item-action bg-light") Prometheus&nbsp;
        b-icon(icon="box-arrow-up-right" aria-hidden="true")

</template>

<script>
export default {
  name: "Sidebar",
  computed: {
    sidebarOpen() {
      return this.$store.state.sidebarOpen;
    }
  }
};
</script>

<style>
#sidebar-wrapper.toggled {
  margin-left: 0;
}

#sidebar-wrapper {
  min-height: 100vh;
  margin-left: -15rem;
  -webkit-transition: margin 0.25s ease-out;
  -moz-transition: margin 0.25s ease-out;
  -o-transition: margin 0.25s ease-out;
  transition: margin 0.25s ease-out;
}

#sidebar-wrapper .sidebar-heading {
  padding: 0.875rem 1.25rem;
  font-size: 1.2rem;
}

#sidebar-wrapper .list-group {
  width: 15rem;
}

@media (min-width: 768px) {
  #sidebar-wrapper {
    margin-left: 0;
  }

  #page-content-wrapper {
    min-width: 0;
    width: 100%;
  }

  #sidebar-wrapper.toggled {
    margin-left: -15rem;
  }
}
</style>