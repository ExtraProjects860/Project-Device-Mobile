import React from "react";
import { View, Image } from "react-native";
import footer from "../../assets/images/footer.png";

/** 
 * Componente responsável pela View do background 
 * 
 * Recebe 2 atributos 
 * O primeiro chamado children responsável pro fazer o carregamento do conteudo que for passado dentro a tag BackGround
 * O Segundo chamdo className responsável por receber estilização para se encrementada ao estilo atual de BackGround
*/

export default function Background({ children, className = "" }) {
  return (
    <View className={`flex-1 bg-light-primary dark:bg-dark-primary -z-10 ${className}`}>
      {children}
      <View className="absolute bottom-0 w-full h-[400px] -z-10">
        <Image source={footer} className="w-full h-full" resizeMode="stretch" />
      </View>
    </View>
  );
}