import Vue from 'vue';

export default {
    state: () => ({
        Deporte: false,
        LoggedIn: false,
        ShowRules: false,
        Wizz: [],
        CanWizz: true,
        WebSocket: null,
        CurrentCountdown: 0,
        SelectedCards: [],
        LostConnection: false,
        PlayersMenuVisible: false,
    }),
    mutations: {
        showRules: (state) => {
            state.ShowRules = !state.ShowRules;
            state.LostConnection = true;
        },

        setDeporte: (state, payload) => {
          state.Deporte = payload;
        },

        togglePlayersMenu: (state) => {
            state.PlayersMenuVisible = !state.PlayersMenuVisible;
        },

        connectionClosed: (state) => {
            state.WebSocket = null;
        },

        setWebsocket: (state, payload) => {
            state.WebSocket = payload;
        },

        SET_GAMESTATE: (state) => {
            state.LoggedIn = true;
            state.SelectedCards = [];
        },

        WIZZ: (state, payload) => {
          if (!state.Wizz.includes(payload))
              state.Wizz.push(payload)
        },

        WIZZ_REFILLED: (state) => {
            state.CanWizz = true;
        },

        canWizz: (state, payload) => {
            state.CanWizz = payload;
        },

        delWizz: (state, payload) => {
            state.Wizz = state.Wizz.filter(e => e !== payload);
        },

        COUNTDOWN: (state, payload) => {
            state.CurrentCountdown = payload;
        },

        toggleSelection: (state, { AmtCardsRequired, payload }) => {
            // If we have only one card, no need to do weird things, just switching the current one
            if (AmtCardsRequired === 1) {
                if (state.SelectedCards.includes(payload)) {
                    state.SelectedCards = [];
                } else {
                    state.SelectedCards = [payload];
                }
                return
            }

            if (state.SelectedCards.includes(payload)) {
                Vue.set(state.SelectedCards, state.SelectedCards.indexOf(payload), -1);
                //state.User.Hand.isSelected = false;
            } else if (AmtCardsRequired > (state.SelectedCards.filter(e => e !== -1).length)) {
                if (state.SelectedCards.includes(-1)) {
                    Vue.set(state.SelectedCards, state.SelectedCards.indexOf(-1), payload);
                } else {
                    state.SelectedCards.push(payload);
                }
                //state.User.Hand.isSelected = true;
            }
        },
        toggleProposalSelection: (state, payload) => {
            state.SelectedCards = [payload];
        },
    },
    actions: {
        sendWizz: ({state, commit}) => {
            state.WebSocket.send(JSON.stringify({
                Command: 'WIZZ',
                Arguments: '{}'
            }));
            commit('canWizz', false)
        },
        reroll: (ctx) => {
            ctx.state.WebSocket.send(JSON.stringify({
                Command: 'REROLL',
                Arguments: '{}'
            }));
        },
        select: ({rootState, commit, dispatch}, payload) => {
            commit('toggleSelection', {AmtCardsRequired: rootState.Room.CurrentBlackCard.AmtCardRequired, payload});
            dispatch('sendSelection');
        },
        selectProposal: (ctx, payload) => {
            ctx.commit('toggleProposalSelection', payload);
            ctx.dispatch('sendSelection');
        },
        sendSelection: ({rootState, state}) => {
            /** @TODO: Rendondant? **/
            if (rootState.User.IsJudge) {
                state.WebSocket.send(JSON.stringify({
                    Command: 'SEND_SELECTION',
                    Arguments: JSON.stringify(state.SelectedCards),
                }))
            } else {
                state.WebSocket.send(JSON.stringify({
                    Command: 'SEND_SELECTION',
                    Arguments: JSON.stringify(state.SelectedCards),
                }))
            }
        },
        skipCountdown: (ctx) => {
            ctx.state.WebSocket.send(JSON.stringify({
                Command: 'SKIP_COUNTDOWN',
                Arguments: '{}'
            }))
        },
    },
    getters: {
        DisplayWizz: state => state.Wizz.length > 0 ? "wizz" : "",
    }
};