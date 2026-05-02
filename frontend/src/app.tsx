/**
 * App — top-level shell component for the Zevaro desktop application.
 *
 * At skeleton stage this renders the "Hello Zevaro" splash screen using the
 * dark-theme palette defined in globals.css (§5 of Zevaro-Architecture.md).
 * Navigation, system tray, and route layout are implemented in ZV-041.
 */
export default function App() {
  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        height: "100%",
        backgroundColor: "var(--bg-canvas)",
      }}
    >
      <h1
        style={{
          color: "var(--text-primary)",
          fontFamily: "system-ui, -apple-system, sans-serif",
          fontSize: "2rem",
          fontWeight: 600,
          margin: 0,
          letterSpacing: "-0.02em",
        }}
      >
        Hello Zevaro
      </h1>
    </div>
  );
}
