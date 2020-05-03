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

let usernames = ["Towel_Bird","Whale_Leash","Water_System","Toolbox_Printer","Comics_Urine","Sink_Robot","Breakfast_Nuclear","Drugs_Towel","Flowers_Plants","Mail_Shoes","Shelf_Hnads","Drugs_Breakfast","Ice_cream_cone_Video_games","Toilet_Mail","Plus_Comics","Ring_Floppy_Disk","Clock_BBQ","Mail_YouTube","Book_Water","Crab_Cat","Printer_Cone","Shelf_Plants","Water_Shower","Post_office_Body","Cat_YouTube","Plus_Robot","Post_office_Sink","Prints_Urine","Kitty_Fence","Dog_Shoe","Trees_Nuclear","Website_Light_saber","Nuclear_Plus","Trees_Settings","System_Hnads","Leash_Light_saber","Website_Settings","Allergies_Horse","Soap_BBQ","Printer_Dislike","Poop_Dog","Plus_Leash","Urine_Fence","Ring_Male","Video_games_Website","System_Fence","Breakfast_Crab","Soap_Plus","Towel_Flowers","Video_games_Flowers","Flowers_Dislike","Hnads_Light_saber","Post_office_Crab","Ice_cream_cone_Running","Shoes_Cone","Shoe_Post_office","Body_Shoes","Post_office_Post_office","Toolbox_Urine","Running_Puppy","Water_Comics","Soap_Leash","Dislike_Fence","Rollers_Soap","Flowers_Horse","Soap_Shoes","Website_Soap","Fence_Settings","Dislike_Floppy_Disk","Toolbox_Robot","Cone_Male","Websites_Urine","Printer_YouTube","Crab_Running","Ring_Mail","Robot_Floppy_Disk","Shoe_Sink","Websites_Toilet","Water_Cone","Nuclear_Nuclear","Book_Elevator","Sink_Elevator","Cone_Trees","Plants_Ice_cream","Robot_Water","Toilet_Fence","Ice_cream_Poop","Allergies_Toolbox","Breakfast_Bird","Dislike_Book","Toolbox_Cone","Shoe_Male","Ring_Dog","Urine_Shoes","BBQ_Mail","Ice_cream_Book","Elevator_Poop","Poop_Hnads","Shoes_Soap","Whale_Trees","Clock_Allergies","Websites_Ice_cream","Allergies_Hnads","Toolbox_Boat","Dislike_Elevator","Body_Body","Websites_Horse","Horse_Horse","Puppy_Hnads","Toolbox_Poop","Towel_Water","Rollers_Mail","Dislike_Solar","Kitty_Solar","Ice_cream_Nuclear","Running_Toilet","Video_games_Dog","YouTube_Mail","Light_saber_Puppy","Puppy_Comics","Book_Poop","Cat_Flowers","Bird_Whale","Towel_Plants","Bird_Male","Urine_Fence","Body_BBQ","Water_Urine","YouTube_Fusion","Plus_Website","Boat_System","Prints_Comics","Ring_Floppy_Disk","Male_Settings","Websites_Nuclear","Monster_Urine","Whale_Book","Puppy_Kitty","Light_saber_Post_office","Puppy_Laptop","Poop_Poop","BBQ_Soap","Cat_Video_games","Plus_Book","Boat_Dislike","Kitty_Robot","Plants_Video_games","Toolbox_Video_games","Horse_Nuclear","System_System"];

export default {
  name: 'HelloWorld',
  data: () => ({
    username: usernames[Math.floor(Math.random() * usernames.length)], // TMP debug only, replace w/ empty str
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
