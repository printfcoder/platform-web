module.exports = {
    // devServer Options don't belong into `configureWebpack`
    devServer: {
        host: "0.0.0.0",
        hot: true,
        disableHostCheck: true,
        proxy: {
            '/': {
                target: 'http://127.0.0.1:9082',
                changeOrigin: true,
                pathRewrite: {
                    '/': '/', // optional, it is just a placeholder of usage.
                },
            },
            '^/websocket': {
                target: 'http://127.0.0.1:8080',
                ws: true,
                changeOrigin: true
            }
        },
    },
    transpileDependencies: [
        'vue-echarts',
        'resize-detector'
    ],
    publicPath: '/platform/',
    outputDir: '../srv/webapp',
    assetsDir: 'static',
};
