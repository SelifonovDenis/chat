import React from 'react'
import ReactDOM from 'react-dom'
import 'core-js/stable'
import 'regenerator-runtime/runtime'

import { BrowserRouter } from "react-router-dom"
import {createBrowserHistory} from 'history'

import App from "./components/App";

const history = createBrowserHistory()

ReactDOM.render((
   <BrowserRouter history={history}>
       <App />
   </BrowserRouter>
   ), document.getElementById('root')
);