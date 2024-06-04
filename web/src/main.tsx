import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'

const rootJsx = (
  <React.StrictMode>
    <App />
  </React.StrictMode>
)

const rootElement = document.getElementById('root')
const root = ReactDOM.createRoot(rootElement)

root.render(rootJsx)
