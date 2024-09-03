import { useState } from 'react';
import Spinner from './components/Spinner';
import './App.css';
import pokemonLogo from './assets/pokemon.png';
import jackpotPokemonLogo from './assets/Jackpot.png';
import questionMarkImage from './assets/question.jpg';
import Collection from './components/Collection';
import { Bounce, toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

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

  interface Error {
    message: string
    error: Error
  }

  const fetchPokemon = async (e: React.MouseEvent<HTMLSpanElement, MouseEvent>) => {
    e.preventDefault();
    try {
      if(user?.id) {
      const response =  await fetch(`http://localhost:8084/v1/pokemon/spin/${spinCounter}`, {
        method: "POST",
        body: JSON.stringify(user),
      })
      if (!response.ok) {
        throw new Error('Error fetching pokemon, perhaps the database was not seeded');
      }
      const data = await response.json();
      setPokemon(data.pokemon);
      setSpinCounter(spinCounter+1);
      if(data.jackpot) {
        setJackpot(data.jackpot);
        toast.success("This pokemon was added to your collection!", {
          position: "top-right",
          autoClose: 5000,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: "light",
          transition: Bounce,
        });
      }
      console.log(data);
    }
    } catch(error) {
      const err = error as Error
      toast.error(err.message, {
        position: "top-right",
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "light",
        transition: Bounce,
      });
    }
  };

  const createAccount = async (e: React.MouseEvent<HTMLSpanElement, MouseEvent>) => {
    e.preventDefault();
    const bodyData = JSON.stringify({email, password});
    try {
      const response = await fetch(`http://localhost:8020/v1/user`, {
        method: "POST",
        body: bodyData,
      });
      if (!response.ok) {
        throw new Error('A user with this email already exists, please login',);
      }
      const data = await response.json();
      setUser(data.user);
    } catch(error) {
      const err = error as Error
      toast.error(err.message, {
        position: "top-right",
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "light",
        transition: Bounce,
        });
    };
  };

    const login = async (e: React.MouseEvent<HTMLSpanElement, MouseEvent>) => {
      e.preventDefault();
      const bodyData = JSON.stringify({email, password});
      try {
        const response = await fetch(`http://localhost:8020/v1/login`, {
          method: "POST",
          headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
          },
          body: bodyData,
        });
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setUser(data);
      } catch(error) {
        const err = error as Error
        console.error(error)
        console.log('There has been a problem with your fetch operation: ', err.message);
        toast(err.message, {
          position: "top-right",
          autoClose: 5000,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: "light",
          transition: Bounce,
          });
      };
    };

  return (
    <>
    {
      <div>
        <ToastContainer />
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
