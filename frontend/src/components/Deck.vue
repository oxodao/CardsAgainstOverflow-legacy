<template>
    <div id="hand" v-if="(TurnState === 0 || TurnState === 1) && IsReady">
        <Card v-for="(card, index) in ((IsDeporte || IsJudge || TurnState === 1) ? Proposal : Cards)" :key="card.ID+card.isSelected" v-bind:IsProposal="IsJudge || TurnState === 1" v-bind:Index="index" v-bind:card="card"/>
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
                IsDeporte: state => state.UI.Deporte,
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
        width: 100%;
        height: 280px;
        flex-direction: row;
        flex-wrap: nowrap;
        overflow: auto;
        align-items: center;
        justify-content: center;
    }

</style>