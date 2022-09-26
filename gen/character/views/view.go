// Code generated by goa v3.8.5, DO NOT EDIT.
//
// character views
//
// Command:
// $ goa gen characters/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// StoredCharacterCollection is the viewed result type that is projected based
// on a view.
type StoredCharacterCollection struct {
	// Type to project
	Projected StoredCharacterCollectionView
	// View to render
	View string
}

// StoredCharacter is the viewed result type that is projected based on a view.
type StoredCharacter struct {
	// Type to project
	Projected *StoredCharacterView
	// View to render
	View string
}

// StoredCharacterCollectionView is a type that runs validations on a projected
// type.
type StoredCharacterCollectionView []*StoredCharacterView

// StoredCharacterView is a type that runs validations on a projected type.
type StoredCharacterView struct {
	// ID is the unique id of the character.
	ID *string
	// Name
	Name *string
	// Description
	Description *string
	// Health
	Health *float64
	// Experience
	Experience *float64
}

var (
	// StoredCharacterCollectionMap is a map indexing the attribute names of
	// StoredCharacterCollection by view name.
	StoredCharacterCollectionMap = map[string][]string{
		"default": {
			"id",
			"name",
			"description",
			"health",
			"experience",
		},
		"tiny": {
			"id",
			"name",
			"health",
			"experience",
		},
	}
	// StoredCharacterMap is a map indexing the attribute names of StoredCharacter
	// by view name.
	StoredCharacterMap = map[string][]string{
		"default": {
			"id",
			"name",
			"description",
			"health",
			"experience",
		},
		"tiny": {
			"id",
			"name",
			"health",
			"experience",
		},
	}
)

// ValidateStoredCharacterCollection runs the validations defined on the viewed
// result type StoredCharacterCollection.
func ValidateStoredCharacterCollection(result StoredCharacterCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredCharacterCollectionView(result.Projected)
	case "tiny":
		err = ValidateStoredCharacterCollectionViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateStoredCharacter runs the validations defined on the viewed result
// type StoredCharacter.
func ValidateStoredCharacter(result *StoredCharacter) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredCharacterView(result.Projected)
	case "tiny":
		err = ValidateStoredCharacterViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateStoredCharacterCollectionView runs the validations defined on
// StoredCharacterCollectionView using the "default" view.
func ValidateStoredCharacterCollectionView(result StoredCharacterCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredCharacterView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredCharacterCollectionViewTiny runs the validations defined on
// StoredCharacterCollectionView using the "tiny" view.
func ValidateStoredCharacterCollectionViewTiny(result StoredCharacterCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredCharacterViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredCharacterView runs the validations defined on
// StoredCharacterView using the "default" view.
func ValidateStoredCharacterView(result *StoredCharacterView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "result"))
	}
	if result.Experience == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("experience", "result"))
	}
	return
}

// ValidateStoredCharacterViewTiny runs the validations defined on
// StoredCharacterView using the "tiny" view.
func ValidateStoredCharacterViewTiny(result *StoredCharacterView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "result"))
	}
	if result.Experience == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("experience", "result"))
	}
	return
}