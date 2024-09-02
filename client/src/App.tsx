import { useState } from 'react'
import Spinner from './components/Spinner'
import './App.css'

function App() {
  const [pokemon, setPokemon] = useState([
    {id: '1', name: '', image_url: "https://i.postimg.cc/254jR6Fb/red-neon-question-mark-on-a-black-background-vector-1925641335.jpg", type: 'Grass'},
    {id: '1', name: '', image_url: "https://i.postimg.cc/254jR6Fb/red-neon-question-mark-on-a-black-background-vector-1925641335.jpg", type: 'Grass'},
    {id: '1', name: '', image_url: "https://i.postimg.cc/254jR6Fb/red-neon-question-mark-on-a-black-background-vector-1925641335.jpg", type: 'Grass'}
  ]);

  return (
    <>
          <div>
            <Spinner pokemon={pokemon}/>
          </div>
    </>
  )
}

export default App
