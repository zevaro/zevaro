import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import App from "./app";

describe("App", () => {
  it('renders "Hello Zevaro" text', () => {
    render(<App />);
    expect(screen.getByText("Hello Zevaro")).toBeInTheDocument();
  });
});
