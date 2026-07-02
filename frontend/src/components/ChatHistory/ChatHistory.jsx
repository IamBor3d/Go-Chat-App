import Message from "../Message/Message";

    const ChatHistory = ({chatHistory})=>{
        
        const messages = chatHistory.map((msg,index) => (<Message text={msg.data} key={index}/>));

        return (
            <div id="ChatHistory" className="p-2 flex flex-col">
                <h2>Chat History</h2>
                {messages}
            </div>
        );
    }

    export default ChatHistory;