import { useState } from 'react'
import Spinner from './components/Spinner'
import './App.css'
import pokemonLogo from './assets/pokemon.png'
import jackpotPokemonLogo from './assets/Jackpot.png'
import questionMarkImage from './assets/question.jpg'
import Collection from './components/Collection'
function App() {
  const [email, setEmail] = useState<string>();
  const [password, setPassword] = useState<string>();
  const [collectionEnabled, setCollectionEnabled] = useState<boolean>(false)
  const [jackpot, setJackpot] = useState<boolean>(false)
  const [pokemon, setPokemon] = useState([
    {id: '', name: '', image_url: questionMarkImage, type: ''},
    {id: '', name: '', image_url: questionMarkImage, type: ''},
    {id: '', name: '', image_url: questionMarkImage, type: ''}
  ]);
  const [spinCounter, setSpinCounter] = useState(0);
  const [user, setUser] = useState<User>();


  interface User {
    id: number
  } 

  const fetchPokemon = (e: React.MouseEvent<HTMLSpanElement, MouseEvent>) => {
    e.preventDefault();
    if(user?.id) {
      setJackpot(false)
      fetch(`http://localhost:8084/v1/pokemon/spin/${spinCounter}`, {
        method: "POST",
        body: JSON.stringify(user),
      })
        .then((response) => {
          return response.json();
        })

        .then((data) => {
          setPokemon(data.pokemon);
          setSpinCounter(spinCounter+1);
          if(data.jackpot) {
            setJackpot(data.jackpot);
          }
        });
    }
  };

  const createAccount = (e: React.MouseEvent<HTMLSpanElement, MouseEvent>) => {
    e.preventDefault();
    const data = JSON.stringify({email, password});
      fetch(`http://localhost:8020/v1/user`, {
        method: "POST",
        body: data,
      })
        .then((response) => {
          return response.json();
        })

        .then((data) => {
          setUser(data.user);
        });
    };

    const login = (e: React.MouseEvent<HTMLSpanElement, MouseEvent>) => {
      e.preventDefault();
      const data = JSON.stringify({email, password});
        fetch(`http://localhost:8020/v1/login`, {
          method: "POST",
          headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
          },
          body: data,
        })
          .then((response) => {
            return response.json();
          })
  
          .then((data) => {
            setUser(data);
          });
      };

  return (
    <>
    {
      <div>
          <div className="logo-container">
            <img src={pokemonLogo} className="logo" alt="Pokemon in pokemon font" />
            <img src={jackpotPokemonLogo} className="logo" alt="Jackpot in pokemon font" />
          </div>
          { user?.id && user?.id > -1 ? (
            <div>
            {
              collectionEnabled ? (
                <div>
                  <button className='back-button button' onClick={() => {setCollectionEnabled(!collectionEnabled)}}>Back</button>
                  <Collection userId={user.id}/>
                </div>
    
              ) : 
              <div>
                <button className='login-button button' onClick={() => {setCollectionEnabled(!collectionEnabled)}}>View Collection</button>
                <Spinner pokemon={pokemon} jackpot={jackpot}/>
                <button className='roll-button button' onClick={fetchPokemon}>Roll</button>
              </div>
            }
            </div>
            ): 
            <div>
              <div className='login-container'>
                  <label htmlFor="email">Email</label>
                  <input type="text" id="email" name="email" onChange={e => setEmail(e.target.value)} placeholder="example@google.com" />
                  <label htmlFor="password">Password</label>
                  <input type="password" id="password" name="password" onChange={e => setPassword(e.target.value)} />
                  <button className='login-button button' onClick={login}>Login</button>
                  <button className='login-button button' onClick={createAccount}>Create Account</button>
                </div>
              </div>
          }
        </div>
    }
    </>
  )
}
export default App
