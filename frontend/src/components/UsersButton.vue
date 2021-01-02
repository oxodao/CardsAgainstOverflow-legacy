<template>
  <button @click="toggleList" v-tooltip="'Joueurs'">
    <img src="../assets/user-friends.png" alt="Users"/>
  </button>
</template>

<script>
export default {
  name: "UsersButton",
  methods: {
    toggleList(e) {
      this.$store.commit('togglePlayersMenu' );
      e.stopPropagation();
    }
  }
}
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