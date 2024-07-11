<template>
    <button
        v-if="!isDeporte"
        v-tooltip="'Wizz'"
        :disabled="!canWizz"
        @click="addWizz"
    >
        <img
            src="../assets/msn_wizz.png"
            alt="wizz"
        />
    </button>
</template>

<script>
import {mapState} from 'vuex';

export default {
    name: 'WizzButton',
    computed: {
        ...mapState({
            canWizz      : state => state.UI.CanWizz,
            isDeporte    : state => state.UI.Deporte,
        })
    },
    methods: {
        addWizz() {
            this.$store.dispatch('sendWizz');
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
}
</style>