import React, { Component } from "react"
import {
    BrowserRouter,
    Route,
    Switch
} from "react-router-dom";

import App from './App';
import Article from "./Article";

class Router extends Component {

    render() {
        return (
            <BrowserRouter>
                <Switch>
                    <Route exact path="/" component={App} />
                    <Route path="/:id" component={Article} />
                </Switch>
            </BrowserRouter>
        );
    }

}

export default Router;