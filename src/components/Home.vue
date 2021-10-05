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
    console.log("1");
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