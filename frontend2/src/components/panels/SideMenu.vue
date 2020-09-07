<template>
    <nav>
        <img src="../../assets/logo.png" alt="Logo"/>

        <button id="endgame">Quitter la partie</button>

        <h2>Joueurs</h2>

        <ul>
            <PlayerName v-for="player in participants" :key="player.Username + player.Score + player.IsJudge + player.IsAdmin + player.HasPlayed" v-bind:player="player" />
        </ul>

        <div id="actions">
            <button :disabled="!canWizz" @click="addWizz" v-tooltip="'Wizz'"><img src="../../assets/msn_wizz.png" alt="wizz"/></button>
            <button v-if="isStarted" :disabled="!canReroll" @click="reroll" v-tooltip="'Re-piocher'"><img src="../../assets/reroll.png" alt="Reroll"/></button>
        </div>
        <h3 v-if="isStarted && currTurn <= maxTurn && !zenMode">Tour: {{currTurn}} / {{maxTurn}}</h3>
        <h3 v-else-if="isStarted && (zenMode || currTurn <= maxTurn)">Tour: {{currTurn}}</h3>
    </nav>
</template>

<script>
    import PlayerName from "../PlayerName";
    import {mapState} from "vuex";

    export default {
        name: "SideMenu",
        components: {
          PlayerName,
        },
        computed: {
            ...mapState({
                participants: state => state.Room.Participants,
                canWizz: state => state.CanWizz,
                currTurn: state => state.Room.Turn,
                maxTurn: state => state.Room.MaxTurn,
                zenMode: state => state.Room.ZenMode,
                isStarted: state => state.Room.Started,
            }),
        }
    }
</script>

<style lang="scss" scoped>
    nav {
        padding: .5em;
        background: #111;
        display: flex;
        flex-direction: column;

        img {
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

            button {
                width: 3em;
                height: 3em;

                img {
                    width: 100%;
                    height: 100%;
                    object-fit: contain;
                }
            }
        }
    }
</style>