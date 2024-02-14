import './App.css'

function App() {

  function ping(){
    fetch("/api/ping")
      .then(res=>res.json())
      .then(data=>{
        console.log(data)
      }).catch(err=>{
        console.log(err)
      })
  }


  return (
    <div className="App">
      {ping()}

    </div>
  )
}

export default App
