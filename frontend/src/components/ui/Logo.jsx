import React from 'react';
import {Image} from 'react-native'
import logo from '../../assets/images/logo.png'

/**
 * Componente Usado para carregar a logo do app de forma simplificada
 * 
 * Recebe 2 atributos.
 * O primeiro chamado className responsável por adicionar estilização a imagem.
 * O segundo resizedMode responsável por modificar o redimensionamento da imagem, possui valor padrão de contain.
 */
export default function Logo({className = " ", resizedMode = "contain"}){
    return (
        <Image source={logo} className={className} resizedMode={resizedMode}/>
    );
}