// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlschema

import (
	"fmt"
	"io"
	"strconv"
)

type CredentialData interface {
	IsCredentialData()
}

//  Every query that implements pagination returns object that implements Pageable interface.
// To specify page details, query specify two parameters: `first` and `after`.
// `first` specify page size, `after` is a cursor for the next page. When requesting first page, set `after` to empty value.
// For requesting next page, set `after` to `pageInfo.endCursor` returned from previous query.
type Pageable interface {
	IsPageable()
}

type APIDefinition struct {
	ID        string   `json:"id"`
	Spec      *APISpec `json:"spec"`
	TargetURL string   `json:"targetURL"`
	//  group allows you to find the same API but in different version
	Group *string `json:"group"`
	// "If runtime does not exist, an error is returned. If runtime exists but Auth for it is not set, defaultAuth is returned if specified.
	Auth *RuntimeAuth `json:"auth"`
	// Returns authentication details for all runtimes, even for a runtime, where Auth is not yet specified.
	Auths []*RuntimeAuth `json:"auths"`
	// If defaultAuth is specified, it will be used for all Runtimes that does not specify Auth explicitly.
	DefaultAuth *Auth    `json:"defaultAuth"`
	Version     *Version `json:"version"`
}

type APIDefinitionInput struct {
	TargetURL   string        `json:"targetURL"`
	Group       *string       `json:"group"`
	Spec        *APISpecInput `json:"spec"`
	Version     *VersionInput `json:"version"`
	DefaultAuth *AuthInput    `json:"defaultAuth"`
}

type APIDefinitionPage struct {
	Data       []*APIDefinition `json:"data"`
	PageInfo   *PageInfo        `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

func (APIDefinitionPage) IsPageable() {}

type APISpec struct {
	// when fetch request specified, data will be automatically populated
	Data         *string       `json:"data"`
	Format       *SpecFormat   `json:"format"`
	Type         APISpecType   `json:"type"`
	FetchRequest *FetchRequest `json:"fetchRequest"`
}

type APISpecInput struct {
	Data         *string            `json:"data"`
	Type         APISpecType        `json:"type"`
	Format       SpecFormat         `json:"format"`
	FetchRequest *FetchRequestInput `json:"fetchRequest"`
}

type Application struct {
	ID             string                `json:"id"`
	Tenant         Tenant                `json:"tenant"`
	Name           string                `json:"name"`
	Description    *string               `json:"description"`
	Labels         Labels                `json:"labels"`
	Annotations    Annotations           `json:"annotations"`
	Status         *ApplicationStatus    `json:"status"`
	Webhooks       []*ApplicationWebhook `json:"webhooks"`
	HealthCheckURL *string               `json:"healthCheckURL"`
	//  group allows to find different versions of the same API
	Apis *APIDefinitionPage `json:"apis"`
	//  group allows to find different versions of the same event API
	EventAPIs *EventAPIDefinitionPage `json:"eventAPIs"`
	Documents *DocumentPage           `json:"documents"`
}

type ApplicationInput struct {
	Name           string                     `json:"name"`
	Description    *string                    `json:"description"`
	Labels         *Labels                    `json:"labels"`
	Annotations    *Annotations               `json:"annotations"`
	Webhooks       []*ApplicationWebhookInput `json:"webhooks"`
	HealthCheckURL *string                    `json:"healthCheckURL"`
	Apis           []*APIDefinitionInput      `json:"apis"`
	Events         []*EventDefinitionInput    `json:"events"`
	Documents      []*DocumentInput           `json:"documents"`
}

type ApplicationPage struct {
	Data       []*Application `json:"data"`
	PageInfo   *PageInfo      `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

func (ApplicationPage) IsPageable() {}

type ApplicationStatus struct {
	Condition ApplicationStatusCondition `json:"condition"`
	Timestamp Timestamp                  `json:"timestamp"`
}

type ApplicationWebhook struct {
	ID   string                 `json:"id"`
	Type ApplicationWebhookType `json:"type"`
	URL  string                 `json:"url"`
	Auth *Auth                  `json:"auth"`
}

type ApplicationWebhookInput struct {
	Type ApplicationWebhookType `json:"type"`
	URL  string                 `json:"url"`
	Auth *AuthInput             `json:"auth"`
}

type Auth struct {
	Credential            CredentialData         `json:"credential"`
	AdditionalHeaders     *HttpHeaders           `json:"additionalHeaders"`
	AdditionalQueryParams *QueryParams           `json:"additionalQueryParams"`
	RequestAuth           *CredentialRequestAuth `json:"requestAuth"`
}

type AuthInput struct {
	Credential            *CredentialDataInput        `json:"credential"`
	AdditionalHeaders     *HttpHeaders                `json:"additionalHeaders"`
	AdditionalQueryParams *QueryParams                `json:"additionalQueryParams"`
	RequestAuth           *CredentialRequestAuthInput `json:"requestAuth"`
}

type BasicCredentialData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (BasicCredentialData) IsCredentialData() {}

type BasicCredentialDataInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CSRFTokenCredentialRequestAuth struct {
	TokenEndpointURL string `json:"tokenEndpointURL"`
	Auth             *Auth  `json:"auth"`
}

type CSRFTokenCredentialRequestAuthInput struct {
	TokenEndpointURL string     `json:"tokenEndpointURL"`
	Auth             *AuthInput `json:"auth"`
}

type CredentialDataInput struct {
	Basic *BasicCredentialDataInput `json:"basic"`
	Oauth *OAuthCredentialDataInput `json:"oauth"`
}

type CredentialRequestAuth struct {
	Csrf *CSRFTokenCredentialRequestAuth `json:"csrf"`
}

type CredentialRequestAuthInput struct {
	Csrf *CSRFTokenCredentialRequestAuthInput `json:"csrf"`
}

type Document struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Format DocumentFormat `json:"format"`
	// for example Service Class, API etc
	Kind         *string       `json:"kind"`
	Data         *string       `json:"data"`
	FetchRequest *FetchRequest `json:"fetchRequest"`
}

