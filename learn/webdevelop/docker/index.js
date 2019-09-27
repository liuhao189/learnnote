const http = require("http");
const hostName = "127.0.0.1";
const port = 8080;

const server = http.createServer(function (req, res) {
  res.setHeader("Content-Type", "text/plain");
  res.end("hello nodejs");
});

server.listen(port, function () {
  console.log(`服务器运行在http://${hostName}:${port}`);
});