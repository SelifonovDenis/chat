import React, {Component} from 'react'
import VWelcomeUser from "./VWelcomeUser";

export class WelcomeUser extends Component{
    constructor(props){
        super(props)
        this.state = {
            login: 'none'
        }
        this.view = VWelcomeUser.bind(this)
    }

    loginWithOther = () =>{
        this.setState(() => ({
            login: 'other'
        }))
    }

    loginWithThis = () =>{
        this.setState(() => ({
            login: 'this'
        }))
    }

    render(){
        return this.view()
    }
}

export default WelcomeUser