<template>
    <div id="board">
        <button id="ruleLink" @click="showRules">Règles / Utilisation</button>
        <header>
            <h1>Cards Against Overflow</h1>
            <div class="gameinfo">
                <span>Code salle: {{room}}</span>
                <div class="buttons">
                    <button @click="exit">Quitter</button>
                </div>
            </div>
        </header>
        <Countdown />
        <div class="game-wrapper">
            <PlayerNav />
            <div v-if="isPlaying" class="game">
                <div id="question">
                    <h1 v-if="currBlackCard !== undefined && currBlackCard !== null && isReady" id="questionText">
                        <span v-for="txt in getCardText" v-bind:key="txt.Question" v-bind:class="txt.Class">{{txt.Question}}</span>
                    </h1>
                    <h1 v-else-if="currBlackCard !== undefined && currBlackCard !== null && !isReady" id="questionText">
                        Il n'y a plus assez de joueurs !
                    </h1>
                    <h1 v-else-if="isStarted" id="questionText">Attente de nouveaux joueurs</h1>
                    {{/* For some reasons the v-else doesn't work here... */ }}
                    <h1 v-if="!currBlackCard && !isStarted " id="questionText">
                        En attente de joueurs...
                        <template v-if="isReady && HasEnoughCards"> <br /> La partie est prête </template>
                        <template v-else-if="isReady && !HasEnoughCards"> <br /> Pas assez de cartes dans les decks sélectionnés! </template>
                    </h1>
                    {{
                    /* @TODO:
                            When a game is started and a pleyer joins the game,
                            @HasEnoughCards
                            Toggle the GUI to prevent players to continue if false
                    */
                    }}
                </div>
                <template>
                    <!-- Quick hack, this component will be rewritten in multiple subcomponents, ugly for now -->
                    <div id="centerizehack" v-if="!isStarted || turnState === 2">
                        <h1 v-if="turnState === 2" class="winner">{{ winner }}</h1>
                        <RoomSettings v-if="!isStarted || (currTurn > maxTurn && !zenMode)"/>
                    </div>
                    <div v-else-if="turnState === 0 && !isJudge && isReady" id="hack">
                        <div id="cards">
                            <Card v-for="(card, index) in getCards" :key="card.ID+card.isSelected" v-bind:isProposal="false" v-bind:index="index" v-bind:currCard="card" />
                        </div>
                    </div>
                    <div v-else-if="turnState === 0 && isJudge && isReady">
                        <h3>Les joueurs jouent!</h3>
                    </div>
                    <div v-else-if="turnState === 1 && isReady && HasEnoughCards">
                        <div id="cards">
                            <Card v-for="(card, index) in getProposals" :key="card.ID+card.isSelected" v-bind:isProposal="true" v-bind:index="index" v-bind:currCard="card" />
                        </div>
                        <div v-if="isJudge" id="validate">
                            <button @click="skipCountdown">Voter !</button>
                        </div>
                    </div>
                </template>
            </div>
            <div v-else class="game winner">
                <h1>{{getWinner}}</h1>
                <RoomSettings v-if="currTurn > maxTurn"/>
            </div>
        </div>
    </div>
</template>

<script>
    import {mapGetters, mapState} from 'vuex';
    import Card from './Card';
    import Countdown from './Countdown';
    import RoomSettings from "./RoomSettings";
    import PlayerNav from "./PlayerNav";

    export default {
        name: 'Board',
        components: {
            PlayerNav,
            RoomSettings,
            Card,
            Countdown
        },
        computed: {
            ...mapState({
                currBlackCard: state => state.Room.CurrentBlackCard,
                room: state => state.Room.RoomID,
                currTurn: state => state.Room.Turn,
                maxTurn: state => state.Room.MaxTurn,
                zenMode: state => state.Room.ZenMode,
                isStarted: state => state.Room.Started,
                isJudge: state => state.User.IsJudge,
                turnState: state => state.Room.TurnState,
                isAdmin: state => state.User.IsAdmin,
                winner: state => state.Room.Winner,
            }),
            ...mapGetters([
                'HasEnoughCards'
            ]),
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
                        let Question = "";
                        let winning = this.$store.state.Room.WinningAnswer;
                        if (winning !== undefined && winning !== null) {
                            Question = winning.Cards[curr].Text
                        } else {
                            Question = this.$store.state.Room.Answers[this.$store.state.Room.JudgeSelection].Cards[curr].Text
                        }
                        values.push({
                            Question,
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
            },
            isPlaying() {
                return this.$store.getters.IsPlaying;
            },
            getWinner() {
                let winners = [""];
                let winnerScore = -1;
                this.$store.state.Room.Participants.forEach(e => {
                    if (e.Score === winnerScore) {
                        winners.push(e.Username);
                        winnerScore = e.Score;
                    } else if (e.Score > winnerScore) {
                        winners = [e.Username];
                        winnerScore = e.Score;
                    }
                })

                if (winnerScore === 0) {
                    return "Personne n'a gagné !";
                }

                if (winners.length === 1) {
                    return "Le gagnant est " + winners[0] + " !";
                }

                if (winners.length === this.$store.state.Room.Participants.length) {
                    return "Match nul !";
                }

                return "Les gagnants sont " + winners.join(", ") + " !";
            }
        },
        methods: {
            skipCountdown() {
                this.$store.dispatch('skipCountdown');
            },
            showRules() {
                this.$store.commit('showRules');
            },
            exit() {
                window.location.reload();
            },
        },
    }
</script>

<style lang="scss" scoped>
    #board {
        width: 100%;
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

                &.winner {
                    justify-content: center;
                }
            }
        }

        #validate {
            padding-bottom: 1em;
            text-align: center;

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

    .winner h1, h1.winner {
        height: 5em;
        color: #3EC480;
        text-transform: uppercase;
        animation: glow 2s ease-in-out infinite alternate;
        text-align: center;
    }

    @keyframes glow {
        from {
            text-shadow: 0 0 15px #3EC480, 0 0 5px #4dbbc7;
        }
        to {
            text-shadow: 0 0 5px #3EC480, 0 0 10px #4dbbc7;
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


    #ruleLink {
        position: absolute;
        top: 1em;
        left: 1em;
    }


    #centerizehack {
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
    }
</style>