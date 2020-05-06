export function connect(e) {
    let store = this.$store;
    let toasted = this.$toasted;
    e.preventDefault();

    let url = ""
    if (process.env.NODE_ENV === "development") {
        url = "localhost:8000"
    } else {
        url = "cao-api.oxodao.fr"
    }

    let ws = new WebSocket("ws://"+url+"/api?username=" + this.username + "&room=" + this.room)
    ws.onmessage = (e) => parseMessage(store, toasted, e.data);
    ws.onerror = (e) => console.log("ERR: ", e);
    ws.onclose = (e) => console.log("Connection closed: ", e)

    window.setInterval(function() {
        ws.send(JSON.stringify({
            Command: 'PING',
            Argments: '{}'
        }))
    }, 5000);

    store.commit('setWebsocket', ws)
}

export function parseMessage(store, toasted, msg) {
    let cmd = JSON.parse(msg)
    cmd.Arguments = JSON.parse(cmd.Arguments)
    
    switch(cmd.Command) {
        case 'ERROR':
            console.log("Error: ", cmd.Arguments)
            //toasted.show(msg.Arguments) // Not working for some reasons, no errors
            alert(cmd.Arguments)
            break;

        case 'CRITICAL_ERROR':
            console.log("Error: ", cmd.Arguments)
            //toasted.show(msg.Arguments) // Not working for some reasons, no errors
            alert(cmd.Arguments)

            break;

        case 'CONNECTED':
            store.commit('connected', cmd.Arguments)
            break;

        case 'SET_GAMESTATE':
            store.commit('setState', cmd.Arguments)
            break;

        case 'PLAYER_LIST':
            store.commit('setPlayerList', cmd.Arguments)
            break;

        case 'GOT_SETTINGS':
            store.commit('gotSettings', cmd.Arguments);
            break;

        case 'COUNTDOWN':
            store.commit('setCountdown', cmd.Arguments)
            break;

        case 'JUDGE_SELECTION':
            store.commit('setJudgeSelection', cmd.Arguments)
            break;

        default:
            console.log("UNHANDLED COMMAND: " + cmd.Command)
    }
}