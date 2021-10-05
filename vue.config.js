module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    port: 8080,
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
