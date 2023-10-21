import React, { useState } from "react";
import ReactDOM from "react-dom/client";

const ReactComponent = () => {
  const [count, setCount] = useState(0);
  return (
    <div>
      <div>This is the React Component</div>
      <div>
        <button onClick={() => setCount((e) => e + 1)}>
          Click me: {count}
        </button>
      </div>
    </div>
  );
};

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <ReactComponent />
);
