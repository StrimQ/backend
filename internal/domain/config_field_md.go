package domain

type CapturedTablesType map[string]map[string][]string

type ConfigFieldType interface {
	int64 | float64 | string | bool | CapturedTablesType
}
type ConfigFieldTypeRepr string

const (
	ConfigFieldTypeRepr_Int64              ConfigFieldTypeRepr = "int64"
	ConfigFieldTypeRepr_Float64            ConfigFieldTypeRepr = "float64"
	ConfigFieldTypeRepr_String             ConfigFieldTypeRepr = "string"
	ConfigFieldTypeRepr_Bool               ConfigFieldTypeRepr = "bool"
	ConfigFieldTypeRepr_CapturedTablesType ConfigFieldTypeRepr = "map[string]map[string][]string"
)

// UIConfigFieldMd holds metadata for user's configuration fields.
type UIConfigFieldMd[T ConfigFieldType] struct {
	Name         string                     // Field name
	Group        string                     // Logical group (e.g., "heartbeat")
	Description  string                     // Human-readable description
	Type         ConfigFieldTypeRepr        // Type of the field
	Default      T                          // Default value if unset
	Required     bool                       // Whether the field is required
	Value        T                          // Actual value
	ValidateFunc *UIConfigFieldValidateFunc // Function to validate the field value
}

// UIConfigFieldValidateFunc is a function type for validating field values.
type UIConfigFieldValidateFunc func(map[string]any) error
