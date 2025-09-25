import { defineConfig } from "eslint/config";
import js from "@eslint/js";
import globals from "globals";
import react from "eslint-plugin-react";
import reactHooks from "eslint-plugin-react-hooks";
import reactNative from "eslint-plugin-react-native";

export default defineConfig([
  {
    // 1. Ignora os diretórios
    ignores: ["node_modules/", "dist/", "build/"],
  },
  {
    // 2. Este é o nosso ÚNICO objeto de configuração para os arquivos do projeto
    files: ["**/*.{js,jsx}"],

    // Definimos os plugins no formato de objeto correto
    plugins: {
      react,
      "react-hooks": reactHooks,
      "react-native": reactNative,
    },

    languageOptions: {
      ecmaVersion: "latest",
      sourceType: "module",
      globals: {
        ...globals.browser,
        ...globals.node,
        __DEV__: "readonly",
      },
      parserOptions: {
        ecmaFeatures: {
          jsx: true,
        },
      },
    },

    settings: {
      react: {
        version: "detect",
      },
    },

    // A MÁGICA ESTÁ AQUI:
    // Nós mesclamos apenas as REGRAS de cada configuração recomendada...
    rules: {
      ...js.configs.recommended.rules,
      ...react.configs.flat.recommended.rules,
      ...reactHooks.configs.recommended.rules,
      ...reactNative.configs.all.rules,

      // ...e então adicionamos nossas próprias regras para ter a palavra final!
      // Suas regras personalizadas
      "no-undef": "off",
      "no-unused-vars": "warn", // Warn about unused variables
      eqeqeq: ["error", "always"], // Enforce strict equality (===)
      "prefer-const": ["error", { ignoreReadBeforeAssign: true }], // Prefer const for variables not reassigned

      // Ajustes que queremos fazer sobre as regras recomendadas
      "react/prop-types": "off",
      "react/react-in-jsx-scope": "off",
      "react-native/no-inline-styles": "warn",
      "react-native/no-color-literals": "off",
    },
  },
]);
