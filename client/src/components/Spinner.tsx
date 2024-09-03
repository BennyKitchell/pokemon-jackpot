import ConfettiExplosion from 'react-confetti-explosion';

interface Pokemon {
    id: string,
    image_url: string,
    name: string,
    type: string
}

interface SpinnerProps {
    pokemon: Pokemon[]
    jackpot: boolean
}
function Spinner(props: SpinnerProps) {
    const pokemon = props.pokemon;
    const jackpot = props.jackpot;
    return(
        <>
        {
            <div className="spinner-container">
            {jackpot && <ConfettiExplosion />}
            {pokemon.length > 0 ? (
                pokemon.map((pokemon, key) => (
                  <div key={key} className={jackpot ? "outer-command-container logo jackpot" :"outer-command-container"}>
                    <div className="pokemon-frame" >
                    <img src={pokemon.image_url} width={250} height={250} alt="React logo" />
                        </div>
                    <div className="pokemon-name">{pokemon.name}</div>
                  </div>
                ))
              ) : (
                <h2>Collection is empty, spin to win!</h2>
              )}
              {jackpot && <ConfettiExplosion />}
              </div>
        }
        </>
    )
}

export default Spinner;