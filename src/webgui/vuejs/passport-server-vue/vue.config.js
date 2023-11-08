// vue.config.js

const theWebAppTitle = 'Passport Server'

const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: true,
  devServer: {
    proxy: 'http://localhost:23405',
    port: 23406,
  },
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = theWebAppTitle
      return args
    })
  },
})
