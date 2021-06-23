import React from 'react'
import ReactDOM from 'react-dom'

import { 
  BrowserRouter as Router,
  Route
} from 'react-router-dom';

import './assets/scss/_main.scss';
import Login from './views/Login';

ReactDOM.render(
  <React.StrictMode>
    <Router>
      <Route>
        <Login/>
      </Route>
    </Router>
  </React.StrictMode>,
  document.getElementById('root')
)
