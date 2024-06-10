import "./App.css";
import "./controls";
import { STREAM_URL } from "./config.ts";
import { directionToCommand, useSwipe } from "./controls/touch.ts";
import { useEffect } from "react";
import { sendMoveCommand } from "./controls/sendMoveCommand.ts";

function App() {
  const swipe = useSwipe();

  useEffect(() => {
    const { direction, distance } = swipe;

    if (!direction || !distance) return;

    console.log(`Swipe ${direction} ${distance}`);

    sendMoveCommand(directionToCommand[direction], distance * 5);
  }, [swipe]);

  return (
    <div className="appContainer">
      <div className="streamContainer">
        <img id="stream" src={STREAM_URL} alt="video feed" />
      </div>
    </div>
  );
}

export default App;
