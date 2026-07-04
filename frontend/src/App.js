import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import ChatRoom from './pages/ChatRoom';
import Header from './components/Header/Header';
import { RoomList } from './pages/RoomList';
import MainMenu from './pages/MainMenu';
import MakeRoom from './pages/MakeRoom';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<MainMenu/>}/>
        <Route path="/rooms" element={<RoomList/>} />
        <Route path="/makeRoom" element={<MakeRoom/>}/>
        {/* The actual chat room page with a dynamic room ID */}
        <Route path="/room/:roomId" element={<ChatRoom />} />
      </Routes>
    </BrowserRouter>
  );
}
export default App;
