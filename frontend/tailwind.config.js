/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'stori-teal': '#003a40',
        'stori-green': '#03d180',
      }
    },
  },
  plugins: [],
}

