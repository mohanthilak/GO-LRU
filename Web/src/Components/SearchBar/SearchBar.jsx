import { useRef, useState } from "react"
import "./searchbar.css"
const SearchBar = () => {
    const inputRef = useRef()
  const [keyValueData, setSetKeyValueData] = useState({})


    const handleSubmit = (e) => {
      e.preventDefault()
      const searchQuery = inputRef.current.value
      fetch(`http://localhost:8080/get/${searchQuery}`).then((res) => res.json()).then((data) => {console.log(data); setSetKeyValueData(data)}).catch(e=>console.log("error while fetching search results:",e));
  }
  return (
    <div className="search-bar-div">
      <div className="search-bar-header">
        <h2>Search Key</h2>
        <p>{keyValueData.key}:{keyValueData.value}</p>
      </div>
      <form onSubmit={handleSubmit}>
        <div className="input-div">
          <input type="text" ref={inputRef} />
        </div>
          <button>Submit</button>
      </form>
    </div>
  )
}

export default SearchBar