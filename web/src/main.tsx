import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import Login from './Login.tsx'

const rootJsx = (
  <React.StrictMode>
    <Login>
      <App />
    </Login>
  </React.StrictMode>
)

const rootElement = document.getElementById('root')!
const root = ReactDOM.createRoot(rootElement)

root.render(rootJsx)
