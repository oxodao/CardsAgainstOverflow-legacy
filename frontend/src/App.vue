<template>
    <div
        id="app"
        :class="DisplayWizz"
        @click="hidePlayersList"
    >
        <Rules />
        <Login v-if="!loggedIn" />
        <Board v-else />
        <audio
            v-for="i in wizz"
            :key="i.user"
            autoPlay
            @ended="delWizz(i)"
        >
            <source
                src="./assets/wizz.mp3"
                type="audio/mpeg"
            />
        </audio>

        <ExitDialog v-if="showExitDialog" />

        <div
            v-if="connectionLost"
            id="lost"
        >
            <h1>Connection perdue!</h1>
            <p>Si vous êtes sur mobile, ne changez pas d'application au cours de la partie!</p>

            <button @click="exit">
                Relancer
            </button>
        </div>
    </div>
</template>

<script>
import Rules from '@/components/panels/Rules.vue';
import Login from '@/views/Login.vue';
import Board from '@/views/Board.vue';
import ExitDialog from '@/components/ExitDialog.vue';

import {mapGetters, mapState} from 'vuex';

export default {
    name: 'App',
    components: {
        ExitDialog,
        Rules,
        Login,
        Board,
    },
    computed: {
        ...mapState({
            wizz: state => state.UI.Wizz,
            loggedIn: state => state.UI.LoggedIn,
            connectionLost: state => state.UI.LostConnection,
            showExitDialog: state => state.UI.QuitGameDialog,
        }),
        ...mapGetters([
            'DisplayWizz'
        ]),
    },
    methods: {
        delWizz(user) {
            this.$store.commit('delWizz', user);
        },
        hidePlayersList(e) {
            if (this.$store.state.UI.PlayersMenuVisible) {
                this.$store.commit('togglePlayersMenu');
                e.stopPropagation();
            }
        },
        exit() {
            window.location.reload();
        }
    },
};
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
