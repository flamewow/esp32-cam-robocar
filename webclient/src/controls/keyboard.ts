import { MoveCommand, sendMoveCommand } from "./sendMoveCommand.ts";

const TIME_CONSTANT = 50;
const LEAP_TIME_CONSTANT = 500;
const TURBO_TIME_CONSTANT = 2000;

const Keyboard = {
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

const turboKeys = [Keyboard.spin, Keyboard.turboForward, Keyboard.turboBack];

const keysMoveCommandMap: Record<string, MoveCommand> = {
  [Keyboard.forward]: "move_forward",
  [Keyboard.turboForward]: "move_forward",
  [Keyboard.leap]: "move_forward",

  [Keyboard.left]: "move_left",

  [Keyboard.right]: "move_right",
  [Keyboard.spin]: "move_right",

  [Keyboard.backward]: "move_backward",
  [Keyboard.turboBack]: "move_backward",
};

addEventListener("keypress", async (e) => {
  const key = e.code;

  let time = TIME_CONSTANT;

  if (turboKeys.includes(key)) {
    time = TURBO_TIME_CONSTANT;
  } else if (key === Keyboard.leap) {
    time = LEAP_TIME_CONSTANT;
  }

  await sendMoveCommand(keysMoveCommandMap[key], time);
});
