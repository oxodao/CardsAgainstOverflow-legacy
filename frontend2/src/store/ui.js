export default {
    state: () => ({
        MenuVisible: false,
        LoggedIn: false,
        ShowRules: false,
        Wizz: [],
        CanWizz: true,
        WebSocket: null,
        CurrentCountdown: 0,
        SelectedCards: [],
    }),
    mutations: {
        showRules: (state) => {
            state.ShowRules = !state.ShowRules;
        },

        toggleMenu: (state) => {
          state.MenuVisible = !state.MenuVisible;
        },

        setWebsocket: (state, payload) => {
            state.WebSocket = payload;
        },

        SET_GAMESTATE: (state) => {
            state.LoggedIn = true;
        },

        WIZZ: (state, payload) => {
          if (!state.Wizz.includes(payload))
              state.Wizz.push(payload)
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
    },
    actions: {
        sendWizz: ({state, commit}) => {
            state.WebSocket.send(JSON.stringify({
                Command: 'WIZZ',
                Arguments: '{}'
            }));
            commit('canWizz', false)
        }
    },
    getters: {
        DisplayWizz: state => state.Wizz.length > 0 ? "wizz" : "",
    }
};