<template>
  <div>
  List of deprecated resources
    <b-form-input class="my-3" v-model="filter" type="search" placeholder="Search..."></b-form-input>
    <b-row cols="1" cols-sm="2" cols-md="3" cols-lg="4" cols-xl="5">
      {{ deprecated }}
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