<template>
  <div>
    <b-form-input class="my-3" v-model="filter" type="search" placeholder="Search..."></b-form-input>
    <b-row cols="1" cols-sm="2" cols-md="3" cols-lg="4" cols-xl="5">
      <b-col class="mb-4" v-for="product in filterNamespaces(namespaces, filter)" :key="product.id">
        <b-card
          border-variant="success"
          :header="product.metadata.name"
          align="center"
          v-if="product.metadata.name.indexOf(filter) !== -1"
        >
          <b-card-text>
            <router-link :to="{ path: '/deprecated/'+ product.metadata.name}">View namespace</router-link>
          </b-card-text>
        </b-card>
      </b-col>
    </b-row>

    <!-- {{namespaces}}
    <b-form-input class="my-3" v-model="filter" type="search" placeholder="Search..."></b-form-input>
    <b-row cols="1" cols-sm="2" cols-md="3" cols-lg="4" cols-xl="5">
          <b-col class="mb-4" v-for="product in namespaces" :key="product.id">
            <b-card border-variant="success" :header="product.metadata.name" align="center" v-if="product.metadata.name.indexOf(filter) !== -1">
              <b-card-text>
                <router-link :to="{ path: '/deprecated/'+ product.metadata.name}">View namespace</router-link>
              </b-card-text>
            </b-card>
          </b-col>
    </b-row>-->
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  data() {
    return {
      filter: "",
    };
  },
  computed: mapState({
    namespaces: (state) => state.namespaces.all,
  }),
  created() {
    this.$store.dispatch("namespaces/loadCoins");
  },
  methods: {
    filterNamespaces(arr, filt) {
      return arr.filter((x) => x.metadata.name.indexOf(filt) !== -1);
    },
  },
};
</script>