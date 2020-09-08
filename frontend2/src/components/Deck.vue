<template>
    <div id="hand" v-if="(TurnState === 0 || IsJudge) && IsReady">
        <Card v-for="(card, index) in (IsJudge ? Proposal : Cards)" :key="card.ID+card.isSelected" v-bind:isProposal="TurnState === 1" v-bind:index="index" v-bind:card="card"/>
    </div>
</template>

<script>
    import {mapGetters, mapState} from "vuex";
    import Card from './Card';

    export default {
        name: "Deck",
        components: {
            Card
        },
        computed: {
            ...mapState({
                Cards: state => state.User.Hand,
                Proposal: state => state.Room.Answers,
                IsJudge: state => state.User.IsJudge,
                TurnState: state => state.Room.TurnState,
            }),
            ...mapGetters([
                'IsReady'
            ])
        },
    }
</script>

<style lang="scss" scoped>

    #hand {
        display: flex;
        height: calc(200px * 1.05); // 200px = height 1 card + 0.05 for the scale animation
        flex-direction: row;
        flex-wrap: nowrap;
        overflow: auto;
        align-items: center;
        justify-content: center;
    }

</style>