type DocumentInput struct {
	Title        string             `json:"title"`
	DisplayName  string             `json:"displayName"`
	Description  string             `json:"description"`
	Format       DocumentFormat     `json:"format"`
	Kind         *string            `json:"kind"`
	Data         *string            `json:"data"`
	FetchRequest *FetchRequestInput `json:"fetchRequest"`
}

type DocumentPage struct {
	Data       []*Document `json:"data"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (DocumentPage) IsPageable() {}

type EventAPIDefinition struct {
	ID string `json:"id"`
	// group allows you to find the same API but in different version
	Group   *string    `json:"group"`
	Spec    *EventSpec `json:"spec"`
	Version *Version   `json:"version"`
}

type EventAPIDefinitionPage struct {
	Data       []*EventAPIDefinition `json:"data"`
	PageInfo   *PageInfo             `json:"pageInfo"`
	TotalCount int                   `json:"totalCount"`
}

func (EventAPIDefinitionPage) IsPageable() {}

type EventDefinitionInput struct {
	Spec    *EventSpecInput `json:"spec"`
	Group   *string         `json:"group"`
	Version *VersionInput   `json:"version"`
}

type EventSpec struct {
	Data         *string       `json:"data"`
	Type         EventSpecType `json:"type"`
	Format       *SpecFormat   `json:"format"`
	FetchRequest *FetchRequest `json:"fetchRequest"`
}

type EventSpecInput struct {
	Data          *string            `json:"data"`
	EventSpecType EventSpecType      `json:"eventSpecType"`
	FetchRequest  *FetchRequestInput `json:"fetchRequest"`
}

//  Compass performs fetch to validate if request is correct and stores a copy
type FetchRequest struct {
	URL    string              `json:"url"`
	Auth   *Auth               `json:"auth"`
	Mode   FetchMode           `json:"mode"`
	Filter *string             `json:"filter"`
	Status *FetchRequestStatus `json:"status"`
}

type FetchRequestInput struct {
	URL    string     `json:"url"`
	Auth   *AuthInput `json:"auth"`
	Mode   *FetchMode `json:"mode"`
	Filter *string    `json:"filter"`
}

type FetchRequestStatus struct {
	Condition FetchRequestStatusCondition `json:"condition"`
	Timestamp Timestamp                   `json:"timestamp"`
}

type HealthCheck struct {
	Type      HealthCheckType            `json:"type"`
	Condition HealthCheckStatusCondition `json:"condition"`
	Origin    *string                    `json:"origin"`
	Message   *string                    `json:"message"`
	Timestamp Timestamp                  `json:"timestamp"`
}

type HealthCheckPage struct {
	Data       []*HealthCheck `json:"data"`
	PageInfo   *PageInfo      `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

func (HealthCheckPage) IsPageable() {}

type LabelFilter struct {
	Label    string          `json:"label"`
	Values   []string        `json:"values"`
	Operator *FilterOperator `json:"operator"`
}

type OAuthCredentialData struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	URL          string `json:"url"`
}

