import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faUser, faKey, faEnvelope } from "@fortawesome/free-solid-svg-icons";
import "./App.css";
import { useState } from "react";

function App() {
  const [showPassword, setShowPassword] = useState(false);

  const togglePassword = () => {
    setShowPassword(!showPassword);
  };

  // Callback to SQL

  // Callback to Cache
  

  return (
    <section className="container">
      <section className="card">
        <header>
          <h1>Docker Project</h1>
        </header>
        <form id="inputs">
          <div className="group">
            <FontAwesomeIcon icon={faUser} className="icon" />
            <input
              type="text"
              name="name"
              id="name"
              placeholder="Type your name"
            />
          </div>
          <div className="group">
            <FontAwesomeIcon icon={faEnvelope} className="icon" />
            <input
              type="text"
              name="email"
              id="email"
              placeholder="Type your email"
            />
          </div>
          <div className="group">
            <FontAwesomeIcon icon={faKey} className="icon" />
            <input
              type={showPassword ? "text" : "password"}
              name="password"
              id="password"
              placeholder="Type your password"
            />
          </div>
          <div className="group">
            <label className="label">
              <div className={showPassword ? "toggle checked" : "toggle"}>
                <input
                  className="toggle-state"
                  type="checkbox"
                  name="check"
                  value="check"
                  onChange={togglePassword}
                />
                <div className="indicator"></div>
              </div>
              <div className="label-text">Show password</div>
            </label>
          </div>

          <div className="group buttons">
            <button type="submit" id="btn-sql" formAction="/sql">
              Save to SQL
            </button>
            {/* <input type="submit" value="Save to SQL" id="btn-sql" /> */}
            <button type="submit" id="btn-cache" formAction="/cache">
              Save to Cache
            </button>
          </div>
        </form>

        <section id="counters">
          <ul>
            <li id="sql-counter">Items stored in SQL: 17</li>
            <li id="cache-counter">Items stored in Cache: 23</li>
          </ul>
        </section>
      </section>
    </section>
  );
}

export default App;
