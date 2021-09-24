// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AverageSessionLength struct {
	Length float64 `json:"length"`
}

type BillingDetails struct {
	Plan               *Plan `json:"plan"`
	Meter              int64 `json:"meter"`
	SessionsOutOfQuota int64 `json:"sessionsOutOfQuota"`
}

type DateRangeInput struct {
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

type ErrorMetadata struct {
	ErrorID         int        `json:"error_id"`
	ErrorSecureID   string     `json:"error_secure_id"`
	SessionID       int        `json:"session_id"`
	SessionSecureID string     `json:"session_secure_id"`
	Environment     *string    `json:"environment"`
	Timestamp       *time.Time `json:"timestamp"`
	Os              *string    `json:"os"`
	Browser         *string    `json:"browser"`
	VisitedURL      *string    `json:"visited_url"`
	Fingerprint     string     `json:"fingerprint"`
	Identifier      *string    `json:"identifier"`
}

type ErrorSearchParamsInput struct {
	DateRange  *DateRangeInput `json:"date_range"`
	Os         *string         `json:"os"`
	Browser    *string         `json:"browser"`
	VisitedURL *string         `json:"visited_url"`
	State      *ErrorState     `json:"state"`
	Event      *string         `json:"event"`
}

type ErrorTrace struct {
	FileName     *string `json:"fileName"`
	LineNumber   *int    `json:"lineNumber"`
	FunctionName *string `json:"functionName"`
	ColumnNumber *int    `json:"columnNumber"`
	Error        *string `json:"error"`
}

type LengthRangeInput struct {
	Min *float64 `json:"min"`
	Max *float64 `json:"max"`
}

type NewUsersCount struct {
	Count int64 `json:"count"`
}

type Plan struct {
	Type  PlanType `json:"type"`
	Quota int      `json:"quota"`
}

type ReferrerTablePayload struct {
	Host    string  `json:"host"`
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
}

type SanitizedAdmin struct {
	ID       int     `json:"id"`
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	PhotoURL *string `json:"photo_url"`
}

type SanitizedAdminInput struct {
	ID    int     `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
}

type SanitizedSlackChannel struct {
	WebhookChannel   *string `json:"webhook_channel"`
	WebhookChannelID *string `json:"webhook_channel_id"`
}

type SanitizedSlackChannelInput struct {
	WebhookChannelName *string `json:"webhook_channel_name"`
	WebhookChannelID   *string `json:"webhook_channel_id"`
}

type SearchParamsInput struct {
	UserProperties          []*UserPropertyInput `json:"user_properties"`
	ExcludedProperties      []*UserPropertyInput `json:"excluded_properties"`
	TrackProperties         []*UserPropertyInput `json:"track_properties"`
	ExcludedTrackProperties []*UserPropertyInput `json:"excluded_track_properties"`
	DateRange               *DateRangeInput      `json:"date_range"`
	LengthRange             *LengthRangeInput    `json:"length_range"`
	Os                      *string              `json:"os"`
	Browser                 *string              `json:"browser"`
	DeviceID                *string              `json:"device_id"`
	VisitedURL              *string              `json:"visited_url"`
	Referrer                *string              `json:"referrer"`
	Identified              *bool                `json:"identified"`
	HideViewed              *bool                `json:"hide_viewed"`
	FirstTime               *bool                `json:"first_time"`
}

type TopUsersPayload struct {
	ID                   int     `json:"id"`
	Identifier           string  `json:"identifier"`
	TotalActiveTime      int     `json:"total_active_time"`
	ActiveTimePercentage float64 `json:"active_time_percentage"`
}

type TrackPropertyInput struct {
	ID    *int   `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type User struct {
	ID int `json:"id"`
}

type UserFingerprintCount struct {
	Count int64 `json:"count"`
}

type UserPropertyInput struct {
	ID    *int   `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ErrorState string

const (
	ErrorStateOpen     ErrorState = "OPEN"
	ErrorStateResolved ErrorState = "RESOLVED"
	ErrorStateIgnored  ErrorState = "IGNORED"
)

var AllErrorState = []ErrorState{
	ErrorStateOpen,
	ErrorStateResolved,
	ErrorStateIgnored,
}

func (e ErrorState) IsValid() bool {
	switch e {
	case ErrorStateOpen, ErrorStateResolved, ErrorStateIgnored:
		return true
	}
	return false
}

func (e ErrorState) String() string {
	return string(e)
}

func (e *ErrorState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ErrorState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ErrorState", str)
	}
	return nil
}

func (e ErrorState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PlanType string

const (
	PlanTypeFree       PlanType = "Free"
	PlanTypeBasic      PlanType = "Basic"
	PlanTypeStartup    PlanType = "Startup"
	PlanTypeEnterprise PlanType = "Enterprise"
)

var AllPlanType = []PlanType{
	PlanTypeFree,
	PlanTypeBasic,
	PlanTypeStartup,
	PlanTypeEnterprise,
}

func (e PlanType) IsValid() bool {
	switch e {
	case PlanTypeFree, PlanTypeBasic, PlanTypeStartup, PlanTypeEnterprise:
		return true
	}
	return false
}

func (e PlanType) String() string {
	return string(e)
}

func (e *PlanType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PlanType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PlanType", str)
	}
	return nil
}

func (e PlanType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SessionCommentType string

const (
	SessionCommentTypeAdmin    SessionCommentType = "Admin"
	SessionCommentTypeFeedback SessionCommentType = "FEEDBACK"
)

var AllSessionCommentType = []SessionCommentType{
	SessionCommentTypeAdmin,
	SessionCommentTypeFeedback,
}

func (e SessionCommentType) IsValid() bool {
	switch e {
	case SessionCommentTypeAdmin, SessionCommentTypeFeedback:
		return true
	}
	return false
}

func (e SessionCommentType) String() string {
	return string(e)
}

func (e *SessionCommentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SessionCommentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SessionCommentType", str)
	}
	return nil
}

func (e SessionCommentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SessionLifecycle string

const (
	SessionLifecycleAll       SessionLifecycle = "All"
	SessionLifecycleLive      SessionLifecycle = "Live"
	SessionLifecycleCompleted SessionLifecycle = "Completed"
)

var AllSessionLifecycle = []SessionLifecycle{
	SessionLifecycleAll,
	SessionLifecycleLive,
	SessionLifecycleCompleted,
}

func (e SessionLifecycle) IsValid() bool {
	switch e {
	case SessionLifecycleAll, SessionLifecycleLive, SessionLifecycleCompleted:
		return true
	}
	return false
}

func (e SessionLifecycle) String() string {
	return string(e)
}

func (e *SessionLifecycle) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SessionLifecycle(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SessionLifecycle", str)
	}
	return nil
}

func (e SessionLifecycle) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
