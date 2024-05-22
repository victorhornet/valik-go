/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./cmd/web/**/*.{templ,html}"],
    theme: {
        extend: {},
    },
    plugins: [require("daisyui")],
};
