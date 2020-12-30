<template>
    <div :class="'card ' + getClassnameIsSelected" @click="toggleSelection">
        <p :class="getClassForTextSize(IsProposal)">{{ IsProposal ? "Proposition #" + (Index+1) : card.Text }}</p>
        <div class="branding">
            <span v-for="t in getDeckName" :key="t">{{ t }}</span>
        </div>
        <div v-if="showPosition && !IsProposal" id="position">
            {{ getSelectedPosition }}
        </div>
    </div>
</template>

<script>
    export default {
        name: "Card",
        props: {
            card: Object,
            Index: Number,
            IsProposal: Boolean,
        },
        computed: {
            getClassnameIsSelected() {
                let state = this.$store.state;
                if (state.Room.TurnState === 0) {
                    return state.UI.SelectedCards.includes(this.Index) ? "isSelected" : "";
                }

                return state.Room.JudgeSelection === this.Index ? "isSelected" : "";
            },

            showPosition() {
                let state = this.$store.state;
                if (state.Room.TurnState !== 0 || state.User.IsJudge)
                    return false

                return state.Room.CurrentBlackCard.AmtCardRequired > 1 && state.UI.SelectedCards.includes(this.Index)
            },
            getSelectedPosition() {
                let state = this.$store.state;
                if (state.Room.TurnState === 0) {
                    return state.UI.SelectedCards.indexOf(this.Index) + 1;
                }

                return state.Room.JudgeSelection
            },
            getDeckName() {
                let decks = this.$store.state.Room.AvailableDecks;

                for (let i = 0; i < decks.length; i++) {
                    if (decks[i].ID === this.card.Deck) {
                        return decks[i].Title.split(" ");
                    }
                }

                return ["Réponses"];
            },
        },
        methods: {
            toggleSelection: function() {
                let state = this.$store.state;
                if (state.Room.TurnState === 0) {
                    this.$store.dispatch('select', this.Index)
                } else if (state.User.IsJudge && state.Room.TurnState === 1) {
                    this.$store.dispatch('selectProposal', this.Index)
                }
            },
            getClassForTextSize() {
                if (this.IsProposal) return '';

                let size = this.card.Text.length || '';

                if (size >= 80) {
                    return 'size2';
                } else if (size >= 55) {
                    return 'size1';
                }

                return '';
            },
        }
    }
</script>

<style lang="scss" scoped>
    /** Stolen from https://designshack.net/tutorialexamples/cardtricks/index.html */
    .card {
        user-select: none;
        //width: 130px;
        height: 200px;
        padding: 5px;
        border-radius: 10px;
        transition: all .1s ease;
        position: relative;

        background: white;
        color: #111;

        flex: 0 0 130px;

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