import React, { useState, useEffect, Fragment } from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import RepairCreate from "./components/Body";
import Show from "./components/BodyShow";
import SignIn from "./components/SignIn";
import Home from "./components/Home";

export default function App() {

  const [token, setToken] = useState<string>("");

  useEffect(() => {
    const token = localStorage.getItem("uid");
    if (token) {
      setToken(token);
    }
  }, []);

  if (!token) {
    return <SignIn />;
  }

  return (

    <Router>

      <div>

        {token && (

          <Fragment>

            <Navbar />

            <Switch>

              <Route exact path="/login" component={SignIn} />

              <Route exact path="/" component={Home} />

              <Route exact path="/show" component={Show} />

              <Route exact path="/create" component={RepairCreate} />

            </Switch>

          </Fragment>

        )}

      </div>

    </Router>

  );

}