import React, { Component } from "react"
import VChat from './VChat'
import Socket from "../../providers/WebSocket";

class Chat extends Component{

    constructor(props){
        super(props)
        this.view = VChat.bind(this)
        this.state = {
            chatHistory: []
        }
        this.socket = new Socket('ws://'+window.location.host+'/chat/socket')
        // if(this.props.user){
        //     connect()
        // }
        this.socket.connect()

    }

    componentDidMount() {
        this.socket.connect((msg) => {
            if(msg.type === 'message') {
                this.setState(prevState => ({
                    chatHistory: [...this.state.chatHistory, JSON.parse(msg.data)]
                }))
            }
        });
        this.scrollToBottom()
    }

    componentDidUpdate() {
        this.scrollToBottom()
    }

    scrollToBottom = () => {
        this.messagesEnd.scrollIntoView({behavior: "smooth"})
    }

    render(){
        return this.view()
    }
}

export default Chat