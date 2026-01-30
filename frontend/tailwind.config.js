const animate = require('tailwindcss-animate')

module.exports = {
  content: [
    './index.html',
    './src/**/*.{js,jsx,ts,tsx,vue}'
  ],
  theme: {
    extend: {},
  },
  plugins: [animate],
}
