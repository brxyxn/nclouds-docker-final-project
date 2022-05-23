import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faUser, faEnvelope } from "@fortawesome/free-solid-svg-icons";
import Counters from "./Counter";
import PasswordField from "./PasswordField";

export default class MyForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      showPassword: false,
      user: {
        username: "",
        email: "",
        password: "",
      },
      counter: {
        sql: 0,
        cache: 0,
      },
    };
  }

  onUsernameChange(value) {
    this.setState({
      user: {
        username: value,
      },
    });
  }

  onEmailChange(value) {
    this.setState({
      user: {
        email: value,
      },
    });
  }

  onPasswordChange(value) {
    this.setState({
      user: {
        password: value,
      },
    });
  }

  mySubmit(e) {
    // e.preventDefault();
  }

  async getCounter(url, path) {
    const response = await fetch(`${url}${path}`);
    const data = await response.json();

    if (path == "/sql/users") {
      this.setState({
        counter: { sql: data.counter, cache: this.state.counter.cache },
      });
      return;
    }
    this.setState({
      counter: { sql: this.state.counter.sql, cache: data.counter },
    });
  }

  async setUser(e, path) {
    e.preventDefault();

    const url = "http://localhost:5000/api/v1";
    const response = await fetch(`${url}${path}`, {
      method: "POST",
      body: JSON.stringify({
        username: this.state.user.username,
        email: this.state.user.email,
        password: this.state.user.password,
      }),
    });
    const data = await response.json();

    if (path == "/sql/users") {
      this.setState((prevState) => ({
        counter: {
          ...prevState.counter,
          sql: data.counter,
        },
      }));
    } else {
      this.setState((prevState) => ({
        counter: {
          ...prevState.counter,
          cache: data.counter,
        },
      }));
    }

    // clear fields
    this.setState({
      user: {
        username: "",
        email: "",
        password: "",
      },
    });
  }

  async componentDidMount() {
    const url = "http://localhost:5000/api/v1";
    const sql = "/sql/users";
    const cache = "/cache/users";

    this.getCounter(url, sql);
    this.getCounter(url, cache);
  }

  render() {
    return (
      <>
        <form id="inputs" onSubmit={(e) => this.mySubmit(e)}>
          <div className="group">
            <FontAwesomeIcon icon={faUser} className="icon" />
            <input
              type="text"
              name="name"
              id="name"
              placeholder="Type your name"
              onChange={(e) =>
                this.setState((prevState) => ({
                  user: {
                    ...prevState.user,
                    username: e.target.value,
                  },
                }))
              }
              value={this.state.user.username}
            />
          </div>
          <div className="group">
            <FontAwesomeIcon icon={faEnvelope} className="icon" />
            <input
              type="text"
              name="email"
              id="email"
              placeholder="Type your email"
              onChange={(e) =>
                this.setState((prevState) => ({
                  user: {
                    ...prevState.user,
                    email: e.target.value,
                  },
                }))
              }
              value={this.state.user.email}
            />
          </div>
          <PasswordField
            onChange={(e) =>
              this.setState((prevState) => ({
                user: {
                  ...prevState.user,
                  password: e.target.value,
                },
              }))
            }
            value={this.state.user.password}
          />

          <div className="group buttons">
            <button
              type="submit"
              id="btn-sql"
              onClick={(e) => this.setUser(e, "/sql/users")}
            >
              Save to SQL
            </button>
            <button
              type="submit"
              id="btn-cache"
              onClick={(e) => this.setUser(e, "/cache/users")}
            >
              Save to Cache
            </button>
          </div>
        </form>
        <Counters
          sql={this.state.counter.sql ?? 0}
          cache={this.state.counter.cache ?? 0}
        />
      </>
    );
  }
}
