// Package pagespeedonline provides access to the PageSpeed Insights API.
//
// See https://developers.google.com/speed/docs/insights/v4/getting-started
//
// Usage example:
//
//   import "google.golang.org/api/pagespeedonline/v4"
//   ...
//   pagespeedonlineService, err := pagespeedonline.New(oauthHttpClient)
package pagespeedonline // import "google.golang.org/api/pagespeedonline/v4"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled

const apiId = "pagespeedonline:v4"
const apiName = "pagespeedonline"
const apiVersion = "v4"
const basePath = "https://www.googleapis.com/pagespeedonline/v4/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Pagespeedapi = NewPagespeedapiService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Pagespeedapi *PagespeedapiService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewPagespeedapiService(s *Service) *PagespeedapiService {
	rs := &PagespeedapiService{s: s}
	return rs
}

type PagespeedapiService struct {
	s *Service
}

type PagespeedApiFormatStringV4 struct {
	// Args: List of arguments for the format string.
	Args []*PagespeedApiFormatStringV4Args `json:"args,omitempty"`

	// Format: A localized format string with {{FOO}} placeholders, where
	// 'FOO' is the key of the argument whose value should be substituted.
	// For HYPERLINK arguments, the format string will instead contain
	// {{BEGIN_FOO}} and {{END_FOO}} for the argument with key 'FOO'.
	Format string `json:"format,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Args") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Args") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiFormatStringV4) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiFormatStringV4
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PagespeedApiFormatStringV4Args struct {
	// Key: The placeholder key for this arg, as a string.
	Key string `json:"key,omitempty"`

	// Rects: The screen rectangles being referred to, with dimensions
	// measured in CSS pixels. This is only ever used for SNAPSHOT_RECT
	// arguments. If this is absent for a SNAPSHOT_RECT argument, it means
	// that that argument refers to the entire snapshot.
	Rects []*PagespeedApiFormatStringV4ArgsRects `json:"rects,omitempty"`

	// SecondaryRects: Secondary screen rectangles being referred to, with
	// dimensions measured in CSS pixels. This is only ever used for
	// SNAPSHOT_RECT arguments.
	SecondaryRects []*PagespeedApiFormatStringV4ArgsSecondaryRects `json:"secondary_rects,omitempty"`

	// Type: Type of argument. One of URL, STRING_LITERAL, INT_LITERAL,
	// BYTES, DURATION, VERBATIM_STRING, PERCENTAGE, HYPERLINK, or
	// SNAPSHOT_RECT.
	Type string `json:"type,omitempty"`

	// Value: Argument value, as a localized string.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiFormatStringV4Args) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiFormatStringV4Args
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PagespeedApiFormatStringV4ArgsRects struct {
	Height int64 `json:"height,omitempty"`

	Left int64 `json:"left,omitempty"`

	Top int64 `json:"top,omitempty"`

	Width int64 `json:"width,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Height") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Height") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiFormatStringV4ArgsRects) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiFormatStringV4ArgsRects
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PagespeedApiFormatStringV4ArgsSecondaryRects struct {
	Height int64 `json:"height,omitempty"`

	Left int64 `json:"left,omitempty"`

	Top int64 `json:"top,omitempty"`

	Width int64 `json:"width,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Height") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Height") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiFormatStringV4ArgsSecondaryRects) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiFormatStringV4ArgsSecondaryRects
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PagespeedApiImageV4 struct {
	// Data: Image data base64 encoded.
	Data string `json:"data,omitempty"`

	// Height: Height of screenshot in pixels.
	Height int64 `json:"height,omitempty"`

	// Key: Unique string key, if any, identifying this image.
	Key string `json:"key,omitempty"`

	// MimeType: Mime type of image data (e.g. "image/jpeg").
	MimeType string `json:"mime_type,omitempty"`

	PageRect *PagespeedApiImageV4PageRect `json:"page_rect,omitempty"`

	// Width: Width of screenshot in pixels.
	Width int64 `json:"width,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Data") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Data") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiImageV4) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiImageV4
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PagespeedApiImageV4PageRect struct {
	Height int64 `json:"height,omitempty"`

	Left int64 `json:"left,omitempty"`

	Top int64 `json:"top,omitempty"`

	Width int64 `json:"width,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Height") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Height") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiImageV4PageRect) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiImageV4PageRect
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PagespeedApiPagespeedResponseV4 struct {
	// CaptchaResult: The captcha verify result
	CaptchaResult string `json:"captchaResult,omitempty"`

	// FormattedResults: Localized PageSpeed results. Contains a ruleResults
	// entry for each PageSpeed rule instantiated and run by the server.
	FormattedResults *PagespeedApiPagespeedResponseV4FormattedResults `json:"formattedResults,omitempty"`

	// Id: Canonicalized and final URL for the document, after following
	// page redirects (if any).
	Id string `json:"id,omitempty"`

	// InvalidRules: List of rules that were specified in the request, but
	// which the server did not know how to instantiate.
	InvalidRules []string `json:"invalidRules,omitempty"`

	// Kind: Kind of result.
	Kind string `json:"kind,omitempty"`

	// LoadingExperience: Metrics of end users' page loading experience.
	LoadingExperience *PagespeedApiPagespeedResponseV4LoadingExperience `json:"loadingExperience,omitempty"`

	// PageStats: Summary statistics for the page, such as number of
	// JavaScript bytes, number of HTML bytes, etc.
	PageStats *PagespeedApiPagespeedResponseV4PageStats `json:"pageStats,omitempty"`

	// ResponseCode: Response code for the document. 200 indicates a normal
	// page load. 4xx/5xx indicates an error.
	ResponseCode int64 `json:"responseCode,omitempty"`

	// RuleGroups: A map with one entry for each rule group in these
	// results.
	RuleGroups map[string]PagespeedApiPagespeedResponseV4RuleGroups `json:"ruleGroups,omitempty"`

	// Screenshot: Base64-encoded screenshot of the page that was analyzed.
	Screenshot *PagespeedApiImageV4 `json:"screenshot,omitempty"`

	// Snapshots: Additional base64-encoded screenshots of the page, in
	// various partial render states.
	Snapshots []*PagespeedApiImageV4 `json:"snapshots,omitempty"`

	// Title: Title of the page, as displayed in the browser's title bar.
	Title string `json:"title,omitempty"`

	// Version: The version of PageSpeed used to generate these results.
	Version *PagespeedApiPagespeedResponseV4Version `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CaptchaResult") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CaptchaResult") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV4) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV4
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PagespeedApiPagespeedResponseV4FormattedResults: Localized PageSpeed
// results. Contains a ruleResults entry for each PageSpeed rule
// instantiated and run by the server.
type PagespeedApiPagespeedResponseV4FormattedResults struct {
	// Locale: The locale of the formattedResults, e.g. "en_US".
	Locale string `json:"locale,omitempty"`

	// RuleResults: Dictionary of formatted rule results, with one entry for
	// each PageSpeed rule instantiated and run by the server.
	RuleResults map[string]PagespeedApiPagespeedResponseV4FormattedResultsRuleResults `json:"ruleResults,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Locale") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Locale") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV4FormattedResults) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV4FormattedResults
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PagespeedApiPagespeedResponseV4FormattedResultsRuleResults: The
// enum-like identifier for this rule. For instance "EnableKeepAlive" or
// "AvoidCssImport". Not localized.
type PagespeedApiPagespeedResponseV4FormattedResultsRuleResults struct {
	// Beta: Whether this rule is in 'beta'. Rules in beta are new rules
	// that are being tested, which do not impact the overall score.
	Beta bool `json:"beta,omitempty"`

	// Groups: List of rule groups that this rule belongs to. Each entry in
	// the list is one of "SPEED", "USABILITY", or "SECURITY".
	Groups []string `json:"groups,omitempty"`

	// LocalizedRuleName: Localized name of the rule, intended for
	// presentation to a user.
	LocalizedRuleName string `json:"localizedRuleName,omitempty"`

	// RuleImpact: The impact (unbounded floating point value) that
	// implementing the suggestions for this rule would have on making the
	// page faster. Impact is comparable between rules to determine which
	// rule's suggestions would have a higher or lower impact on making a
	// page faster. For instance, if enabling compression would save 1MB,
	// while optimizing images would save 500kB, the enable compression rule
	// would have 2x the impact of the image optimization rule, all other
	// things being equal.
	RuleImpact float64 `json:"ruleImpact,omitempty"`

	// Summary: A brief summary description for the rule, indicating at a
	// high level what should be done to follow the rule and what benefit
	// can be gained by doing so.
	Summary *PagespeedApiFormatStringV4 `json:"summary,omitempty"`

	// UrlBlocks: List of blocks of URLs. Each block may contain a heading
	// and a list of URLs. Each URL may optionally include additional
	// details.
	UrlBlocks []*PagespeedApiPagespeedResponseV4FormattedResultsRuleResultsUrlBlocks `json:"urlBlocks,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Beta") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Beta") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV4FormattedResultsRuleResults) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV4FormattedResultsRuleResults
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *PagespeedApiPagespeedResponseV4FormattedResultsRuleResults) UnmarshalJSON(data []byte) error {
	type NoMethod PagespeedApiPagespeedResponseV4FormattedResultsRuleResults
	var s1 struct {
		RuleImpact gensupport.JSONFloat64 `json:"ruleImpact"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.RuleImpact = float64(s1.RuleImpact)
	return nil
}

type PagespeedApiPagespeedResponseV4FormattedResultsRuleResultsUrlBlocks struct {
	// Header: Heading to be displayed with the list of URLs.
	Header *PagespeedApiFormatStringV4 `json:"header,omitempty"`

	// Urls: List of entries that provide information about URLs in the url
	// block. Optional.
	Urls []*PagespeedApiPagespeedResponseV4FormattedResultsRuleResultsUrlBlocksUrls `json:"urls,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Header") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Header") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV4FormattedResultsRuleResultsUrlBlocks) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV4FormattedResultsRuleResultsUrlBlocks
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PagespeedApiPagespeedResponseV4FormattedResultsRuleResultsUrlBlocksUrls struct {
	// Details: List of entries that provide additional details about a
	// single URL. Optional.
	Details []*PagespeedApiFormatStringV4 `json:"details,omitempty"`

	// Result: A format string that gives information about the URL, and a
	// list of arguments for that format string.
	Result *PagespeedApiFormatStringV4 `json:"result,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Details") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Details") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV4FormattedResultsRuleResultsUrlBlocksUrls) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV4FormattedResultsRuleResultsUrlBlocksUrls
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PagespeedApiPagespeedResponseV4LoadingExperience: Metrics of end
// users' page loading experience.
type PagespeedApiPagespeedResponseV4LoadingExperience struct {
	// Id: The url, pattern or origin which the metrics are on.
	Id string `json:"id,omitempty"`

	InitialUrl string `json:"initial_url,omitempty"`

	Metrics map[string]PagespeedApiPagespeedResponseV4LoadingExperienceMetrics `json:"metrics,omitempty"`

	OverallCategory string `json:"overall_category,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV4LoadingExperience) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV4LoadingExperience
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PagespeedApiPagespeedResponseV4LoadingExperienceMetrics: The type of
// the metric.
type PagespeedApiPagespeedResponseV4LoadingExperienceMetrics struct {
	Category string `json:"category,omitempty"`

	Distributions []*PagespeedApiPagespeedResponseV4LoadingExperienceMetricsDistributions `json:"distributions,omitempty"`

	Median int64 `json:"median,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Category") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Category") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV4LoadingExperienceMetrics) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV4LoadingExperienceMetrics
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PagespeedApiPagespeedResponseV4LoadingExperienceMetricsDistributions struct {
	Max int64 `json:"max,omitempty"`

	Min int64 `json:"min,omitempty"`

	Proportion float64 `json:"proportion,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Max") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Max") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV4LoadingExperienceMetricsDistributions) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV4LoadingExperienceMetricsDistributions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *PagespeedApiPagespeedResponseV4LoadingExperienceMetricsDistributions) UnmarshalJSON(data []byte) error {
	type NoMethod PagespeedApiPagespeedResponseV4LoadingExperienceMetricsDistributions
	var s1 struct {
		Proportion gensupport.JSONFloat64 `json:"proportion"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Proportion = float64(s1.Proportion)
	return nil
}

// PagespeedApiPagespeedResponseV4PageStats: Summary statistics for the
// page, such as number of JavaScript bytes, number of HTML bytes, etc.
type PagespeedApiPagespeedResponseV4PageStats struct {
	// Cms: Content management system (CMS) used for the page.
	Cms string `json:"cms,omitempty"`

	// CssResponseBytes: Number of uncompressed response bytes for CSS
	// resources on the page.
	CssResponseBytes int64 `json:"cssResponseBytes,omitempty,string"`

	// FlashResponseBytes: Number of response bytes for flash resources on
	// the page.
	FlashResponseBytes int64 `json:"flashResponseBytes,omitempty,string"`

	// HtmlResponseBytes: Number of uncompressed response bytes for the main
	// HTML document and all iframes on the page.
	HtmlResponseBytes int64 `json:"htmlResponseBytes,omitempty,string"`

	// ImageResponseBytes: Number of response bytes for image resources on
	// the page.
	ImageResponseBytes int64 `json:"imageResponseBytes,omitempty,string"`

	// JavascriptResponseBytes: Number of uncompressed response bytes for JS
	// resources on the page.
	JavascriptResponseBytes int64 `json:"javascriptResponseBytes,omitempty,string"`

	// NumRenderBlockingRoundTrips: The needed round trips to load render
	// blocking resources
	NumRenderBlockingRoundTrips int64 `json:"numRenderBlockingRoundTrips,omitempty"`

	// NumTotalRoundTrips: The needed round trips to load the full page
	NumTotalRoundTrips int64 `json:"numTotalRoundTrips,omitempty"`

	// NumberCssResources: Number of CSS resources referenced by the page.
	NumberCssResources int64 `json:"numberCssResources,omitempty"`

	// NumberHosts: Number of unique hosts referenced by the page.
	NumberHosts int64 `json:"numberHosts,omitempty"`

	// NumberJsResources: Number of JavaScript resources referenced by the
	// page.
	NumberJsResources int64 `json:"numberJsResources,omitempty"`

	// NumberResources: Number of HTTP resources loaded by the page.
	NumberResources int64 `json:"numberResources,omitempty"`

	// NumberRobotedResources: Number of roboted resources.
	NumberRobotedResources int64 `json:"numberRobotedResources,omitempty"`

	// NumberStaticResources: Number of static (i.e. cacheable) resources on
	// the page.
	NumberStaticResources int64 `json:"numberStaticResources,omitempty"`

	// NumberTransientFetchFailureResources: Number of transient-failed
	// resources.
	NumberTransientFetchFailureResources int64 `json:"numberTransientFetchFailureResources,omitempty"`

	// OtherResponseBytes: Number of response bytes for other resources on
	// the page.
	OtherResponseBytes int64 `json:"otherResponseBytes,omitempty,string"`

	// OverTheWireResponseBytes: Number of over-the-wire bytes, uses the
	// default gzip compression strategy as an estimation.
	OverTheWireResponseBytes int64 `json:"overTheWireResponseBytes,omitempty,string"`

	// RobotedUrls: List of roboted urls.
	RobotedUrls []string `json:"robotedUrls,omitempty"`

	// TextResponseBytes: Number of uncompressed response bytes for text
	// resources not covered by other statistics (i.e non-HTML, non-script,
	// non-CSS resources) on the page.
	TextResponseBytes int64 `json:"textResponseBytes,omitempty,string"`

	// TotalRequestBytes: Total size of all request bytes sent by the page.
	TotalRequestBytes int64 `json:"totalRequestBytes,omitempty,string"`

	// TransientFetchFailureUrls: List of transient fetch failure urls.
	TransientFetchFailureUrls []string `json:"transientFetchFailureUrls,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Cms") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Cms") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV4PageStats) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV4PageStats
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PagespeedApiPagespeedResponseV4RuleGroups: The name of this rule
// group: one of "SPEED", "USABILITY", or "SECURITY".
type PagespeedApiPagespeedResponseV4RuleGroups struct {
	Pass bool `json:"pass,omitempty"`

	// Score: The score (0-100) for this rule group, which indicates how
	// much better a page could be in that category (e.g. how much faster,
	// or how much more usable, or how much more secure). A high score
	// indicates little room for improvement, while a lower score indicates
	// more room for improvement.
	Score int64 `json:"score,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Pass") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Pass") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV4RuleGroups) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV4RuleGroups
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PagespeedApiPagespeedResponseV4Version: The version of PageSpeed used
// to generate these results.
type PagespeedApiPagespeedResponseV4Version struct {
	// Major: The major version number of PageSpeed used to generate these
	// results.
	Major int64 `json:"major,omitempty"`

	// Minor: The minor version number of PageSpeed used to generate these
	// results.
	Minor int64 `json:"minor,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Major") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Major") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PagespeedApiPagespeedResponseV4Version) MarshalJSON() ([]byte, error) {
	type NoMethod PagespeedApiPagespeedResponseV4Version
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "pagespeedonline.pagespeedapi.runpagespeed":

type PagespeedapiRunpagespeedCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Runpagespeed: Runs PageSpeed analysis on the page at the specified
// URL, and returns PageSpeed scores, a list of suggestions to make that
// page faster, and other information.
func (r *PagespeedapiService) Runpagespeed(url string) *PagespeedapiRunpagespeedCall {
	c := &PagespeedapiRunpagespeedCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("url", url)
	return c
}

// FilterThirdPartyResources sets the optional parameter
// "filter_third_party_resources": Indicates if third party resources
// should be filtered out before PageSpeed analysis.
func (c *PagespeedapiRunpagespeedCall) FilterThirdPartyResources(filterThirdPartyResources bool) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("filter_third_party_resources", fmt.Sprint(filterThirdPartyResources))
	return c
}

// Locale sets the optional parameter "locale": The locale used to
// localize formatted results
func (c *PagespeedapiRunpagespeedCall) Locale(locale string) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("locale", locale)
	return c
}

