export default {
    state: () => ({
        Answers: null,
        AvailableDecks: null,
        CurrentBlackCard: null,
        DefaultCountdown: 80,
        DefaultRerollTimeout: 4,
        JudgeSelection: -1,
        MaxTurn: 10,
        Participants: [],
        RoomID: "",
        Started: false,
        Turn: 0,
        TurnState: 0,
        Winner: "",
        WinningAnswer: null,
        ZenMode: false,
    }),
    mutations: {
        SET_GAMESTATE: (state, payload) => {
            Object.assign(state, payload.Room);
        }
    },
    actions: {

    },
    getters: {

    }
};