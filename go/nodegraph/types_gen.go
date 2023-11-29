// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package nodegraph

type ArcOption struct {
	// Field from which to get the value. Values should be less than 1, representing fraction of a circle.
	Field *string `json:"field,omitempty"`
	// The color of the arc.
	Color *string `json:"color,omitempty"`
}

type NodeOptions struct {
	// Unit for the main stat to override what ever is set in the data frame.
	MainStatUnit *string `json:"mainStatUnit,omitempty"`
	// Unit for the secondary stat to override what ever is set in the data frame.
	SecondaryStatUnit *string `json:"secondaryStatUnit,omitempty"`
	// Define which fields are shown as part of the node arc (colored circle around the node).
	Arcs []ArcOption `json:"arcs,omitempty"`
}

type EdgeOptions struct {
	// Unit for the main stat to override what ever is set in the data frame.
	MainStatUnit *string `json:"mainStatUnit,omitempty"`
	// Unit for the secondary stat to override what ever is set in the data frame.
	SecondaryStatUnit *string `json:"secondaryStatUnit,omitempty"`
}

type Options struct {
	Nodes *NodeOptions `json:"nodes,omitempty"`
	Edges *EdgeOptions `json:"edges,omitempty"`
}
