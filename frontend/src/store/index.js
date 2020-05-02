import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        CurrentState: {
            ShowLogin: true,
            SendAnswersAllowed: false,
        },
        User: {
            Username: '',
            Color: '',
            Hand: [],
            Answer: [],
            HasPlayed: false,
            IsJudge: false,
            IsAdmin: false,
        },
        Room: {
            ID: '',
            BlackCard: '',
            isStarted: false,
            Participants: []
        },
        Judge: {
            Answers: [],
            SelectedAnswer: -1
        },
        Websocket: null
    },
    mutations: {
        setWebsocket: (state, payload) => {
            state.Websocket = payload;
        },
        connected: (state, payload) => {
            state.CurrentState.ShowLogin = false;

            state.User.Username = payload.User.Username;
            state.Room.ID = payload.Room;
            state.User.IsAdmin = payload.User.IsAdmin;
            state.User.IsJudge = payload.User.IsJudge;
        },
        startGame: (state) => {
            state.Room.Started = true
        },
        setPlayerList: (state, payload) => {
            state.Room.Participants = payload;
            payload.forEach(e => {
                if (e.Username == state.User.Username) {
                    state.User.IsAdmin = e.IsAdmin
                    state.User.IsJudge = e.IsJudge
                }
            });
        },
        setState: (state, payload) => {
            console.log(payload)
            state = payload
        },
        updateCards: (state, payload) => {
            state.Room.BlackCard = payload.BlackCard
            state.Judge.Answers = []
            state.Judge.SelectedAnswer = -1

            // This array is a fixed-sized array
            // This let us remove card and put them back without causing issues
            state.User.Answer = new Array(payload.BlackCard.AmtCardRequired)
            //Object.seal(state.User.Answer)

            state.User.Hand = payload.Hand.map(a => ({ ...a, answerPosition: -1 }))
            state.User.HasPlayed = false;
        },
        toggleSelection: (state, payload) => {
            if (state.User.HasPlayed) {
                return
            }

            if (payload.answerPosition == -1) {
                for (let i = 0; i < state.User.Answer.length; i++) {
                    // We find the first empty place
                    if (state.User.Answer[i] === undefined || state.User.Answer[i] === null) {
                        state.User.Answer[i] = payload;
                        payload.answerPosition = i
                        break
                    }
                }
            } else {
                for (let i = 0; i < state.User.Answer.length; i++) {
                    if (state.User.Answer[i] !== undefined && state.User.Answer[i] !== null && state.User.Answer[i].ID == payload.ID) {
                        state.User.Answer[i] = null;
                        payload.answerPosition = -1
                    }
                }
            }

            state.CurrentState.SendAnswersAllowed= state.User.Answer.filter(e => e !== undefined && e !== null).length == state.Room.BlackCard.AmtCardRequired;
        },
        toggleAnswerSelection: (state, payload) => {
            if (state.User.HasPlayed)
                return

            let selected = state.Judge.SelectedAnswer;

            if (selected !== -1) {
                state.Judge.Answers[selected].IsSelected = false;
                state.Judge.SelectedAnswer = -1;
            }

            if (payload.ID !== selected) {
                state.Judge.Answers.forEach((e, i) => {
                    if (e.ID == payload.ID) {
                        state.Judge.SelectedAnswer = i;
                        state.Judge.Answers[i].IsSelected = true;
                    }
                })
            }
        },
        addToAnswersList: (state, payload) => {
            Vue.set(state.Judge.Answers, state.Judge.Answers.length, {
                Text: "Proposition #" + (state.Judge.Answers.length+1),
                Cards: payload.map(e => e.Text),
                ID: payload[0].ID,
                IsSelected: false,
            });
        },
        played: (state) => {
            state.User.HasPlayed = true
        }
    },
    actions: {
        startGame(ctx) {
            ctx.state.Websocket.send(JSON.stringify({ Command: "START_GAME", Arguments: "{}" }))
        },
        answer: (ctx) => {
            if (ctx.state.User.IsJudge) {
                ctx.state.Websocket.send(JSON.stringify({
                    Command: 'SEND_ANSWERS',
                    Arguments: JSON.stringify([
                        ctx.state.Judge.Answers[ctx.state.Judge.SelectedAnswer]
                    ])
                }))
            } else {
                ctx.state.Websocket.send(JSON.stringify({
                    Command: 'SEND_ANSWERS',
                    Arguments: JSON.stringify(ctx.state.User.Answer.map(e => ({
                            ID: e.ID,
                            Text: e.Text,
                            IsBlackCard: e.IsBlackCard
                        })))
                }))
            }
            ctx.commit('played')
        },
    },
    getters: {
        IsPlayerJudge: state => state.User.IsJudge,
        IsPlayerAdmin: state => state.User.IsAdmin
    }
})