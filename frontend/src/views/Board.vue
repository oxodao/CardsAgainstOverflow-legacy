<template>
    <div id="game">
        <BoardHeader />
        <SideMenu />
        <div id="board">
            <GameStatus />
            <RoomSettings v-if="!Started" />
            <template v-else>
                <BlackCard />
                <Deck />
                <div
                    v-if="!IsDeporte && IsReady && TurnState === 1 && IsJudge"
                    id="buttonbar"
                >
                    <button @click="sendProposalSelection">
                        Valider
                    </button>
                </div>
            </template>

            <Countdown />
        </div>
        <BottomMenu />
    </div>
</template>

<script>
import SideMenu               from '@/components/panels/SideMenu.vue';
import BoardHeader            from '@/components/panels/Header.vue';
import GameStatus             from '@/components/GameStatus.vue';
import RoomSettings           from '@/components/panels/RoomSettings.vue';
import Countdown              from '@/components/Countdown.vue';
import BlackCard              from '@/components/BlackCard.vue';
import Deck                   from '@/components/Deck.vue';
import BottomMenu             from '@/components/panels/BottomMenu.vue';

import {mapGetters, mapState} from 'vuex';

export default {
    name      : 'Board',
    components: {BottomMenu, Deck, BlackCard, Countdown, RoomSettings, GameStatus, BoardHeader, SideMenu},
    computed  : {
        ...mapState({
            Started  : state => state.Room.Started,
            IsJudge  : state => state.User.IsJudge,
            TurnState: state => state.Room.TurnState,
            IsDeporte: state => state.UI.Deporte,
        }),
        ...mapGetters(['IsReady'])
    },
    methods   : {
        sendProposalSelection() {
            this.$store.dispatch('skipCountdown');
        }
    }
};
</script>

<style lang="scss" scoped>
#game {
  width: 100%;
  height: 100%;
  display: grid;
  grid-template-areas: "a a"
                       "b c";
  grid-template-rows: 3em 1fr;
  grid-template-columns: 350px 1fr;

  @media (max-width: 650px) {
    grid-template-areas: "a" "c" "d";
    grid-template-rows: 3em 1fr 64px;
    grid-template-columns: 1fr;
  }

  #board {
    grid-area: c;

    display: grid;
    align-items: center;
    justify-content: space-around;

    grid-template-rows: auto min-content min-content;

    #buttonbar {
      text-align: center;
      margin: .5em 0;

      button {
        padding: 1em;
      }
    }

    @media (max-width: 650px) {
      max-width: 100%;
    }
  }
}
</style>