import Vue from 'vue';

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
            // Only set the User state if the device is not in deporte mode
            // We keep the dummy data on top in this case
            if (payload.User !== undefined && payload.User !== null)
                Object.assign(state, payload.User);
        },

        toggleSelection: (state, { AmtCardsRequired, payload }) => {
            if (AmtCardsRequired > 1) {
                Vue.set(state.Hand, payload, { ...state.Hand[payload], isSelected: !state.Hand[payload].isSelected });
            }
        }
    },
    actions: {

    },
    getters: {

    }
};