<template>
  <div class="home">
    <HelloWorld msg="Kubernetes Dashboard" />
    <p>
      <b-card
        img-src="../assets/kubernetes-logo.png"
        img-alt="Image"
        img-top
        class="mb-2"
        padding-top="30px"
      >
        <b-card-text>
          This dashboard notifies the cluster's administrators about the status of their namespaces.<br />
          It displays critical and/or relevant information about potential actions to be undertaken in order to guarantee the smooth running of the various namespaces. <br />
          
          <div class="card" style="width: 18rem; margin-top:20px;">
            <div class="card-header">
              Features
            </div>
            <ul class="list-group list-group-flush">
              <li class="list-group-item"><router-link :to="{ path: '/namespaces'}">View namespaces statuses</router-link></li>
              <li class="list-group-item"><router-link :to="{ path: '/deprecated'}">View deprecated resources</router-link></li>
            </ul>
          </div>

        </b-card-text>
      </b-card>

    </p>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";

export default {
  name: "Home",
  components: {
    HelloWorld
  }
};
</script>
