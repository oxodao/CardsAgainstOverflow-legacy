import React from 'react';
import MainReducer from '../reducers/MainReducer';

export const CtxInitialState = {
    Main: {
        CurrentRoom: ''
    },
    User: {
        Username: '',
    },
    Websocket: null
}

export const AppContext = React.createContext(CtxInitialState);
export const useStateValue = () => React.useContext(AppContext);
export const useReducers = () => React.useReducer(MainReducer, CtxInitialState);