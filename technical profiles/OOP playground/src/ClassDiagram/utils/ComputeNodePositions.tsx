import type { UMLClass } from "../../ClassVisualizer/UMLClass";

export default function ComputeNodePositions(classes: UMLClass[]) {
  const roots = classes.filter(cls => !cls.parentClass);
  const positions = new Map<string, { x: number, y: number }>();
  const horizontalSpacing = 300;
  const verticalSpacing = 200;
  let globalX = 0;

  function layoutTree(className: string, depth: number) {
    const children = classes.filter(c => c.parentClass === className);

    for (const child of children) {
      layoutTree(child.className, depth + 1);
    }

    if (children.length === 0) {
      positions.set(className, {
        x: globalX * horizontalSpacing,
        y: depth * verticalSpacing
      });
      globalX++;
    } else {
      const childPositions = children.map(c => positions.get(c.className)!);
      const minX = Math.min(...childPositions.map(p => p.x));
      const maxX = Math.max(...childPositions.map(p => p.x));
      const centerX = (minX + maxX) / 2;

      positions.set(className, {
        x: centerX,
        y: depth * verticalSpacing
      });
    }
  }

  for (const root of roots) {
    layoutTree(root.className, 0);
  }

  return positions;
}
