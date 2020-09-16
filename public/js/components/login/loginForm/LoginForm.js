import React, { Component } from "react"
import VLoginForm from './VLoginForm'
import Request from '../../providers/Request'
import ReactDOM from 'react-dom'

class LoginForm extends Component {
    constructor(props){
        super(props)
        this.state = {
            value:'',
            auth: false,
            error: '',
        }
        this.request = new Request()
	    this.view = VLoginForm.bind(this)
    }

    nicknameOnChange = (event) => {
        this.setState({value:event.target.value})
    }

    login = () => {
        this.request.Put('login', {nickname: this.state.value})
            .then((res)=>{
                if(res.type === "data") {
                    this.setState(() => ({
                        auth: true
                    }))
                }
                else {
		    this.setState(() => ({
                        error: res.data
                    }))
                    console.log(res)
                }
            }).catch((err) => {
            console.log(err)
        })
    }

    render() {
        return this.view()
    }
}

export default LoginForm