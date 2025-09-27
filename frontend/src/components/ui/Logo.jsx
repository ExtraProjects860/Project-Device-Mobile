import React from 'react';
import {Image, View} from 'react-native'
import logo from '../../assets/images/logo.png'

export default function Logo({className = " ", resizedMode = "contain"}){
    return (
        <Image source={logo} className={className} resizedMode={resizedMode}/>
    );
}