<template>
  <div class="main">
    <Header></Header>
    <div id="pages">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Header from "./components/shared/Header.vue";
export default {
  components: {
    Header,
  },
};
</script>

<style>
</style>