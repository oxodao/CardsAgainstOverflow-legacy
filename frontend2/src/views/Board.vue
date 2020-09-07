<template>
    <div id="game">
        <Header />
        <div id="inner-game">
            <SideMenu />
        </div>
    </div>
</template>

<script>
    import SideMenu from "../components/panels/SideMenu";
    import Header from "../components/panels/Header";

    export default {
        name: "Board",
        components: {Header, SideMenu},

    }
</script>

<style lang="scss" scoped>
    #game {
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: column;

        #inner-game {
            flex: 1;
            display: flex;
            flex-direction: row;
        }
    }
</style>