import { useState } from 'react'
import Spinner from './components/Spinner'
import './App.css'
import pokemonLogo from './assets/pokemon.png'
import jackpotPokemonLogo from './assets/Jackpot.png'
import questionMarkImage from './assets/question.jpg'
function App() {
  const [email, setEmail] = useState<string>();
  const [password, setPassword] = useState<string>();
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
      fetch(`http://localhost:8084/pokemon/roll/${spinCounter}`, {
        method: "POST",
        body: JSON.stringify(user),
      })
        .then((response) => {
          return response.json();
        })

        .then((data) => {
          console.log(data);
          setPokemon(data.pokemon);
          setSpinCounter(spinCounter+1);
        });
    }
  };

  const createAccount = (e: React.MouseEvent<HTMLSpanElement, MouseEvent>) => {
    e.preventDefault();
    const data = JSON.stringify({email, password});
    console.log(data)
      fetch(`http://localhost:8020/user`, {
        method: "POST",
        body: data,
      })
        .then((response) => {
          return response.json();
        })

        .then((data) => {
          console.log(data);
          setUser(data.user);
        });
    };

    const login = (e: React.MouseEvent<HTMLSpanElement, MouseEvent>) => {
      e.preventDefault();
      const data = JSON.stringify({email, password});
        fetch(`http://localhost:8020/login`, {
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
            console.log(data);
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
              <Spinner pokemon={pokemon}/>
              <button className='roll-button button' onClick={fetchPokemon}>Spin</button>
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
