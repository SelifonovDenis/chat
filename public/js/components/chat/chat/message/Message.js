import React, { Component } from "react"
import VMessage from './VMessage'

class Message extends Component{

    constructor(props){
        super(props)
        this.view = VMessage.bind(this)
    }

    render(){
        return this.view(this.props.message.sendler.id === this.props.user.id)
    }
}

export default Message