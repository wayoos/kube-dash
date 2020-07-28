<template>
  <div>
  List of deprecated resources 
    <b-form-input class="my-3" v-model="filter" type="search" placeholder="Search..."></b-form-input>
    <b-row cols="1">
      <b-col class="mb-4" v-for="ns in deprecated" :key="ns">
        <b-card border-variant="success" :header="ns.Namespace" align="left">    
          <b-table hover :fields="fields" :items="ns.StateHelm.items">
            <template v-slot:cell(api.version)="data">{{ data.value }}</template>
          </b-table>
          <!--<div v-for="item in ns.State.items" :key="item">
            {{ item }}
          </div>
        
         -->
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
    deprecated: state => state.deprecated.all
  }),
  created() {
    this.$store.dispatch("deprecated/loadCoins");
  }
};
</script>