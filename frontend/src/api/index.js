
//Create websocket
var socket = new WebSocket("ws://localhost:8080/ws");

//Listen for events on a socket 
let connect = () => {

    console.log("Attempting Connection");

    socket.onopen = () => {
        console.log("Successfully opened");
    };

    socket.onmessage = msg => {
        console.log(msg);
    };

    socket.onclose = event => {
        console.log("Closed Connection: ", event);
    };

    socket.onerror = err =>{
        console.log("Server Error", err);
    };

};

//send a message to backend
let sendMsg = msg => {
    console.log("sending msg:", msg);
    socket.send(msg);
};

export { connect, sendMsg}