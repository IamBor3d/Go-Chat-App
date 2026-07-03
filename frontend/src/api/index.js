
//Create websocket
let socket = null //new WebSocket("ws://localhost:8080/ws");

//Listen for events on a socket 
let connect = (roomId, cb) => {

    if (socket) {
        
        socket.close();
    }
    console.log("Attempting Connection");

    socket = new WebSocket(`ws://localhost:8080/room/${roomId}`)
    socket.onopen = () => {
        console.log("Successfully opened");
    };

    socket.onmessage = msg => {
        console.log(msg);
        cb(msg);
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