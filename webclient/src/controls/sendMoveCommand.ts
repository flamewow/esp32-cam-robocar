import { STREAM_BE_BASE_URL } from "../config.ts";

export type MoveCommand =
  | "move_forward"
  | "move_left"
  | "move_right"
  | "move_backward";

async function _sendMoveCommand(command: MoveCommand, time: number) {
  const url = `${STREAM_BE_BASE_URL}/ctl/control?var=${command}&val=${time}`;
  await fetch(url, { method: "GET" }).catch((e) => console.error(e));
}

const createSingleExecutionMovementCommand = () => {
  let requestMutex = false;

  return async (command: MoveCommand, time: number) => {
    if (requestMutex) {
      return false;
    }

    requestMutex = true;

    await _sendMoveCommand(command, time);

    requestMutex = false;
  };
};

export const sendMoveCommand = createSingleExecutionMovementCommand();
