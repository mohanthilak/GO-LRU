import { useState } from "react";
import "./navbar.css"
const Navbar = () => {
    const [seeInfo, setSeeInfo] = useState(false);
    const toggleSeeInfo = (e) => {
        e.preventDefault();
        console.log("toggle")
        setSeeInfo(prev => !prev)
    }
  return (
    <div className="navbar-div">
        <header>
            <h2>Go LRU</h2>
            <button onClick={toggleSeeInfo}> <em>Info</em></button>
        </header>
        <div className={`info-content-div ${seeInfo ? "visible":"not-visible"}`}>
            <div className="left-content-div">
                <p>Go LRU is a LRU cache with time and capacity based eviction policy built with GO a very nice, easy to use and high performance language. The application manages an in-memory key-value cache with very little memory and CPU overhead. It can handle parallel workloads as it is thread-safe. The LRU can be accessed with this react-client, where you can add your data to the cache and get the data from the cache by using the key.</p>
            </div>
            <div className="right-content-div">
                <p className="text-left">Uses a <em>HashMap</em> for constant look up time , it’s value is a pointer to a node in a <em>Linked List</em> to maintain  order.</p>
                <p className="text-right"><em>Initial capacity</em> is set to 4 and the <em>expiration duration</em> is set to 1 minute. Both can be configured. Once either of these rules are violated, the least recently used cache is removed.</p>
                <p className="text-left"><em>Key</em> is of type String and can store any <em>value</em>(uses interface{"{}"}).</p>
            </div>
        </div>

    </div>
  )
}

export default Navbar