'use strict'
const path = require('path')

module.exports = {
    outputDir: "dist",
    publicPath: './',
    productionSourceMap: true,
    configureWebpack: {
        devtool: "cheap-module-source-map",
        // devtool: 'source-map',
        // devServer: {
        //     disableHostCheck: true,
        //     open: false,
        //     // host: 'exmaple.com',
        //     port: 443,
        //     https: true,
        //     hotOnly: false,
        // },
    },
    pages: {
        index: {
            entry: "index.js",
            template: 'public/index.html',
            filename: 'index.html',
            chunks: ['chunk-vendors', 'chunk-common', "index"]
        }
    },
    filenameHashing: true,
    // devServer: {
    //     host: '172.20.109.196',
    //     port: 8089,
    //     // https: true,
    //     open: true,
    //     proxy: {
    //         // detail: https://cli.vuejs.org/config/#devserver-proxy
    //         '/': {
    //             target: `https://172.20.109.222:8090/`,
    //             // target: `http://localhost:801/`,
    //             changeOrigin: true,
    //             pathRewrite: {
    //                 ['^/']: ''
    //             }
    //         }
    //     },
    //     disableHostCheck: true
    // },
}
