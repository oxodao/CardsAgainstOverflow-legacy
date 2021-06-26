import React from 'react';

type Props = {
    isDeported: boolean;
};

export default function JoinGame({isDeported}: Props) {

    return <div>
        <h1>Code salle</h1>
        {
            !isDeported &&
            <p>Pour rejoindre une salle, remplissez le code donné par l'administrateur de la partie.</p>
        }
        {
            isDeported &&
            <p>Pour utiliser l'affichage déporté, complétez avec le code salle de la partie.</p>
        }
        <input type="text"/>

        <input type="submit"/>
    </div>;
}