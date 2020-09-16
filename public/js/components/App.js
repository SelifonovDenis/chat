import React, { Component } from 'react';

import {
    Route,
    Switch,
    Redirect,
    withRouter
} from "react-router-dom"
import '../../css/main.css'
import LoginPage from './login/LoginPage'
import ChatPage from './chat/ChatPage'

class App extends Component {
    render() {
        const { history } = this.props
        return (
            <div className="App">
                <Switch>
                    <Route history={history} exact path='/' component={LoginPage} />
                    <Route history={history} path='/chat' component={ChatPage} />
                </Switch>
            </div>
        );
    }
}

export default withRouter(App)