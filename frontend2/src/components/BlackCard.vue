<template>
    <h1 v-if="Started && currBlackCard !== undefined && currBlackCard !== null && IsReady">
        {{currBlackCard.Text}}
    </h1>
</template>

<script>
    import {mapGetters, mapState} from "vuex";

    export default {
        name: "BlackCard",
        computed: {
            ...mapState({
                Started: state => state.Room.Started,
                currBlackCard: state => state.Room.CurrentBlackCard,
            }),
            ...mapGetters([
                'IsReady'
            ])
        }
    }
</script>

<style scoped>
    h1 {
        text-align: center;
    }
</style>