<template>
  <div id="ExitDialog">
    <p>Vous allez quitter la partie!</p>
    <button @click="closeDialog">Retour au jeu</button>
    <button id="quit" @click="exit">Je me barre!</button>
  </div>
</template>

<script>
export default {
  name: "ExitDialog",
  methods: {
    closeDialog() {
      this.$store.commit('hideCloseDialog');
    },
    exit(e) {
      e.preventDefault();
      window.location.reload();
      return false;
    }
  }
}
</script>

<style lang="scss" scoped>

#ExitDialog {
  position: absolute;
  width: 300px;
  background: rgba(black, .8);
  border: 1px solid #3EC480;
  top: 50%;
  left: 50%;
  transform: translateX(-50%) translateY(-50%);

  text-align: center;

  padding: 2em;

  p {
    padding: 0;
    margin: 0 0 1em;
  }

  #quit {
    background: #ff3333;
  }
}

</style>