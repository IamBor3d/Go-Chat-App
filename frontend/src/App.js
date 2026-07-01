import logo from './logo.svg';
import { sendMsg, connect } from './api';
import { useEffect } from 'react';
import './App.css';


let send = ()=> {
  let msg = "Hello";
  console.log(msg);
  sendMsg(msg);
}

function App() {
  useEffect(() => {
    connect();
  }, []);
  return (
    <div className="text-center border border-red-400">
      <button onClick={send}> Click Me!</button>
    </div>
  );
}

export default App;
