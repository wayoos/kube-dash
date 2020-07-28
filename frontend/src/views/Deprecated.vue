<template>
  <div>
  List of deprecated resources
  {{ deprecated[0] }}
  {{ namespaces }}
    <b-form-input class="my-3" v-model="filter" type="search" placeholder="Search..."></b-form-input>
    <b-row cols="1" cols-sm="2" cols-md="3" cols-lg="4" cols-xl="5">
      <b-col class="mb-4" v-for="product in deprecated" :key="product">
        <b-card border-variant="success" :header="product.Namespace" align="center">
          <b-card-text>{{product.State}}</b-card-text>
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
      filter: ""
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