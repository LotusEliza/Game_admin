// const PurgecssPlugin = require('purgecss-webpack-plugin');
// const glob = require('glob-all');
const path = require('path');

module.exports = {
  chainWebpack: config => {
    config.module
        .rule('vue')
        .use('vue-loader')
        .loader('vue-loader')
        .tap(options => {
          options['transformAssetUrls'] = {
            img: 'src',
            image: 'xlink:href',
            'b-carousel-slide': 'img-src',
            'b-embed': 'src'
          }

          return options
        })
  },

  configureWebpack: {
    // plugins: [
    //   new PurgecssPlugin({
    //     paths: glob.sync([
    //       path.join(__dirname, './src/index.html'),
    //       path.join(__dirname, './**/*.vue'),
    //       path.join(__dirname, './src/**/*.js')
    //     ])
    //   })
    // ]
  },

  css: {
    loaderOptions: {
      scss: {
        prependData: '@import "@/styles/_variables.scss";'
      }
    }
  },

  pwa: {
    name: 'Acme',
    themeColor: '#5f4884',
    msTileColor: '#5f4884'
  }
};
