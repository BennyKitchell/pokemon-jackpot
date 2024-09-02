
import './Spinner.css'
interface Pokemon {
    id: string,
    image_url: string,
    name: string,
    type: string
}

interface SpinnerProps {
    pokemon: Pokemon[]
}
function Spinner(props: SpinnerProps) {
    const pokemon = props.pokemon;
    return(
        <>
        {
            <div className="spinner-container">
            {pokemon.length > 0 ? (
                pokemon.map((pokemon, key) => (
                  <div key={key} className={"outer-command-container"}>
                    <div className="pokemon-frame" >
                    <img src={pokemon.image_url} width={250} height={250} alt="Pokemon Icon" />
                        </div>
                    <div className="pokemon-name">{pokemon.name}</div>
                  </div>
                ))
              ) : (
                <h2>Syncing Pokemon...</h2>
              )}
              </div>
        }
        </>
    )
}

export default Spinner;