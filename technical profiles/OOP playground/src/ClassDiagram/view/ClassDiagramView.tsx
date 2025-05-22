
type Props = {
  className: string;
  fields: string[];
};

export default function ClassDiagramView({ className, fields }: Props) {
  return (
    <div
      style={{
        border: "2px solid #aaa",
        borderRadius: "8px",
        width: "200px",
        backgroundColor: "#f9f9f9",
        fontFamily: "monospace",
        overflow: "hidden"
      }}
    >
      <div
        style={{
          backgroundColor: "#444",
          color: "white",
          padding: "0.5rem",
          fontWeight: "bold"
        }}
      >
        {className}
      </div>
      <div style={{ padding: "0.5rem", color: "#000" }}>
        {fields.length > 0 ? (
          fields.map((field, idx) => <div key={idx}>{field}</div>)
        ) : (
          <div style={{ color: "#aaa" }}>(no fields)</div>
        )}
      </div>
    </div>
  );
}
