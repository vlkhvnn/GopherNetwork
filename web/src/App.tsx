import { useState } from 'react'
import './App.css'

export const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080/v1"

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
    <div>App home screen</div>
    </>
  )
}

export default App
