import "./App.css";
import "./controls";
import { STREAM_URL } from "./config.ts";
import { useSwipe } from "./controls/touch.tsx";
import { useEffect } from "react";
import { sendMoveCmd } from "./controls/api.ts";

function App() {
  const swipe = useSwipe();

  useEffect(() => {
    const { direction, distance } = swipe;

    if (distance < 150) {
      console.log(`Slight swipe ${direction}`);
    }

    if (distance > 150) {
      console.log(`Strong swipe ${direction}`);
    }
    const directionToCommand = {
      up: "move_forward",
      down: "move_backward",
      left: "move_left",
      right: "move_right",
    };

    sendMoveCmd(directionToCommand[direction], distance * 5);
  }, [swipe]);

  return (
    <div className="appContainer">
      <div className="streamContainer">
        <img id="stream" src={STREAM_URL} alt="video feed" />
      </div>
      {/*<div className="controlsContainer">*/}
      {/*  <button id="start" className="controlButton">Left</button>*/}
      {/*  <button id="start" className="controlButton">Right</button>*/}
      {/*</div>*/}
    </div>
  );
}

export default App;
