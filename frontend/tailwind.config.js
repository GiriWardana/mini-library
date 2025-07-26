// tailwind.config.js
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        serif: ['"Merriweather"', 'serif'],
        sans: ['"Inter"', 'sans-serif'],
      },
      colors: {
        primary: '#264653',       // Deep navy for header
        secondary: '#e9c46a',     // Warm yellow accent
        background: '#f8f4ee',    // Soft paper-like background
      },
    },
  },
  plugins: [],
}
