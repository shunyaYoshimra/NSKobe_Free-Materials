<template>
  <div id="materials">
    <v-app>
      <div class="search-group">
        <div class="search_bar">
          <v-btn icon>
            <v-icon>mdi-magnify</v-icon>
          </v-btn>
          <input id="text2" type="text" placeholder="キーワードを入力" />
          <i class="fas fa-times search_icon"></i>
        </div>
      </div>
      <div class="materials-wrapper">
        <template>
          <v-card
            :loading="loading"
            class="mx-auto my-12 each-card"
            max-width="374"
            v-for="material in materials"
            :key="material"
          >
            <template slot="progress">
              <v-progress-linear
                color="deep-purple"
                height="10"
                indeterminate
              ></v-progress-linear>
            </template>

            <v-img class="card-img" height="250" :src="material.url"></v-img>

            <v-card-title>
              フリー素材
              <span
                id="download"
                class="download-icon"
                @click="downloadMedia('/dist/sample.gif', 'sample.gif')"
              >
                <v-tooltip top>
                  <template v-slot:activator="{ on, attrs }">
                    <v-icon v-bind="attrs" v-on="on"> mdi-cloud-upload </v-icon>
                  </template>
                  <span>ダウンロードする!!</span>
                </v-tooltip>
              </span></v-card-title
            >

            <v-card-text>
              <div>
                <!-- Small plates, salads & sandwiches - an intimate setting with 12
                indoor seats plus patio seating. -->
                {{ material.tags }}
              </div>
            </v-card-text>
          </v-card>
        </template>
      </div>
    </v-app>
  </div>
</template>

<script>
import axios from "axios";
import { mdiAccount } from "@mdi/js";
export default {
  data: () => ({
    loading: false,
    selection: 1,
    icons: {
      mdiAccount,
    },
    awsPath: "https://golang-s3-test.s3.ap-northeast-1.",
    materials: [],
  }),
  mounted() {
    axios.get("/api/materials").then((res) => {
      for (let i = 0; i < res.data.length; i++) {
        this.materials.push(res.data[i]);
      }
      console.log(res.data);
    });
  },
  methods: {
    downloadMedia(url, label) {
      axios
        .get(url, { responseType: "blob" })
        .then((response) => {
          const blob = new Blob([response.data], { type: "application/pdf" });
          const link = document.createElement("a");
          link.href = URL.createObjectURL(blob);
          link.download = label;
          link.click();
          URL.revokeObjectURL(link.href);
        })
        .catch(console.error);
    },
  },
};
</script>

<style>
.search-group {
  width: 400px;
  margin: auto;
  position: fixed;
  z-index: 100;
  left: 10%;
  top: 80px;
}
.search_bar {
  display: flex; /*アイコン、テキストボックスを調整する*/
  align-items: center; /*アイコン、テキストボックスを縦方向の中心に*/
  justify-content: center; /*アイコン、テキストボックスを横方向の中心に*/
  height: 50px;
  width: 100%;
  background-color: #ddd;
  border-radius: 10px;
}

.search_icon {
  /*アイコンに一定のスペースを設ける*/
  height: 15px;
  width: 15px;
  padding: 5px 5px;
}

#text2 {
  font-size: 16px;
  width: 100%; /*flexの中で100%広げる*/
  background-color: #ddd;
  border: none; /*枠線非表示*/
  outline: none; /*フォーカス時の枠線非表示*/
  box-sizing: border-box; /*横幅の解釈をpadding, borderまでとする*/
}
.search-wrapper {
  text-align: center;
}
.search-form {
  height: 40px;
  width: 500px;
  margin: auto;
}
.materials-wrapper {
  width: 80%;
  margin: auto;
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
  margin-top: 6%;
}
.each-card {
  margin-bottom: 20px;
}
.card-img:hover {
  opacity: 0.7;
}
.download-icon {
  color: #999;
  position: absolute;
  right: 10px;
  cursor: pointer;
}
.download-icon:hover {
  opacity: 0.7;
}
.v-responsive__content {
  width: 1512px !important;
}
</style>