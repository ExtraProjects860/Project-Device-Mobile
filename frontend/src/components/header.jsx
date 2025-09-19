import { View, StyleSheet, Image, TouchableOpacity, StatusBar } from "react-native";
import Icon from 'react-native-vector-icons/MaterialCommunityIcons';

const logo = require('../assets/images/logo.png');

/**
 * * @param {object} props
 * @param {function} props.onMenuPress 
 */

export default function Header({ onMenuPress }) { 
    return(
        <View style={styles.container}>
            <StatusBar className=''/>

            <Image source={logo} style={styles.logo}/> 

            <TouchableOpacity onPress={onMenuPress} style={styles.menuButton}>
              <Icon name="menu" size={30} color='#FFF'/>  
            </TouchableOpacity>
        </View>
    );
};
