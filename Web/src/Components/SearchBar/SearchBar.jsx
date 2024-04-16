import { useRef } from "react"

const SearchBar = () => {
    const inputRef = useRef()

    const handleSubmit = (e) => {
        e.preventDefault()
        const searchQuery = inputRef.current.value
        fetch(`http://localhost:8000/${searchQuery}`).then((res) => res.json()).then((data) => {console.log(data)}).catch(e=>console.log("error while fetching search results:",e));
    }
  return (
    <div>
        <form onSubmit={handleSubmit}>
            <input type="text" ref={inputRef} />
            <button>Submit</button>
        </form>
    </div>
  )
}

export default SearchBar