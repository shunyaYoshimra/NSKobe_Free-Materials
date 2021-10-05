<template>
  <div id="materials">
    <div class="materials">
      <template>
        <v-card
          :loading="loading"
          class="mx-auto my-12 each-card"
          max-width="374"
          v-for="material in 10"
          :key="material"
        >
          <template slot="progress">
            <v-progress-linear
              color="deep-purple"
              height="10"
              indeterminate
            ></v-progress-linear>
          </template>

          <v-img
            class="card-img"
            height="250"
            src="https://cdn.vuetifyjs.com/images/cards/cooking.png"
          ></v-img>

          <v-card-title>
            フリー素材
            <span
              id="download"
              class="download-icon"
              @click="downloadMedia(1, 'logo.png')"
            >
              <v-icon dark> mdi-cloud-upload </v-icon></span
            ></v-card-title
          >

          <v-card-text>
            <div>
              Small plates, salads & sandwiches - an intimate setting with 12
              indoor seats plus patio seating.
            </div>
          </v-card-text>
        </v-card>
      </template>
    </div>
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
  }),
  methods: {
    downloadMedia(id, filename) {
      axios
        .post(
          "https://cdn.vuetifyjs.com/images/cards/cooking.png",
          {
            id: id,
            filename: filename,
          },
          {
            responseType: "blob",
          }
        )
        .then((response) => {
          const fileURL = window.URL.createObjectURL(new Blob([response.data]));
          const fileLink = document.createElement("a");
          fileLink.href = fileURL;
          fileLink.setAttribute("download", filename);
          //fileLink.appendChild(document.createTextNode('test'));
          document.body.appendChild(fileLink);
          fileLink.click();
        });
    },
  },
};
</script>

<style>
.materials {
  width: 80%;
  margin: auto;
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
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
</style>