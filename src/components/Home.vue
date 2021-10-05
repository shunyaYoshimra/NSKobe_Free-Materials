<template>
  <div>{{ messageFromGo }}</div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      messageFromGo: "",
    };
  },
  mounted() {
    axios
      .get("/api/message")
      .then((res) => {
        console.log(res);
        this.messageFromGo = res.data.message;
      })
      .catch((err) => {
        console.log(err);
      });
  },
};
</script>