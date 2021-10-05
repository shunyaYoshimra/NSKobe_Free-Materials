module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy: {
      '/api': {
        target: "http://localhost:3000",
        //     changeOrigin: true,
        //     pathRewrite: p => {
        //       console.log(p);
        //       r = p.replace(/^\/postal\/*/, '?zipcode=');
        //       console.log("result:" + r);
        //       return r;
        //     },
      }
    }
  }
}

console.log("this is from vue.config.js")