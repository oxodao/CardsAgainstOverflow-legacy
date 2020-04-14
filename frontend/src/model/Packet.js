export default class Packet {
    
    constructor (type, value) {
        this.type = type;
        this.value = value;
    }

}

// Packet Login is sent when the user clicked on "Rejoindre"
// Value: { "username": "pseudo", "room": "salle" }
export const PACKET_LOGIN = "LOGIN";