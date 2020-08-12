import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Deprecated from '../views/Deprecated.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/namespaces',
    name: 'Namespaces',
    component: () => import(/* webpackChunkName: "about" */ '../views/Namespaces.vue')
  },
  {
    path: '/deprecated',
    name: 'Deprecated',
    component: () => import(/* webpackChunkName: "about" */ '../views/Deprecated.vue'),
    children: [
      {
        // `UserProfile` va être rendu à l'intérieur du `<router-view>` de `User`
        // quand `/utilisateur/:id/profil` concorde
        path: ':id',
        component: Deprecated
      }
    ]
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
