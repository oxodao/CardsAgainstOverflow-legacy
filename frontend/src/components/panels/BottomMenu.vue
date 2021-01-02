<template>
  <div id="bottom-menu">
    <UsersButton/>
    <WizzButton/>
    <RerollButton/>
    <span v-if="isStarted && currTurn <= maxTurn && !zenMode">Tour: {{ currTurn }} / {{ maxTurn }}</span>
    <span v-else-if="isStarted && (zenMode || currTurn <= maxTurn)">Tour: {{ currTurn }}</span>

    <UserList v-if="isPlayerListShown"/>
  </div>
</template>

<script>
import UsersButton  from "@/components/UsersButton";
import RerollButton from "@/components/RerollButton";
import WizzButton   from "@/components/WizzButton";
import {mapState}   from "vuex";
import UserList     from "@/components/panels/UserList";

export default {
  name      : "BottomMenu",
  components: {UserList, WizzButton, RerollButton, UsersButton},
  computed  : {
    ...mapState({
      isStarted        : state => state.Room.Started,
      zenMode          : state => state.Room.ZenMode,
      currTurn         : state => state.Room.Turn,
      maxTurn          : state => state.Room.MaxTurn,
      isPlayerListShown: state => state.UI.PlayersMenuVisible,
      //isPlayerListShown: () => true
    })
  }
}
</script>

<style lang="scss" scoped>

#bottom-menu {
  height: 64px;
  background: #111;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-around;

  grid-area: d;

  @media(min-width: 650px) {
    display: none;
  }

  .ui-button {
    width: 48px;
    height: 48px;

  }
}

</style>