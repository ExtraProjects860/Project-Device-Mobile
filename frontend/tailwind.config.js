/** @type {import('tailwindcss').Config} */ module.exports = {
  // NOTE: Update this to include the paths to all files that contain Nativewind classes.
  content: [
    "./App.{js,jsx,ts,tsx}",
    "./src/components/**/*.{js,jsx,ts,tsx}",
    "./src/screens/**/*.{js,jsx,ts,tsx}",
  ],
  presets: [require("nativewind/preset")],
  theme: {
    extend: {
      colors: {
        'magenta': "#E91D62",
        'teal': "#087975",
        'white': "#ffffff",
        'black': "#000000",
      },
    },
  },
  plugins: [],
};
