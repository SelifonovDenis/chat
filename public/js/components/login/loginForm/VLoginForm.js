import React from "react";
import {
    Route,
    Switch,
    Redirect,
    withRouter
} from "react-router-dom"

function VLoginForm() {
    if (this.state.auth === true){
        return (<Redirect to='/chat'/>)
    }

    return (
        <div className="loginForm">
            <p>Введите никнейм, чтобы войти в чат:</p>
            <input type='text' placeholder='Никнейм' value={this.state.value} onChange={this.nicknameOnChange}/>
            <button onClick={this.login}>Войти</button>
            <div className='error'>{this.state.error}</div>
        </div>
    )
}

export default VLoginForm;