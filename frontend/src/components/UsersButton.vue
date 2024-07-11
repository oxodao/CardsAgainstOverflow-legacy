<template>
    <button
        v-tooltip="'Joueurs'"
        @click="toggleList"
    >
        <img
            src="../assets/user-friends.png"
            alt="Users"
        />
        <span class="amtUsers">{{ amtUsers }}</span>
    </button>
</template>

<script>
import {mapState} from 'vuex';

export default {
    name: 'UsersButton',
    computed: {
        ...mapState({
            amtUsers: state => state.Room.Participants.length,
        })
    },
    methods: {
        toggleList(e) {
            this.$store.commit('togglePlayersMenu' );
            e.stopPropagation();
        }
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

  .amtUsers {
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