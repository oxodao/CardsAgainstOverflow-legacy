<template>
    <div class="card" @click="toggleSelection(currCard)" v-bind:class="getClassnameIsSelected">
        <p>{{currCard.Text}}</p>
        <div class="branding">
            <span>Cards</span>
            <span>Against</span>
            <span>Overflow</span>
        </div>
        <div v-if="showPosition" id="position">
            {{ getSelectedPosition+1 }}
        </div>
    </div>
</template>

<script>
export default {
    name: 'Card',
    props: [
        'currCard',
        'isJudge'
    ],
    data: function() {
        return {
            currentCard: this.currCard,
            isPlayerJudge: this.isJudge
        }
    },
    computed: {
        getClassnameIsSelected() {
            if (this.isPlayerJudge)
                return this.currentCard.IsSelected ? "isSelected" : ""

            return this.currentCard.answerPosition !== -1 ? "isSelected" : ""
        },
        getTextCard() {
            return ""
        },
        getSelectedPosition() {
            return this.currentCard.answerPosition
        },
        showPosition() {
            return this.currentCard.answerPosition !== -1 && !this.isPlayerJudge && this.$store.state.Room.BlackCard.AmtCardRequired> 1 
        }
    },
    methods: {
        toggleSelection: function(card) {
            if (this.$store.getters.IsPlayerJudge)
                this.$store.commit('toggleAnswerSelection', card)
            else
                this.$store.commit('toggleSelection', card)
        }
    }
}
</script>

<style lang="scss" scoped>
    /** Stolen from https://designshack.net/tutorialexamples/cardtricks/index.html */
    .card {
        box-sizing: border-box;
        width: 130px;
        height: 200px;
        padding: 5px;
        border-radius: 10px;
        margin-right: 10px;
        transition: all .1s ease;
        position: relative;
        
        background: white;
        color: #111;

        &.isBlackCard {
            background: #111;
            color: white;
        }

        &:not(.isSelected):hover {
            transform: scale(1.05, 1.05);
            box-shadow: 1px 1px 7px rgba(0, 0, 0, .9);
        }

        .branding {
            left: 0;
            width: 100%;
            color: white;
            background: #111;
            position: absolute;
            bottom: 10px;
            text-align: right;
            font-size: .55em;
            padding: 2px 0 2px 0;

            span {
                display: block;
                margin-right: 5px;
            }
        }
    }

    .isSelected {
        transform: translateY(-20px);
        box-shadow: 1px 1px 7px rgba(0, 0, 0, .9);
    }

    #position {
        position: absolute;
        width: 2em;
        height: 2em;
        border-radius: 50%;
        background: #3EC480;
        border: 1px solid darken(#3EC480, 5%);

        top: -1em;
        right: -1em;

        text-align: center;
        line-height: 2em;

    }

</style>