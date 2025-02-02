import { useState } from "react";
import { useEffect } from "react";
import reactLogo from "./assets/react.svg";
import { invoke } from "@tauri-apps/api/core";
import "./App.css";

function App() {
  const [msg, setMsg] = useState("");

  useEffect(() => {
    let ws: WebSocket;

    invoke("start_go_server").then(() => {
      ws = new WebSocket("ws://localhost:8080/ws");
      ws.onopen = () => console.log("Connected to WebSocket");
      ws.onmessage = (event) => {
        setMsg(event.data);
      };
      ws.onerror = (error) => console.error("WebSocket error:", error);
      ws.onclose = () => console.log("WebSocket closed");
    });
    return () => {
      if (ws) {
        ws.close();
      }
    };
  }, []);

  return (
    <main className="container">
      <h1>Welcome to Tauri + React</h1>

      <div className="row">
        <a href="https://vitejs.dev" target="_blank">
          <img src="/vite.svg" className="logo vite" alt="Vite logo" />
        </a>
        <a href="https://tauri.app" target="_blank">
          <img src="/tauri.svg" className="logo tauri" alt="Tauri logo" />
        </a>
        <a href="https://reactjs.org" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <p>{msg}</p>
    </main>
  );
}

export default App;
