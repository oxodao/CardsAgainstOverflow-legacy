<template>
    <h1
        v-if="Started && currBlackCard !== undefined && currBlackCard !== null && IsReady"
        :class="'size ' + Size"
    >
        <span
            v-for="txt in getCardText()"
            :key="txt.Question"
            :class="txt.Class"
        >{{ txt.Question }}</span>
        <span
            v-if="TurnState === 2"
            class="colored-blue"
        ><br />{{ winner }}</span>
    </h1>
</template>

<script>
import {mapGetters, mapState} from 'vuex';

export default {
    name    : 'BlackCard',
    data() {
        return {
            Size: 'Size1',
        };
    },
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
            let length = this.currBlackCard.Text.length || 100;
            if (length < 150) {
                this.Size = 'Size1';
            } else if (length >= 150) {
                this.Size = 'Size2';
            }

            if (this.TurnState === 0) {
                return [ 
                    { 'Question': this.currBlackCard.Text, 'Class': '' }
                ]; 
            }

            let txt = this.currBlackCard.Text;
            let txtSplitted = txt.split(/(____)/);
            let curr = 0;
            let values = [];

            txtSplitted.forEach(e => {
                if (e === '____') {
                    let Question = '';
                    let winning = this.$store.state.Room.WinningAnswer;

                    if (winning !== undefined && winning !== null) {
                        Question = winning.Cards[curr].Text;
                    } else {
                        Question = this.$store.state.Room.Answers[this.$store.state.Room.JudgeSelection].Cards[curr].Text;
                    }

                    curr++;

                    values.push({
                        Question,
                        Class: 'colored-green'
                    });
                } else
                {values.push({Question: e, Class: ''});}
            });

            return values;
        },
    }
};
</script>

<style lang="scss" scoped>
h1 {
  text-align: center;
  margin: .5em;
}

$sizes: ("Size1" 2em), ("Size2" 1em);
@each $i, $val in $sizes {
  .#{$i} {
    @media(max-width: 650px) {
      font-size: $val;
    }
  }
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