import { useRef, useState } from 'react'
import './App.css'
import AddKeyValuePair from './Components/AddKeyValuePair/AddKeyValuePair'
function App() {
  const [keyValueData, setSetKeyValueData] = useState({})
  const inputRef = useRef()

    const handleSubmit = (e) => {
        e.preventDefault()
        const searchQuery = inputRef.current.value
        fetch(`http://localhost:8080/get/${searchQuery}`).then((res) => res.json()).then((data) => {console.log(data); setSetKeyValueData(data)}).catch(e=>console.log("error while fetching search results:",e));
    }
  return (
    <div>
      <AddKeyValuePair />
      <form onSubmit={handleSubmit}>
          <label htmlFor="search-key">Enter key to find value</label>
          <input type="text" ref={inputRef} />
          <button>Submit</button>
      </form>

      <div>
        <p>{keyValueData.key}:{keyValueData.value}</p>
      </div>
    </div>
  )
}

export default App
