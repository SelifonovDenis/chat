// const socket = new WebSocket('ws://'+window.location.host+'/chat/socket')
//
// let connect = cb => {
//     socket.onopen = () => {
//         console.log("Successfully Connected");
//     };
//
//     socket.onmessage = msg => {
//         console.log(msg);
//         cb(msg);
//     };
//
//     socket.onclose = event => {
//         console.log("Socket Closed Connection: ", event);
//     };
//
//     socket.onerror = error => {
//         console.log("Socket Error: ", error);
//     };
// };
//
// let sendMsg = msg => {
//     console.log("sending msg: ", msg);
//     socket.send(msg);
// };
//
// export { connect, sendMsg };

export default class Socket {

    constructor(url){
        this.ws = new WebSocket(url)

        this.connect = cb =>{
            this.ws.onopen = () => {
                console.log("Successfully Connected");
            };

            this.ws.onmessage = msg => {
                console.log(msg);
                cb(msg);
            };

            this.ws.onclose = event => {
                console.log("Socket Closed Connection: ", event);
            };

            this.ws.onerror = error => {
                console.log("Socket Error: ", error);
            };
        }

        this.sendMsg = msg =>{
            console.log("sending msg: ", msg);
            this.ws.send(msg);
        }
    }

}