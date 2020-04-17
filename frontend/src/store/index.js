import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        CurrentState: {
            ShowLogin: true,
        },
        User: {
            Username: '',
            Color: '',
            Hand: [],
            SelectedCard: -1
        },
        Room: {
            ID: '',
            BlackCard: '',
            Participants: [],
            VotingHand: []
        },
        Websocket: null
    },
    mutations: {
        setWebsocket: (state, payload) => {
            state.Websocket = payload;
        },
        connected: (state, payload) => {
            state.CurrentState.ShowLogin = false;

            state.User.Username = payload.Username;
            state.User.Color = payload.Color;
            state.Room.ID = payload.Room;
        },
        setReady: (state, payload) => {
            state.Room.RoomReady = payload.IsReady
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
        updateCards: (state, payload) => {
            state.Room.BlackCard = payload.BlackCard
            state.User.Hand = payload.Hand.map(a => ({ ...a, isSelected: false }))
            state.User.HasPlayed = false
        },
        setSelected: (state, payload) => {
            if (state.Room.Started && !state.User.HasPlayed) {
                state.User.SelectedCard = payload
                for (let i = 0; i < state.User.length; i++) {
                    if (state.User.SelectedCard[i].ID == payload) {
                        state.User.SelectedCard[i].isSelected = true
                    }
                }
                state.User.HasPlayed = true
            }
        },
        setVotingCards: (state, payload) => {
            state.Room.VotingHand = payload.map(a => ({ ...a, isSelected: false }))
            console.log("Vote list: ", state.Room.VotingHand)
        },
    },
    actions: {
        setSelected(ctx, SelectedCard) {
            if (!ctx.state.User.HasPlayed) {
                ctx.state.Websocket.send(JSON.stringify({
                    Command: "SELECT_CARD",
                    Arguments: JSON.stringify({
                        SelectedCard
                    })
                }))
                ctx.commit('setSelected', SelectedCard)
            }
        },
        startGame(ctx) {
            ctx.state.Websocket.send(JSON.stringify({ Command: "START_GAME", Arguments: "{}" }))
        },
        setVote(ctx, card) {
            ctx.state.Websocket.send(JSON.stringify({
                Command: "VOTE",
                Arguments: JSON.stringify({
                    Vote: card
                })
            }))
        }
    }
})