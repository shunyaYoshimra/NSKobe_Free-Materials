<template>
  <v-card class="overflow-hidden">
    <v-app-bar
      absolute
      color="white"
      elevate-on-scroll
      scroll-target="#scrolling-techniques-7"
    >
      <v-app-bar-nav-icon></v-app-bar-nav-icon>

      <v-toolbar-title @click="changeRoute('/')"
        ><span class="logo-title">Kobe CP Free Materials</span></v-toolbar-title
      >

      <v-spacer></v-spacer>

      <v-btn icon @click="changeRoute('/')">
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <v-btn icon @click="changeRoute('/upload')">
        <v-icon>mdi-cloud-upload</v-icon>
      </v-btn>
    </v-app-bar>
    <v-sheet id="scrolling-techniques-7" class="overflow-y-auto">
      <v-container style="height: 64px"> </v-container>
    </v-sheet>
  </v-card>
</template>

<script>
export default {
  methods: {
    changeRoute(path) {
      this.$router.push(path);
    },
  },
};
</script>

<style>
.logo-title {
  font-family: "Philosopher", sans-serif;
}
</style>
