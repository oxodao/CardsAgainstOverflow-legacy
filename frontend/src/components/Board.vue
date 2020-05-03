<template>
    <div id="board">
        <header>
            <h1>Cards Against Overflow</h1>
            <div class="gameinfo">
                <span>Code salle: {{room}}</span>
                <div class="buttons">
                    <template v-if="isAdmin && isStarted">
                        <button>Changer la partie</button>
                    </template>
                    <template v-else-if="isAdmin">
                        <button v-if="isReady" @click="startGame">Démarrer</button>
                        <button>Paramètres</button>
                    </template>
                    <button @click="exit">Quitter</button>
                </div>
            </div>
        </header>
        <Countdown />
        <div class="game-wrapper">
            <nav>
                <img src="../assets/logo.png"/>

                <h2>Joueurs</h2>
                <ul>
                    <PlayerName v-for="player in participants" :key="player.Username + player.Score +player.IsJudge+player.IsAdmin" v-bind:username="player.Username" v-bind:score="player.Score" v-bind:isAdmin="player.IsAdmin" v-bind:isJudge="player.IsJudge" />
                </ul>
                <h3>Tour: {{currTurn}} / {{maxTurn}}</h3>
            </nav>
            <div class="game">
                <div id="question">
                    <h1 v-if="currBlackCard !== undefined && currBlackCard !== null" id="question">
                        <span v-for="txt in getCardText" v-bind:key="txt.Question" v-bind:class="txt.Class">{{txt.Question}}</span>
                    </h1>
                    <h1 v-else-if="isStarted" id="question">Attente de nouveaux joueurs</h1>
                    {{/* For some reasons the v-else doesn't work here... */ }}
                    <h1 v-if="!currBlackCard && !isStarted " id="question">
                        En attente de joueurs...
                        <template v-if="isReady"> <br /> La partie est prête </template>
                    </h1>
                </div>
                <template>
                    <div v-if="turnState === 0 && !isJudge" id="hack">
                        <div id="cards">
                            <Card v-for="(card, index) in getCards" :key="card.ID+card.isSelected" v-bind:isProposal="false" v-bind:index="index" v-bind:currCard="card" />
                        </div>
                    </div>
                    <div v-else-if="turnState === 0 && isJudge">
                        <h3>Les joueurs jouent!</h3>
                    </div>
                    <div v-else-if="turnState === 1">
                        <div id="cards">
                            <Card v-for="(card, index) in getProposals" :key="card.ID+card.isSelected" v-bind:isProposal="true" v-bind:index="index" v-bind:currCard="card" />
                        </div>
                    </div>
                    <div v-else-if="turnState === 2">
                        <h1 class="winner">{{ winner }}</h1>
                    </div>
                </template>
            </div>
        </div>
    </div>
</template>

<script>
import {mapState} from 'vuex';
import Card from './Card';
import Countdown from './Countdown';
import PlayerName from './PlayerName';

export default {
    name: 'Board',
    components: {
        PlayerName,
        Card,
        Countdown
    },
    computed: {
        ...mapState({
            buttonSendAnswers: state => state.CurrentState.SendAnswersAllowed,
            hasPlayed: state => state.User.HasPlayed,
            currBlackCard: state => state.Room.CurrentBlackCard,
            room: state => state.Room.RoomID,
            currTurn: state => state.Room.Turn,
            maxTurn: state => state.Room.MaxTurn,
            participants: state => state.Room.Participants,
            isStarted: state => state.Room.Started,
            isJudge: state => state.User.IsJudge,
            turnState: state => state.Room.TurnState,
            isAdmin: state => state.User.IsAdmin,
            winner: state => state.Room.Winner
        }),
        getCards() {
            return this.$store.state.User.Hand.filter((card) => card !== undefined && card !== null)
        },
        getProposals() {
            return this.$store.state.Room.Answers
        },
        getCardText() {
            if (this.turnState === 0)
                return [ 
                    {
                    "Question": this.currBlackCard.Text,
                    "Class": "",
                    }
                ];

            let txt = this.currBlackCard.Text;
            let txtSplitted = txt.split(/(____)/)
            let curr = 0;
            let values = [];
            txtSplitted.forEach(e => {
                if (e === "____") {
                    values.push({
                        Question: this.$store.state.Room.Answers[this.$store.state.Room.JudgeSelection].Cards[curr].Text,//"ANSW" + curr, //this.selectedAnswers[this.selectedAnswersIndex].Cards[curr],
                        Class: 'colored'
                    })
                    curr++
                } else
                    values.push({ Question: e, Class: '' })
            });

            return values
        },
        isReady() {
            return this.$store.state.Room.Participants.length >= 3
        }
    },
    methods: {
        startGame() {
            this.$store.dispatch('startGame');
        },
        sendAnswers() {
            this.$store.dispatch('answer');
        },
        exit() {
            window.location.reload();
        }
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
                font-size: 2em;
            }

            .gameinfo {
                display: flex; 
                flex-direction: column;

                position: absolute;
                top: 0;
                right: 0;
                margin: .25em .5em 0 0;

                span {
                    text-align: right;
                }

                .buttons {
                    text-align: right;
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
                
                h3 {
                    text-align: center;
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
                    display: flex;
                    flex-direction: row;
                    align-items: center;
                    justify-content: space-around;
                    padding: 2em;
                }
            }
        }

        #validate {
            padding-bottom: 1em;

            button {
                padding: 1em;
                color: #111;
            }
        }

        button {
            display: inline-block;
            background: #3EC480;
            border-radius: .25em;
            padding: .25em;
            margin-left: 5px;
            border: none;
        }

        button:disabled {
            background: darken(#3EC480, 20%);
            color: #333;
        }
    }

    @media (max-width: 900px){
        header .h1 {
            margin-left: .25em;
            text-align: left !important;
        }

        .game-wrapper h1 {
            font-size: 1.5em;
        }
    }

    .colored {
        color: #3EC480;
    }

    .winner {
        height: 5em;
        color: #3EC480;
        text-transform: uppercase;
        animation: glow 2s ease-in-out infinite alternate;
    }

    @keyframes glow {
        from {
            text-shadow: 0px 0px 15px #3EC480, 0 0 5px #4dbbc7;
        }
        to {
            text-shadow: 0px 0px 5px #3EC480, 0 0 10px #4dbbc7;
        }
    }

    @media (max-width: 666px) {
        header h1 {
            opacity: 0; // ugly hack so we keep the reserved space
        }
    }

        @media (max-width: 1300px){
        h1 {
            font-size: 1.7em;
        }
    }

    @media (max-width: 900px){
        h1 {
            font-size: 1.5em;
        }
    }
</style>