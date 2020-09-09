export function connect(e) {
    let store = this.$store;
    let toasted = this.$toasted;
    let deporte = this.affichageDeporte.length > 0;

    console.log("DEPORTE: ", deporte);
    store.commit('setDeporte', deporte);

    e.preventDefault();

    // Building the URL
    let url = "ws://" + location.host + "/";

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development')
        url = "ws://192.168.1.12:8000/"

    if (!deporte) {
        url += "api?username=" + this.username + "&";
    } else {
        url += "deporte?";
    }

    url += "room=" + this.room;

    // Connecting to it
    let ws = new WebSocket(url)
    ws.onmessage = (e) => parseMessage(store, toasted, e.data);
    ws.onerror = (e) => console.log("ERR: ", e);
    ws.onclose = (e) => {
        console.log("Connection closed: ", e)
        store.commit("connectionClosed")
    }

    window.setInterval(function() {
        ws.send(JSON.stringify({
            Command: 'PING',
            Arguments: '{}'
        }))
    }, 5000);

    store.commit('setWebsocket', ws)
}

export function parseMessage(store, toasted, msg) {
    let cmd = JSON.parse(msg);
    cmd.Arguments = JSON.parse(cmd.Arguments);

    switch(cmd.Command) {
        case "ERROR":
        case "CRITICAL_ERROR":
            console.log(cmd.Command + ": ", cmd.Arguments);
            alert(cmd.Command + ": " + cmd.Arguments);
            break

        default:
            store.commit(cmd.Command, cmd.Arguments)
            break
    }
}