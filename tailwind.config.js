/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: "class", // or 'media' or 'class'
    content: ["./**/*.{templ,html}"],
    theme: {
        extend: {},
    },
    plugins: [require("daisyui")],
};
