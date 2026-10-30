//Initializes React
import React from "react";
import { createRoot } from "react-dom/client";
import App from "./App"

//Renders App.txt
createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
