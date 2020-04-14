export const SET_WEBSOCKET = "SET_WEBSOCKET";
export const CONNECTED = "CONNECTED";

export const setWebSocketAction = (ws) => {
    return { type: SET_WEBSOCKET, payload: ws };
}

export const connectedAction = (payload) => {
    return { type: CONNECTED, payload };
}