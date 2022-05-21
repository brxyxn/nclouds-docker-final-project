import "./App.css";

function App() {
  return (
    <>
      <section class="container">
        <header>
          <h1>Docker Project</h1>
        </header>
        <form id="inputs">
          <div class="group">
            <label for="name">Name</label>
            <input
              type="text"
              name="name"
              id="name"
              placeholder="Type your name"
            />
          </div>
          <div class="group">
            <label for="email">Email</label>
            <input
              type="text"
              name="email"
              id="email"
              placeholder="Type your email"
            />
          </div>
          <div class="group">
            <label for="password">Password</label>
            <input
              type="text"
              name="password"
              id="password"
              placeholder="Type your password"
            />
          </div>
          <div class="buttons">
            <input type="submit" value="Save to SQL" id="btn-sql" />
            <input
              type="submit"
              value="Save to Cache"
              id="btn-cache"
              formaction="/action_page2"
            />
          </div>
        </form>
        <section id="counters">
          <ul>
            <li id="sql-counter">Items stored in SQL: 17</li>
            <li id="cache-counter">Items stored in Cache: 23</li>
          </ul>
        </section>
      </section>
    </>
  );
}

export default App;
