import { toast } from "react-toastify"

import {connectedAction} from '../actions/ws';

export default function onMessage(dispatch, msg) {
    let obj = JSON.parse(msg.data)
    obj.Arguments = JSON.parse(obj.Arguments)

    switch (obj.Command) {
        case 'ERROR': 
            toast.error(obj.Arguments)
            break;
        case 'CONNECTED':
            dispatch(connectedAction(obj.Arguments))
            break;
        default:
            console.log("UNHANDLED COMMAND: " + obj.Command)
    }
}