const path = require('path');
const { VueLoaderPlugin } = require('vue-loader');
const webpack = require('webpack');

const workspace = path.join(__dirname, '..');

function resolve(dir) {
  return path.join(workspace, dir);
}

const appRoot = 'app';

// Documentation - https://balm.js.org/docs/config/
// 中文文档 - https://balm.js.org/docs/zh/config/
module.exports = {
  server: {
    // proxyConfig: {
    //   context: '/api',
    //   options: {
    //     target: 'http://your.project.dev', // Target host
    //     changeOrigin: true // Needed for virtual hosted sites
    //   }
    // }
  },
  roots: {
    source: appRoot
  },
  styles: {
    extname: 'scss'
  },
  scripts: {
    entry: {
      main: `./${appRoot}/scripts/main.js` // Entry js file
    },
    loaders: [
      {
        test: /\.vue$/,
        loader: 'vue-loader'
      }
    ],
    plugins: [
      new VueLoaderPlugin(),
      // feature flags <http://link.vuejs.org/feature-flags>
      new webpack.DefinePlugin({
        __VUE_OPTIONS_API__: JSON.stringify(true),
        __VUE_PROD_DEVTOOLS__: JSON.stringify(false)
      })
    ],
    alias: {
      vue$: 'vue/dist/vue.esm-bundler.js',
      '@': resolve(`${appRoot}/scripts`)
    }
  },
  assets: {
    root: 'assets', // Replace 'assets' to your remote project root
    mainDir: 'public',
    cache: true
  }
  // More Config
};
