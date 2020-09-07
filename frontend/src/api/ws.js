export function connect(e) {
    let store = this.$store;
    let toasted = this.$toasted;
    e.preventDefault();

    let url = location.host;

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development')
        url = "localhost:8000"

    let ws = new WebSocket("ws://"+url+"/api?username=" + this.username + "&room=" + this.room)
    ws.onmessage = (e) => parseMessage(store, toasted, e.data);
    ws.onerror = (e) => console.log("ERR: ", e);
    ws.onclose = (e) => console.log("Connection closed: ", e)

    //if (process.env.NODE_ENV !== "development") {
        //localStorage.setItem('username', this.username);
    //}

    window.setInterval(function() {
        ws.send(JSON.stringify({
            Command: 'PING',
            Arguments: '{}'
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

        case 'WIZZ':
            store.commit('addWizz', cmd.Arguments)
            break;

        case 'WIZZ_REFILLED':
            store.commit('canWizz', true);
            break;

        case 'HAS_PLAYED':
            store.commit('hasPlayed', cmd.Arguments)
            break;

        default:
            console.log("UNHANDLED COMMAND: " + cmd.Command)
    }
}