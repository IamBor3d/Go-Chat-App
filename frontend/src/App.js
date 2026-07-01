import logo from './logo.svg';
import { sendMsg, connect } from './api';
import { useEffect } from 'react';
import './App.css';
import Header from './components/Header/Header';


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
    <div>
      <Header/>
      <div className="text-center p-3">
        <button onClick={send} className=' border border-black p-5'> Click Me!</button>
      </div>
    </div>
  );
}

export default App;
