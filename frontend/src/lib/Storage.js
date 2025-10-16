import AsyncStorage from "@react-native-async-storage/async-storage";

export const Storage = {
  async setItem(key, value) {
    try {
      await AsyncStorage.setItem(key, JSON.stringify(value));
    } catch (error) {
      // TODO vai precisar tratar isso melhor depois
      console.error(`Erro ao salvar ${key}:`, error);
    }
  },

  async getItem(key) {
    try {
      const value = await AsyncStorage.getItem(key);
      return value ? JSON.parse(value) : null;
    } catch (error) {
      console.error(`Erro ao ler ${key}:`, error);
      return null;
    }
  },

  async removeItem(key) {
    try {
      await AsyncStorage.removeItem(key);
    } catch (error) {
      // TODO vai precisar tratar isso melhor depois
      console.error(`Erro ao remover ${key}:`, error);
    }
  },
};
