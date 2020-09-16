import React, { Component } from "react"
import Welcome from './Welcome/Welcome'
import LoginForm from './loginForm/LoginForm'
import WelcomeUser from './welcomeUser/welcomeUser'
import Request from '../providers/Request'

import '../../../css/login.css'

class LoginPage extends Component {
    constructor(props){
        super(props)
        this.state = {
            user: null
        }
        this.request = new Request()
        this.getUser()
    }

	getUser(){
	    this.request.Get('user')
            .then((res) => {
                if(res.type === "data"){
                    this.setState(()=>({user: res.data}))
                }
                else {
                    console.log(res)
                }
            })
            .catch((err) => {
            console.log(err)
        })
    }


    render() {
        if(this.state.user && this.state.user.id !== 0){
            return (
                <div id="page">
                    <Welcome/>
                    <WelcomeUser user={this.state.user}/>
                </div>
            )
        }
        
        return (
            <div id="page">
                <Welcome/>
		        <LoginForm/>
            </div>
        )
    }
}

export default LoginPage;