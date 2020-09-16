import React from "react";

function VSendForm() {
    return (
        <div className="sendForm">
            <input type="text" placeholder="Введите сообщение..."  value={this.state.value} onChange={this.messageOnChange}/>
            <button  onClick={this.send}>Отправить</button>
        </div>
    )
}

export default VSendForm;