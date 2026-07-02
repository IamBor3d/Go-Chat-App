
const ChatInput = ({send}) => {

    return (
        <div>
           <input className="border border-black" onKeyDown={send}></input> 
        </div>
        
    );
}

export default ChatInput; 