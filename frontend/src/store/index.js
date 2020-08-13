import Vue from 'vue'
import Vuex from 'vuex'
import namespaces from './modules/namespaces'
import deprecated from './modules/deprecated'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    sidebarOpen: true
  },
  mutations: {
    toggleSidebar(state) {
      state.sidebarOpen = !state.sidebarOpen
    }
  },
  actions: {
    toggleSidebar: ({ commit }) => commit('toggleSidebar')
  },
  modules: {
    namespaces,
    deprecated
  }
})
