export default {
    state: () => ({
        Hand: [],
        HasPlayed: false,
        IsAdmin: false,
        IsJudge: false,
        RerollTimeout: 0,
        Score: 0,
        Username: "",
    }),
    mutations: {
        SET_GAMESTATE: (state, payload) => {
            Object.assign(state, payload.User);
        }
    },
    actions: {

    },
    getters: {

    }
};