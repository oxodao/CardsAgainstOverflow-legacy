<template>
    <h1 v-if="Started && currBlackCard !== undefined && currBlackCard !== null && !IsReady">Il n'y a plus assez de joueurs !</h1>

    <h1 v-else-if="!Started">
        En attente de joueurs...
        <template v-if="IsReady && HasEnoughCards"><br />La partie est prête.</template>
        <template v-else-if="IsReady && !HasEnoughCards">Pas assez de cartes dans les decks sélectionnés!</template>
    </h1>
</template>

<script>
    import {mapGetters, mapState} from "vuex";

    export default {
        name: "GameStatus",
        computed: {
            ...mapState({
                currBlackCard: state => state.Room.CurrentBlackCard,
                Started: state => state.Room.Started,
            }),
            ...mapGetters([
                'IsReady',
                'IsPlaying',
                'HasEnoughCards'
            ])
        },
    }
</script>

<style scoped>

</style>