import { useState, useEffect } from "react";
import ClassDiagram from "./ClassDiagram/ClassDiagram";
import type { UMLClass } from "./ClassVisualizer/UMLClass";

const initialCode = ``;

function App() {
  const [code, setCode] = useState(initialCode);
  const [classes, setClasses] = useState<UMLClass[]>([]);

  useEffect(() => {
    const classRegex = /class\s+([A-Za-z0-9_]+)(?:\s+extends\s+([A-Za-z0-9_]+))?\s*\{([\s\S]*?)\}/g;
    const matches = Array.from(code.matchAll(classRegex));
    const parsed: UMLClass[] = matches.map((match) => {
      const className = match[1];
      const parentClass = match[2];
      const body = match[3];

      const fields = Array.from(
        body.matchAll(/private\s+\w+\s+(\w+);/g)
      ).map((m) => m[1]);

      return {
        className,
        fields,
        parentClass,
      };
    });

    setClasses(parsed);
  }, [code]);

  return (
    <div style={{ display: "flex", height: "100vh", backgroundColor: "#1e1e1e" }}>
      {/* Left: Code input */}
      <div style={{ flex: 1, padding: "1rem", borderRight: "1px solid #555" }}>
        <h2 style={{ color: "white" }}>Java Code</h2>
        <textarea
          value={code}
          onChange={(e) => setCode(e.target.value)}
          style={{
            width: "100%",
            height: "90%",
            fontFamily: "monospace",
            fontSize: "1rem",
            padding: "1rem",
            backgroundColor: "#2d2d2d",
            color: "#eee",
            borderRadius: "8px",
            border: "1px solid #444",
            resize: "none"
          }}
        />
      </div>

      {/* Right: Visualization */}
      <div style={{ flex: 1, padding: "1rem", color: "white" }}>
        <h2>Visualization</h2>
        <ClassDiagram classes={classes} />
      </div>
    </div>
  );
}

export default App;
