<template>
    <li v-bind:class="isMe">
        <div>
            <img v-if="player.HasPlayed" src="../assets/card_icon.png" alt="Has played"/>
            <img v-if="isWizzing" src="../assets/msn_wizz.png" alt="Wizzing"/>
            <img v-if="player.IsAdmin" src="../assets/admin.png" alt="Admin"/>
            <img v-if="player.IsJudge" src="../assets/gavel.png" alt="Judge"/>
            <span>{{ player.Username }}</span>
        </div>
        <span class="score">{{ player.Score }}</span>
    </li>
</template>

<script>
    import {mapState} from "vuex";

    export default {
        name: 'PlayerName',
        props: [
            'player'
        ],
        computed: {
            ...mapState({
                IsDeporte: state => state.UI.Deporte,
            }),
            isMe() {
                return (!this.IsDeporte && (this.player.Username === this.$store.state.User.Username)) ? "isMe" : "";
            },
            isWizzing() {
                return this.$store.state.UI.Wizz.map(e => e.user).includes(this.player.Username);
            }
        }
    }
</script>

<style lang="scss" scoped>
    li {
        display: flex;
        justify-content: center;
        align-items: center;
        background: #282828;
        font-size: 1.15em;
        border-radius: 3px;
        margin-bottom: 5px;
        padding: 5px;

        div {
            flex: 1;
            text-align: center;
            img {
                height: 1.15em;
                margin-right: 10px;
            }
        }

        &.isMe {
            background: #444;
        }

        .score {
            width: 2.5em;
            border-left: 1px solid #111;
            text-align: center;
            margin-left: .75em;
        }
    }
</style>