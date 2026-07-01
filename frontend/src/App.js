import logo from './logo.svg';
import { sendMsg, connect } from './api';
import { useEffect, useState } from 'react';
import './App.css';
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';


let send = ()=> {
  let msg = "Hello";
  console.log(msg);
  sendMsg(msg);
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
        <button onClick={send} className=' border border-black p-5'> Click Me!</button>
      </div>
    </div>
  );
}

export default App;