func (OAuthCredentialData) IsCredentialData() {}

type OAuthCredentialDataInput struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	URL          string `json:"url"`
}

type PageInfo struct {
	StartCursor string `json:"startCursor"`
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

type Runtime struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description *string        `json:"description"`
	Tenant      Tenant         `json:"tenant"`
	Labels      Labels         `json:"labels"`
	Annotations Annotations    `json:"annotations"`
	Status      *RuntimeStatus `json:"status"`
	// directive for checking auth
	AgentAuth *Auth `json:"agentAuth"`
}

type RuntimeAuth struct {
	RuntimeID string `json:"runtimeID"`
	Auth      *Auth  `json:"auth"`
}

type RuntimeInput struct {
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	Labels      *Labels      `json:"labels"`
	Annotations *Annotations `json:"annotations"`
}

type RuntimePage struct {
	Data       []*Runtime `json:"data"`
	PageInfo   *PageInfo  `json:"pageInfo"`
	TotalCount int        `json:"totalCount"`
}

func (RuntimePage) IsPageable() {}

type RuntimeStatus struct {
	Condition RuntimeStatusCondition `json:"condition"`
	Timestamp Timestamp              `json:"timestamp"`
}

type Version struct {
	// for example 4.6
	Value      string `json:"value"`
	Deprecated *bool  `json:"deprecated"`
	// for example 4.5
	DeprecatedSince *string `json:"deprecatedSince"`
	// if true, will be removed in the next version
	ForRemoval *bool `json:"forRemoval"`
}

type VersionInput struct {
	Value           string  `json:"value"`
	Deprecated      *bool   `json:"deprecated"`
	DeprecatedSince *string `json:"deprecatedSince"`
	ForRemoval      *bool   `json:"forRemoval"`
}

type APISpecType string

const (
	APISpecTypeOdata   APISpecType = "ODATA"
	APISpecTypeOpenAPI APISpecType = "OPEN_API"
)

var AllAPISpecType = []APISpecType{
	APISpecTypeOdata,
	APISpecTypeOpenAPI,
}

func (e APISpecType) IsValid() bool {
	switch e {
	case APISpecTypeOdata, APISpecTypeOpenAPI:
		return true
	}
	return false
}

func (e APISpecType) String() string {
	return string(e)
}

func (e *APISpecType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = APISpecType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid APISpecType", str)
	}
	return nil
}

