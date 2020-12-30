<template>
  <div id="app" @click="hideMenu" v-bind:class="DisplayWizz">
    <Rules/>
    <Login v-if="!loggedIn"/>
    <Board v-else />
    <audio autoPlay v-for="i in wizz" v-bind:key="i" @ended="delWizz(i)">
      <source src="./assets/wizz.mp3" type="audio/mpeg" />
    </audio>

    <div id="lost" v-if="connectionLost">
      <h1>Connection perdue!</h1>
      <p>Si vous êtes sur mobile, ne changez pas d'application au cours de la partie!</p>

      <button @click="exit">Relancer</button>
    </div>
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
        connectionLost: state => state.UI.LostConnection,
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
        if (this.$store.state.UI.MenuVisible) {
          this.$store.commit('toggleMenu');
          e.stopPropagation();
        }
      },
      exit() {
        window.location.reload();
      }
    },
  }
</script>

<style lang="scss">
  * {
    box-sizing: border-box;
  }

  html {
    width: 100%;
    height: 100%;
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

    width: 100%;
    height: 100%;
  }

  #lost {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 9999999999999;
    background: rgba(black, .8);
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
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
