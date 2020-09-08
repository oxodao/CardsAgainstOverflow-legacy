<template>
    <div>
        {{ countdown }}
    </div>
</template>

<script>
    import {mapState} from "vuex";

    export default {
        name: "Countdown",
        computed: mapState({ countdown: state => state.UI.CurrentCountdown })
    }
</script>

<style lang="scss" scoped>
    div {
        position: absolute;
        width: 2.5em;
        height: 2.5em;

        right: .5em;
        bottom: .5em;

        display: flex;
        align-items: center;
        justify-content: center;

        border-radius: 50%;
        border: 2px solid white;
        background: rgba(black, .5);
    }
</style>