func (e APISpecType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ApplicationStatusCondition string

const (
	ApplicationStatusConditionInitial ApplicationStatusCondition = "INITIAL"
	ApplicationStatusConditionUnknown ApplicationStatusCondition = "UNKNOWN"
	ApplicationStatusConditionReady   ApplicationStatusCondition = "READY"
	ApplicationStatusConditionFailed  ApplicationStatusCondition = "FAILED"
)

var AllApplicationStatusCondition = []ApplicationStatusCondition{
	ApplicationStatusConditionInitial,
	ApplicationStatusConditionUnknown,
	ApplicationStatusConditionReady,
	ApplicationStatusConditionFailed,
}

func (e ApplicationStatusCondition) IsValid() bool {
	switch e {
	case ApplicationStatusConditionInitial, ApplicationStatusConditionUnknown, ApplicationStatusConditionReady, ApplicationStatusConditionFailed:
		return true
	}
	return false
}

func (e ApplicationStatusCondition) String() string {
	return string(e)
}

func (e *ApplicationStatusCondition) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ApplicationStatusCondition(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ApplicationStatusCondition", str)
	}
	return nil
}

func (e ApplicationStatusCondition) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ApplicationWebhookType string

const (
	ApplicationWebhookTypeConfigurationChanged ApplicationWebhookType = "CONFIGURATION_CHANGED"
)

var AllApplicationWebhookType = []ApplicationWebhookType{
	ApplicationWebhookTypeConfigurationChanged,
}

func (e ApplicationWebhookType) IsValid() bool {
	switch e {
	case ApplicationWebhookTypeConfigurationChanged:
		return true
	}
	return false
}

func (e ApplicationWebhookType) String() string {
	return string(e)
}

func (e *ApplicationWebhookType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ApplicationWebhookType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ApplicationWebhookType", str)
	}
	return nil
}

func (e ApplicationWebhookType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DocumentFormat string

const (
	DocumentFormatMarkdown DocumentFormat = "MARKDOWN"
)

var AllDocumentFormat = []DocumentFormat{
	DocumentFormatMarkdown,
}

func (e DocumentFormat) IsValid() bool {
	switch e {
	case DocumentFormatMarkdown:
		return true
	}
	return false
}

func (e DocumentFormat) String() string {
	return string(e)
}

func (e *DocumentFormat) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DocumentFormat(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DocumentFormat", str)
	}
	return nil
}

func (e DocumentFormat) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventSpecType string

const (
	EventSpecTypeAsyncAPI EventSpecType = "ASYNC_API"
)

var AllEventSpecType = []EventSpecType{
	EventSpecTypeAsyncAPI,
}

func (e EventSpecType) IsValid() bool {
	switch e {
	case EventSpecTypeAsyncAPI:
		return true
	}
	return false
}

func (e EventSpecType) String() string {
	return string(e)
}

func (e *EventSpecType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventSpecType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventSpecType", str)
	}
	return nil
}

func (e EventSpecType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FetchMode string

const (
	FetchModeSingle  FetchMode = "SINGLE"
	FetchModePackage FetchMode = "PACKAGE"
	FetchModeIndex   FetchMode = "INDEX"
)

var AllFetchMode = []FetchMode{
	FetchModeSingle,
	FetchModePackage,
	FetchModeIndex,
}

func (e FetchMode) IsValid() bool {
	switch e {
	case FetchModeSingle, FetchModePackage, FetchModeIndex:
		return true
	}
	return false
}

func (e FetchMode) String() string {
	return string(e)
}

func (e *FetchMode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FetchMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FetchMode", str)
	}
	return nil
}

func (e FetchMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FetchRequestStatusCondition string

const (
	FetchRequestStatusConditionInitial   FetchRequestStatusCondition = "INITIAL"
	FetchRequestStatusConditionSucceeded FetchRequestStatusCondition = "SUCCEEDED"
	FetchRequestStatusConditionFailed    FetchRequestStatusCondition = "FAILED"
)

var AllFetchRequestStatusCondition = []FetchRequestStatusCondition{
	FetchRequestStatusConditionInitial,
	FetchRequestStatusConditionSucceeded,
	FetchRequestStatusConditionFailed,
}

func (e FetchRequestStatusCondition) IsValid() bool {
	switch e {
	case FetchRequestStatusConditionInitial, FetchRequestStatusConditionSucceeded, FetchRequestStatusConditionFailed:
		return true
	}
	return false
}

func (e FetchRequestStatusCondition) String() string {
	return string(e)
}

func (e *FetchRequestStatusCondition) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FetchRequestStatusCondition(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FetchRequestStatusCondition", str)
	}
	return nil
}

