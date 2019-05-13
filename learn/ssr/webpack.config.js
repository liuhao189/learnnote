const VueLoaderPlugin = require('vue-loader/lib/plugin');
const path = require('path');

module.exports = {
    mode: 'production',
    entry: {
        client: path.resolve(__dirname, './src/entry-client.js')
    },
    output: {
        filename: '[name].bundle.js',
        path: __dirname + '/dist'
    },
    resolve: {
        extensions: [".js", ".json", ".vue"]
    },
    module: {
        rules: [
            {
                test: /\.vue$/,
                use: 'vue-loader'
            }
        ]
    },
    plugins: [
        new VueLoaderPlugin()
    ]
}