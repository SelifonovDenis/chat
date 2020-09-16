import React from "react";
import SendForm from "./sendForm/SendForm";
import Message from "./message/Message";

function VChat() {
    return (
        <div className="chat">
            <div className="head"><p>Вопрос дня: Навальный или Путин?</p></div>
            <div className="messages">
                {this.state.chatHistory.map(message => {
                    return <Message message={message} user={this.props.user}/>
                })}
                <div style={{float:"left", clear:"both"}} ref={(el)=>{this.messagesEnd = el;}}></div>
            </div>
            <SendForm user={this.props.user} socket={this.socket}/>
        </div>
    )
}

export default VChat;