<template>
    <div id="game">
        <Header />
        <div id="inner-game">
            <SideMenu />
            <div id="board">
                <GameStatus />
                <RoomSettings v-if="!Started" />
                <template v-else>
                    <BlackCard />
                    <Deck />
                    <button v-if="!IsDeporte && IsReady && TurnState === 1 && IsJudge" @click="sendProposalSelection">Valider</button>
                </template>

                <Countdown />
            </div>
        </div>
    </div>
</template>

<script>
    import SideMenu from "../components/panels/SideMenu";
    import Header from "../components/panels/Header";
    import GameStatus from "../components/GameStatus";
    import RoomSettings from "../components/panels/RoomSettings";
    import Countdown from "../components/Countdown";
    import {mapGetters, mapState} from "vuex";
    import BlackCard from "../components/BlackCard";
    import Deck from "../components/Deck";

    export default {
        name: "Board",
        components: {Deck, BlackCard, Countdown, RoomSettings, GameStatus, Header, SideMenu},
        computed: {
            ...mapState({
                Started: state => state.Room.Started,
                IsJudge: state => state.User.IsJudge,
                TurnState: state => state.Room.TurnState,
                IsDeporte: state => state.UI.Deporte,
            }),
            ...mapGetters([ 'IsReady' ])
        },
        methods: {
            sendProposalSelection() {
                this.$store.dispatch('skipCountdown');
            }
        }
    }
</script>

<style lang="scss" scoped>
    #game {
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: column;

        #inner-game {
            flex: 1;
            display: flex;
            flex-direction: row;

            #board {
                flex: 1;

                display: flex;
                flex-direction: column;
                align-items: center;
                justify-content: space-around;

                > button {
                    padding: 1em;
                }

                @media (max-width: 650px) {
                    max-width: 100%;
                }
            }
        }
    }
</style>