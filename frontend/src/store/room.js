import Vue from 'vue';

export default {
    state: () => ({
        Answers: null,
        AvailableDecks: [],
        CurrentBlackCard: null,
        DefaultCountdown: 80,
        DefaultRerollTimeout: 4,
        JudgeSelection: -1,
        MaxTurn: 10,
        Participants: [],
        RoomID: '',
        Started: false,
        Turn: 0,
        TurnState: 0,
        Winner: '',
        WinningAnswer: null,
        ZenMode: false,
    }),
    mutations: {
        SET_GAMESTATE: (state, payload) => {
            Object.assign(state, payload.Room);
        },

        HAS_PLAYED: (state, payload) => {
            for (let i = 0; i < state.Participants.length; i++){
                if (state.Participants[i].Username === payload) {
                    Vue.set(state.Participants, i, { ...state.Participants[i], HasPlayed: true });
                }
            }
        },

        JUDGE_SELECTION: (state, payload) => {
            state.JudgeSelection = payload;
        },

        /**
         * Settings mutations
         */
        GOT_SETTINGS: (state, payload) => {
            state.MaxTurn = payload.MaxTurn;
            state.ZenMode = payload.ZenMode;
            state.DefaultCountdown = payload.DefaultCountdown;
            state.DefaultRerollTimeout = payload.DefaultRerollTimeout;

            for (let i = 0; i < state.AvailableDecks.length; i++) {
                let e = state.AvailableDecks[i];

                if (payload.SelectedDecks.includes(e.ID)) {
                    Vue.set(state.AvailableDecks[i], 'IsSelected', true);
                } else {
                    Vue.set(state.AvailableDecks[i], 'IsSelected', false);
                }
            }
        },
        updateTurns: (state, payload) => {
            state.MaxTurn = parseInt(payload) ?? 30;
        },
        updateZenMode: (state, payload) => {
            state.ZenMode = payload;
        },
        updateCountdown: (state, payload) => {
            state.DefaultCountdown = parseInt(payload) ?? 80;
        },
        updateRerollTimeout: (state, payload) => {
            state.DefaultRerollTimeout = parseInt(payload) ?? 6;
        },
        updateSelectedDecks: (state, payload) => {
            payload.ID = parseInt(payload.ID);

            for (let i = 0; i < state.AvailableDecks.length; i++) {
                if (state.AvailableDecks[i].ID === payload.ID) {
                    Vue.set(state.AvailableDecks, i, { ...state.AvailableDecks[i], IsSelected: payload.Selected });
                }
            }
        },
    },
    actions: {
        startGame({rootState}) {
            rootState.UI.WebSocket.send(JSON.stringify({ Command: 'START_GAME', Arguments: '{}' }));
        },
        sendSettings({rootState, state}) {
            let settings = {
                SelectedDecks: state.AvailableDecks.filter(e => e.IsSelected).map(e => e.ID),
                MaxTurn: state.MaxTurn,
                ZenMode: state.ZenMode,
                DefaultCountdown: state.DefaultCountdown,
                DefaultRerollTimeout: state.DefaultRerollTimeout,
            };

            rootState.UI.WebSocket.send(JSON.stringify({
                Command: 'SET_SETTINGS',
                Arguments: JSON.stringify(settings)
            }));
        },
    },
    getters: {
        IsReady: state => state.Participants.length >= 3,
        IsPlaying: state => state.ZenMode || state.Turn < state.MaxTurn,
        AmtCards: state => {
            let sumBlack = 0;
            let sumWhite = 0;

            let decks = state.AvailableDecks;

            if (!decks) {
                return { sumBlack: 0, sumWhite: 0 };
            }

            for (let i = 0; i < decks.length; i++) {
                if (decks[i].IsSelected) {
                    sumBlack += decks[i].AmtBlack;
                    sumWhite += decks[i].AmtWhite;
                }
            }

            return { sumBlack, sumWhite };
        },
        HasEnoughCards: (state, getters) => {
            return getters.AmtCards.sumWhite >= ((state.Participants.length+1) * 7);
        }
    }
};