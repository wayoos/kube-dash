<template lang="pug">
  #sidebar-wrapper(class="bg-light border-right")
    .sidebar-heading Kubedash
    div(class="list-group list-group-flush")
      router-link(class="list-group-item list-group-item-action bg-light" to="/") Dashboard
      router-link(class="list-group-item list-group-item-action bg-light" to="/") Namespaces
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
  props: {
    sidebarOpen: {
      type: Boolean,
      default: true
    }
  }
};
</script>

<style>
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

  #wrapper.toggled #sidebar-wrapper {
    margin-left: -15rem;
  }
}
</style>