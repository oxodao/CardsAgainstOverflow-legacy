<template>
  <div id="app" v-bind:class="showWizz">
    <Rules/>
    <LoginModal v-if="ShowModal"/>
    <Board v-else />
      <audio autoPlay v-for="i in wizz" v-bind:key="i" @ended="remWiz(i)">
        <source src="./assets/wizz.mp3" type="audio/mpeg" />
      </audio>
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
  computed: {
    ...mapState({
      ShowModal: state => state.ShowLogin,
      wizz: state => state.Wizz,
    }),
    showWizz() {
      return this.wizz.length > 0 ? "wizz" : "";
    }
  },
  methods: {
    remWiz(user) {
      this.$store.commit('delWizz', user)
    }
  }
}
</script>

<style>
  #app {
    height: 100vh;
  }

  button {
    display: inline-block;
    background: #3EC480;
    border-radius: .25em;
    padding: .25em;
    margin-left: 5px;
    border: none;
  }

  .wizz {
    animation: wizz .3s ease-in-out;
    animation-iteration-count: infinite;
  }

  @keyframes wizz {
    0% {
      transform: translateX(-1vw) rotate(-2deg);
    }
    15% {
      transform: translateX(1vw) rotate(0deg);
    }
    30% {
      transform: translateX(-1vw) rotate(2deg);
    }
    45% {
      transform: translateX(1vw) rotate(0deg);
    }
    60% {
      transform: translateX(-1vw) rotate(-2deg);
    }
    75% {
      transform: translateX(1vw) rotate(0deg);
    }
    90% {
      transform: translateX(-1vw) rotate(2deg);
    }
    100% {
      transform: translateX(0) rotate(0deg);
    }
  }
</style>
