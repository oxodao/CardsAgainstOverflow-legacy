import { SET_WEBSOCKET, CONNECTED } from "../actions/ws"

export default (state, action) => {

    switch(action.type) {
        case SET_WEBSOCKET:
            return { ...state, Websocket: action.payload };
        case CONNECTED:
            console.log({ ...state, User: { ...state.User, ...action.payload}})
            return { ...state, User: { ...state.User, ...action.payload}}
    }

    return state
}