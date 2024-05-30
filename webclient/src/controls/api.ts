export type MoveCommand =
  | "move_forward"
  | "move_left"
  | "move_right"
  | "move_backward";

async function sendMoveCommand(command: MoveCommand, time: number) {
  const url = `/ctl/control?var=${command}&val=${time}`;
  await fetch(url, { method: "GET" });
}

const useMoveCommand = () => {
  let requestMutex = false;

  return async (command: MoveCommand, time: number) => {
    if (requestMutex) {
      return false;
    }

    requestMutex = true;

    await sendMoveCommand(command, time);

    requestMutex = false;
  };
};

export const sendMoveCmd = useMoveCommand();
