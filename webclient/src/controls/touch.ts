import { useEffect, useState } from "react";
import { MoveCommand } from "./sendMoveCommand.ts";

type Direction = "up" | "down" | "left" | "right";

export const directionToCommand: Record<Direction, MoveCommand> = {
  up: "move_forward",
  down: "move_backward",
  left: "move_left",
  right: "move_right",
};

function absDistance(x: number, y: number) {
  const xSquared = Math.pow(x, 2);
  const ySquared = Math.pow(y, 2);

  const sumOfSquares = xSquared + ySquared;

  return Math.sqrt(sumOfSquares);
}

function calcDirection(xDiff: number, yDiff: number): Direction {
  if (Math.abs(xDiff) > Math.abs(yDiff)) {
    return xDiff < 0 ? "right" : "left";
  } else {
    return yDiff < 0 ? "down" : "up";
  }
}

export const useSwipe = () => {
  const touches = {
    touchstart: {
      x: -1,
      y: -1,
    },
    touchmove: {
      x: -1,
      y: -1,
    },
  };

  const setTouches = (nextState: { touchstart: any; touchmove: any }) => {
    touches.touchstart = nextState.touchstart;
    touches.touchmove = nextState.touchmove;
  };

  const [swipe, setSwipe] = useState({
    direction: null,
    distance: 0,
  });

  const touchHandler = (event) => {
    const [touch] = event.touches || event.originalEvent.touches;

    switch (event.type) {
      case "touchstart":
      case "touchmove":
        const nextState = {
          ...touches,
          [event.type]: {
            x: touch.pageX,
            y: touch.pageY,
          },
        };

        setTouches(nextState);
        break;
      case "touchend":
        const xDiff = touches.touchstart.x - touches.touchmove.x;
        const yDiff = touches.touchstart.y - touches.touchmove.y;

        setSwipe({
          direction: calcDirection(xDiff, yDiff),
          distance: absDistance(xDiff, yDiff),
        });
        break;
      default:
        break;
    }
  };

  useEffect(() => {
    window.addEventListener("touchstart", touchHandler);
    window.addEventListener("touchmove", touchHandler);
    window.addEventListener("touchend", touchHandler);

    return () => {
      window.removeEventListener("touchstart", touchHandler);
      window.removeEventListener("touchmove", touchHandler);
      window.removeEventListener("touchend", touchHandler);
    };
  }, []);

  return swipe;
};
