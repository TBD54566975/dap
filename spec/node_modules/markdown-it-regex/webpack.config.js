const path = require('path')

const loaders = [
  {
    test: /\.json$/,
    loader: 'json-loader'
  },
  {
    test: /\.js$/,
    loader: 'babel-loader',
    exclude: /node_modules/,
    query: {
      presets: [
        ['env', {
          'targets': {
            'node': 'current'
          }
        }]
      ]
    }
  }
]

const configurations = [
  {
    entry: {
      'index': './src/index.js'
    },
    output: {
      path: path.join(__dirname, 'src'),
      filename: '[name].bundle.js',
      libraryTarget: 'commonjs2'
    },
    module: { loaders },
    target: 'node'
  },
  {
    entry: {
      'index': './test/index.js'
    },
    output: {
      path: path.join(__dirname, 'test'),
      filename: '[name].bundle.js',
      libraryTarget: 'commonjs2'
    },
    module: { loaders },
    target: 'node'
  }
]

module.exports = configurations