// Rule sets the optional parameter "rule": A PageSpeed rule to run; if
// none are given, all rules are run
func (c *PagespeedapiRunpagespeedCall) Rule(rule ...string) *PagespeedapiRunpagespeedCall {
	c.urlParams_.SetMulti("rule", append([]string{}, rule...))
	return c
}

// Screenshot sets the optional parameter "screenshot": Indicates if
// binary data containing a screenshot should be included
func (c *PagespeedapiRunpagespeedCall) Screenshot(screenshot bool) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("screenshot", fmt.Sprint(screenshot))
	return c
}

// Snapshots sets the optional parameter "snapshots": Indicates if
// binary data containing snapshot images should be included
func (c *PagespeedapiRunpagespeedCall) Snapshots(snapshots bool) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("snapshots", fmt.Sprint(snapshots))
	return c
}

// Strategy sets the optional parameter "strategy": The analysis
// strategy (desktop or mobile) to use, and desktop is the default
//
// Possible values:
//   "desktop" - Fetch and analyze the URL for desktop browsers
//   "mobile" - Fetch and analyze the URL for mobile devices
func (c *PagespeedapiRunpagespeedCall) Strategy(strategy string) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("strategy", strategy)
	return c
}

// UtmCampaign sets the optional parameter "utm_campaign": Campaign name
// for analytics.
func (c *PagespeedapiRunpagespeedCall) UtmCampaign(utmCampaign string) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("utm_campaign", utmCampaign)
	return c
}

