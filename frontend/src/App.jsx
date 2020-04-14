import React from 'react';
import { ToastContainer } from 'react-toastify';

import {useReducers} from './context';
import Home from './components/home';

import './assets/App.css';

function App() {
  const [global, dispatch] = useReducers()

  console.log(global.User.Username)
  console.log(global.User.Username.length)
  
  if (global.User.Username.length === 0) {
    return <div className="App" onClick={() => console.log(global)}>
      <Home />
      <ToastContainer autoClose={5000} />
    </div>
  }

  return (
    <div className="App">
      Ingame
      <ToastContainer autoClose={5000} />
    </div>
  );
}

export default App;
