<template>
    <div v-if="countdown > 0 && IsReady">
        {{ countdown }}
    </div>
</template>

<script>
    import {mapGetters, mapState} from "vuex";

    export default {
        name: "Countdown",
        computed: {
            ...mapState({
                         countdown: state => state.UI.CurrentCountdown
            }),
            ...mapGetters([
                'IsReady'
            ])
        }
    }
</script>

<style lang="scss" scoped>
    div {
        position: fixed;
        width: 2.5em;
        height: 2.5em;

        right: .5em;
        bottom: .5em;

        @media(max-width: 650px) {
          bottom: calc(.5em + 64px + .25em);
        }

        display: flex;
        align-items: center;
        justify-content: center;

        border-radius: 50%;
        border: 2px solid white;
        background: rgba(black, .5);
    }
</style>