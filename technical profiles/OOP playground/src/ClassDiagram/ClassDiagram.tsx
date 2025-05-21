import type { UMLClass } from "../ClassVisualizer/UMLClass";
import ReactFlow, {
  Background,
  Controls,
  MarkerType,
} from "reactflow";
import type {Edge, Node} from "reactflow";
import "reactflow/dist/style.css";

type Props = {
  classes: UMLClass[];
};

export default function ClassDiagram({ classes }: Props) {
  const nodes: Node[] = classes.map((cls, index) => ({
    id: cls.className,
    data: {
      label: (
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
            {cls.className}
          </div>
          <div style={{ padding: "0.5rem", color: "#000" }}>
            {cls.fields.length > 0 ? (
              cls.fields.map((field, idx) => <div key={idx}>{field}</div>)
            ) : (
              <div style={{ color: "#aaa" }}>(no fields)</div>
            )}
          </div>
        </div>
      )
    },
    position: { x: 100 + index * 300, y: 100 },
    style: { width: 200 },
    type: "default"
  }));

  const edges: Edge[] = classes
    .filter((cls) => cls.parentClass)
    .map((cls) => ({
      id: `${cls.className}->${cls.parentClass}`,
      source: cls.className,
      target: cls.parentClass!,
      label: "extends",
      markerEnd: {
        type: MarkerType.ArrowClosed
      },
      style: {
        stroke: "#222",
        strokeWidth: 2
      }
    }));

  return (
    <div style={{ width: "100%", height: "100%" }}>
      <ReactFlow
        nodes={nodes}
        edges={edges}
        fitView
        nodesDraggable={true}
        nodesConnectable={false}
        elementsSelectable={false}
      >
        <Background />
        <Controls />
      </ReactFlow>
    </div>
  );
}
