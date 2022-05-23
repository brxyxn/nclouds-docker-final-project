import React from "react";

const Counters = (props) => {
  return (
    <>
      <section id="counters">
        <ul>
          <li id="sql-counter">Items stored in SQL: {props.sql}</li>
          <li id="cache-counter">Items stored in Cache: {props.cache}</li>
        </ul>
      </section>
    </>
  );
};

export default Counters;
