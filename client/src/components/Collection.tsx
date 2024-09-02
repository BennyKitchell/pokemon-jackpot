import { useEffect, useState } from "react";
import './Collection.css';

interface Pokemon {
    id: string,
    image_url: string,
    name: string,
    type: string
}

interface collectionProps {
    userId: number
}
function Collection(props: collectionProps) {
    const userId = props.userId;
    const [pokemon, setPokemon] = useState<Pokemon[]>([]);

    useEffect(() => {
        fetch(`http://localhost:8084/v1/collection/${userId}`, {
            method: "GET",
        })
        .then((response) => {
            return response.json();
        })

        .then((data) => {
            setPokemon(data.pokemon);
        });
    }, []);

    return(
        <>
        {
            <div className="grid-container">
            {pokemon.length > 0 ? (
                pokemon.map((pokemon, key) => (
                  <div key={key} className="outer-command-container grid-item">
                    <div className="pokemon-frame" >
                    <img src={pokemon.image_url} width={250} height={250} alt="React logo" />
                        </div>
                    <div className="pokemon-name">{pokemon.name}</div>
                  </div>
                ))
              ) : (
                <h2>Collection is Syncing...</h2>
              )}
              </div>
        }
        </>
    )
}

export default Collection;