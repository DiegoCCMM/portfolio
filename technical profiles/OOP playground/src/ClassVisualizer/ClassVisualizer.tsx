import type { UMLClass } from "./UMLClass";

type Props = {
  umlClass: UMLClass;
};

export default function ClassVisualizer({ umlClass }: Props) {
  return (
    <div
      style={{
        border: "2px solid #aaa",
        borderRadius: "8px",
        width: "220px",
        backgroundColor: "#f9f9f9",
        boxShadow: "0 4px 8px rgba(0,0,0,0.1)",
        fontFamily: "monospace",
        overflow: "hidden"
      }}
    >
      <div
        style={{
          backgroundColor: "#444",
          color: "white",
          padding: "0.5rem 1rem",
          fontWeight: "bold"
        }}
      >
        {umlClass.className}
      </div>
      <div style={{ padding: "0.5rem 1rem", color: "black" }}>
        {umlClass.fields.length > 0 ? (
          umlClass.fields.map((field, idx) => <div key={idx}>{field}</div>)
        ) : (
          <div style={{ color: "#aaa" }}>(no fields)</div>
        )}
      </div>
    </div>
  );
}
