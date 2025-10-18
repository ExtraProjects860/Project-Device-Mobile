/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./App.{js,jsx,ts,tsx}",
    "./src/components/**/*.{js,jsx,ts,tsx}",
    "./src/screens/**/*.{js,jsx,ts,tsx}",
  ],
  presets: [require("nativewind/preset")],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        "gray-strong": "#374151",
        "gray-soft": "#e5e7eb",

        light: {
          primary: "#087975",
          secondary: "#E91D62",
          card: "#ffffff",
          "text-primary": "#1e293b",
          "text-secondary": "#475569",
          "text-inverted": "#ffffff",
        },
        dark: {
          primary: "#0f172a",
          secondary: "#E91D62",
          card: "#1e293b",
          "text-primary": "#f1f5f9",
          "text-secondary": "#94a3b8",
        },
      },
    },
  },
  plugins: [],
};
