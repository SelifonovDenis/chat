import React from "react";

function VMessage(myMessage) {
    if (myMessage){
        return (
            <div className="Row">
                <div className="myMessage">
                    <div className="sendler">{this.props.message.sendler.nickname}</div>
                    <div className="text">{this.props.message.text}</div>
                </div>
            </div>
        )
    } else{
        return (
            <div className="Row">
                <div className="message">
                    <div className="sendler">{this.props.message.sendler.nickname}</div>
                    <div className="text">{this.props.message.text}</div>
                </div>
            </div>
        )
    }
}

export default VMessage;