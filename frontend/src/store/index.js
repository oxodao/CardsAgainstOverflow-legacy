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
            JudgeSelection: 0,
            Turn: 0,
            MaxTurn: 30,
            ZenMode: false,
            DefaultCountdown: 80,
            AvailableDecks: [],
        },
        SelectedCards: [],
        ShowLogin: true,
        ShowRules: false,
        Websocket: null,
    },
    mutations: {
        setWebsocket: (state, payload) => {
            state.Websocket = payload;
        },
        setPlayerList: (state, payload) => {
            state.Room.Participants = payload;
        },
        showRules: (state) => {
            state.ShowRules = !state.ShowRules;
            console.log("Showing rules; " +state.ShowRules)
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
            Vue.set(state.Room, 'CurrentCountdown', payload);
        },
        toggleSelection: (state, payload) => {
            // If we have only one card, no need to do weird things, just switching the current one
            if (state.Room.CurrentBlackCard.AmtCardRequired === 1) {
                if (state.SelectedCards.includes(payload)) {
                    state.SelectedCards = [];
                } else {
                    state.SelectedCards = [payload];
                }
                return
            }

            if (state.SelectedCards.includes(payload)) {
                Vue.set(state.SelectedCards, state.SelectedCards.indexOf(payload), -1);
                state.User.Hand.isSelected = false;
            } else if (state.Room.CurrentBlackCard.AmtCardRequired > (state.SelectedCards.filter(e => e !== -1).length)) {
                if (state.SelectedCards.includes(-1)) {
                    Vue.set(state.SelectedCards, state.SelectedCards.indexOf(-1), payload);
                } else {
                    state.SelectedCards.push(payload);
                }
                state.User.Hand.isSelected = true;
            }
        },
        toggleProposalSelection: (state, payload) => {
            state.SelectedCards = [payload];
        },
        setJudgeSelection: (state, payload) => {
            state.Room.JudgeSelection = payload;
        },

        /**
         * Settings mutations
         */
        gotSettings: (state, payload) => {
            state.Room.MaxTurn = payload.MaxTurn;
            state.Room.ZenMode = payload.ZenMode;
            state.Room.DefaultCountdown = payload.DefaultCountdown;

            for(let i = 0; i < state.Room.AvailableDecks.length; i++) {
                let e = state.Room.AvailableDecks[i]
                if (payload.SelectedDecks.includes(e.ID)) {
                    Vue.set(state.Room.AvailableDecks[i], 'IsSelected', true)
                } else {
                    Vue.set(state.Room.AvailableDecks[i], 'IsSelected', false)
                }
            }
        },
        updateTurns: (state, payload) => {
            state.Room.MaxTurn = parseInt(payload) ?? 30;
        },
        updateZenMode: (state, payload) => {
            state.Room.ZenMode = payload;
        },
        updateCountdown: (state, payload) => {
            state.Room.DefaultCountdown = parseInt(payload) ?? 80;
        },
        updateSelectedDecks: (state, payload) => {
            payload.ID = parseInt(payload.ID);

            for (let i = 0; i < state.Room.AvailableDecks.length; i++) {
                if (state.Room.AvailableDecks[i].ID === payload.ID) {
                    Vue.set(state.Room.AvailableDecks, i, { ...state.Room.AvailableDecks[i], IsSelected: payload.Selected })
                }
            }
        }
    },
    actions: {
        startGame(ctx) {
            ctx.state.Websocket.send(JSON.stringify({ Command: "START_GAME", Arguments: "{}" }))
        },
        sendSettings(ctx) {
            let settings = {
                SelectedDecks: ctx.state.Room.AvailableDecks.filter(e => e.IsSelected).map(e => e.ID),
                MaxTurn: ctx.state.Room.MaxTurn,
                ZenMode: ctx.state.Room.ZenMode,
                DefaultCountdown: ctx.state.Room.DefaultCountdown,
            };

            ctx.state.Websocket.send(JSON.stringify({
                Command: 'SET_SETTINGS',
                Arguments: JSON.stringify(settings)
            }))
        },
        select: (ctx, payload) => {
            ctx.commit('toggleSelection', payload);
            ctx.dispatch('sendSelection');
        },
        selectProposal: (ctx, payload) => {
            ctx.commit('toggleProposalSelection', payload);
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
        },
        skipCountdown: (ctx) => {
            ctx.state.Websocket.send(JSON.stringify({
                Command: 'SKIP_COUNTDOWN',
                Arguments: ''
            }))
        }
    },
    getters: {
        IsPlayerJudge: state => state.User.IsJudge,
        IsPlayerAdmin: state => state.User.IsAdmin,
        IsPlaying: state => (state.Room.ZenMode || state.Room.Turn <= state.Room.MaxTurn),
    }
})