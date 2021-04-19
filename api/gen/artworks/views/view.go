// Code generated by goa v3.3.1, DO NOT EDIT.
//
// artworks views
//
// Command:
// $ goa gen github.com/pastelnetwork/walletnode/api/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// RegisterResult is the viewed result type that is projected based on a view.
type RegisterResult struct {
	// Type to project
	Projected *RegisterResultView
	// View to render
	View string
}

// Job is the viewed result type that is projected based on a view.
type Job struct {
	// Type to project
	Projected *JobView
	// View to render
	View string
}

// Image is the viewed result type that is projected based on a view.
type Image struct {
	// Type to project
	Projected *ImageView
	// View to render
	View string
}

// RegisterResultView is a type that runs validations on a projected type.
type RegisterResultView struct {
	// Job ID of the registration process
	JobID *int
}

// JobView is a type that runs validations on a projected type.
type JobView struct {
	// JOb ID of the registration process
	ID *int
	// Status of the registration process
	Status *string
	// txid
	Txid *string
}

// ImageView is a type that runs validations on a projected type.
type ImageView struct {
	// Uploaded image ID
	ImageID *string
	// Image expiration
	ExpiresIn *string
}

var (
	// RegisterResultMap is a map of attribute names in result type RegisterResult
	// indexed by view name.
	RegisterResultMap = map[string][]string{
		"default": []string{
			"job_id",
		},
	}
	// JobMap is a map of attribute names in result type Job indexed by view name.
	JobMap = map[string][]string{
		"default": []string{
			"id",
			"status",
			"txid",
		},
	}
	// ImageMap is a map of attribute names in result type Image indexed by view
	// name.
	ImageMap = map[string][]string{
		"default": []string{
			"image_id",
			"expires_in",
		},
	}
)

// ValidateRegisterResult runs the validations defined on the viewed result
// type RegisterResult.
func ValidateRegisterResult(result *RegisterResult) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateRegisterResultView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateJob runs the validations defined on the viewed result type Job.
func ValidateJob(result *Job) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateJobView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateImage runs the validations defined on the viewed result type Image.
func ValidateImage(result *Image) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateImageView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateRegisterResultView runs the validations defined on
// RegisterResultView using the "default" view.
func ValidateRegisterResultView(result *RegisterResultView) (err error) {
	if result.JobID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("job_id", "result"))
	}
	return
}

// ValidateJobView runs the validations defined on JobView using the "default"
// view.
func ValidateJobView(result *JobView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.Txid != nil {
		if utf8.RuneCountInString(*result.Txid) < 64 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.txid", *result.Txid, utf8.RuneCountInString(*result.Txid), 64, true))
		}
	}
	if result.Txid != nil {
		if utf8.RuneCountInString(*result.Txid) > 64 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.txid", *result.Txid, utf8.RuneCountInString(*result.Txid), 64, false))
		}
	}
	return
}

// ValidateImageView runs the validations defined on ImageView using the
// "default" view.
func ValidateImageView(result *ImageView) (err error) {
	if result.ImageID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image_id", "result"))
	}
	if result.ExpiresIn == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("expires_in", "result"))
	}
	if result.ImageID != nil {
		if utf8.RuneCountInString(*result.ImageID) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.image_id", *result.ImageID, utf8.RuneCountInString(*result.ImageID), 8, true))
		}
	}
	if result.ImageID != nil {
		if utf8.RuneCountInString(*result.ImageID) > 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.image_id", *result.ImageID, utf8.RuneCountInString(*result.ImageID), 8, false))
		}
	}
	if result.ExpiresIn != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.expires_in", *result.ExpiresIn, goa.FormatDateTime))
	}
	return
}
