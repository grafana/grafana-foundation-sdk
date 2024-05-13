// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"encoding/json"
	"errors"
	"fmt"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

func VariantConfig() cogvariants.DataqueryConfig {
	return cogvariants.DataqueryConfig{
		Identifier: "__expr__",
		DataqueryUnmarshaler: func(raw []byte) (cogvariants.Dataquery, error) {
			dataquery := Expr{}

			if err := json.Unmarshal(raw, &dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
}

func (resource TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) MarshalJSON() ([]byte, error) {
	if resource.TypeMath != nil {
		return json.Marshal(resource.TypeMath)
	}
	if resource.TypeReduce != nil {
		return json.Marshal(resource.TypeReduce)
	}
	if resource.TypeResample != nil {
		return json.Marshal(resource.TypeResample)
	}
	if resource.TypeClassicConditions != nil {
		return json.Marshal(resource.TypeClassicConditions)
	}
	if resource.TypeThreshold != nil {
		return json.Marshal(resource.TypeThreshold)
	}
	if resource.TypeSql != nil {
		return json.Marshal(resource.TypeSql)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}
func (resource *TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "classic_conditions":
		var typeClassicConditions TypeClassicConditions
		if err := json.Unmarshal(raw, &typeClassicConditions); err != nil {
			return err
		}

		resource.TypeClassicConditions = &typeClassicConditions
		return nil
	case "math":
		var typeMath TypeMath
		if err := json.Unmarshal(raw, &typeMath); err != nil {
			return err
		}

		resource.TypeMath = &typeMath
		return nil
	case "reduce":
		var typeReduce TypeReduce
		if err := json.Unmarshal(raw, &typeReduce); err != nil {
			return err
		}

		resource.TypeReduce = &typeReduce
		return nil
	case "resample":
		var typeResample TypeResample
		if err := json.Unmarshal(raw, &typeResample); err != nil {
			return err
		}

		resource.TypeResample = &typeResample
		return nil
	case "sql":
		var typeSql TypeSql
		if err := json.Unmarshal(raw, &typeSql); err != nil {
			return err
		}

		resource.TypeSql = &typeSql
		return nil
	case "threshold":
		var typeThreshold TypeThreshold
		if err := json.Unmarshal(raw, &typeThreshold); err != nil {
			return err
		}

		resource.TypeThreshold = &typeThreshold
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}
