<template>
  <h1 v-if="Started && currBlackCard !== undefined && currBlackCard !== null && IsReady">
    <span v-for="txt in getCardText()" v-bind:key="txt.Question" v-bind:class="txt.Class">{{ txt.Question }}</span>
    <span v-if="TurnState === 2" class="colored-blue"><br/>{{ winner }}</span>
  </h1>
</template>

<script>
import {mapGetters, mapState} from "vuex";

export default {
  name    : "BlackCard",
  computed: {
    ...mapState({
      Started      : state => state.Room.Started,
      currBlackCard: state => state.Room.CurrentBlackCard,
      TurnState    : state => state.Room.TurnState,
      winner       : state => state.Room.Winner,
    }),
    ...mapGetters([
      'IsReady'
    ])
  },
  methods : {
    getCardText() {
      if (this.TurnState === 0)
        return [
          {
            "Question": this.currBlackCard.Text,
            "Class"   : "",
          }
        ];

      let txt = this.currBlackCard.Text;
      let txtSplitted = txt.split(/(____)/)
      let curr = 0;
      let values = [];
      txtSplitted.forEach(e => {
        if (e === "____") {
          let Question = "";
          let winning = this.$store.state.Room.WinningAnswer;
          if (winning !== undefined && winning !== null) {
            Question = winning.Cards[curr].Text
          } else {
            Question = this.$store.state.Room.Answers[this.$store.state.Room.JudgeSelection].Cards[curr].Text
          }
          values.push({
            Question,
            Class: 'colored-green'
          })
          curr++
        } else
          values.push({Question: e, Class: ''})
      });

      return values
    },
  }
}
</script>

<style scoped>
h1 {
  text-align: center;
  margin: .5em;
}

.colored-green {
  color: #3EC480;
  text-shadow: #3EC480 0px 0px 10px;
}

.colored-blue {
  color: #4dbbc7;
  text-shadow: #4dbbc7 0px 0px 10px;
}
</style>