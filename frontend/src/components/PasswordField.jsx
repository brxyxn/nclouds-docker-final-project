import React, { useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faKey } from "@fortawesome/free-solid-svg-icons";

const PasswordField = (props) => {
  const [showPassword, setShowPassword] = useState(false);
  const togglePassword = () => {
    setShowPassword(!showPassword);
  };

  return (
    <>
      <div className="group">
        <FontAwesomeIcon icon={faKey} className="icon" />
        <input
          type={showPassword ? "text" : "password"}
          name="password"
          id="password"
          placeholder="Type your password"
          onChange={props.onChange}
          value={props.value}
        />
      </div>
      <div className="group">
        <label className="label">
          <div className={showPassword ? "toggle checked" : "toggle"}>
            {" "}
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
    </>
  );
};

export default PasswordField;