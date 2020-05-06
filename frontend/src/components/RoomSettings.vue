<template>
    <div id="settings">
        <ul>
            <li v-for="deck in decks" v-bind:key="'deck'+deck.ID+deck.IsSelected"><input type="checkbox" v-bind:disabled="!IsAdmin" v-bind:checked="deck.IsSelected" v-bind:id="'cb'+deck.ID" @input="updateDeckSelection" :data-id="deck.ID"/><Label v-bind:for="'cb'+deck.ID">{{deck.Title}}</Label></li>
        </ul>
        <div id="subsettings">
            <div id="subsubsettings">
                <Label for="amtTurns">Nombre de tours:</Label>
                <div>
                    <input id="amtTurns" :disabled="!IsAdmin || zenMode" type="number" min="1" :value="maxTurn" @input="updateTurns" />
                    <input id="cbZenmode" :disabled="!IsAdmin" type="checkbox" :checked="zenMode" @input="updateZenMode"/>
                    <Label for="cbZenmode">Mode Zen</Label>
                </div>
                <Label for="countdownDuration">Durée de choix (en secondes)</Label>
                <div>
                    <input id="countdownDuration" type="number" min="10" :disabled="!IsAdmin" :value="countdown" @input="updateCountdown" />
                </div>
                <Label for="rerollTimeout">Nombre de tour entre re-roll</Label>
                <div>
                    <input id="rerollTimeout" type="number" min="0" :disabled="!IsAdmin" :value="rerollTimeout" @input="updateReroll" />
                </div>
            </div>
            <button v-if="IsAdmin && isReady" @click="startGame">Démarrer</button>
        </div>
    </div>
</template>

<script>
    import {mapState} from "vuex";

    export default {
        name: 'RoomSettings',
        computed: {
            ...mapState({
                IsAdmin: state => state.User.IsAdmin,

                decks: state => state.Room.AvailableDecks,
                maxTurn: state => state.Room.MaxTurn,
                zenMode: state => state.Room.ZenMode,
                countdown: state => state.Room.DefaultCountdown,
                rerollTimeout: state => state.Room.DefaultRerollTimeout,

                selectedDecks: state => state.SelectedDecks,
            }),
            isReady() {
                return this.$store.state.Room.Participants.length >= 3
            },
        },
        methods: {
            startGame() {
                this.$store.dispatch('startGame');
            },
            updateTurns(e) {
                this.$store.commit('updateTurns', e.target.value);
                this.$store.dispatch('sendSettings');
            },
            updateZenMode(e) {
                this.$store.commit('updateZenMode', e.target.checked);
                this.$store.dispatch('sendSettings');
            },
            updateCountdown(e) {
                this.$store.commit('updateCountdown', e.target.value);
                this.$store.dispatch('sendSettings');
            },
            updateReroll(e) {
              this.$store.commit('updateRerollTimeout', e.target.value);
              this.$store.dispatch('sendSettings');
            },
            updateDeckSelection(e) {
                this.$store.commit('updateSelectedDecks', { ID: e.target.getAttribute('data-id'), Selected: e.target.checked });
                this.$store.dispatch('sendSettings');
            },
        }
    };
</script>

<style lang="scss" scoped>
    #settings {
        margin-bottom: 2em;
        padding: 1em;

        width: 80%;
        height: 300px;
        background: #333;
        border-radius: 16px;

        display: flex;
        flex-direction: row;
        align-items: stretch;

        ul {
            margin: 0;
            padding: .25em;
            background: #444;

            li {
                list-style: none;
            }
        }

        #subsettings {
            margin-left: 1em;

            display: flex;
            flex-direction: column;

            #subsubsettings{
                flex: 1;
                div {
                    margin: .5em 0 0 1em ;
                    display: flex;
                    flex-direction: row;
                    align-items: center;
                }
            }

            button {
                background: #3EC480;
                border: none;
                color: #111;
                padding: 1em;
                border-radius: 10px;
            }
        }

    }
</style>
