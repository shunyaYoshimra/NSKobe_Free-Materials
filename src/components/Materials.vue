<template>
  <div id="materials">
    <v-app>
      <div class="search-group">
        <div class="search_bar">
          <v-btn icon>
            <v-icon>mdi-magnify</v-icon>
          </v-btn>
          <input
            v-model="keyword"
            id="text2"
            type="text"
            placeholder="キーワードを入力"
            @keyup.enter="search()"
          />
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
              <span
                id="download"
                class="download-icon icon"
                @click="downloadMedia(material.file_name)"
              >
                <v-tooltip top>
                  <template v-slot:activator="{ on, attrs }">
                    <v-icon v-bind="attrs" v-on="on"> mdi-cloud-upload </v-icon>
                  </template>
                  <span>ダウンロードする!!</span>
                </v-tooltip>
              </span>
              <v-icon class="delete-icon icon" @click="dialogID = material.id">
                mdi-delete
              </v-icon>
              <template v-if="dialogID === material.id">
                <div class="delete-dialog" @click="dialogID = -1"></div>
                <v-card class="mx-auto delete-card" max-width="344" outlined>
                  <v-list-item three-line>
                    <v-list-item-content>
                      <div class="text-overline mb-4">
                        <input
                          placeholder="Enter Password to Delete"
                          class="delete-password-field"
                          type="text"
                          id="password"
                          v-model="password"
                        />
                        <p>{{ errMessage }}</p>
                      </div>

                      <v-list-item-subtitle
                        >削除したい場合にはTAに頼んでみましょう！</v-list-item-subtitle
                      >
                    </v-list-item-content>

                    <v-list-item-avatar tile size="80" color="grey"
                      ><img :src="material.url" alt=""
                    /></v-list-item-avatar>
                  </v-list-item>

                  <v-card-actions>
                    <v-btn
                      @click="deleteMaterial(material.id)"
                      outlined
                      rounded
                      text
                      class="delete-btn"
                    >
                      削除</v-btn
                    >
                  </v-card-actions>
                </v-card>
              </template>
            </v-card-title>

            <v-card-text>
              <div>
                <!-- Small plates, salads & sandwiches - an intimate setting with 12
                indoor seats plus patio seating. -->
                #{{ material.tags }}
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
    dialogID: -1,
    loading: false,
    selection: 1,
    icons: {
      mdiAccount,
    },
    awsPath: "https://golang-s3-test.s3.ap-northeast-1.",
    materials: [],
    keyword: "",
    password: "",
    errMessage: "",
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
    downloadMedia(fileName) {
      axios.post("/api/download/" + fileName).then((res) => {
        console.log(res);
        axios
          .get("/app/images/" + fileName, { responseType: "blob" })
          .then((response) => {
            const blob = new Blob([response.data], { type: "application/pdf" });
            const link = document.createElement("a");
            link.href = URL.createObjectURL(blob);
            link.download = fileName;
            link.click();
            URL.revokeObjectURL(link.href);
            axios.post("/api/delete-file/" + fileName).then((res) => {
              console.log(res);
            });
          })
          .catch(console.error);
      });
    },
    search() {
      axios.post("/api/search/" + this.keyword).then((res) => {
        this.materials = [];
        for (let i = 0; i < res.data.length; i++) {
          this.materials.push(res.data[i]);
        }
      });
    },
    deleteMaterial(id) {
      if (this.password === "ntutors") {
        this.password = "";
        this.dialog = false;
        console.log(`ID: ${id}`);
        axios.delete(`/api/materials/${id}`);
        this.$router.go({ path: "/", force: true });
      } else {
        this.errMessage = "パスワードが違います";
      }
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
.icon {
  color: #999;
  position: absolute;
  cursor: pointer;
}
.icon:hover {
  opacity: 0.7;
}
.download-icon {
  right: 10px;
}
.delete-icon {
  left: 290px;
}
.v-responsive__content {
  width: 1512px !important;
}
.delete-dialog {
  padding: 20px;
}
.delete-password-field {
  margin-top: 15px;
  width: 250px;
  outline: 0;
}
.delete-err-message {
  color: tomato;
}
.delete-dialog {
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  z-index: 1000;
  background-color: #000;
  opacity: 0.4;
}
.delete-btn {
  color: #1565c0 !important;
}
.delete-card {
  position: absolute;
  top: 70px;
  z-index: 1001;
}
</style>