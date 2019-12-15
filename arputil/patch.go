/*
  This go file is the basis for the structure of a patch json
*/

package arputil

type Patch struct {
	Version    Version    						 `json:"version"`
	Children   []Children 						 `json:"children"`
	Identifier string     						 `json:"identifier"`
	ValuesMap  map[string]interface{}  `json:"valuesMap"`
	ModelName  string    							 `json:"modelName"`
}

type Version struct {
	Patch int `json:"patch"`
	Dev   int `json:"dev"`
	Prod  int `json:"prod"`
}

type DeepChildren struct {
	RawSubtype           int           `json:"rawSubtype,omitempty"`
	Type                 int           `json:"type,omitempty"`
	Identifier           string        `json:"identifier"`
	ModelName            string        `json:"modelName"`
	WidgetStyle          int           `json:"widgetStyle,omitempty"`
	GatekeeperName       string        `json:"gatekeeperName,omitempty"`
	EnumOptions          []interface{} `json:"enumOptions,omitempty"`
	PortName             string        `json:"portName,omitempty"`
	IsParameter          bool          `json:"isParameter,omitempty"`
	PreferValueOverTitle bool          `json:"preferValueOverTitle,omitempty"`
	HiddenValue          bool          `json:"hiddenValue,omitempty"`
	Hidden               bool          `json:"hidden,omitempty"`
	HiddenConnection     bool          `json:"hiddenConnection,omitempty"`
	IsVariant            bool          `json:"isVariant,omitempty"`
}

type DefaultValue struct {
	Value int `json:"value"`
	Type  int `json:"type"`
}

type ExtendedChildren struct {
	ModelName    string       		`json:"modelName"`
	Identifier   string       		`json:"identifier"`
	Name         string       		`json:"name"`
	Children     []DeepChildren   `json:"children"`
	PortTag      int          		`json:"portTag"`
	DefaultValue DefaultValue 		`json:"defaultValue"`
}

type ComponentChildren struct {
	Identifier           string     				`json:"identifier"`
	ModelName            string     				`json:"modelName"`
	Children             []ExtendedChildren `json:"children,omitempty"`
	PreferredGenericType int       				  `json:"preferredGenericType,omitempty"`
	InternalGroup        bool       				`json:"internalGroup,omitempty"`
	Hidden               bool      				 	`json:"hidden,omitempty"`
	Position             []int      				`json:"position,omitempty"`
}
type PropertyChildren struct {
	ModelName           string     					`json:"modelName"`
	Identifier          string     					`json:"identifier"`
	Name                string     					`json:"name"`
	Prototype           string     					`json:"prototype"`
	Children            []ComponentChildren `json:"children"`
	DataModelIdentifier string     					`json:"dataModelIdentifier,omitempty"`
	PropertyIdentifier  string     					`json:"propertyIdentifier,omitempty"`
}
type PatchModelChildren struct {
	Children             []PropertyChildren `json:"children,omitempty"`
	Identifier           string     				`json:"identifier"`
	ModelName            string     				`json:"modelName"`
	PreferredGenericType int        				`json:"preferredGenericType,omitempty"`
	InternalGroup        bool       				`json:"internalGroup,omitempty"`
	Hidden               bool       				`json:"hidden,omitempty"`
	Position             []int      				`json:"position,omitempty"`
}
type Children struct {
	ModelName  string     					`json:"modelName"`
	Identifier string     					`json:"identifier"`
	Name       string     					`json:"name"`
	Children   []PatchModelChildren `json:"children"`
}
