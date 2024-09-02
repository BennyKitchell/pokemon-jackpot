import { useState } from 'react'
import Spinner from './components/Spinner'
import './App.css'
import pokemonLogo from './assets/pokemon.png'
import jackpotPokemonLogo from './assets/Jackpot.png'
import questionMarkImage from './assets/question.jpg'
function App() {
  const [pokemon, setPokemon] = useState([
    {id: '', name: '', image_url: questionMarkImage, type: ''},
    {id: '', name: '', image_url: questionMarkImage, type: ''},
    {id: '', name: '', image_url: questionMarkImage, type: ''}
  ]);

  return (
    <>
          <div className="logo-container">
            <img src={pokemonLogo} className="logo" alt="Pokemon in pokemon font" />
            <img src={jackpotPokemonLogo} className="logo" alt="Jackpot in pokemon font" />
          </div>
          <div>
            <Spinner pokemon={pokemon}/>
          </div>
    </>
  )
}

export default App
