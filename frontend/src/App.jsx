import React from "react";
import MyForm from "./components/Form";
import "./App.css";

function App() {
  return (
    <section className="container">
      <section className="card">
        <header>
          <h1>Docker Project</h1>
        </header>
        <MyForm />
      </section>
    </section>
  );
}

export default App;
