<template>
    <nav>
        <img src="../assets/logo.png" alt="Logo"/>

        <h2>Joueurs</h2>
        <ul>
            <PlayerName v-for="player in participants" :key="player.Username + player.Score + player.IsJudge + player.IsAdmin + player.HasPlayed" v-bind:username="player.Username" v-bind:score="player.Score" v-bind:isAdmin="player.IsAdmin" v-bind:isJudge="player.IsJudge" v-bind:hasPlayed="player.HasPlayed" />
        </ul>

        <div id="actions">
            <button :disabled="!canWizz" @click="addWizz" v-tooltip="'Wizz'"><img src="../assets/msn_wizz.png" alt="wizz"/></button>
            <button v-if="isStarted" :disabled="!canReroll" @click="reroll" v-tooltip="'Re-piocher'"><img src="../assets/reroll.png" alt="Reroll"/></button>
        </div>
        <h3 v-if="isStarted && currTurn <= maxTurn && !zenMode">Tour: {{currTurn}} / {{maxTurn}}</h3>
        <h3 v-else-if="isStarted && (zenMode || currTurn <= maxTurn)">Tour: {{currTurn}}</h3>
    </nav>
</template>

<script>
    import {mapState} from "vuex";
    import PlayerName from './PlayerName';

    export default {
        name: "PlayerNav",
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
            canReroll() {
                return !this.isJudge && this.turnState === 0 && this.$store.state.User.RerollTimeout  === 0;
            },
        },
        methods: {
            reroll() {
                this.$store.dispatch('reroll');
            },
            addWizz() {
                this.$store.dispatch('sendWizz')
            },
        }
    }
</script>

<style lang="scss" scoped>
    nav {
        padding: .5em;
        flex: 0 0 calc(300px - 1em);
        background: #111;
        display: flex;
        flex-direction: column;

        img {
            width: 75%;
            display: block;
            margin: auto;
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
            list-style-type: none;
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
            }
        }
    }
</style>