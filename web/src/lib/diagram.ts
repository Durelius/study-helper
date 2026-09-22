// Diagram specs for the visual questions.
//
// The exam shows a model and asks you to circle what is wrong with it, so these are
// authored as data and drawn as SVG with clickable regions: answering means pointing
// at the offending element, not picking a sentence that describes it.
//
// A hotspot's `label` is also the question's choice text, so the same answer can be
// given by clicking the diagram or by choosing from the list underneath — which keeps
// the question answerable by keyboard, and survives a diagram that renders oddly.

export type Hotspot = {
  /** Which element is clickable: a node id, or "from>to" for a flow. */
  target: string
  label: string
}

export type BpmnNodeType =
  | 'start'
  | 'startMessage'
  | 'intermediate'
  | 'end'
  | 'endMessage'
  | 'task'
  | 'subprocess'
  | 'xor'
  | 'and'
  | 'or'
  /** Draws nothing: a point for a message flow to touch on a black-box pool's edge. */
  | 'anchor'

export type BpmnNode = {
  id: string
  type: BpmnNodeType
  lane: string
  /** Horizontal position in the lane's content area, in pixels. */
  x: number
  /** Vertical nudge from the lane's centre line. */
  dy?: number
  label?: string
  /** Draws the node straddling the lane below it — the deck's "activity not in a lane" error. */
  straddle?: boolean
}

export type BpmnFlow = {
  from: string
  to: string
  label?: string
  /** message: dashed with an open arrowhead, for exchanges between pools. */
  kind?: 'sequence' | 'message'
  /** Route via a horizontal detour at this y offset, for flows that double back. */
  detour?: number
}

/** `h` overrides the default lane height, for lanes holding parallel branches. */
export type BpmnLane = { id: string; label: string; h?: number }
export type BpmnPool = { label: string; lanes: BpmnLane[]; blackBox?: boolean }

export type BpmnDiagram = {
  kind: 'bpmn'
  width: number
  pools: BpmnPool[]
  nodes: BpmnNode[]
  flows: BpmnFlow[]
  hotspots: Hotspot[]
}

export type WbsNode = {
  id: string
  /** WBS numbering as the deck writes it: 1.0, 1.1, 1.2.3. */
  code?: string
  label: string
  level: number
  parent?: string
}

export type WbsDiagram = {
  kind: 'wbs'
  width: number
  nodes: WbsNode[]
  hotspots: Hotspot[]
}

export type MatrixDiagram = {
  kind: 'matrix'
  columns: string[]
  rows: { id: string; label: string; cells: string[] }[]
  /** Cell ids are "rowId:columnIndex". */
  hotspots: Hotspot[]
  caption?: string
}

export type DiagramSpec = BpmnDiagram | WbsDiagram | MatrixDiagram

// Geometry, kept here so the renderers and the hotspot overlays agree.
export const LANE_H = 96
export const LABEL_W = 22
export const TASK_W = 116
export const TASK_H = 50
export const EVENT_R = 17
export const GATE = 20
