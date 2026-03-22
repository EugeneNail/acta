package validation

import "fmt"

// Validator applies field rulesets to input data.
type Validator struct {
	data     map[string]any
	rulesets map[string][]Rule
}

// NewValidator constructs a validator for the provided data and rulesets.
func NewValidator(data map[string]any, rulesets map[string][]Rule) (*Validator, error) {
	for field, rules := range rulesets {
		if len(rules) == 0 {
			return nil, fmt.Errorf("ruleset for field %q is empty", field)
		}
	}

	return &Validator{
		data:     data,
		rulesets: rulesets,
	}, nil
}

// Validate executes rulesets and returns either validation violations or a rule execution error.
func (validator *Validator) Validate() error {
	validationError := NewError()

	for field, rules := range validator.rulesets {
	ruleLoop:
		for i, rule := range rules {
			message, err := rule(validator.data, field)
			if err != nil {
				return fmt.Errorf("applying %dth rule to field %q: %w", i, field, err)
			}

			if len(message) > 0 {
				validationError.AddViolation(field, message)
				break ruleLoop
			}
		}
	}

	if len(validationError.violations) > 0 {
		return validationError
	}

	return nil
}
