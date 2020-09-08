export function connect(e) {
    let store = this.$store;
    let toasted = this.$toasted;

    e.preventDefault();

    let url = location.host;

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development')
        url = "192.168.1.12:8000"
        //url = "localhost:8000"

    let ws = new WebSocket("ws://"+url+"/api?username=" + this.username + "&room=" + this.room)
    ws.onmessage = (e) => parseMessage(store, toasted, e.data);
    ws.onerror = (e) => console.log("ERR: ", e);
    ws.onclose = (e) => console.log("Connection closed: ", e)

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