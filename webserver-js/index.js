const express = require("express");
const serveStatic = require("serve-static");
const proxy = require("express-http-proxy");
const basicAuth = require("express-basic-auth");

const localIp = "192.168.1.71";
const app = express();

app.use(
  basicAuth({
    users: { admin: "robot" },
    challenge: true,
  })
);

app.use("/stream", proxy(`http://${localIp}:81`));
app.use("/ctl", proxy(`http://${localIp}:80`));

app.use(serveStatic("static", { index: ["index.html"] }));

app.listen(3000);
console.log(`started listening at http://0.0.0.0:3000`);
