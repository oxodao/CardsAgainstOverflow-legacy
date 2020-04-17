<template>
    <div class="card" @click="setSelected(currCard.ID)" v-bind:class="getClassnameIsSelected">
        <p>{{currCard.Text}}</p>
        <div class="branding">
            <span>Cards</span>
            <span>Against</span>
            <span>Overflow</span>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Card',
    props: [
        'currCard',
    ],
    data: function() {
        return {
            currentCard: this.currCard,
        }
    },
    computed: {
        getClassnameIsSelected: function() {
            return this.currentCard.isSelected ? "isSelected" : ""
        }
    },
    methods: {
        setSelected(id) {
            this.$store.dispatch('setSelected', id);
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

</style>