// UtmSource sets the optional parameter "utm_source": Campaign source
// for analytics.
func (c *PagespeedapiRunpagespeedCall) UtmSource(utmSource string) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("utm_source", utmSource)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PagespeedapiRunpagespeedCall) Fields(s ...googleapi.Field) *PagespeedapiRunpagespeedCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PagespeedapiRunpagespeedCall) IfNoneMatch(entityTag string) *PagespeedapiRunpagespeedCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PagespeedapiRunpagespeedCall) Context(ctx context.Context) *PagespeedapiRunpagespeedCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PagespeedapiRunpagespeedCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PagespeedapiRunpagespeedCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "runPagespeed")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "pagespeedonline.pagespeedapi.runpagespeed" call.
// Exactly one of *PagespeedApiPagespeedResponseV4 or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *PagespeedApiPagespeedResponseV4.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PagespeedapiRunpagespeedCall) Do(opts ...googleapi.CallOption) (*PagespeedApiPagespeedResponseV4, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &PagespeedApiPagespeedResponseV4{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Runs PageSpeed analysis on the page at the specified URL, and returns PageSpeed scores, a list of suggestions to make that page faster, and other information.",
	//   "httpMethod": "GET",
	//   "id": "pagespeedonline.pagespeedapi.runpagespeed",
	//   "parameterOrder": [
	//     "url"
	//   ],
	//   "parameters": {
	//     "filter_third_party_resources": {
	//       "default": "false",
	//       "description": "Indicates if third party resources should be filtered out before PageSpeed analysis.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "locale": {
	//       "description": "The locale used to localize formatted results",
	//       "location": "query",
	//       "pattern": "[a-zA-Z]+(_[a-zA-Z]+)?",
	//       "type": "string"
	//     },
	//     "rule": {
	//       "description": "A PageSpeed rule to run; if none are given, all rules are run",
	//       "location": "query",
	//       "pattern": "[a-zA-Z]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "screenshot": {
	//       "default": "false",
	//       "description": "Indicates if binary data containing a screenshot should be included",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "snapshots": {
	//       "default": "false",
	//       "description": "Indicates if binary data containing snapshot images should be included",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "strategy": {
	//       "description": "The analysis strategy (desktop or mobile) to use, and desktop is the default",
	//       "enum": [
	//         "desktop",
	//         "mobile"
	//       ],
	//       "enumDescriptions": [
	//         "Fetch and analyze the URL for desktop browsers",
	//         "Fetch and analyze the URL for mobile devices"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "url": {
	//       "description": "The URL to fetch and analyze",
	//       "location": "query",
	//       "pattern": "(?i)(site:|origin:)?http(s)?://.*",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "utm_campaign": {
	//       "description": "Campaign name for analytics.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "utm_source": {
	//       "description": "Campaign source for analytics.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "runPagespeed",
	//   "response": {
	//     "$ref": "PagespeedApiPagespeedResponseV4"
	//   }
	// }

}
