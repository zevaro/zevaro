/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{ts,tsx}"],
  darkMode: ["class", '[data-theme="dark"]'],
  theme: {
    extend: {
      colors: {
        canvas: "var(--bg-canvas)",
        surface: "var(--bg-surface)",
        "surface-elevated": "var(--bg-surface-elevated)",
      },
      borderColor: {
        subtle: "var(--border-subtle)",
        strong: "var(--border-strong)",
      },
      textColor: {
        primary: "var(--text-primary)",
        secondary: "var(--text-secondary)",
        muted: "var(--text-muted)",
      },
      backgroundColor: {
        canvas: "var(--bg-canvas)",
        surface: "var(--bg-surface)",
        "surface-elevated": "var(--bg-surface-elevated)",
      },
    },
  },
  plugins: [],
};
