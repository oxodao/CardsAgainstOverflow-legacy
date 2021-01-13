<template>
  <nav :class="MenuVisible ? 'visible' : ''">
    <img src="../../assets/logo.png" alt="Logo"/>

    <button id="endgame" @click="exit">Quitter la partie</button>

    <h2>Joueurs</h2>

    <ul>
      <PlayerName v-for="player in participants"
                  :key="player.Username + player.Score + player.IsJudge + player.IsAdmin + player.HasPlayed"
                  v-bind:player="player"/>
    </ul>

    <div id="actions">
      <WizzButton/>
      <RerollButton/>
    </div>
    <h3 v-if="isStarted && currTurn <= maxTurn && !zenMode">Tour: {{ currTurn }} / {{ maxTurn }}</h3>
    <h3 v-else-if="isStarted && (zenMode || currTurn <= maxTurn)">Tour: {{ currTurn }}</h3>
  </nav>
</template>

<script>
import PlayerName   from "../PlayerName";
import {mapState}   from "vuex";
import RerollButton from "@/components/RerollButton";
import WizzButton   from "@/components/WizzButton";

export default {
  name      : "SideMenu",
  components: {
    WizzButton,
    RerollButton,
    PlayerName,
  },
  computed  : {
    ...mapState({
      participants : state => state.Room.Participants,
      currTurn     : state => state.Room.Turn,
      maxTurn      : state => state.Room.MaxTurn,
      zenMode      : state => state.Room.ZenMode,
      isStarted    : state => state.Room.Started,
      MenuVisible  : state => state.UI.MenuVisible,
    })
  },
  methods   : {
    exit(e) {
      e.stopPropagation();
      this.$store.commit('showCloseDialog')
      return false;
    }
  },

}
</script>

<style lang="scss" scoped>
nav {
  grid-area: b;
  padding: .5em;
  background: #111;
  display: flex;
  flex-direction: column;
  align-items: center;
  z-index: 999999999;

  @media(max-width: 650px) {
    position: absolute;
    top: 0;
    left: -100%;
    bottom: 0;

    transition: left .25s linear;

    &.visible {
      left: 0;
    }
  }

  img {
    text-align: center;
    width: 215px;
  }

  #endgame {
    margin-top: 1em;
  }

  h2 {
    text-align: center;
    text-decoration: underline;
  }

  h3 {
    text-align: center;
  }

  ul {
    flex: 1;
    list-style: none;
    padding: 0;
  }

  #actions {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
  }
}
</style>