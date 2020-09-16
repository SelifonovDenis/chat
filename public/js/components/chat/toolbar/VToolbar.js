import React from "react";
import {
    Route,
    Switch,
    Redirect,
    withRouter
} from "react-router-dom"

function VToolbar() {
    if (this.state.logout === true){
        return (<Redirect to='/'/>)
    }
    return (
        <div className="menu">
            <div className="menuHeader">Вы вошли как:<br/>{this.props.user.nickname}</div>
            <button className="menuItem" onClick={this.logout}>Выход</button>
        </div>
    )
}

export default VToolbar;