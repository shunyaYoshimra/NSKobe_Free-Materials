<template>
  <div>
    <form
      name="myForm"
      ref="myForm"
      :action="getPath"
      method="post"
      enctype="multipart/form-data"
      autocomplete="off"
    >
      <input type="hidden" name="time" :value="now" />
      <input type="hidden" name="extension" :value="extension" />

      <div
        id="upload"
        class="form-group commonStyle"
        v-bind:class="{ styleA: styleA, styleB: styleB }"
        @dragover.prevent="changeStyle($event, 'ok')"
        @dragleave.prevent="changeStyle($event, 'no')"
        @drop.prevent="uploadFile($event)"
      >
        <label for="upload_image" class="button">
          <p>画像を選択</p>
          <input
            id="upload_image"
            type="file"
            name="img"
            @change="uploadFile($event)"
            style="display: none"
            accept="image/*"
          />
        </label>

        <!-- ここからプレビュー機能の部分 -->
        <p>またはここに画像ファイルをドラッグ＆ドロップ</p>
        <img v-show="preview" v-bind:src="preview" style="width: 300px" />
        <p v-show="preview">{{ name }}</p>
        <!-- ここまでプレビュー機能の部分 -->
      </div>
      <p id="err-message">{{ errMessage }}</p>
      <v-combobox
        v-model="chips"
        chips
        clearable
        label="Your favorite hobbies"
        multiple
        solo
        id="chip"
      >
        <template v-slot:selection="{ attrs, item, select, selected }">
          <v-chip
            v-bind="attrs"
            :input-value="selected"
            close
            @click="select"
            @click:close="remove(item)"
          >
            <strong>{{ item }}</strong>
          </v-chip>
        </template>
      </v-combobox>
      <input
        placeholder="Enter Password to Upload"
        class="password-field"
        type="password"
        v-model="password"
        autocomplete="off"
      />
      <p id="err-message">{{ errMessagePass }}</p>
      <div class="button-wrapper">
        <v-btn
          :loading="loading"
          :disabled="loading"
          color="blue-grey"
          class="ma-2 white--text"
          @click="clickUpload()"
        >
          アップロード
          <v-icon right dark> mdi-cloud-upload </v-icon>
        </v-btn>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loader: null,
      loading: false,
      preview: "",
      name: "",
      styleA: true,
      styleB: false,
      chips: ["Kobe", "神戸", "フリー素材"],
      now: "00:00:00:00:00:000",
      extension: "",
      attached: false,
      errMessage: "",
      password: "",
      errMessagePass: "",
    };
  },
  mounted() {
    const date = new Date();
    this.now =
      date.getMonth() +
      ":" +
      date.getDate() +
      ":" +
      date.getHours() +
      ":" +
      date.getMinutes() +
      ":" +
      date.getSeconds() +
      ":" +
      date.getMilliseconds();
  },
  watch: {
    loader() {
      const l = this.loader;
      this[l] = !this[l];

      setTimeout(() => (this[l] = false), 3000);

      this.loader = null;
    },
  },
  computed: {
    getPath() {
      console.log("/api/materials/" + this.chips);
      return "/api/materials/" + this.chips;
    },
  },
  methods: {
    uploadFile(event) {
      this.attached = true;
      this.styleA = true;
      this.styleB = false;
      const files = event.target.files
        ? event.target.files
        : event.dataTransfer.files;
      const file = files[0];
      const reader = new FileReader();
      reader.onload = (event) => {
        this.preview = event.target.result;
      };
      reader.readAsDataURL(file);
      this.name = files[0].name;
      this.extension = this.name.split(".").pop();
      document.getElementById("upload_image").files = files;
    },
    changeStyle(event, flag) {
      if (flag == "ok") {
        this.styleA = false;
        this.styleB = true;
      } else {
        this.styleA = true;
        this.styleB = false;
      }
    },
    remove(item) {
      this.chips.splice(this.chips.indexOf(item), 1);
      this.chips = [...this.chips];
    },
    clickUpload() {
      if (this.password === "ntutors") {
        this.errMessagePass = "";
        if (this.attached === true) {
          this.loader = "loading";
          setTimeout(() => document.myForm.submit(), 1000);
        } else {
          this.errMessage = "ファイルを選択して下さい";
        }
      } else {
        this.errMessagePass = "パスワードが違います";
      }
    },
  },
};
</script>

<style>
.commonStyle {
  padding: 30px;
  text-align: center;
  margin: auto;
  margin-top: 10%;
  width: 60%;
}
.styleA {
  border: 3px dotted gray;
}
.styleB {
  border: 3px dotted rgba(0, 200, 0, 0.7);
}
.button {
  /* border: 1px solid green;
  padding: 3px;
  border-radius: 5px; */
  cursor: pointer;
}
.button p {
  color: green;
  margin-top: 10px;
  margin-left: 10px;
  margin-right: 10px;
}
.v-input__slot {
  box-shadow: none !important;
  width: 60% !important;
  margin: auto !important;
  margin-top: 20px !important;
}
.custom-loader {
  animation: loader 1s infinite;
  display: flex;
}
@-moz-keyframes loader {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}
@-webkit-keyframes loader {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}
@-o-keyframes loader {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}
@keyframes loader {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}
.button-wrapper {
  width: 60%;
  margin: auto;
  margin-top: 20px;
}
#err-message {
  color: tomato;
  margin-left: 20%;
  transform: translateY(20%);
}
.password-field {
  margin-left: 20%;
  width: 250px;
  outline: 0;
}
</style>
