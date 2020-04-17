<template>
    <div id="board">
        <header>
            <h1>Cards Against Overflow</h1>
            <div class="gameinfo">
                <span>Code salle: {{room}}</span>
                <div class="buttons">
                    <template v-if="isStarted">
                        <button>Changer la partie</button>
                    </template>
                    <template v-else>
                        <button @click="startGame">Démarrer</button>
                        <button>Paramètres</button>
                    </template>
                </div>
            </div>
        </header>
        <div class="game-wrapper">
            <nav>
                <img src="../assets/logo.png"/>

                <h2>Joueurs</h2>
                <ul>
                    <PlayerName v-for="player in participants" :key="player.Username" v-bind:username="player.Username" v-bind:isAdmin="player.IsAdmin" v-bind:isJudge="player.IsJudge" />
                </ul>
            </nav>
            <div class="game">
                <div id="question">
                    <h1 v-if="currBlackCard != null" id="question">{{currBlackCard.Text}}</h1>
                    <h1 v-else-if="isStarted" id="question">Attente de nouveaux joueurs</h1>
                    <h1 v-else id="question">
                        En attente de joueurs...
                        <template v-if="isReady"> <br /> La partie est prête </template>
                    </h1>
                </div>
                <div id="cards">
                    <Card v-for="card in getCards" :key="card.ID+card.isSelected" v-bind:currCard="card" />
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import {mapState} from 'vuex';
import Card from './Card';
import PlayerName from './PlayerName';

export default {
    name: 'Board',
    components: {
        PlayerName,
        Card
    },
    computed: {
        ...mapState({
            currBlackCard: state => state.Room.BlackCard,
            room: state => state.Room.ID,
            participants: state => state.Room.Participants,
            selectedCards: state => state.User.SelectedCard,
            isStarted: state => state.Room.isStarted,
            isJudge: state => state.User.IsJudge
        }),
        getCards() {
            if (this.$store.state.User.IsJudge) {
                return this.$store.state.Room.selectedCards.filter((card) => card !== undefined && card !== null)
            }

            return this.$store.state.User.Hand.filter((card) => card !== undefined && card !== null)
        },
        isReady() {
            // Not working...
            return this.$store.Room && this.$store.Room.Participants.length >= 3
        }
    },
    methods: {
        startGame() {
            this.$store.dispatch('startGame')
        },
    },
}
</script>

<style lang="scss" scoped>
    #board {
        height: 100%;
        display: flex;
        flex-direction: column;

        header {
            background: #111;
            width: 100%;

            h1 {
                text-align: center;
                margin: 0;
                padding: .25em 0 .25em 0;
            }

            .gameinfo {
                display: flex; 
                flex-direction: column;

                position: absolute;
                top: 0;
                right: 0;
                margin: .25em .5em 0 0;

                .buttons {
                    text-align: right;

                    button {
                        display: inline-block;
                        background: #3EC480;
                        border-radius: .25em;
                        padding: .25em;
                        margin-left: 5px;
                        border: none;
                    }
                }

            }
        }

        .game-wrapper {
            flex: 1;
            display: flex;
            flex-direction: row;

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

                ul {
                    flex: 1;
                    list-style-type: none;
                    padding: 0;
                }
            }

            .game {
                flex: 1;
                display: flex;
                flex-direction: column;
                align-items: center;

                #question {
                    text-align: center;
                    flex: 1;
                    width: 100%;
                }

                #cards {
                    flex: 0 0 100px;
                    width: 1000px;
                    display: flex;
                    flex-direction: row;
                    align-items: center;
                    justify-content: space-around;
                    margin-bottom: 2em;
                }
            }
        }
    }
</style>