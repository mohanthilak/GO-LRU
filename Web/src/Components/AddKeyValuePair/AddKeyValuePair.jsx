import { useRef } from 'react'
import "./style.css"
const AddKeyValuePair = () => {
    const InputKeyRef = useRef();
    const InputValueRef = useRef();

    const handleSubmit = (e) => {
        e.preventDefault();
        const keyText = InputKeyRef.current.value;
        const valueText = InputValueRef.current.value;
        if(keyText.length == 0 || valueText.length == 0) {
            alert("Both Key and Value required")
            return;
        }
        fetch("http://localhost:8080/set", {
            method: 'POST',
            body:JSON.stringify({
                key: keyText,
                value: valueText  
            }),
            headers: { 'Content-Type': 'application/json' }
        }).then(res=>res.json()).then(res=>{
            if(res.status)alert(res.status)
        }).catch(e=>{
            alert("could not set the key value pair")
            console.log("error while setting key value pair:", e)
        })
    }
  return (
    <div className='add-component-div'>
        <div className='header'>
            <h2>Add to cache</h2>
            <p></p>

        </div>
        <form onSubmit={handleSubmit}>
            <div className='input-div'>
                <label htmlFor="add-key">Add Key</label>
                <input ref={InputKeyRef} type="text" id='add-key' />
            </div>
            <div className='input-div'>
                <label htmlFor="add-value">Add Value</label>
                <input ref={InputValueRef} type="text" id='add-value' />
            </div>
            <button>Submit</button>
        </form>
    </div>
  )
}

export default AddKeyValuePair