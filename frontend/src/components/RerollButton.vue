<template>
    <button
        v-if="isStarted && !isDeporte"
        v-tooltip="'Re-piocher'"
        :disabled="!canReroll"
        @click="reroll"
    >
        <img
            src="../assets/reroll.png"
            alt="Reroll"
        />
        <span
            v-if="rerollTimeout > 0"
            class="rerollTimeout"
        >{{ rerollTimeout }}</span>
    </button>
</template>

<script>
import {mapState} from 'vuex';

export default {
    name    : 'RerollButton',
    computed: {
        ...mapState({
            isStarted    : state => state.Room.Started,
            isDeporte    : state => state.UI.Deporte,
            rerollTimeout: state => state.User.RerollTimeout,
        }),
        canReroll() {
            let state = this.$store.state;
            return !state.User.IsJudge && state.Room.TurnState === 0 && state.User.RerollTimeout === 0;
        },
    },
    methods: {
        reroll() {
            this.$store.dispatch('reroll');
        },
    }
};
</script>

<style lang="scss" scoped>
button {
  position: relative;
  width: 3em;
  height: 3em;

  img {
    width: 100%;
    height: 100%;
    object-fit: contain;
  }

  &:disabled {
    background: darken(#3EC480, 10%);
  }

  .rerollTimeout {
    $size: 2em;
    position: absolute;
    top: -25%;
    right: -25%;

    width: $size;
    height: $size;
    border-radius: 50%;
    background: rgba(darken(#3EC480, 25%), .8);
    border: 1px solid lighten(#3EC480, 25%);
    color: white;

    line-height: $size;
    text-align: center;
  }
}
</style>