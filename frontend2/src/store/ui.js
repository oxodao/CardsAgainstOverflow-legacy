export default {
    state: () => ({
        LoggedIn: false,
        ShowRules: false,
        Wizz: [],
        CanWizz: true,
        WebSocket: null,
    }),
    mutations: {
        showRules: (state) => {
            state.ShowRules = !state.ShowRules;
        },

        setWebsocket: (state, payload) => {
            state.WebSocket = payload;
        },

        SET_GAMESTATE: (state) => {
            state.LoggedIn = true;
        }
    },
    actions: {

    },
    getters: {
        DisplayWizz: state => state.Wizz.length > 0 ? "wizz" : "",
    }
};