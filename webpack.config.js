const path = require("path");

module.exports = {
  entry: {
    "./bundle.js": "./public/js/index.js",
  },
  output: {
    filename: '[name]',
    path: path.resolve(__dirname, "public/dist"),
    publicPath: 'public/dist/'
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader",
        },
      },
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"]
      },
      {
        test: /\.png$/,
        use: {
          loader: "file-loader",
        },
      }
    ]
  },
};