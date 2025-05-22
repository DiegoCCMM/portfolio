import { useState } from "react";
import ClassDiagram from "./ClassDiagram/ClassDiagram";
import type { UMLClass } from "./ClassVisualizer/UMLClass";

const initialCode = ``;

function App() {
  const [code, setCode] = useState(initialCode);
  const [classes, setClasses] = useState<UMLClass[]>([]);


  const handleVisualize = () => {
    const classRegex = /class\s+(\w+)(?:\s+extends\s+(\w+))?/g;
    const fieldRegex = /private\s+\w+\s+(\w+);/g;

    const newClasses: UMLClass[] = [];

    let match;
    while ((match = classRegex.exec(code)) !== null) {
      const [_, className, parentClass] = match;

      // Busca campos solo dentro de esta clase (esto puede mejorarse si usas un parser real)
      const classBodyMatch = new RegExp(
        `class\\s+${className}(?:\\s+extends\\s+\\w+)?\\s*{([\\s\\S]*?)}`,
        "m"
      ).exec(code);

      const body = classBodyMatch?.[1] ?? "";
      const fields = Array.from(body.matchAll(fieldRegex)).map((m) => m[1]);

      newClasses.push({
        className,
        fields,
        parentClass,
      });
    }

    setClasses(newClasses);
  };

  return (
  <div style={{ display: "flex", height: "100vh", width: "100vw", backgroundColor: "#1e1e1e" }}>
    {/* Left: Code input */}
    <div
      style={{
        flex: 1,
        padding: "1rem",
        borderRight: "1px solid #555",
        boxSizing: "border-box"
      }}
    >
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
          resize: "none",
          boxSizing: "border-box"
        }}
      />
      <button
          onClick={handleVisualize}
          style={{
            marginTop: "1rem",
            padding: "0.5rem 1rem",
            fontSize: "1rem",
            backgroundColor: "#4caf50",
            color: "white",
            border: "none",
            borderRadius: "4px",
            cursor: "pointer"
          }}
        >
          Visualize
        </button>
    </div>

    {/* Right: Visualization */}
    <div
      style={{
        flex: 1,
        padding: "1rem",
        color: "white",
        boxSizing: "border-box"
      }}
    >
      <h2>Visualization</h2>
      <ClassDiagram classes={classes} />
    </div>
  </div>
);

}

export default App;
