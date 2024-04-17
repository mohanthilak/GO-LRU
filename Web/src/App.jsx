// import { useRef, useState } from 'react'
import './App.css'
import AddKeyValuePair from './Components/AddKeyValuePair/AddKeyValuePair'
import Navbar from './Components/Navbar/Navbar'
import SearchBar from './Components/SearchBar/SearchBar'
function App() {

  return (
    <div>
      <Navbar />
      <div className='main-playground'>

        <div className="add-component">
          <AddKeyValuePair />
        </div>
        <div className="add-component">
          <SearchBar />
        </div>
      </div>
    </div>
  )
}

export default App

