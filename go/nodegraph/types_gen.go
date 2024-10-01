// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package nodegraph

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type ArcOption struct {
	// Field from which to get the value. Values should be less than 1, representing fraction of a circle.
	Field *string `json:"field,omitempty"`
	// The color of the arc.
	Color *string `json:"color,omitempty"`
}

func (resource ArcOption) Equals(other ArcOption) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Color == nil && other.Color != nil || resource.Color != nil && other.Color == nil {
		return false
	}

	if resource.Color != nil {
		if *resource.Color != *other.Color {
			return false
		}
	}

	return true
}

type NodeOptions struct {
	// Unit for the main stat to override what ever is set in the data frame.
	MainStatUnit *string `json:"mainStatUnit,omitempty"`
	// Unit for the secondary stat to override what ever is set in the data frame.
	SecondaryStatUnit *string `json:"secondaryStatUnit,omitempty"`
	// Define which fields are shown as part of the node arc (colored circle around the node).
	Arcs []ArcOption `json:"arcs,omitempty"`
}

func (resource NodeOptions) Equals(other NodeOptions) bool {
	if resource.MainStatUnit == nil && other.MainStatUnit != nil || resource.MainStatUnit != nil && other.MainStatUnit == nil {
		return false
	}

	if resource.MainStatUnit != nil {
		if *resource.MainStatUnit != *other.MainStatUnit {
			return false
		}
	}
	if resource.SecondaryStatUnit == nil && other.SecondaryStatUnit != nil || resource.SecondaryStatUnit != nil && other.SecondaryStatUnit == nil {
		return false
	}

	if resource.SecondaryStatUnit != nil {
		if *resource.SecondaryStatUnit != *other.SecondaryStatUnit {
			return false
		}
	}

	if len(resource.Arcs) != len(other.Arcs) {
		return false
	}

	for i1 := range resource.Arcs {
		if !resource.Arcs[i1].Equals(other.Arcs[i1]) {
			return false
		}
	}

	return true
}

type EdgeOptions struct {
	// Unit for the main stat to override what ever is set in the data frame.
	MainStatUnit *string `json:"mainStatUnit,omitempty"`
	// Unit for the secondary stat to override what ever is set in the data frame.
	SecondaryStatUnit *string `json:"secondaryStatUnit,omitempty"`
}

func (resource EdgeOptions) Equals(other EdgeOptions) bool {
	if resource.MainStatUnit == nil && other.MainStatUnit != nil || resource.MainStatUnit != nil && other.MainStatUnit == nil {
		return false
	}

	if resource.MainStatUnit != nil {
		if *resource.MainStatUnit != *other.MainStatUnit {
			return false
		}
	}
	if resource.SecondaryStatUnit == nil && other.SecondaryStatUnit != nil || resource.SecondaryStatUnit != nil && other.SecondaryStatUnit == nil {
		return false
	}

	if resource.SecondaryStatUnit != nil {
		if *resource.SecondaryStatUnit != *other.SecondaryStatUnit {
			return false
		}
	}

	return true
}

type Options struct {
	Nodes *NodeOptions `json:"nodes,omitempty"`
	Edges *EdgeOptions `json:"edges,omitempty"`
}

func (resource Options) Equals(other Options) bool {
	if resource.Nodes == nil && other.Nodes != nil || resource.Nodes != nil && other.Nodes == nil {
		return false
	}

	if resource.Nodes != nil {
		if !resource.Nodes.Equals(*other.Nodes) {
			return false
		}
	}
	if resource.Edges == nil && other.Edges != nil || resource.Edges != nil && other.Edges == nil {
		return false
	}

	if resource.Edges != nil {
		if !resource.Edges.Equals(*other.Edges) {
			return false
		}
	}

	return true
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "nodegraph",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := json.Unmarshal(raw, options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
