
const ChatHistory = ({chatHistory})=>{
    const messages = chatHistory.map((msg,index) => (<p key={index}>msg</p>));

    return (
        <div id="ChatHistory">
            <h2>Chat History</h2>
            {messages}
        </div>
    )
}

export default ChatHistory;