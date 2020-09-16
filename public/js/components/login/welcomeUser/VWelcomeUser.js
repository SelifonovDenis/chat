import React from "react";
import {
    Route,
    Switch,
    Redirect,
    withRouter
} from "react-router-dom"
import LoginForm from "../loginForm/LoginForm";

function VWelcomeUser() {
    if (this.state.login === 'this'){
        return (<Redirect to='/chat'/>)
    }
    if (this.state.login === 'other'){
        return (<LoginForm/>)
    }
    return (
        <div className="loginWelcome">
            <p className="nickname">{this.props.user.nickname}</p>
            <button onClick={this.loginWithThis}>Войти</button>
            <button onClick={this.loginWithOther}>Войти под другим никнеймом</button>
        </div>
    )
}

export default VWelcomeUser;