<template>
  <div id="app" @click="hideMenu" v-bind:class="DisplayWizz">
    <Rules/>
    <Login v-if="!loggedIn"/>
    <Board v-else />
    <audio autoPlay v-for="i in wizz" v-bind:key="i" @ended="delWizz(i)">
      <source src="./assets/wizz.mp3" type="audio/mpeg" />
    </audio>
  </div>
</template>

<script>
  import Rules from './components/panels/Rules.vue'
  import Login from "./views/Login";
  import Board from "./views/Board";

  import {mapGetters, mapState} from "vuex";

  export default {
    name: 'App',
    components: {
      Rules,
      Login,
      Board,
    },
    computed: {
      ...mapState({
        wizz: state => state.UI.Wizz,
        loggedIn: state => state.UI.LoggedIn,
      }),
      ...mapGetters([
        'DisplayWizz'
      ]),
    },
    methods: {
      delWizz(user) {
        this.$store.commit('delWizz', user)
      },
      hideMenu(e) {
          console.log(this.$store.state.UI.MenuVisible)
        if (this.$store.state.UI.MenuVisible) {
          this.$store.commit('toggleMenu');
          e.stopPropagation();
        }
      }
    },
  }
</script>

<style lang="scss">
  * {
    box-sizing: border-box;
  }

  body {
    padding: 0;
    margin: 0;
    font-family: Helvetica, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    font-weight: bold;

    background: #222;
    color: #fff;
  }

  #app {
    width: 100vw;
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

  .tooltip {
    padding: .25em;
    margin: 1em;
    border-radius: 5px;
    background: rgba(10, 10, 10, .75);
  }
</style>
