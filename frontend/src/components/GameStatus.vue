<template>
    <h1 v-if="Started && currBlackCard !== undefined && currBlackCard !== null && !IsReady">
        Il n'y a plus assez de joueurs !
    </h1>

    <h1 v-else-if="!Started">
        <template v-if="TurnState !== 2">
            <!-- @TODO: This does not work, TurnState = 0 when enough players are here -->
            En attente de joueurs...
        </template>
        <template v-else>
            {{ getWinner() }}<br />
        </template>
        <template v-if="IsReady && HasEnoughCards">
            <br />La partie est prête.
        </template>
        <template v-else-if="IsReady && !HasEnoughCards">
            <br />Pas assez de cartes dans les decks sélectionnés!
        </template>
    </h1>
</template>

<script>
import {mapGetters, mapState} from 'vuex';

export default {
    name: 'GameStatus',
    computed: {
        ...mapState({
            currBlackCard: state => state.Room.CurrentBlackCard,
            Started: state => state.Room.Started,
            TurnState: state => state.Room.TurnState,
        }),
        ...mapGetters([
            'IsReady',
            'IsPlaying',
            'HasEnoughCards'
        ])
    },
    methods: {
        getWinner() {
            let winners = [''];
            let winnerScore = -1;

            this.$store.state.Room.Participants.forEach(e => {
                if (e.Score === winnerScore) {
                    winners.push(e.Username);
                    winnerScore = e.Score;
                } else if (e.Score > winnerScore) {
                    winners = [e.Username];
                    winnerScore = e.Score;
                }
            });

            if (winnerScore === 0) {
                return "Personne n'a gagné !";
            }

            if (winners.length === 1) {
                return 'Le gagnant est ' + winners[0] + ' !';
            }

            if (winners.length === this.$store.state.Room.Participants.length) {
                return 'Match nul !';
            }

            return 'Les gagnants sont ' + winners.join(', ') + ' !';
        }
    }
};
</script>

<style scoped>
h1 {
    text-align: center;
}
</style>