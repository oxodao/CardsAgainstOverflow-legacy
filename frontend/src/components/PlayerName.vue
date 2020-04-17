<template>
    <li v-bind:class="isMe">
        <img v-if="isAdmin" src="@/assets/admin.png" alt="Admin"/>
        <img v-if="isJudge" src="@/assets/gavel.png" alt="Judge"/>
        <span>{{ username }}</span>
    </li>
</template>

<script>
export default {
    name: 'PlayerName',
    props: [
        'username',
        'isAdmin',
        'isJudge'
    ],
    data: function() {
        return {
            currentUsername: this.username,
        }
    },
    computed: {
        isMe() {
            return this.currentUsername == this.$store.state.User.Username ? "isMe" : ""
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

        img {
            height: 1.15em;
            margin-right: 10px;
        }

        &.isMe {
            background: #333;
        }
    }
</style>