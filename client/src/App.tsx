import { useState } from 'react'
import Spinner from './components/Spinner'
import './App.css'
import pokemonLogo from './assets/pokemon.png'
import jackpotPokemonLogo from './assets/Jackpot.png'
import questionMarkImage from './assets/question.jpg'
function App() {
  const [spinCounter, setSpinCounter] = useState(0);
  const [User, setUser] = useState<User>();
  const [pokemon, setPokemon] = useState([
    {id: '', name: '', image_url: questionMarkImage, type: ''},
    {id: '', name: '', image_url: questionMarkImage, type: ''},
    {id: '', name: '', image_url: questionMarkImage, type: ''}
  ]);

  interface User {
    id: number
  } 

  const fetchPokemon = (e: React.MouseEvent<HTMLSpanElement, MouseEvent>) => {
    e.preventDefault();
    if(User?.id) {
      fetch(`http://localhost:8084/pokemon/roll/${spinCounter}`, {
        method: "POST",
        body: JSON.stringify(User),
      })
        .then((response) => {
          return response.json();
        })

        .then((data) => {
          console.log(data)
          setPokemon(data.pokemon);
          setSpinCounter(spinCounter+1);
        });
    }
  };

  return (
    <>
          <div className="logo-container">
            <img src={pokemonLogo} className="logo" alt="Pokemon in pokemon font" />
            <img src={jackpotPokemonLogo} className="logo" alt="Jackpot in pokemon font" />
          </div>
          <div>
            <Spinner pokemon={pokemon}/>
            <button className='roll-button button' onClick={fetchPokemon}>Spin</button>
          </div>
    </>
  )
}

export default App
