// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: signals.proto

package contractv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetStrategiesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetStrategiesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStrategiesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStrategiesRequestMultiError, or nil if none found.
func (m *GetStrategiesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStrategiesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.StrategyId != nil {
		// no validation rules for StrategyId
	}

	if len(errors) > 0 {
		return GetStrategiesRequestMultiError(errors)
	}

	return nil
}

// GetStrategiesRequestMultiError is an error wrapping multiple validation
// errors returned by GetStrategiesRequest.ValidateAll() if the designated
// constraints aren't met.
type GetStrategiesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStrategiesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStrategiesRequestMultiError) AllErrors() []error { return m }

// GetStrategiesRequestValidationError is the validation error returned by
// GetStrategiesRequest.Validate if the designated constraints aren't met.
type GetStrategiesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStrategiesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStrategiesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStrategiesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStrategiesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStrategiesRequestValidationError) ErrorName() string {
	return "GetStrategiesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetStrategiesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStrategiesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStrategiesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStrategiesRequestValidationError{}

// Validate checks the field values on GetStrategiesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetStrategiesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStrategiesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStrategiesResponseMultiError, or nil if none found.
func (m *GetStrategiesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStrategiesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetStrategies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetStrategiesResponseValidationError{
						field:  fmt.Sprintf("Strategies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetStrategiesResponseValidationError{
						field:  fmt.Sprintf("Strategies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetStrategiesResponseValidationError{
					field:  fmt.Sprintf("Strategies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetStrategiesResponseMultiError(errors)
	}

	return nil
}

// GetStrategiesResponseMultiError is an error wrapping multiple validation
// errors returned by GetStrategiesResponse.ValidateAll() if the designated
// constraints aren't met.
type GetStrategiesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStrategiesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStrategiesResponseMultiError) AllErrors() []error { return m }

// GetStrategiesResponseValidationError is the validation error returned by
// GetStrategiesResponse.Validate if the designated constraints aren't met.
type GetStrategiesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStrategiesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStrategiesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStrategiesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStrategiesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStrategiesResponseValidationError) ErrorName() string {
	return "GetStrategiesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetStrategiesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStrategiesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStrategiesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStrategiesResponseValidationError{}

// Validate checks the field values on Strategy with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Strategy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Strategy with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StrategyMultiError, or nil
// if none found.
func (m *Strategy) ValidateAll() error {
	return m.validate(true)
}

func (m *Strategy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StrategyId

	// no validation rules for StrategyName

	// no validation rules for StrategyType

	// no validation rules for ActiveSignals

	// no validation rules for TotalSignals

	// no validation rules for TimeInPosition

	if all {
		switch v := interface{}(m.GetAverageSignalYield()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StrategyValidationError{
					field:  "AverageSignalYield",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StrategyValidationError{
					field:  "AverageSignalYield",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAverageSignalYield()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StrategyValidationError{
				field:  "AverageSignalYield",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAverageSignalYieldYear()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StrategyValidationError{
					field:  "AverageSignalYieldYear",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StrategyValidationError{
					field:  "AverageSignalYieldYear",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAverageSignalYieldYear()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StrategyValidationError{
				field:  "AverageSignalYieldYear",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetYield()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StrategyValidationError{
					field:  "Yield",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StrategyValidationError{
					field:  "Yield",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetYield()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StrategyValidationError{
				field:  "Yield",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetYieldYear()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StrategyValidationError{
					field:  "YieldYear",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StrategyValidationError{
					field:  "YieldYear",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetYieldYear()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StrategyValidationError{
				field:  "YieldYear",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.StrategyDescription != nil {
		// no validation rules for StrategyDescription
	}

	if m.StrategyUrl != nil {
		// no validation rules for StrategyUrl
	}

	if len(errors) > 0 {
		return StrategyMultiError(errors)
	}

	return nil
}

// StrategyMultiError is an error wrapping multiple validation errors returned
// by Strategy.ValidateAll() if the designated constraints aren't met.
type StrategyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyMultiError) AllErrors() []error { return m }

// StrategyValidationError is the validation error returned by
// Strategy.Validate if the designated constraints aren't met.
type StrategyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyValidationError) ErrorName() string { return "StrategyValidationError" }

// Error satisfies the builtin error interface
func (e StrategyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyValidationError{}

// Validate checks the field values on GetSignalsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetSignalsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSignalsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSignalsRequestMultiError, or nil if none found.
func (m *GetSignalsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSignalsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.SignalId != nil {
		// no validation rules for SignalId
	}

	if m.StrategyId != nil {
		// no validation rules for StrategyId
	}

	if m.StrategyType != nil {
		// no validation rules for StrategyType
	}

	if m.InstrumentUid != nil {
		// no validation rules for InstrumentUid
	}

	if m.From != nil {

		if all {
			switch v := interface{}(m.GetFrom()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetSignalsRequestValidationError{
						field:  "From",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetSignalsRequestValidationError{
						field:  "From",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetFrom()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetSignalsRequestValidationError{
					field:  "From",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.To != nil {

		if all {
			switch v := interface{}(m.GetTo()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetSignalsRequestValidationError{
						field:  "To",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetSignalsRequestValidationError{
						field:  "To",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetSignalsRequestValidationError{
					field:  "To",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Direction != nil {
		// no validation rules for Direction
	}

	if m.Active != nil {
		// no validation rules for Active
	}

	if m.Paging != nil {

		if all {
			switch v := interface{}(m.GetPaging()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetSignalsRequestValidationError{
						field:  "Paging",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetSignalsRequestValidationError{
						field:  "Paging",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPaging()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetSignalsRequestValidationError{
					field:  "Paging",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetSignalsRequestMultiError(errors)
	}

	return nil
}

// GetSignalsRequestMultiError is an error wrapping multiple validation errors
// returned by GetSignalsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetSignalsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSignalsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSignalsRequestMultiError) AllErrors() []error { return m }

// GetSignalsRequestValidationError is the validation error returned by
// GetSignalsRequest.Validate if the designated constraints aren't met.
type GetSignalsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSignalsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSignalsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSignalsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSignalsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSignalsRequestValidationError) ErrorName() string {
	return "GetSignalsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetSignalsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSignalsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSignalsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSignalsRequestValidationError{}

// Validate checks the field values on GetSignalsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSignalsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSignalsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSignalsResponseMultiError, or nil if none found.
func (m *GetSignalsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSignalsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetSignals() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetSignalsResponseValidationError{
						field:  fmt.Sprintf("Signals[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetSignalsResponseValidationError{
						field:  fmt.Sprintf("Signals[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetSignalsResponseValidationError{
					field:  fmt.Sprintf("Signals[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetPaging()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetSignalsResponseValidationError{
					field:  "Paging",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetSignalsResponseValidationError{
					field:  "Paging",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaging()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetSignalsResponseValidationError{
				field:  "Paging",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetSignalsResponseMultiError(errors)
	}

	return nil
}

// GetSignalsResponseMultiError is an error wrapping multiple validation errors
// returned by GetSignalsResponse.ValidateAll() if the designated constraints
// aren't met.
type GetSignalsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSignalsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSignalsResponseMultiError) AllErrors() []error { return m }

// GetSignalsResponseValidationError is the validation error returned by
// GetSignalsResponse.Validate if the designated constraints aren't met.
type GetSignalsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSignalsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSignalsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSignalsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSignalsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSignalsResponseValidationError) ErrorName() string {
	return "GetSignalsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetSignalsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSignalsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSignalsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSignalsResponseValidationError{}

// Validate checks the field values on Signal with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Signal) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Signal with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SignalMultiError, or nil if none found.
func (m *Signal) ValidateAll() error {
	return m.validate(true)
}

func (m *Signal) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SignalId

	// no validation rules for StrategyId

	// no validation rules for StrategyName

	// no validation rules for InstrumentUid

	if all {
		switch v := interface{}(m.GetCreateDt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SignalValidationError{
					field:  "CreateDt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SignalValidationError{
					field:  "CreateDt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateDt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignalValidationError{
				field:  "CreateDt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Direction

	if all {
		switch v := interface{}(m.GetInitialPrice()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SignalValidationError{
					field:  "InitialPrice",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SignalValidationError{
					field:  "InitialPrice",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInitialPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignalValidationError{
				field:  "InitialPrice",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetTargetPrice()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SignalValidationError{
					field:  "TargetPrice",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SignalValidationError{
					field:  "TargetPrice",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTargetPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignalValidationError{
				field:  "TargetPrice",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEndDt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SignalValidationError{
					field:  "EndDt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SignalValidationError{
					field:  "EndDt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndDt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignalValidationError{
				field:  "EndDt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.Info != nil {
		// no validation rules for Info
	}

	if m.Probability != nil {
		// no validation rules for Probability
	}

	if m.Stoploss != nil {

		if all {
			switch v := interface{}(m.GetStoploss()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SignalValidationError{
						field:  "Stoploss",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SignalValidationError{
						field:  "Stoploss",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStoploss()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SignalValidationError{
					field:  "Stoploss",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.ClosePrice != nil {

		if all {
			switch v := interface{}(m.GetClosePrice()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SignalValidationError{
						field:  "ClosePrice",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SignalValidationError{
						field:  "ClosePrice",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetClosePrice()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SignalValidationError{
					field:  "ClosePrice",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.CloseDt != nil {

		if all {
			switch v := interface{}(m.GetCloseDt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SignalValidationError{
						field:  "CloseDt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SignalValidationError{
						field:  "CloseDt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCloseDt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SignalValidationError{
					field:  "CloseDt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SignalMultiError(errors)
	}

	return nil
}

// SignalMultiError is an error wrapping multiple validation errors returned by
// Signal.ValidateAll() if the designated constraints aren't met.
type SignalMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignalMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignalMultiError) AllErrors() []error { return m }

// SignalValidationError is the validation error returned by Signal.Validate if
// the designated constraints aren't met.
type SignalValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignalValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignalValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignalValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignalValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignalValidationError) ErrorName() string { return "SignalValidationError" }

// Error satisfies the builtin error interface
func (e SignalValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignal.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignalValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignalValidationError{}
