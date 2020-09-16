import React, { Component } from "react"
import Toolbar from './toolbar/Toolbar'
import Chat from './chat/Chat'
import '../../../css/chat.css'
import Header from "./header/Header"
import Request from '../providers/Request'

class ChatPage extends Component {
    constructor(props){
        super(props)
        this.state = {
            user: null,
        }
        this.request = new Request()
        this.getUser()
    }

    getUser(){
        this.request.Get('user')
            .then((res) => {
                if(res.type === "data"){
                    this.setState(prevState => ({
                        user: res.data
                    }))
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
	    if(!this.state.user){
            return(
                <div></div>
            )
	    }

        return (
            <div id="page">
                <Header/>
                <Toolbar user={this.state.user}/>
                <Chat user={this.state.user}/>
            </div>
        );
    }
}

export default ChatPage;