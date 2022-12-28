/** @type {import('tailwindcss').Config} */
module.exports = {
  daisyui: {
    themes: [
      {
        mytheme: {

          "primary": "#fbcfe8",


          "secondary": "#e9d5ff",


          "accent": "#B4E9D6",


          "neutral": "#a5f3fc",


          "base-100": "#FFFFFF",


          "info": "#3ABFF8",


          "success": "#36D399",


          "warning": "#FBBD23",


          "error": "#F87272",

        },
      },
    ],
  },
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {

    },
  },
  plugins: [require("daisyui")],
}
