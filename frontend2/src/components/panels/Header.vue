<template>
    <header>
        <h1 id="game-title">Cards Against Overflow</h1>

        <div>
            <span>Code salle: {{ RoomID }}</span>
        </div>
    </header>
</template>

<script>
    import {mapState} from "vuex";

    export default {
        name: "Header",
        computed: {
            ...mapState({
                RoomID: state => state.Room.RoomID,
            })
        }
    }
</script>

<style lang="scss" scoped>
    header {
        position: relative;
        background: #111;
        text-align: center;

        div {
            position: absolute;
            top: 0;
            bottom: 0;
            right: 0;

            padding-right: .5em;

            display: flex;
            flex-direction: column;

            align-items: flex-end;
            justify-content: center;

            span, button {
                display: block;
            }
        }
    }

    #game-title {
        font-size: 1.5em;

        @media (max-width: 650px) {
            font-size: 1.2em;
            text-align: left;
        }
    }
</style>