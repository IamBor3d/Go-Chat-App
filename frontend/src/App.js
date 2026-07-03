import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import ChatRoom from './pages/ChatRoom';
import Header from './components/Header/Header';
import { RoomList } from './pages/RoomList';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        {/* The landing page to pick or create a room */}
        <Route path="/" element={<RoomList/>} />
        
        {/* The actual chat room page with a dynamic room ID */}
        <Route path="/room/:roomId" element={<ChatRoom />} />
      </Routes>
    </BrowserRouter>
  );
}
export default App;
