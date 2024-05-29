const TIME_CONSTANT = 50;
const LEAP_TIME_CONSTANT = 500;
const TURBO_TIME_CONSTANT = 2000;
const TIME_DELTA = 100;

const Keys = {
  forward: "KeyW",
  left: "KeyA",
  right: "KeyD",
  backward: "KeyS",
  // long actions
  leap: "Space",
  spin: "KeyG",
  turboForward: "KeyF",
  turboBack: "KeyR",
};

const turboKeys = [Keys.spin, Keys.turboForward, Keys.turboBack];

const Keys2Cmd = {
  [Keys.forward]: "move_forward",
  [Keys.turboForward]: "move_forward",
  [Keys.leap]: "move_forward",

  [Keys.left]: "move_left",

  [Keys.right]: "move_right",
  [Keys.spin]: "move_right",

  [Keys.backward]: "move_backward",
  [Keys.turboBack]: "move_backward",
};

let controlLocked = false;
addEventListener("keypress", async (e) => {
  const key = e.code;
  console.log(key);

  if (controlLocked) {
    return;
  }

  controlLocked = true;

  let time = TIME_CONSTANT;
  if (turboKeys.includes(key)) {
    time = TURBO_TIME_CONSTANT;
  } else if (key === Keys.leap) {
    time = LEAP_TIME_CONSTANT;
  }

  await sendMoveRequest(key, time);

  controlLocked = false;
});

async function sendMoveRequest(key, time) {
  const command = Keys2Cmd[key];
  if (!command) {
    console.error(`wrong key ${key}`);
    return false;
  }

  console.log(`sending ${key} => ${command}`);
  const url = `/ctl/control?var=${command}&val=${time}`;
  await fetch(url, { method: "GET" });
  return true;
}

const lampInputRange = document.querySelector(".lamp");

lampInputRange.addEventListener("change", async (e) => {
  const value = e.target.value;
  console.log(`lamp brightness changed ${value}`);
  await changeLight(value);
});

async function changeLight(value) {
  const command = "lamp";
  const url = `/ctl/control?var=${command}&val=${value}`;
  await fetch(url, { method: "GET" });
  return true;
}