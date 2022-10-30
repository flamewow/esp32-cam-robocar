const express = require("express");
const serveStatic = require("serve-static");
const proxy = require("express-http-proxy");
const basicAuth = require("express-basic-auth");

const app = express();

app.use(
  basicAuth({
    users: { admin: "robot" },
    challenge: true,
  })
);

app.use("/stream", proxy("http://192.168.0.108:81"));
app.use("/ctl", proxy("http://192.168.0.108:80"));

app.use(serveStatic("client", { index: ["fpv.html"] }));

app.listen(3000);
console.log(`started listening at http://0.0.0.0:3000`);
