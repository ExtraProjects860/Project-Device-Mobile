import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, Image } from 'react-native';
import { MaterialIcons, Feather } from '@expo/vector-icons';
import Background from '../components/ui/Background';
import Logo from '../components/ui/Logo'

export default function LoginScreen(){
    const [showPassword, setShowPassword] = useState(false);
        return (
            <Background>
                
                <View className="items-center mt-4">
                    <Logo className="w-24 h-24" />
                    <View>
                        <Text className="text-white font-bold text-3xl mb-1">Olá! Seja Bem Vindo!</Text>
                    </View>

                    <View className="w-4/5 h-0.5 bg-white my-3" />

                    <View>
                        <Text className="text-white font-bold text-base mb-1">E-email: </Text>
                        <View className="flex-row items-center bg-white rounded-full px-4 py-2 w-full">
                            <MaterialIcons name="email" size={20} color="teal" />
                            <TextInput
                                className="ml-2 flex-1 text-gray-700 font-semibold"
                                placeholder="Exemplo@gmail.com"
                                placeholderTextColor="gray"
                                keyboardType="email-address"
                            />
                        </View>
                    </View>
                    
                    <Text className="w-full text-white font-bold text-base mb-1">Senha:</Text>
                    <View className="flex-row items-center bg-white rounded-full px-4 py-2 w-full">
                        <MaterialIcons name="lock" size={20} color="teal" />
                        <TextInput
                            className="ml-2 flex-1 text-gray-700 font-semibold"
                            secureTextEntry={!showPassword}
                            placeholder="*********"
                            placeholderTextColor="gray"
                        />
                        <TouchableOpacity onPress={() => setShowPassword(!showPassword)}>
                            <Feather
                                name={showPassword ? 'eye-off' : 'eye'}
                                size={20}
                                color="teal"
                            />
                        </TouchableOpacity>  
                    </View>

                    <View> 
                    <TouchableOpacity className="mt-20 p-2 bg-magenta rounded-xl items-center">
                        <Text className="text-white text-xl">Entrar</Text>
                    </TouchableOpacity>
                    <TouchableOpacity>
                        <Text className="text-white text-s underline underline-offset-1">Esqueceu a Senha? Clique Aqui!</Text>
                    </TouchableOpacity>
                    </View>
                   

                    
                </View>
            </Background>

        );
}