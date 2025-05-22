import type { UMLClass } from "../ClassVisualizer/UMLClass";
import ReactFlow, {
  Background,
  Controls,
  MarkerType,
  Position
} from "reactflow";
import type {Edge, Node} from "reactflow";
import "reactflow/dist/style.css";
import ClassDiagramView from "./view/ClassDiagramView";
import ComputeNodePositions from "./utils/ComputeNodePositions";

type Props = {
  classes: UMLClass[];
};

export default function ClassDiagram({ classes }: Props) {
  const nodes: Node[] = classes.map((cls) => ({
    id: cls.className,
    data: {
      label: (
        <ClassDiagramView className={cls.className} fields={cls.fields} />
      )
    },
    position: ComputeNodePositions(classes).get(cls.className) || { x: 0, y: 0 },
    style: { width: 200 },
    type: "default",
    sourcePosition: Position.Top,
    targetPosition: Position.Bottom
  }));

  const edges: Edge[] = classes
    .filter((cls) => cls.parentClass)
    .map((cls) => ({
      id: `${cls.className}->${cls.parentClass}`,
      source: cls.className,
      target: cls.parentClass!,
      type: "smoothstep",
      label: "extends",
      markerEnd: {
        type: MarkerType.ArrowClosed
      },
      style: {
        stroke: "#ccc",
        strokeWidth: 2
      }
    }));
    
    console.log("nodes", nodes);
    console.log("edges", edges);

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
