<template>
    <div id="rules" v-bind:class="isShown">
        <div id="rules-inner">
            <div class="icon" v-bind:class="cardGame ? 'selected' : ''" @click="selectCard">
                <img src="../assets/card_icon.png"/>
                <span>Règles du jeu</span>
            </div>
            <div class="icon" v-bind:class="usage ? 'selected' : ''" @click="selectUsage">
                <img src="../assets/game_icon.png"/>
                <span>Utilisation du site</span>
            </div>
        </div>
        <div id="rules-content" v-if="cardGame">
            <h2>Règles du jeu de cartes</h2>
            <p>Si vous avez déjà joué à Cards Against Humanity ou Blanc Manger Coco, les règles sont similaires. Ces règles considèrent que les joueurs ont un moyen de discussion vocal (Skype, appel messenger, etc...)</p>
            <p>Le jeu comporte des cartes blanches (réponses) et des cartes à trou.</p>
            <p>À chaque tour, un joueur est désigné <img class="mini-icon" src="../assets/gavel.png"/> Juge . Il va lire à haute voix la carte à trou.</p>
            <p>Chacun des joueurs va ensuite compléter avec les cartes qu'il possède dans son jeu et tenter de faire la combinaison la plus drôle.</p>
            <p>Le juge va ensuite lire chaque proposition à haute voix et choisir celle qu'il considère la plus drôle.</p>
            <p>Si un joueur n'aime pas sa main, il peut décider d'effectuer un "Re-roll" une fois tout les 6 tours (personnalisable). Il va alors abadonner toutes ses cartes et en repiocher.</p>
        </div>
        <div id="rules-content" v-else-if="usage">
            <h2>Utilisation du site</h2>
            <p>Pour créer une partie, rien de plus simple. Choisissez votre pseudo et cliquez sur "Rejoindre".</p>
            <p>En haut à droite de l'écran, le code de la salle est affiché. Partagez-le avec vos amis afin qu'ils vous rejoignent en le renseignant dans la case "Code salle"</p>
            <br />
            <p>Personnalisez la salle selon vos préférenceset lancez la partie !</p>
            <p>Le mode zen permet de ne pas avoir de limite de tours et jouer indéfiniement</p>
            <p>Un chronomètre permet de limiter le temps de chaque tour. Si tous les joueurs ont joués leurs cartes, le compteur va diminuer à 10 secondes afin de laisser le temps de changer d'avis.</p>
            <p>Lors du vote, seul le juge peut changer la carte affichée en cours.</p>
            <br />
            <p>Ce jeu n'est pas conçu pour être joué sur mobile mais uniquement sur ordinateur.</p>
            <p>Testé et fonctionnel dans Firefox et Google Chrome.</p>
            <br />
            <p>Si vous trouvez une faille, n'hésitez pas à me contacter pour me l'expliquer afin que je puisse la corriger.</p>
            <p>Si vous me connaissez, contactez moi par votre moyen habituel sinon n'hésitez pas à m'envoyer un email à <a href="mailto:nathan@janczewski.fr">nathan@janczewski.fr</a></p>
        </div>
        <div id="closeBt" @click="close">X</div>
    </div>
</template>

<script>
    export default {
        name: "Rules",
        props: [ 'showRules' ],
        computed: {
            isShown() {
                return this.$store.state.ShowRules ? "shown" : "";
            }
        },
        data: () => ({
            cardGame: true,
            usage: false,
        }),
        methods: {
            selectCard() {
                this.cardGame = true;
                this.usage = false;
            },
            selectUsage() {
                this.cardGame = false;
                this.usage = true;
            },
            close() {
                this.$store.commit('showRules');
            }
        }
    }
</script>

<style lang="scss" scoped>

    #rules {
        display: none;
        position: absolute;

        flex-direction: column;

        $width: 500px;
        $height: 700px;

        width: $width;
        height: $height;
        left: calc(50% - #{$width/2});
        top: calc(50% - #{$height/2});

        background: rgba(0, 0, 0, .8);
        border-radius: 2em;

        #rules-inner {
            display: flex;
            flex-direction: row;
            justify-content: center;
        }

        #rules-content {
            flex: 1;
            padding: 1em;
            text-align: justify;

            h2 {
                font-size: 1.1em;
            }

            p {
                font-size: .75em;
            }

            a {
                color: #4dbbc7;
            }
        }

        .icon {
            width: 192px;

            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            text-align: center;

            img {
                display: inline-block;
            }

            &.selected {
                span {
                    text-shadow: 0 0 15px #3EC480;
                }
            }
        }

        .mini-icon {
            height: 1.2em;
        }

        #closeBt {
            width: 2.5em;
            height: 2.5em;
            position: absolute;
            top: 1em;
            left: 1em;

            text-align: center;
            line-height: 2.5em;

            color: black;
            background: rgba(white, .75);
            border-radius: 50%;
        }
    }

    .shown {
        display: flex !important;
    }

</style>