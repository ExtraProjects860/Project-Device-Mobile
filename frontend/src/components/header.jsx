import {
  View,
  StyleSheet,
  Image,
  TouchableOpacity,
  StatusBar,
} from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
const logo = require("../assets/images/logo.png");
/*** * @param {object} props* @param {function} props.onMenuPress*/ export default function Header({
  onMenuPress,
}) {
  return (
    <View className="flex-row items-center justify-between bg-magenta px-4 py-3 shadow-md">
      <StatusBar barStyle={"light-content"} className="bg-magenta" />
      <Image source={logo} className="px-230 py-200" resizeMode="contain" />
      <TouchableOpacity onPress={onMenuPress} className="p-[5px]">
        <Icon name="menu" size={60} color="white" />
      </TouchableOpacity>
    </View>
  );
}
