import React, { Component } from "react"
import VSendForm from './VSendForm'

class SendForm extends Component{

    constructor(props){
        super(props)
        this.state = {
            value:'',
        }
        this.view = VSendForm.bind(this)
    }

    messageOnChange = (event) => {
        this.setState({value:event.target.value})
    }

    send = () => {
        let message = {
            sendler: this.props.user,
            text: this.state.value
        }
        this.props.socket.sendMsg(JSON.stringify(message))
        this.setState({value:''})
    }

    render(){
        return this.view()
    }
}

export default SendForm