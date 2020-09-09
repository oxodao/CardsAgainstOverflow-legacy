import Vue from 'vue'
import Vuex from 'vuex'

import ui from "./ui";
import user from "./user";
import room from "./room";

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    UI: ui,
    Room: room,
    User: user,
  }
})
