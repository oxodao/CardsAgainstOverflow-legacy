<template>
    <nav :class="MenuVisible ? 'visible' : ''">
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
                canWizz: state => state.UI.CanWizz,
                currTurn: state => state.Room.Turn,
                maxTurn: state => state.Room.MaxTurn,
                zenMode: state => state.Room.ZenMode,
                isStarted: state => state.Room.Started,
                MenuVisible: state => state.UI.MenuVisible,
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
        },
        mounted() {
            document.addEventListener('resize', () => {
                console.log("resize", this.$store.UI.MenuVisible)
                let w = document.documentElement.clientWidth;
                if (this.$store.UI.MenuVisible && w > 650) {
                    this.$store.commit('toggleMenu');
                }
            })
        }
    }
</script>

<style lang="scss" scoped>
    nav {
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