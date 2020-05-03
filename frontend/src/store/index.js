import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        User: {
            Username: "",
            IsAdmin: false,
            IsJudge: false,
            Hand: [],
        },
        Room: {
            IsStarted: false,
            RoomID: "",
            Participants: [],
            TurnState: 0,
            CurrentCountdown: "",
        },
        SelectedCards: [],
        ShowLogin: true,
        Websocket: null,
    },
    mutations: {
        setWebsocket: (state, payload) => {
            state.Websocket = payload;
        },
        setPlayerList: (state, payload) => {
            state.Room.Participants = payload
        },
        setState: (state, payload) => {
            let ws = state.Websocket;

            state.User = payload.User;
            state.Room = payload.Room;

            state.Websocket = ws;
            state.ShowLogin = false;
            state.SelectedCards = [ ];
        },
        setCountdown: (state, payload) => {
            Vue.set(state.Room, 'CurrentCountdown', payload)
        },
        toggleSelection: (state, payload) => {
            if (state.SelectedCards.includes(payload)) {
                Vue.set(state.SelectedCards, state.SelectedCards.indexOf(payload), -1);
                state.User.Hand.isSelected = false;
            } else if (state.Room.CurrentBlackCard.AmtCardRequired > (state.SelectedCards.filter(e => e != -1).length)) {
                if (state.SelectedCards.includes(-1)) {
                    Vue.set(state.SelectedCards, state.SelectedCards.indexOf(-1), payload);
                } else {
                    state.SelectedCards.push(payload);
                }
                state.User.Hand.isSelected = true;
            }
        }
    },
    actions: {
        startGame(ctx) {
            ctx.state.Websocket.send(JSON.stringify({ Command: "START_GAME", Arguments: "{}" }))
        },
        select: (ctx, payload) => {
            ctx.commit('toggleSelection', payload);
            ctx.dispatch('sendSelection');
        },
        sendSelection: (ctx) => {
            if (ctx.state.User.IsJudge) {
                ctx.state.Websocket.send(JSON.stringify({
                    Command: 'SEND_SELECTION',
                    Arguments: JSON.stringify(ctx.state.SelectedCards),
                }))
            } else {
                console.log("Sending selection")
                ctx.state.Websocket.send(JSON.stringify({
                    Command: 'SEND_SELECTION',
                    Arguments: JSON.stringify(ctx.state.SelectedCards),
                }))
            }
        }
    },
    getters: {
        IsPlayerJudge: state => state.User.IsJudge,
        IsPlayerAdmin: state => state.User.IsAdmin,
    }
})