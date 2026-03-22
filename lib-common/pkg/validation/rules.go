package validation

import (
	"fmt"
	"reflect"
	"regexp"
)

// LettersPattern allows only Latin and Cyrillic letters.
const LettersPattern = `^[\p{Latin}\p{Cyrillic}]+$`

// AlphaNumericPattern allows only Latin and Cyrillic letters with integer digits.
const AlphaNumericPattern = `^[\p{Latin}\p{Cyrillic}\d]+$`

// SlugPattern allows slug-like text with Latin and Cyrillic letters, integer digits, hyphens, and dots.
const SlugPattern = `^[\p{Latin}\p{Cyrillic}\d.-]+$`

// SlugWithSpacesPattern allows slug-like text with Latin and Cyrillic letters, integer digits, hyphens, dots, and spaces.
const SlugWithSpacesPattern = `^[\p{Latin}\p{Cyrillic}\d. -]+$`

// SentencePattern allows sentence-like text with letters, digits, punctuation, and spaces.
const SentencePattern = `^[\p{L}\p{N}\p{P}\p{Zs}]+$`

// Rule validates a field within the full data payload.
type Rule func(data map[string]any, field string) (string, error)

// Required validates that the field exists, is not nil, and differs from its type zero value.
func Required() Rule {
	return func(data map[string]any, field string) (string, error) {
		value, exists := data[field]
		if !exists || isNil(value) || isZeroValue(value) {
			return "is required", nil
		}

		return "", nil
	}
}

// Min validates that the field value or size is at least the provided minimum.
func Min(minimum int) Rule {
	return withSkipIfNotPresent(func(value any, field string) (string, error) {
		current, err := extractMeasurableValue(value)
		if err != nil {
			return "", fmt.Errorf("measuring minimum for field %q: %w", field, err)
		}

		if current < float64(minimum) {
			return fmt.Sprintf("must be at least %d", minimum), nil
		}

		return "", nil
	})
}

// Max validates that the field value or size does not exceed the provided maximum.
func Max(maximum int) Rule {
	return withSkipIfNotPresent(func(value any, field string) (string, error) {
		current, err := extractMeasurableValue(value)
		if err != nil {
			return "", fmt.Errorf("measuring maximum for field %q: %w", field, err)
		}

		if current > float64(maximum) {
			return fmt.Sprintf("must be at most %d", maximum), nil
		}

		return "", nil
	})
}

// Gte validates that the field value or size is greater than or equal to the provided threshold.
func Gte(threshold int) Rule {
	return withSkipIfNotPresent(func(value any, field string) (string, error) {
		current, err := extractMeasurableValue(value)
		if err != nil {
			return "", fmt.Errorf("measuring threshold for field %q: %w", field, err)
		}

		if current < float64(threshold) {
			return fmt.Sprintf("must be greater than or equal to %d", threshold), nil
		}

		return "", nil
	})
}

// Lte validates that the field value or size is less than or equal to the provided threshold.
func Lte(threshold int) Rule {
	return withSkipIfNotPresent(func(value any, field string) (string, error) {
		current, err := extractMeasurableValue(value)
		if err != nil {
			return "", fmt.Errorf("measuring threshold for field %q: %w", field, err)
		}

		if current > float64(threshold) {
			return fmt.Sprintf("must be less than or equal to %d", threshold), nil
		}

		return "", nil
	})
}

// Regex validates that the field string matches the provided regular expression mask.
func Regex(mask string) Rule {
	pattern, err := regexp.Compile(mask)

	return withSkipIfNotPresent(func(value any, field string) (string, error) {
		if err != nil {
			return "", fmt.Errorf("compiling regex mask %q: %w", mask, err)
		}

		text, ok := value.(string)
		if !ok {
			return "", fmt.Errorf("matching regex for field %q: type %T is not supported", field, value)
		}

		if !pattern.MatchString(text) {
			return "has invalid format", nil
		}

		return "", nil
	})
}

// withSkipIfNotPresent executes the callback only when the field exists in the data payload.
func withSkipIfNotPresent(rule func(value any, field string) (string, error)) Rule {
	return func(data map[string]any, field string) (string, error) {
		value, exists := data[field]
		if !exists {
			return "", nil
		}

		return rule(value, field)
	}
}

// extractMeasurableValue converts supported scalar values to numbers and collection-like values to their length.
func extractMeasurableValue(value any) (float64, error) {
	switch typedValue := value.(type) {
	case int:
		return float64(typedValue), nil
	case int8:
		return float64(typedValue), nil
	case int16:
		return float64(typedValue), nil
	case int32:
		return float64(typedValue), nil
	case int64:
		return float64(typedValue), nil
	case uint:
		return float64(typedValue), nil
	case uint8:
		return float64(typedValue), nil
	case uint16:
		return float64(typedValue), nil
	case uint32:
		return float64(typedValue), nil
	case uint64:
		return float64(typedValue), nil
	case uintptr:
		return float64(typedValue), nil
	case float32:
		return float64(typedValue), nil
	case float64:
		return typedValue, nil
	case string:
		return float64(len(typedValue)), nil
	}

	reflectedValue := reflect.ValueOf(value)
	if reflectedValue.Kind() == reflect.Array || reflectedValue.Kind() == reflect.Slice || reflectedValue.Kind() == reflect.Map {
		return float64(reflectedValue.Len()), nil
	}

	return 0, fmt.Errorf("type %T is not supported", value)
}

// isNil reports whether the value is nil or wraps a nil reference-like Go value.
func isNil(value any) bool {
	if value == nil {
		return true
	}

	reflectedValue := reflect.ValueOf(value)
	switch reflectedValue.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return reflectedValue.IsNil()
	default:
		return false
	}
}

// isZeroValue reports whether the value is the zero value for its type.
func isZeroValue(value any) bool {
	if value == nil {
		return true
	}

	switch typedValue := value.(type) {
	case bool:
		return !typedValue
	case string:
		return typedValue == ""
	case int:
		return typedValue == 0
	case int8:
		return typedValue == 0
	case int16:
		return typedValue == 0
	case int32:
		return typedValue == 0
	case int64:
		return typedValue == 0
	case uint:
		return typedValue == 0
	case uint8:
		return typedValue == 0
	case uint16:
		return typedValue == 0
	case uint32:
		return typedValue == 0
	case uint64:
		return typedValue == 0
	case uintptr:
		return typedValue == 0
	case float32:
		return typedValue == 0
	case float64:
		return typedValue == 0
	}

	return reflect.ValueOf(value).IsZero()
}
