<template>
    <div class="card" @click="toggleSelection(index)" v-bind:class="getClassnameIsSelected">
        <p :class="getClassForTextSize">{{ !isProposal ? currCard.Text : "Proposition #" + (index+1)}}</p>
        <div class="branding">
            <span v-for="t in getDeckName" :key="t">{{ t }}</span>
        </div>
        <div v-if="showPosition && !isProposal" id="position">
            {{ getSelectedPosition }}
        </div>
    </div>
</template>

<script>
export default {
    name: 'Card',
    props: [
        'currCard',
        'isProposal',
        'index',
    ],
    data: function() {
        return {
            currentCard: this.index,
        }
    },
    computed: {
        getClassnameIsSelected() {
            if (this.$store.state.Room.TurnState === 0) {
                return this.$store.state.SelectedCards.includes(this.currentCard) ? "isSelected" : "";
            }

            return this.$store.state.Room.JudgeSelection === this.index ? "isSelected" : "";
        },
        getSelectedPosition() {
            if (this.$store.state.Room.TurnState === 0) {
                return this.$store.state.SelectedCards.indexOf(this.currentCard) + 1;
            }

            return this.$store.state.Room.JudgeSelection
        },
        showPosition() {
            if (this.$store.state.Room.TurnState !== 0 || this.$store.state.User.IsPlayerJudge)
                return false

            return this.$store.state.Room.CurrentBlackCard.AmtCardRequired > 1 && this.$store.state.SelectedCards.includes(this.index)
        },
        getDeckName() {
            let decks = this.$store.state.Room.AvailableDecks
            for (let i = 0; i < decks.length; i++) {
                if (decks[i].ID === this.currCard.Deck) {
                    return decks[i].Title.split(" ")
                }
            }

            return "Réponses".split(" ")
        },
        getClassForTextSize() {
            // not seems to be working
            if (this.isProposal) return '';

            let size = this.currCard.Text.length;

            if (size >= 80) {
                return 'size2';
            } else if (size >= 55) {
                return 'size1';
            }

            return '';
        }
    },
    methods: {
        toggleSelection: function(card) {
            if (this.$store.state.Room.TurnState === 0) {
                this.$store.dispatch('select', card)
            } else if (this.$store.state.User.IsJudge && this.$store.state.Room.TurnState === 1) {
                this.$store.dispatch('selectProposal', card)
            }
        }
    }
}
</script>

<style lang="scss" scoped>
    /** Stolen from https://designshack.net/tutorialexamples/cardtricks/index.html */
    .card {
        user-select: none;
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


    .size1{
        font-size: .8em;
    }

    .size2 {
        font-size: .5em;
    }

</style>