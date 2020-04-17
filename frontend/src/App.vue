<template>
  <div id="app">
    <LoginModal v-if="ShowModal"/>
    <Board v-else />
  </div>
</template>

<script>
import {mapState} from 'vuex';
import LoginModal from './components/LoginModal.vue'
import Board from './components/Board.vue'

export default {
  name: 'App',
  components: {
    LoginModal,
    Board
  },
  computed: mapState({
    ShowModal: state => state.CurrentState.ShowLogin
  })
}
</script>

<style>
  #app {
    height: 100vh;
  }
</style>
