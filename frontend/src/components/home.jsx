import React, { useState } from 'react';
import { toast } from 'react-toastify';

import onMessage from '../commands/commands';
import { useReducers } from '../context';
import { setWebSocketAction } from '../actions/ws';

import logo from '../assets/logo.png';
import '../assets/home.css';
import 'react-toastify/dist/ReactToastify.css';

const WSURL = "ws://localhost:8080";

const initialState = {
    username: "",
    room: "",
}

export default function () {
    const [global, dispatch] = useReducers()
    const [state, setState] = useState(initialState);

    const submit = (e) => {
        e.preventDefault();

        console.log(e)
        let ws = new WebSocket(WSURL + "/?username=" + encodeURI(state.username) + "&room=" + encodeURI(state.room));

        ws.onclose = () => {
            // @TODO display connection lost modal
            console.log("Connection lost")
        }

        ws.onmessage = (e) => onMessage(dispatch, e);
        dispatch(setWebSocketAction(ws));

        return false;
    };

    return <div className="homepage">
        <img src={logo} />
        <h1>Cards Against Overflow</h1>
        <form onSubmit={submit}>
            <input type="text" id="username" name="username" placeholder="Pseudo" value={state.username} onChange={(e) => setState({ ...state, username: e.target.value })} required />
            <input type="text" id="room" name="room" placeholder="Code salle (Optionnel)" value={state.room} onChange={(e) => setState({ ...state, room: e.target.value })} />

            <div className="buttons">
                <input type="submit" value="Rejoindre" />
            </div>
        </form>
    </div>;
}