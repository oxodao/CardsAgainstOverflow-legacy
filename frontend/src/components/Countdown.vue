<template>
    <div v-if="countdown !== '?'" v-bind:class="isNearEnd ? 'near' : ''">
        {{ countdown }}
    </div>
</template>

<script>
export default {
    name: 'Countdown',
    computed: {
        countdown() {
            return this.$store.state.Room.CurrentCountdown ?? '?';
        },
        isNearEnd() {
            return this.$store.state.Room.CurrentCountdown < 10;
        }
    }
}
</script>

<style lang="scss" scoped>
    div {
        z-index: 99999;
        width: 3em;
        height: 3em;
        border-radius: 50%;
        background: rgba(0, 0, 0, .75);
        border: 1px solid rgba(#cacaca, .25);
        text-align: center;
        line-height: 3em;

        position: absolute;
        bottom: 1em;
        right: 1em;
    }

    .near {
        background: rgba(#eb5454, .75);
        border: 1px solid rgba(#ff0000, .25);
    }
</style>