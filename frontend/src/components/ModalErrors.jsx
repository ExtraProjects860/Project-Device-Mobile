import { View, Modal } from "react-native";

export default function ModalErrors({ errorMessage }) {
  // TODO personalizar a buceta do modal e colocar a mensagem de erro, além de chamar nos locais necessários
  // Geralmente nos tratamento de erro, ele precisa ser genérico
  // No caso de locais como login, pode ser melhor colocar legendas embaixo dos inputs, sacou?
  return (
    <View>
      <Modal></Modal>
    </View>
  );
}
