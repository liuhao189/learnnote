const path = require('path');
const webpack = require('webpack')

module.exports = {
    entry: {
        example: './example.js'
    },
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: '[name].[hash:8].js'
    },
    plugins: [
        new webpack.DllReferencePlugin({
            context: path.join(__dirname, '..', 'dll'),
            manifest: require('./../dll/dist/alpha-manifest.json')
        }),
        new webpack.DllReferencePlugin({
            scope: 'beta',
            manifest: require('./../dll/dist/beta-manifest.json'),
            extensions: ['.js', ',jsx']
        })
    ]
}