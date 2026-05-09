import React from "react";
import ReactDOM from "react-dom/client";
import { LoginButton } from "./components/Login";
import Settings from "./pages/Settings";

function App() {
  return (
    <main style={{ fontFamily: "system-ui, sans-serif", padding: 24 }}>
      <h1>PRoctor Fixtures</h1>
      <section>
        <h2>Login</h2>
        <div style={{ display: "flex", gap: 16, alignItems: "center" }}>
          <LoginButton />
          <LoginButton loading />
        </div>
      </section>
      <section>
        <h2>Settings</h2>
        <Settings />
      </section>
    </main>
  );
}

ReactDOM.createRoot(document.getElementById("root")!).render(<App />);
