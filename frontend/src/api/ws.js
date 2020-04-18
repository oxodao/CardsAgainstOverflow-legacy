export function connect(e) {
    let store = this.$store;
    let toasted = this.$toasted;
    e.preventDefault();

    let ws = new WebSocket("ws://localhost:8000/api?username=" + this.username + "&room=" + this.room)
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

        case 'DISCONNECTED':
            console.log("Add to logs")
            break;

        case 'PLAYER_LIST':
            store.commit('setPlayerList', cmd.Arguments)
            break;

        case 'GAME_STARTED':
            store.commit('startGame', cmd.Arguments)
            break;

        case 'UPDATE_CARDS':
            store.commit('updateCards', cmd.Arguments)
            break;

        case 'ANSWERS_LIST':
            store.commit('addToAnswersList', cmd.Arguments)
            break;

        default:
            console.log("UNHANDLED COMMAND: " + cmd.Command)
    }
}