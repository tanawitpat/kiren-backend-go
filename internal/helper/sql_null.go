package helper

import "database/sql"

func ConvertNullStringToString(input sql.NullString) string {
	var output string
	if input.Valid {
		output = input.String
	}
	return output
}

func ConvertNullBoolToBool(input sql.NullBool) bool {
	var output bool
	if input.Valid {
		output = input.Bool
	}
	return output
}

func ConvertNullFloat64ToFloat64(input sql.NullFloat64) float64 {
	var output float64
	if input.Valid {
		output = input.Float64
	}
	return output
}
