<template>
  <div id="app">
    <Rules/>
    <LoginModal v-if="ShowModal"/>
    <Board v-else />
  </div>
</template>

<script>
import {mapState} from 'vuex';
import LoginModal from './components/LoginModal.vue';
import Rules from './components/Rules.vue';
import Board from './components/Board.vue';

export default {
  name: 'App',
  components: {
    LoginModal,
    Rules,
    Board
  },
  computed: mapState({
    ShowModal: state => state.ShowLogin
  })
}
</script>

<style>
  #app {
    height: 100vh;
  }
</style>
