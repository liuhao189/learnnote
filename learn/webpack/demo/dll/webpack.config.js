const path = require('path');
const webpack = require('webpack');

module.exports = {
    resolve: {
        extensions: [".js", ".jsx"]
    },
    entry: {
        alpha: ['./alpha', './a'],
        beta: ['./beta.js', './b.js']
    },
    output: {
        path: path.join(__dirname, 'dist'),
        filename: 'MyDll.[name].js',
        library: '[name]_[hash]'
    },
    plugins: [
        new webpack.DllPlugin({
            path: path.join(__dirname, 'dist', '[name]-manifest.json'),
            name: '[name]_[hash]'
        })
    ]
};