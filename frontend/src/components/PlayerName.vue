<template>
    <li v-bind:class="isMe">
        <div>
            <img v-if="hasPlayed" src="../assets/card_icon.png" alt="Has played"/>
            <img v-if="isWizzing" src="../assets/msn_wizz.png" alt="Wizzing"/>
            <img v-if="isAdmin" src="../assets/admin.png" alt="Admin"/>
            <img v-if="isJudge" src="../assets/gavel.png" alt="Judge"/>
            <span>{{ username }}</span>
        </div>
        <span class="score">{{ score }}</span>
    </li>
</template>

<script>
export default {
    name: 'PlayerName',
    props: [
        'username',
        'score',
        'isAdmin',
        'isJudge',
        'hasPlayed'
    ],
    computed: {
        isMe() {
            console.log(this.username, this.$store.state.User.Username, this.username === this.$store.state.User.Username)
            return this.username === this.$store.state.User.Username ? "isMe" : ""
        },
        isWizzing() {
            return this.$store.state.Wizz.includes(this.username);
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
            width: 5em;
            border-left: 1px solid #111;
            text-align: center;
        }
    }
</style>