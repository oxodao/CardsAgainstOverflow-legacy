<template>
  <div id="modal-login">
    <img src="../assets/logo.png"/>
    <h1>Cards Against Overflow</h1>
    <form @submit="auth">
        <input v-model="username" id="username" type="text" name="username" placeholder="Pseudo" required/>
        <input v-model="room" id="room" type="text" name="room" placeholder="Code salle (Optionnel)"/>
        <input type="submit" value="Rejoindre" v-on:submit="auth" />
    </form>
  </div>
</template>

<script>
import {connect} from '../api/ws';

export default {
  name: 'HelloWorld',
  data: () => ({
    username: '',
    room: '',
  }),
  methods: {
    auth: connect
  }
}
</script>

<style lang="scss" scoped>
  #modal-login {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
    text-align: center;

    img {
      height: 3em;
    }

    form {
      border: 1px solid #045757;
      border-radius: .5em;
      padding: 1em;
      background: #111;
      display: flex;
      flex-direction: column;

      #username {
        display: block;
        margin-bottom: .5em;
      }

      #room {
        text-transform: uppercase;
        display: block;
        margin-bottom: 1em;

        &::placeholder {
          text-transform: none;
        }
      }

      input[type="submit"] {
        border: none;
        background: #045757;
        color: #111;
        padding: 1em;
        font-weight: bold;
      }
    }
  }
</style>
