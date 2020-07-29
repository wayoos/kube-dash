<template>
  <div>
  <h1>List of deprecated resources </h1>
    <b-row cols="1">
      <b-col class="mb-4" v-for="ns in deprecated" :key="ns">
        <b-card v-bind:class="{ 'border-warning' : ns.StateHelm.items, 'border-success' : !ns.StateHelm.items }" align="left"> 
          <template v-slot:header>
            namespace : <strong>{{ns.Namespace}}</strong>
          </template>
          <!-- No resources to list -->
          <p v-if="!ns.StateHelm.items">There were no resources found with known deprecated apiVersions.</p>

          <!-- Display resources as table -->
          <b-table v-else hover :fields="fields" :items="ns.StateHelm.items" ></b-table>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  data() {
    return {
      filter: "",
      fields: [
        { key: 'name', label: 'name' },
        { key: 'namespace', label: 'namespace' },
        { key: 'api.kind', label: 'kind' },
        { key: 'api.version', label: 'version' },
        { key: 'api.replacement-api', label: 'replacement-api' },
        { key: 'deprecated', label: 'Deprecated' },
        { key: 'api.component', label: 'component' },
        { key: 'api.deprecated-in', label: 'deprecated-in' }, 
        { key: 'api.removed-in', label: 'removed-in' },
      ]
    };
  },
  computed: mapState({
    deprecated: state => state.deprecated.all,
  }),
  created() {
    this.$store.dispatch("deprecated/loadCoins");
  }
};
</script>