func (e FetchRequestStatusCondition) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FilterOperator string

const (
	FilterOperatorAll FilterOperator = "ALL"
	FilterOperatorAny FilterOperator = "ANY"
)

var AllFilterOperator = []FilterOperator{
	FilterOperatorAll,
	FilterOperatorAny,
}

func (e FilterOperator) IsValid() bool {
	switch e {
	case FilterOperatorAll, FilterOperatorAny:
		return true
	}
	return false
}

func (e FilterOperator) String() string {
	return string(e)
}

func (e *FilterOperator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilterOperator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilterOperator", str)
	}
	return nil
}

func (e FilterOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type HealthCheckStatusCondition string

const (
	HealthCheckStatusConditionSucceeded HealthCheckStatusCondition = "SUCCEEDED"
	HealthCheckStatusConditionFailed    HealthCheckStatusCondition = "FAILED"
)

var AllHealthCheckStatusCondition = []HealthCheckStatusCondition{
	HealthCheckStatusConditionSucceeded,
	HealthCheckStatusConditionFailed,
}

func (e HealthCheckStatusCondition) IsValid() bool {
	switch e {
	case HealthCheckStatusConditionSucceeded, HealthCheckStatusConditionFailed:
		return true
	}
	return false
}

func (e HealthCheckStatusCondition) String() string {
	return string(e)
}

func (e *HealthCheckStatusCondition) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HealthCheckStatusCondition(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HealthCheckStatusCondition", str)
	}
	return nil
}

func (e HealthCheckStatusCondition) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type HealthCheckType string

const (
	HealthCheckTypeManagementPlaneApplicationHealthcheck HealthCheckType = "MANAGEMENT_PLANE_APPLICATION_HEALTHCHECK"
)

var AllHealthCheckType = []HealthCheckType{
	HealthCheckTypeManagementPlaneApplicationHealthcheck,
}

func (e HealthCheckType) IsValid() bool {
	switch e {
	case HealthCheckTypeManagementPlaneApplicationHealthcheck:
		return true
	}
	return false
}

func (e HealthCheckType) String() string {
	return string(e)
}

func (e *HealthCheckType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HealthCheckType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HealthCheckType", str)
	}
	return nil
}

func (e HealthCheckType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RuntimeStatusCondition string

const (
	RuntimeStatusConditionInitial RuntimeStatusCondition = "INITIAL"
	RuntimeStatusConditionReady   RuntimeStatusCondition = "READY"
	RuntimeStatusConditionFailed  RuntimeStatusCondition = "FAILED"
)

var AllRuntimeStatusCondition = []RuntimeStatusCondition{
	RuntimeStatusConditionInitial,
	RuntimeStatusConditionReady,
	RuntimeStatusConditionFailed,
}

func (e RuntimeStatusCondition) IsValid() bool {
	switch e {
	case RuntimeStatusConditionInitial, RuntimeStatusConditionReady, RuntimeStatusConditionFailed:
		return true
	}
	return false
}

func (e RuntimeStatusCondition) String() string {
	return string(e)
}

func (e *RuntimeStatusCondition) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RuntimeStatusCondition(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RuntimeStatusCondition", str)
	}
	return nil
}

func (e RuntimeStatusCondition) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SpecFormat string

const (
	SpecFormatYaml SpecFormat = "YAML"
	SpecFormatJSON SpecFormat = "JSON"
)

var AllSpecFormat = []SpecFormat{
	SpecFormatYaml,
	SpecFormatJSON,
}

func (e SpecFormat) IsValid() bool {
	switch e {
	case SpecFormatYaml, SpecFormatJSON:
		return true
	}
	return false
}

func (e SpecFormat) String() string {
	return string(e)
}

func (e *SpecFormat) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SpecFormat(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SpecFormat", str)
	}
	return nil
}

func (e SpecFormat) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
