<template>
  <ul>
    <PlayerName v-for="player in participants"
                :key="player.Username + player.Score + player.IsJudge + player.IsAdmin + player.HasPlayed"
                v-bind:player="player"/>
  </ul>
</template>

<script>
import {mapState} from "vuex";
import PlayerName from "@/components/PlayerName";

export default {
  name: "UserList",
  components: {PlayerName},
  computed  : {
    ...mapState({
      participants: state => state.Room.Participants,
    })
  },
  mounted() {
    document.addEventListener('resize', () => {
      let w = document.documentElement.clientWidth;
      if (this.$store.UI.PlayersMenuVisible && w > 650) {
        this.$store.commit('togglePlayersMenu');
      }
    })
  }
}
</script>

<style lang="scss" scoped>
ul {
  position: absolute;
  bottom: 80px;
  left: 10px;

  padding: 1em;
  margin: 0;
  list-style: none;

  background: rgba(0, 0, 0, .8);
  border-radius: 1em;
  border: 1px solid black;
}
</style>