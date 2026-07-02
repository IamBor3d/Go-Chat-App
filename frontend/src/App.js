import logo from './logo.svg';
import { sendMsg, connect } from './api';
import { useEffect, useState } from 'react';
import './App.css';
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';
import ChatInput from './components/ChatInput/ChatInput';


let send = (event) => {
  if (event.keyCode === 13) {
    sendMsg(event.target.value);
    event.target.value = "";
  }

}

function App() {
  const [history, setHistory] = useState([]);
  
  
  useEffect(() => {
    connect((msg) => {
      console.log("New Message");
      setHistory(prevHistory => [...prevHistory, msg]);
    });
  }, []);

  

  return (
    <div>
      <Header/>
      <div className="text-center p-3">
        <ChatHistory chatHistory={history}></ChatHistory>
        <ChatInput send={send}/>
      </div>
    </div>
  );
}

export default App;
