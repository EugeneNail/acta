package validation

import "testing"

func TestRequired(t *testing.T) {
	t.Run("returning a violation when the field is missing", func(t *testing.T) {
		message, err := Required()(map[string]any{}, "name")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "is required" {
			t.Fatalf("expected required message, got %q", message)
		}
	})

	t.Run("returning a violation when the field is a zero value", func(t *testing.T) {
		message, err := Required()(map[string]any{"name": ""}, "name")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "is required" {
			t.Fatalf("expected required message, got %q", message)
		}
	})

	t.Run("returning no violation when the field is set", func(t *testing.T) {
		message, err := Required()(map[string]any{"name": "acta"}, "name")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "" {
			t.Fatalf("expected no message, got %q", message)
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("validating numbers", func(t *testing.T) {
		message, err := Min(10)(map[string]any{"value": 9}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "must be at least 10" {
			t.Fatalf("expected minimum violation, got %q", message)
		}
	})

	t.Run("validating strings by length", func(t *testing.T) {
		message, err := Min(4)(map[string]any{"value": "abc"}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "must be at least 4" {
			t.Fatalf("expected minimum violation, got %q", message)
		}
	})

	t.Run("returning an error for unsupported types", func(t *testing.T) {
		_, err := Min(1)(map[string]any{"value": true}, "value")
		if err == nil {
			t.Fatal("expected an error for unsupported type")
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("validating slices by length", func(t *testing.T) {
		message, err := Max(2)(map[string]any{"value": []int{1, 2, 3}}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "must be at most 2" {
			t.Fatalf("expected maximum violation, got %q", message)
		}
	})

	t.Run("validating maps by length", func(t *testing.T) {
		message, err := Max(1)(map[string]any{"value": map[string]int{"a": 1, "b": 2}}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "must be at most 1" {
			t.Fatalf("expected maximum violation, got %q", message)
		}
	})
}

func TestGte(t *testing.T) {
	t.Run("validating floats", func(t *testing.T) {
		message, err := Gte(10)(map[string]any{"value": 9.0}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "must be greater than or equal to 10" {
			t.Fatalf("expected gte violation, got %q", message)
		}
	})

	t.Run("allowing equal arrays by length", func(t *testing.T) {
		message, err := Gte(3)(map[string]any{"value": [3]int{1, 2, 3}}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "" {
			t.Fatalf("expected no message, got %q", message)
		}
	})
}

func TestLte(t *testing.T) {
	t.Run("validating slices by length", func(t *testing.T) {
		message, err := Lte(2)(map[string]any{"value": []int{1, 2, 3}}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "must be less than or equal to 2" {
			t.Fatalf("expected lte violation, got %q", message)
		}
	})

	t.Run("allowing equal numbers", func(t *testing.T) {
		message, err := Lte(10)(map[string]any{"value": 10}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "" {
			t.Fatalf("expected no message, got %q", message)
		}
	})
}

func TestRegex(t *testing.T) {
	t.Run("validating letters pattern", func(t *testing.T) {
		message, err := Regex(LettersPattern)(map[string]any{"value": "ПриветHello"}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "" {
			t.Fatalf("expected no message, got %q", message)
		}
	})

	t.Run("returning a violation for unmatched text", func(t *testing.T) {
		message, err := Regex(AlphaNumericPattern)(map[string]any{"value": "abc-123"}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "has invalid format" {
			t.Fatalf("expected regex violation, got %q", message)
		}
	})

	t.Run("validating slug pattern", func(t *testing.T) {
		message, err := Regex(SlugPattern)(map[string]any{"value": "release-1.2"}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "" {
			t.Fatalf("expected no message, got %q", message)
		}
	})

	t.Run("validating slug with spaces pattern", func(t *testing.T) {
		message, err := Regex(SlugWithSpacesPattern)(map[string]any{"value": "release 1.2 - beta"}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "" {
			t.Fatalf("expected no message, got %q", message)
		}
	})

	t.Run("validating sentence pattern", func(t *testing.T) {
		message, err := Regex(SentencePattern)(map[string]any{"value": "Привет, world! 123..."}, "value")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if message != "" {
			t.Fatalf("expected no message, got %q", message)
		}
	})

	t.Run("returning an error for unsupported types", func(t *testing.T) {
		_, err := Regex(LettersPattern)(map[string]any{"value": 10}, "value")
		if err == nil {
			t.Fatal("expected an error for unsupported type")
		}
	})
}
