<template>
    <header>
        <button @click="quitGame">
            <img
                src="../../assets/exit.png"
                alt="Quitter la partie"
            />
        </button>
        <h1 id="game-title">
            Cards Against Overflow
        </h1>

        <div>
            <span>Code salle: {{ RoomID }}</span>
        </div>
    </header>
</template>

<script>
import {mapState} from 'vuex';

export default {
    name    : 'BoardHeader',
    computed: {
        ...mapState({
            RoomID: state => state.Room.RoomID,
        })
    },
    methods : {
        quitGame(e) {
            e.stopPropagation();
            this.$store.commit('showCloseDialog');
            return false;
        }
    }
};
</script>

<style lang="scss" scoped>
header {
  grid-area: a;
  position: relative;
  background: #111;
  text-align: center;
  display: flex;
  flex-direction: row;
  align-items: center;

  button {
    width: 34px;
    height: 34px;
    padding: 5px;

    @media (min-width: 651px) {
      display: none;
    }

    img {
      width: 24px;
      height: 24px;
    }
  }

  div {
    position: absolute;
    top: 0;
    bottom: 0;
    right: 0;

    padding-right: .5em;

    display: flex;
    flex-direction: column;

    align-items: flex-end;
    justify-content: center;

    span, button {
      display: block;
    }
  }
}

#game-title {
  font-size: 1.5em;
  flex: 1;

  @media (max-width: 650px) {
    font-size: 1.2em;
    text-align: left;
    margin-left: .5em;
  }
}
</style>