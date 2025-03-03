// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    coordinate, err := UnmarshalCoordinate(bytes)
//    bytes, err = coordinate.Marshal()

package testly

import (
	"bytes"
	"encoding/json"
	"errors"
)

func UnmarshalCoordinate(data []byte) (Coordinate, error) {
	var r Coordinate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Coordinate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Coordinate struct {
	After      *After            `json:"after,omitempty"`
	Before     *Before           `json:"before,omitempty"`
	Components *Components       `json:"components,omitempty"`
	Config     *Config           `json:"config,omitempty"`
	Env        map[string]string `json:"env,omitempty"`
	Include    []string          `json:"include,omitempty"`
	Name       string            `json:"name"`
	Tests      map[string]Test   `json:"tests"`
	Version    string            `json:"version"`
}

type After struct {
	Env      map[string]interface{} `json:"env,omitempty"`
	Name     *string                `json:"name,omitempty"`
	Steps    []AfterStep            `json:"steps"`
	Testdata *AfterTestdata         `json:"testdata,omitempty"`
}

type AfterStep struct {
	Delay   *string        `json:"delay,omitempty"`
	Graphql *GraphQLStep   `json:"graphql,omitempty"`
	Grpc    *PurpleGrpc    `json:"grpc,omitempty"`
	HTTP    *HTTPStep      `json:"http,omitempty"`
	ID      *string        `json:"id,omitempty"`
	If      *string        `json:"if,omitempty"`
	Name    *string        `json:"name,omitempty"`
	Plugin  *PurplePlugin  `json:"plugin,omitempty"`
	Retries *PurpleRetries `json:"retries,omitempty"`
	SSE     *PurpleSSE     `json:"sse,omitempty"`
	Trpc    *TRPCStep      `json:"trpc,omitempty"`
}

type GraphQLStep struct {
	Query           string                        `json:"query"`
	Variables       map[string]interface{}        `json:"variables"`
	Auth            *GraphQLStepAuth              `json:"auth,omitempty"`
	Captures        map[string]GraphQLStepCapture `json:"captures,omitempty"`
	Check           *GraphQLStepCheck             `json:"check,omitempty"`
	Cookies         map[string]string             `json:"cookies,omitempty"`
	FollowRedirects *bool                         `json:"followRedirects,omitempty"`
	Headers         map[string]string             `json:"headers,omitempty"`
	Method          string                        `json:"method"`
	Params          map[string]string             `json:"params,omitempty"`
	Retries         *float64                      `json:"retries,omitempty"`
	Timeout         *Timeout                      `json:"timeout"`
	URL             string                        `json:"url"`
}

type GraphQLStepAuth struct {
	Basic       *PurpleBasic       `json:"basic,omitempty"`
	Bearer      *PurpleBearer      `json:"bearer,omitempty"`
	Certificate *PurpleCertificate `json:"certificate,omitempty"`
	Oauth       *PurpleOauth       `json:"oauth,omitempty"`
	TLS         *PurpleTLS         `json:"tls,omitempty"`
}

type PurpleBasic struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type PurpleBearer struct {
	Token string `json:"token"`
}

type PurpleCertificate struct {
	CA         *AmbitiousCA   `json:"ca"`
	CERT       *AmbitiousCERT `json:"cert"`
	Key        *AmbitiousKey  `json:"key"`
	Passphrase *string        `json:"passphrase,omitempty"`
}

type PurpleCA struct {
	File string `json:"file"`
}

type PurpleCERT struct {
	File string `json:"file"`
}

type PurpleKey struct {
	File string `json:"file"`
}

type PurpleOauth struct {
	Audience     *string `json:"audience,omitempty"`
	ClientID     string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	Endpoint     string  `json:"endpoint"`
}

type PurpleTLS struct {
	CERTChain  *FriskyCERTChain  `json:"certChain"`
	PrivateKey *FriskyPrivateKey `json:"privateKey"`
	RootCerts  *FriskyRootCerts  `json:"rootCerts"`
}

type PurpleCERTChain struct {
	File string `json:"file"`
}

type PurplePrivateKey struct {
	File string `json:"file"`
}

type PurpleRootCerts struct {
	File string `json:"file"`
}

type GraphQLStepCapture struct {
	Body     *bool   `json:"body,omitempty"`
	Cookie   *string `json:"cookie,omitempty"`
	Header   *string `json:"header,omitempty"`
	Jsonpath *string `json:"jsonpath,omitempty"`
	Regex    *string `json:"regex,omitempty"`
	Selector *string `json:"selector,omitempty"`
	Xpath    *string `json:"xpath,omitempty"`
}

type GraphQLStepCheck struct {
	Body        *AmbitiousBody                   `json:"body"`
	BodySize    *StickyBodySize                  `json:"bodySize"`
	Captures    map[string]interface{}           `json:"captures,omitempty"`
	Co2         *HilariousCo2                    `json:"co2"`
	Cookies     map[string]*StickyCooky          `json:"cookies,omitempty"`
	Headers     map[string]*StickyHeader         `json:"headers,omitempty"`
	JSON        map[string]interface{}           `json:"json,omitempty"`
	Jsonpath    map[string]interface{}           `json:"jsonpath,omitempty"`
	Md5         *string                          `json:"md5,omitempty"`
	Performance map[string]*HilariousPerformance `json:"performance,omitempty"`
	Redirected  *bool                            `json:"redirected,omitempty"`
	Redirects   []string                         `json:"redirects,omitempty"`
	RequestSize *StickyRequestSize               `json:"requestSize"`
	Schema      map[string]interface{}           `json:"schema,omitempty"`
	Selectors   map[string]*StickySelector       `json:"selectors,omitempty"`
	Sha256      *string                          `json:"sha256,omitempty"`
	Size        *HilariousSize                   `json:"size"`
	SSL         *PurpleSSL                       `json:"ssl,omitempty"`
	Status      *StickyStatus                    `json:"status"`
	StatusText  *StickyStatusText                `json:"statusText"`
	Xpath       map[string]*StickyXpath          `json:"xpath,omitempty"`
}

type PurpleBody struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleBodySize struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleCo2 struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleCooky struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleHeader struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurplePerformance struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleRequestSize struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleSSL struct {
	DaysUntilExpiration *StickyDaysUntilExpiration `json:"daysUntilExpiration"`
	Signed              *bool                      `json:"signed,omitempty"`
	Valid               *bool                      `json:"valid,omitempty"`
}

type PurpleDaysUntilExpiration struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleSelector struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleSize struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleStatus struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleStatusText struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleXpath struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleGrpc struct {
	Auth     *PurpleAuth                 `json:"auth,omitempty"`
	Captures map[string]PurpleCapture    `json:"captures,omitempty"`
	Check    *PurpleCheck                `json:"check,omitempty"`
	Data     *Data                       `json:"data"`
	Host     string                      `json:"host"`
	Metadata map[string]*StickyMetadatum `json:"metadata,omitempty"`
	Method   string                      `json:"method"`
	Proto    *Proto                      `json:"proto"`
	Service  string                      `json:"service"`
	Timeout  *Timeout                    `json:"timeout"`
}

type PurpleAuth struct {
	TLS *FluffyTLS `json:"tls,omitempty"`
}

type FluffyTLS struct {
	CERTChain  *MischievousCERTChain  `json:"certChain"`
	PrivateKey *MischievousPrivateKey `json:"privateKey"`
	RootCerts  *MischievousRootCerts  `json:"rootCerts"`
}

type FluffyCERTChain struct {
	File string `json:"file"`
}

type FluffyPrivateKey struct {
	File string `json:"file"`
}

type FluffyRootCerts struct {
	File string `json:"file"`
}

type PurpleCapture struct {
	Jsonpath *string `json:"jsonpath,omitempty"`
}

type PurpleCheck struct {
	Captures    map[string]interface{}           `json:"captures,omitempty"`
	Co2         *AmbitiousCo2                    `json:"co2"`
	JSON        map[string]interface{}           `json:"json,omitempty"`
	Jsonpath    map[string]interface{}           `json:"jsonpath,omitempty"`
	Performance map[string]*AmbitiousPerformance `json:"performance,omitempty"`
	Schema      map[string]interface{}           `json:"schema,omitempty"`
	Size        *AmbitiousSize                   `json:"size"`
}

type FluffyCo2 struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FluffyPerformance struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FluffySize struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type PurpleMetadatum struct {
	ToStringTag23   MetadatumToStringTag23 `json:"__@toStringTag@23"`
	Buffer          ArrayBufferLike        `json:"buffer"`
	ByteLength      float64                `json:"byteLength"`
	ByteOffset      float64                `json:"byteOffset"`
	BytesPerElement float64                `json:"BYTES_PER_ELEMENT"`
	Length          float64                `json:"length"`
}

type ArrayBufferLike struct {
	ToStringTag23 string            `json:"__@toStringTag@23"`
	ByteLength    float64           `json:"byteLength"`
	Species469    *Species469_Class `json:"__@species@469,omitempty"`
}

type Species469_Class struct {
	Species469    *Species469_Class          `json:"__@species@469"`
	ToStringTag23 Species469___ToStringTag23 `json:"__@toStringTag@23"`
	ByteLength    float64                    `json:"byteLength"`
}

type HTTPStep struct {
	Body            *HTTPStepBody              `json:"body"`
	Form            map[string]string          `json:"form,omitempty"`
	FormData        map[string]*FormDatumValue `json:"formData,omitempty"`
	Graphql         *Graphql                   `json:"graphql,omitempty"`
	JSON            map[string]interface{}     `json:"json,omitempty"`
	Trpc            *Trpc                      `json:"trpc,omitempty"`
	Auth            *HTTPStepAuth              `json:"auth,omitempty"`
	Captures        map[string]HTTPStepCapture `json:"captures,omitempty"`
	Check           *HTTPStepCheck             `json:"check,omitempty"`
	Cookies         map[string]string          `json:"cookies,omitempty"`
	FollowRedirects *bool                      `json:"followRedirects,omitempty"`
	Headers         map[string]string          `json:"headers,omitempty"`
	Method          string                     `json:"method"`
	Params          map[string]string          `json:"params,omitempty"`
	Retries         *float64                   `json:"retries,omitempty"`
	Timeout         *Timeout                   `json:"timeout"`
	URL             string                     `json:"url"`
}

type HTTPStepAuth struct {
	Basic       *FluffyBasic       `json:"basic,omitempty"`
	Bearer      *FluffyBearer      `json:"bearer,omitempty"`
	Certificate *FluffyCertificate `json:"certificate,omitempty"`
	Oauth       *FluffyOauth       `json:"oauth,omitempty"`
	TLS         *TentacledTLS      `json:"tls,omitempty"`
}

type FluffyBasic struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type FluffyBearer struct {
	Token string `json:"token"`
}

type FluffyCertificate struct {
	CA         *CunningCA   `json:"ca"`
	CERT       *CunningCERT `json:"cert"`
	Key        *CunningKey  `json:"key"`
	Passphrase *string      `json:"passphrase,omitempty"`
}

type FluffyCA struct {
	File string `json:"file"`
}

type FluffyCERT struct {
	File string `json:"file"`
}

type FluffyKey struct {
	File string `json:"file"`
}

type FluffyOauth struct {
	Audience     *string `json:"audience,omitempty"`
	ClientID     string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	Endpoint     string  `json:"endpoint"`
}

type TentacledTLS struct {
	CERTChain  *BraggadociousCERTChain  `json:"certChain"`
	PrivateKey *BraggadociousPrivateKey `json:"privateKey"`
	RootCerts  *BraggadociousRootCerts  `json:"rootCerts"`
}

type TentacledCERTChain struct {
	File string `json:"file"`
}

type TentacledPrivateKey struct {
	File string `json:"file"`
}

type TentacledRootCerts struct {
	File string `json:"file"`
}

type FluffyBody struct {
	File string `json:"file"`
}

type HTTPStepCapture struct {
	Body     *bool   `json:"body,omitempty"`
	Cookie   *string `json:"cookie,omitempty"`
	Header   *string `json:"header,omitempty"`
	Jsonpath *string `json:"jsonpath,omitempty"`
	Regex    *string `json:"regex,omitempty"`
	Selector *string `json:"selector,omitempty"`
	Xpath    *string `json:"xpath,omitempty"`
}

type HTTPStepCheck struct {
	Body        *CunningBody                   `json:"body"`
	BodySize    *IndigoBodySize                `json:"bodySize"`
	Captures    map[string]interface{}         `json:"captures,omitempty"`
	Co2         *CunningCo2                    `json:"co2"`
	Cookies     map[string]*IndigoCooky        `json:"cookies,omitempty"`
	Headers     map[string]*IndigoHeader       `json:"headers,omitempty"`
	JSON        map[string]interface{}         `json:"json,omitempty"`
	Jsonpath    map[string]interface{}         `json:"jsonpath,omitempty"`
	Md5         *string                        `json:"md5,omitempty"`
	Performance map[string]*CunningPerformance `json:"performance,omitempty"`
	Redirected  *bool                          `json:"redirected,omitempty"`
	Redirects   []string                       `json:"redirects,omitempty"`
	RequestSize *IndigoRequestSize             `json:"requestSize"`
	Schema      map[string]interface{}         `json:"schema,omitempty"`
	Selectors   map[string]*IndigoSelector     `json:"selectors,omitempty"`
	Sha256      *string                        `json:"sha256,omitempty"`
	Size        *CunningSize                   `json:"size"`
	SSL         *FluffySSL                     `json:"ssl,omitempty"`
	Status      *IndigoStatus                  `json:"status"`
	StatusText  *IndigoStatusText              `json:"statusText"`
	Xpath       map[string]*IndigoXpath        `json:"xpath,omitempty"`
}

type TentacledBody struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FluffyBodySize struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledCo2 struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FluffyCooky struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FluffyHeader struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledPerformance struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FluffyRequestSize struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FluffySSL struct {
	DaysUntilExpiration *IndigoDaysUntilExpiration `json:"daysUntilExpiration"`
	Signed              *bool                      `json:"signed,omitempty"`
	Valid               *bool                      `json:"valid,omitempty"`
}

type FluffyDaysUntilExpiration struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FluffySelector struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledSize struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FluffyStatus struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FluffyStatusText struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FluffyXpath struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FormDatumClass struct {
	File string `json:"file"`
}

type Graphql struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type Trpc struct {
	Mutation map[string]map[string]interface{} `json:"mutation,omitempty"`
	Query    *Query                            `json:"query"`
}

type PurplePlugin struct {
	Check  map[string]interface{} `json:"check,omitempty"`
	ID     string                 `json:"id"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type PurpleRetries struct {
	Count    float64  `json:"count"`
	Interval *Timeout `json:"interval"`
}

type PurpleSSE struct {
	Auth    *FluffyAuth            `json:"auth,omitempty"`
	Check   *FluffyCheck           `json:"check,omitempty"`
	Headers map[string]string      `json:"headers,omitempty"`
	JSON    map[string]interface{} `json:"json,omitempty"`
	Params  map[string]string      `json:"params,omitempty"`
	Timeout *float64               `json:"timeout,omitempty"`
	URL     string                 `json:"url"`
}

type FluffyAuth struct {
	Basic       *TentacledBasic       `json:"basic,omitempty"`
	Bearer      *TentacledBearer      `json:"bearer,omitempty"`
	Certificate *TentacledCertificate `json:"certificate,omitempty"`
	Oauth       *TentacledOauth       `json:"oauth,omitempty"`
	TLS         *StickyTLS            `json:"tls,omitempty"`
}

type TentacledBasic struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type TentacledBearer struct {
	Token string `json:"token"`
}

type TentacledCertificate struct {
	CA         *MagentaCA   `json:"ca"`
	CERT       *MagentaCERT `json:"cert"`
	Key        *MagentaKey  `json:"key"`
	Passphrase *string      `json:"passphrase,omitempty"`
}

type TentacledCA struct {
	File string `json:"file"`
}

type TentacledCERT struct {
	File string `json:"file"`
}

type TentacledKey struct {
	File string `json:"file"`
}

type TentacledOauth struct {
	Audience     *string `json:"audience,omitempty"`
	ClientID     string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	Endpoint     string  `json:"endpoint"`
}

type StickyTLS struct {
	CERTChain  *CERTChain1  `json:"certChain"`
	PrivateKey *PrivateKey1 `json:"privateKey"`
	RootCerts  *RootCerts1  `json:"rootCerts"`
}

type StickyCERTChain struct {
	File string `json:"file"`
}

type StickyPrivateKey struct {
	File string `json:"file"`
}

type StickyRootCerts struct {
	File string `json:"file"`
}

type FluffyCheck struct {
	Messages []PurpleMessage `json:"messages,omitempty"`
}

type PurpleMessage struct {
	Body     *MagentaBody           `json:"body"`
	ID       string                 `json:"id"`
	JSON     map[string]interface{} `json:"json,omitempty"`
	Jsonpath map[string]interface{} `json:"jsonpath,omitempty"`
	Schema   map[string]interface{} `json:"schema,omitempty"`
}

type StickyBody struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TRPCStep struct {
	Mutation        map[string]map[string]interface{} `json:"mutation,omitempty"`
	Query           *Query                            `json:"query"`
	Auth            *TRPCStepAuth                     `json:"auth,omitempty"`
	Captures        map[string]TRPCStepCapture        `json:"captures,omitempty"`
	Check           *TRPCStepCheck                    `json:"check,omitempty"`
	Cookies         map[string]string                 `json:"cookies,omitempty"`
	FollowRedirects *bool                             `json:"followRedirects,omitempty"`
	Headers         map[string]string                 `json:"headers,omitempty"`
	Method          string                            `json:"method"`
	Params          map[string]string                 `json:"params,omitempty"`
	Retries         *float64                          `json:"retries,omitempty"`
	Timeout         *Timeout                          `json:"timeout"`
	URL             string                            `json:"url"`
}

type TRPCStepAuth struct {
	Basic       *StickyBasic       `json:"basic,omitempty"`
	Bearer      *StickyBearer      `json:"bearer,omitempty"`
	Certificate *StickyCertificate `json:"certificate,omitempty"`
	Oauth       *StickyOauth       `json:"oauth,omitempty"`
	TLS         *IndigoTLS         `json:"tls,omitempty"`
}

type StickyBasic struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type StickyBearer struct {
	Token string `json:"token"`
}

type StickyCertificate struct {
	CA         *FriskyCA   `json:"ca"`
	CERT       *FriskyCERT `json:"cert"`
	Key        *FriskyKey  `json:"key"`
	Passphrase *string     `json:"passphrase,omitempty"`
}

type StickyCA struct {
	File string `json:"file"`
}

type StickyCERT struct {
	File string `json:"file"`
}

type StickyKey struct {
	File string `json:"file"`
}

type StickyOauth struct {
	Audience     *string `json:"audience,omitempty"`
	ClientID     string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	Endpoint     string  `json:"endpoint"`
}

type IndigoTLS struct {
	CERTChain  *CERTChain2  `json:"certChain"`
	PrivateKey *PrivateKey2 `json:"privateKey"`
	RootCerts  *RootCerts2  `json:"rootCerts"`
}

type IndigoCERTChain struct {
	File string `json:"file"`
}

type IndigoPrivateKey struct {
	File string `json:"file"`
}

type IndigoRootCerts struct {
	File string `json:"file"`
}

type TRPCStepCapture struct {
	Body     *bool   `json:"body,omitempty"`
	Cookie   *string `json:"cookie,omitempty"`
	Header   *string `json:"header,omitempty"`
	Jsonpath *string `json:"jsonpath,omitempty"`
	Regex    *string `json:"regex,omitempty"`
	Selector *string `json:"selector,omitempty"`
	Xpath    *string `json:"xpath,omitempty"`
}

type TRPCStepCheck struct {
	Body        *FriskyBody                    `json:"body"`
	BodySize    *IndecentBodySize              `json:"bodySize"`
	Captures    map[string]interface{}         `json:"captures,omitempty"`
	Co2         *MagentaCo2                    `json:"co2"`
	Cookies     map[string]*IndecentCooky      `json:"cookies,omitempty"`
	Headers     map[string]*IndecentHeader     `json:"headers,omitempty"`
	JSON        map[string]interface{}         `json:"json,omitempty"`
	Jsonpath    map[string]interface{}         `json:"jsonpath,omitempty"`
	Md5         *string                        `json:"md5,omitempty"`
	Performance map[string]*MagentaPerformance `json:"performance,omitempty"`
	Redirected  *bool                          `json:"redirected,omitempty"`
	Redirects   []string                       `json:"redirects,omitempty"`
	RequestSize *IndecentRequestSize           `json:"requestSize"`
	Schema      map[string]interface{}         `json:"schema,omitempty"`
	Selectors   map[string]*IndecentSelector   `json:"selectors,omitempty"`
	Sha256      *string                        `json:"sha256,omitempty"`
	Size        *MagentaSize                   `json:"size"`
	SSL         *TentacledSSL                  `json:"ssl,omitempty"`
	Status      *IndecentStatus                `json:"status"`
	StatusText  *IndecentStatusText            `json:"statusText"`
	Xpath       map[string]*IndecentXpath      `json:"xpath,omitempty"`
}

type IndigoBody struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledBodySize struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type StickyCo2 struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledCooky struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledHeader struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type StickyPerformance struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledRequestSize struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledSSL struct {
	DaysUntilExpiration *IndecentDaysUntilExpiration `json:"daysUntilExpiration"`
	Signed              *bool                        `json:"signed,omitempty"`
	Valid               *bool                        `json:"valid,omitempty"`
}

type TentacledDaysUntilExpiration struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledSelector struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type StickySize struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledStatus struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledStatusText struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledXpath struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type AfterTestdata struct {
	Content *string        `json:"content,omitempty"`
	File    *string        `json:"file,omitempty"`
	Options *PurpleOptions `json:"options,omitempty"`
}

type PurpleOptions struct {
	Delimiter    *string  `json:"delimiter,omitempty"`
	Escape       *string  `json:"escape,omitempty"`
	Headers      *Headers `json:"headers"`
	Quote        *string  `json:"quote,omitempty"`
	WorkflowPath *string  `json:"workflowPath,omitempty"`
}

type Before struct {
	Env      map[string]interface{} `json:"env,omitempty"`
	Name     *string                `json:"name,omitempty"`
	Steps    []BeforeStep           `json:"steps"`
	Testdata *BeforeTestdata        `json:"testdata,omitempty"`
}

type BeforeStep struct {
	Delay   *string        `json:"delay,omitempty"`
	Graphql *GraphQLStep   `json:"graphql,omitempty"`
	Grpc    *FluffyGrpc    `json:"grpc,omitempty"`
	HTTP    *HTTPStep      `json:"http,omitempty"`
	ID      *string        `json:"id,omitempty"`
	If      *string        `json:"if,omitempty"`
	Name    *string        `json:"name,omitempty"`
	Plugin  *FluffyPlugin  `json:"plugin,omitempty"`
	Retries *FluffyRetries `json:"retries,omitempty"`
	SSE     *FluffySSE     `json:"sse,omitempty"`
	Trpc    *TRPCStep      `json:"trpc,omitempty"`
}

type FluffyGrpc struct {
	Auth     *TentacledAuth              `json:"auth,omitempty"`
	Captures map[string]FluffyCapture    `json:"captures,omitempty"`
	Check    *TentacledCheck             `json:"check,omitempty"`
	Data     *Data                       `json:"data"`
	Host     string                      `json:"host"`
	Metadata map[string]*IndigoMetadatum `json:"metadata,omitempty"`
	Method   string                      `json:"method"`
	Proto    *Proto                      `json:"proto"`
	Service  string                      `json:"service"`
	Timeout  *Timeout                    `json:"timeout"`
}

type TentacledAuth struct {
	TLS *IndecentTLS `json:"tls,omitempty"`
}

type IndecentTLS struct {
	CERTChain  *CERTChain3  `json:"certChain"`
	PrivateKey *PrivateKey3 `json:"privateKey"`
	RootCerts  *RootCerts3  `json:"rootCerts"`
}

type IndecentCERTChain struct {
	File string `json:"file"`
}

type IndecentPrivateKey struct {
	File string `json:"file"`
}

type IndecentRootCerts struct {
	File string `json:"file"`
}

type FluffyCapture struct {
	Jsonpath *string `json:"jsonpath,omitempty"`
}

type TentacledCheck struct {
	Captures    map[string]interface{}        `json:"captures,omitempty"`
	Co2         *FriskyCo2                    `json:"co2"`
	JSON        map[string]interface{}        `json:"json,omitempty"`
	Jsonpath    map[string]interface{}        `json:"jsonpath,omitempty"`
	Performance map[string]*FriskyPerformance `json:"performance,omitempty"`
	Schema      map[string]interface{}        `json:"schema,omitempty"`
	Size        *FriskySize                   `json:"size"`
}

type IndigoCo2 struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type IndigoPerformance struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type IndigoSize struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type FluffyMetadatum struct {
	ToStringTag23   MetadatumToStringTag23 `json:"__@toStringTag@23"`
	Buffer          ArrayBufferLike        `json:"buffer"`
	ByteLength      float64                `json:"byteLength"`
	ByteOffset      float64                `json:"byteOffset"`
	BytesPerElement float64                `json:"BYTES_PER_ELEMENT"`
	Length          float64                `json:"length"`
}

type FluffyPlugin struct {
	Check  map[string]interface{} `json:"check,omitempty"`
	ID     string                 `json:"id"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type FluffyRetries struct {
	Count    float64  `json:"count"`
	Interval *Timeout `json:"interval"`
}

type FluffySSE struct {
	Auth    *StickyAuth            `json:"auth,omitempty"`
	Check   *StickyCheck           `json:"check,omitempty"`
	Headers map[string]string      `json:"headers,omitempty"`
	JSON    map[string]interface{} `json:"json,omitempty"`
	Params  map[string]string      `json:"params,omitempty"`
	Timeout *float64               `json:"timeout,omitempty"`
	URL     string                 `json:"url"`
}

type StickyAuth struct {
	Basic       *IndigoBasic       `json:"basic,omitempty"`
	Bearer      *IndigoBearer      `json:"bearer,omitempty"`
	Certificate *IndigoCertificate `json:"certificate,omitempty"`
	Oauth       *IndigoOauth       `json:"oauth,omitempty"`
	TLS         *HilariousTLS      `json:"tls,omitempty"`
}

type IndigoBasic struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type IndigoBearer struct {
	Token string `json:"token"`
}

type IndigoCertificate struct {
	CA         *MischievousCA   `json:"ca"`
	CERT       *MischievousCERT `json:"cert"`
	Key        *MischievousKey  `json:"key"`
	Passphrase *string          `json:"passphrase,omitempty"`
}

type IndigoCA struct {
	File string `json:"file"`
}

type IndigoCERT struct {
	File string `json:"file"`
}

type IndigoKey struct {
	File string `json:"file"`
}

type IndigoOauth struct {
	Audience     *string `json:"audience,omitempty"`
	ClientID     string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	Endpoint     string  `json:"endpoint"`
}

type HilariousTLS struct {
	CERTChain  *CERTChain4  `json:"certChain"`
	PrivateKey *PrivateKey4 `json:"privateKey"`
	RootCerts  *RootCerts4  `json:"rootCerts"`
}

type HilariousCERTChain struct {
	File string `json:"file"`
}

type HilariousPrivateKey struct {
	File string `json:"file"`
}

type HilariousRootCerts struct {
	File string `json:"file"`
}

type StickyCheck struct {
	Messages []FluffyMessage `json:"messages,omitempty"`
}

type FluffyMessage struct {
	Body     *MischievousBody       `json:"body"`
	ID       string                 `json:"id"`
	JSON     map[string]interface{} `json:"json,omitempty"`
	Jsonpath map[string]interface{} `json:"jsonpath,omitempty"`
	Schema   map[string]interface{} `json:"schema,omitempty"`
}

type IndecentBody struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type BeforeTestdata struct {
	Content *string        `json:"content,omitempty"`
	File    *string        `json:"file,omitempty"`
	Options *FluffyOptions `json:"options,omitempty"`
}

type FluffyOptions struct {
	Delimiter    *string  `json:"delimiter,omitempty"`
	Escape       *string  `json:"escape,omitempty"`
	Headers      *Headers `json:"headers"`
	Quote        *string  `json:"quote,omitempty"`
	WorkflowPath *string  `json:"workflowPath,omitempty"`
}

type Components struct {
	Credentials map[string]Credential  `json:"credentials,omitempty"`
	Schemas     map[string]interface{} `json:"schemas,omitempty"`
}

type Credential struct {
	Basic       *CredentialBasic       `json:"basic,omitempty"`
	Bearer      *CredentialBearer      `json:"bearer,omitempty"`
	Certificate *CredentialCertificate `json:"certificate,omitempty"`
	Oauth       *CredentialOauth       `json:"oauth,omitempty"`
	TLS         *CredentialTLS         `json:"tls,omitempty"`
}

type CredentialBasic struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type CredentialBearer struct {
	Token string `json:"token"`
}

type CredentialCertificate struct {
	CA         *BraggadociousCA   `json:"ca"`
	CERT       *BraggadociousCERT `json:"cert"`
	Key        *BraggadociousKey  `json:"key"`
	Passphrase *string            `json:"passphrase,omitempty"`
}

type IndecentCA struct {
	File string `json:"file"`
}

type IndecentCERT struct {
	File string `json:"file"`
}

type IndecentKey struct {
	File string `json:"file"`
}

type CredentialOauth struct {
	Audience     *string `json:"audience,omitempty"`
	ClientID     string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	Endpoint     string  `json:"endpoint"`
}

type CredentialTLS struct {
	CERTChain  *CERTChain5  `json:"certChain"`
	PrivateKey *PrivateKey5 `json:"privateKey"`
	RootCerts  *RootCerts5  `json:"rootCerts"`
}

type AmbitiousCERTChain struct {
	File string `json:"file"`
}

type AmbitiousPrivateKey struct {
	File string `json:"file"`
}

type AmbitiousRootCerts struct {
	File string `json:"file"`
}

type Config struct {
	Concurrency    *float64    `json:"concurrency,omitempty"`
	ContinueOnFail *bool       `json:"continueOnFail,omitempty"`
	Grpc           *ConfigGrpc `json:"grpc,omitempty"`
	HTTP           *HTTP       `json:"http,omitempty"`
	LoadTest       *LoadTest   `json:"loadTest,omitempty"`
}

type ConfigGrpc struct {
	Proto *Proto `json:"proto"`
}

type HTTP struct {
	BaseURL            *string `json:"baseURL,omitempty"`
	Http2              *bool   `json:"http2,omitempty"`
	RejectUnauthorized *bool   `json:"rejectUnauthorized,omitempty"`
}

type LoadTest struct {
	Check  *LoadTestCheck `json:"check,omitempty"`
	Phases []Phase        `json:"phases"`
}

type LoadTestCheck struct {
	Avg *AvgUnion `json:"avg"`
	Max *MaxUnion `json:"max"`
	Med *MedUnion `json:"med"`
	Min *MinUnion `json:"min"`
	P95 *P95Union `json:"p95"`
	P99 *P99Union `json:"p99"`
}

type AvgElement struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type MaxElement struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type MedElement struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type MinElement struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type P95Element struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type P99Element struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type Phase struct {
	ArrivalRate float64 `json:"arrivalRate"`
	Duration    float64 `json:"duration"`
}

type Test struct {
	Env      map[string]interface{} `json:"env,omitempty"`
	Name     *string                `json:"name,omitempty"`
	Steps    []TestStep             `json:"steps"`
	Testdata *TestTestdata          `json:"testdata,omitempty"`
}

type TestStep struct {
	Delay   *string           `json:"delay,omitempty"`
	Graphql *GraphQLStep      `json:"graphql,omitempty"`
	Grpc    *TentacledGrpc    `json:"grpc,omitempty"`
	HTTP    *HTTPStep         `json:"http,omitempty"`
	ID      *string           `json:"id,omitempty"`
	If      *string           `json:"if,omitempty"`
	Name    *string           `json:"name,omitempty"`
	Plugin  *TentacledPlugin  `json:"plugin,omitempty"`
	Retries *TentacledRetries `json:"retries,omitempty"`
	SSE     *TentacledSSE     `json:"sse,omitempty"`
	Trpc    *TRPCStep         `json:"trpc,omitempty"`
}

type TentacledGrpc struct {
	Auth     *IndigoAuth                   `json:"auth,omitempty"`
	Captures map[string]TentacledCapture   `json:"captures,omitempty"`
	Check    *IndigoCheck                  `json:"check,omitempty"`
	Data     *Data                         `json:"data"`
	Host     string                        `json:"host"`
	Metadata map[string]*IndecentMetadatum `json:"metadata,omitempty"`
	Method   string                        `json:"method"`
	Proto    *Proto                        `json:"proto"`
	Service  string                        `json:"service"`
	Timeout  *Timeout                      `json:"timeout"`
}

type IndigoAuth struct {
	TLS *AmbitiousTLS `json:"tls,omitempty"`
}

type AmbitiousTLS struct {
	CERTChain  *CERTChain6  `json:"certChain"`
	PrivateKey *PrivateKey6 `json:"privateKey"`
	RootCerts  *RootCerts6  `json:"rootCerts"`
}

type CunningCERTChain struct {
	File string `json:"file"`
}

type CunningPrivateKey struct {
	File string `json:"file"`
}

type CunningRootCerts struct {
	File string `json:"file"`
}

type TentacledCapture struct {
	Jsonpath *string `json:"jsonpath,omitempty"`
}

type IndigoCheck struct {
	Captures    map[string]interface{}             `json:"captures,omitempty"`
	Co2         *MischievousCo2                    `json:"co2"`
	JSON        map[string]interface{}             `json:"json,omitempty"`
	Jsonpath    map[string]interface{}             `json:"jsonpath,omitempty"`
	Performance map[string]*MischievousPerformance `json:"performance,omitempty"`
	Schema      map[string]interface{}             `json:"schema,omitempty"`
	Size        *MischievousSize                   `json:"size"`
}

type IndecentCo2 struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type IndecentPerformance struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type IndecentSize struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TentacledMetadatum struct {
	ToStringTag23   MetadatumToStringTag23 `json:"__@toStringTag@23"`
	Buffer          ArrayBufferLike        `json:"buffer"`
	ByteLength      float64                `json:"byteLength"`
	ByteOffset      float64                `json:"byteOffset"`
	BytesPerElement float64                `json:"BYTES_PER_ELEMENT"`
	Length          float64                `json:"length"`
}

type TentacledPlugin struct {
	Check  map[string]interface{} `json:"check,omitempty"`
	ID     string                 `json:"id"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type TentacledRetries struct {
	Count    float64  `json:"count"`
	Interval *Timeout `json:"interval"`
}

type TentacledSSE struct {
	Auth    *IndecentAuth          `json:"auth,omitempty"`
	Check   *IndecentCheck         `json:"check,omitempty"`
	Headers map[string]string      `json:"headers,omitempty"`
	JSON    map[string]interface{} `json:"json,omitempty"`
	Params  map[string]string      `json:"params,omitempty"`
	Timeout *float64               `json:"timeout,omitempty"`
	URL     string                 `json:"url"`
}

type IndecentAuth struct {
	Basic       *IndecentBasic       `json:"basic,omitempty"`
	Bearer      *IndecentBearer      `json:"bearer,omitempty"`
	Certificate *IndecentCertificate `json:"certificate,omitempty"`
	Oauth       *IndecentOauth       `json:"oauth,omitempty"`
	TLS         *CunningTLS          `json:"tls,omitempty"`
}

type IndecentBasic struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type IndecentBearer struct {
	Token string `json:"token"`
}

type IndecentCertificate struct {
	CA         *CA1    `json:"ca"`
	CERT       *CERT1  `json:"cert"`
	Key        *Key1   `json:"key"`
	Passphrase *string `json:"passphrase,omitempty"`
}

type HilariousCA struct {
	File string `json:"file"`
}

type HilariousCERT struct {
	File string `json:"file"`
}

type HilariousKey struct {
	File string `json:"file"`
}

type IndecentOauth struct {
	Audience     *string `json:"audience,omitempty"`
	ClientID     string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	Endpoint     string  `json:"endpoint"`
}

type CunningTLS struct {
	CERTChain  *CERTChain7  `json:"certChain"`
	PrivateKey *PrivateKey7 `json:"privateKey"`
	RootCerts  *RootCerts7  `json:"rootCerts"`
}

type MagentaCERTChain struct {
	File string `json:"file"`
}

type MagentaPrivateKey struct {
	File string `json:"file"`
}

type MagentaRootCerts struct {
	File string `json:"file"`
}

type IndecentCheck struct {
	Messages []TentacledMessage `json:"messages,omitempty"`
}

type TentacledMessage struct {
	Body     *BraggadociousBody     `json:"body"`
	ID       string                 `json:"id"`
	JSON     map[string]interface{} `json:"json,omitempty"`
	Jsonpath map[string]interface{} `json:"jsonpath,omitempty"`
	Schema   map[string]interface{} `json:"schema,omitempty"`
}

type HilariousBody struct {
	Eq        interface{}            `json:"eq"`
	Gt        *float64               `json:"gt,omitempty"`
	Gte       *float64               `json:"gte,omitempty"`
	In        map[string]interface{} `json:"in,omitempty"`
	IsArray   *bool                  `json:"isArray,omitempty"`
	IsBoolean *bool                  `json:"isBoolean,omitempty"`
	IsDefined *bool                  `json:"isDefined,omitempty"`
	IsNull    *bool                  `json:"isNull,omitempty"`
	IsNumber  *bool                  `json:"isNumber,omitempty"`
	IsObject  *bool                  `json:"isObject,omitempty"`
	IsString  *bool                  `json:"isString,omitempty"`
	Lt        *float64               `json:"lt,omitempty"`
	LTE       *float64               `json:"lte,omitempty"`
	Match     *string                `json:"match,omitempty"`
	Ne        interface{}            `json:"ne"`
	Nin       map[string]interface{} `json:"nin,omitempty"`
}

type TestTestdata struct {
	Content *string           `json:"content,omitempty"`
	File    *string           `json:"file,omitempty"`
	Options *TentacledOptions `json:"options,omitempty"`
}

type TentacledOptions struct {
	Delimiter    *string  `json:"delimiter,omitempty"`
	Escape       *string  `json:"escape,omitempty"`
	Headers      *Headers `json:"headers"`
	Quote        *string  `json:"quote,omitempty"`
	WorkflowPath *string  `json:"workflowPath,omitempty"`
}

type Species469___ToStringTag23 string

const (
	SharedArrayBuffer Species469___ToStringTag23 = "SharedArrayBuffer"
)

type MetadatumToStringTag23 string

const (
	Uint8Array MetadatumToStringTag23 = "Uint8Array"
)

type AmbitiousCA struct {
	PurpleCA *PurpleCA
	String   *string
}

func (x *AmbitiousCA) UnmarshalJSON(data []byte) error {
	x.PurpleCA = nil
	var c PurpleCA
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleCA = &c
	}
	return nil
}

func (x *AmbitiousCA) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleCA != nil, x.PurpleCA, false, nil, false, nil, false)
}

type AmbitiousCERT struct {
	PurpleCERT *PurpleCERT
	String     *string
}

func (x *AmbitiousCERT) UnmarshalJSON(data []byte) error {
	x.PurpleCERT = nil
	var c PurpleCERT
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleCERT = &c
	}
	return nil
}

func (x *AmbitiousCERT) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleCERT != nil, x.PurpleCERT, false, nil, false, nil, false)
}

type AmbitiousKey struct {
	PurpleKey *PurpleKey
	String    *string
}

func (x *AmbitiousKey) UnmarshalJSON(data []byte) error {
	x.PurpleKey = nil
	var c PurpleKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleKey = &c
	}
	return nil
}

func (x *AmbitiousKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleKey != nil, x.PurpleKey, false, nil, false, nil, false)
}

type FriskyCERTChain struct {
	PurpleCERTChain *PurpleCERTChain
	String          *string
}

func (x *FriskyCERTChain) UnmarshalJSON(data []byte) error {
	x.PurpleCERTChain = nil
	var c PurpleCERTChain
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleCERTChain = &c
	}
	return nil
}

func (x *FriskyCERTChain) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleCERTChain != nil, x.PurpleCERTChain, false, nil, false, nil, false)
}

type FriskyPrivateKey struct {
	PurplePrivateKey *PurplePrivateKey
	String           *string
}

func (x *FriskyPrivateKey) UnmarshalJSON(data []byte) error {
	x.PurplePrivateKey = nil
	var c PurplePrivateKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurplePrivateKey = &c
	}
	return nil
}

func (x *FriskyPrivateKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurplePrivateKey != nil, x.PurplePrivateKey, false, nil, false, nil, false)
}

type FriskyRootCerts struct {
	PurpleRootCerts *PurpleRootCerts
	String          *string
}

func (x *FriskyRootCerts) UnmarshalJSON(data []byte) error {
	x.PurpleRootCerts = nil
	var c PurpleRootCerts
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleRootCerts = &c
	}
	return nil
}

func (x *FriskyRootCerts) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleRootCerts != nil, x.PurpleRootCerts, false, nil, false, nil, false)
}

type AmbitiousBody struct {
	PurpleBodyArray []PurpleBody
	String          *string
}

func (x *AmbitiousBody) UnmarshalJSON(data []byte) error {
	x.PurpleBodyArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.PurpleBodyArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *AmbitiousBody) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.PurpleBodyArray != nil, x.PurpleBodyArray, false, nil, false, nil, false, nil, false)
}

type StickyBodySize struct {
	Double              *float64
	PurpleBodySizeArray []PurpleBodySize
}

func (x *StickyBodySize) UnmarshalJSON(data []byte) error {
	x.PurpleBodySizeArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.PurpleBodySizeArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StickyBodySize) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.PurpleBodySizeArray != nil, x.PurpleBodySizeArray, false, nil, false, nil, false, nil, false)
}

type HilariousCo2 struct {
	Double         *float64
	PurpleCo2Array []PurpleCo2
}

func (x *HilariousCo2) UnmarshalJSON(data []byte) error {
	x.PurpleCo2Array = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.PurpleCo2Array, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *HilariousCo2) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.PurpleCo2Array != nil, x.PurpleCo2Array, false, nil, false, nil, false, nil, false)
}

type StickyCooky struct {
	PurpleCookyArray []PurpleCooky
	String           *string
}

func (x *StickyCooky) UnmarshalJSON(data []byte) error {
	x.PurpleCookyArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.PurpleCookyArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StickyCooky) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.PurpleCookyArray != nil, x.PurpleCookyArray, false, nil, false, nil, false, nil, false)
}

type StickyHeader struct {
	PurpleHeaderArray []PurpleHeader
	String            *string
}

func (x *StickyHeader) UnmarshalJSON(data []byte) error {
	x.PurpleHeaderArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.PurpleHeaderArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StickyHeader) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.PurpleHeaderArray != nil, x.PurpleHeaderArray, false, nil, false, nil, false, nil, false)
}

type HilariousPerformance struct {
	Double                 *float64
	PurplePerformanceArray []PurplePerformance
}

func (x *HilariousPerformance) UnmarshalJSON(data []byte) error {
	x.PurplePerformanceArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.PurplePerformanceArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *HilariousPerformance) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.PurplePerformanceArray != nil, x.PurplePerformanceArray, false, nil, false, nil, false, nil, false)
}

type StickyRequestSize struct {
	Double                 *float64
	PurpleRequestSizeArray []PurpleRequestSize
}

func (x *StickyRequestSize) UnmarshalJSON(data []byte) error {
	x.PurpleRequestSizeArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.PurpleRequestSizeArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StickyRequestSize) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.PurpleRequestSizeArray != nil, x.PurpleRequestSizeArray, false, nil, false, nil, false, nil, false)
}

type StickyDaysUntilExpiration struct {
	Double                         *float64
	PurpleDaysUntilExpirationArray []PurpleDaysUntilExpiration
}

func (x *StickyDaysUntilExpiration) UnmarshalJSON(data []byte) error {
	x.PurpleDaysUntilExpirationArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.PurpleDaysUntilExpirationArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StickyDaysUntilExpiration) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.PurpleDaysUntilExpirationArray != nil, x.PurpleDaysUntilExpirationArray, false, nil, false, nil, false, nil, false)
}

type StickySelector struct {
	PurpleSelectorArray []PurpleSelector
	String              *string
}

func (x *StickySelector) UnmarshalJSON(data []byte) error {
	x.PurpleSelectorArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.PurpleSelectorArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StickySelector) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.PurpleSelectorArray != nil, x.PurpleSelectorArray, false, nil, false, nil, false, nil, false)
}

type HilariousSize struct {
	Double          *float64
	PurpleSizeArray []PurpleSize
}

func (x *HilariousSize) UnmarshalJSON(data []byte) error {
	x.PurpleSizeArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.PurpleSizeArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *HilariousSize) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.PurpleSizeArray != nil, x.PurpleSizeArray, false, nil, false, nil, false, nil, false)
}

type StickyStatus struct {
	Double            *float64
	PurpleStatusArray []PurpleStatus
	String            *string
}

func (x *StickyStatus) UnmarshalJSON(data []byte) error {
	x.PurpleStatusArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, &x.String, true, &x.PurpleStatusArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StickyStatus) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, x.String, x.PurpleStatusArray != nil, x.PurpleStatusArray, false, nil, false, nil, false, nil, false)
}

type StickyStatusText struct {
	PurpleStatusTextArray []PurpleStatusText
	String                *string
}

func (x *StickyStatusText) UnmarshalJSON(data []byte) error {
	x.PurpleStatusTextArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.PurpleStatusTextArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StickyStatusText) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.PurpleStatusTextArray != nil, x.PurpleStatusTextArray, false, nil, false, nil, false, nil, false)
}

type StickyXpath struct {
	PurpleXpathArray []PurpleXpath
	String           *string
}

func (x *StickyXpath) UnmarshalJSON(data []byte) error {
	x.PurpleXpathArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.PurpleXpathArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StickyXpath) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.PurpleXpathArray != nil, x.PurpleXpathArray, false, nil, false, nil, false, nil, false)
}

type Timeout struct {
	Double *float64
	String *string
}

func (x *Timeout) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, &x.Double, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Timeout) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

type MischievousCERTChain struct {
	FluffyCERTChain *FluffyCERTChain
	String          *string
}

func (x *MischievousCERTChain) UnmarshalJSON(data []byte) error {
	x.FluffyCERTChain = nil
	var c FluffyCERTChain
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyCERTChain = &c
	}
	return nil
}

func (x *MischievousCERTChain) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyCERTChain != nil, x.FluffyCERTChain, false, nil, false, nil, false)
}

type MischievousPrivateKey struct {
	FluffyPrivateKey *FluffyPrivateKey
	String           *string
}

func (x *MischievousPrivateKey) UnmarshalJSON(data []byte) error {
	x.FluffyPrivateKey = nil
	var c FluffyPrivateKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyPrivateKey = &c
	}
	return nil
}

func (x *MischievousPrivateKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyPrivateKey != nil, x.FluffyPrivateKey, false, nil, false, nil, false)
}

type MischievousRootCerts struct {
	FluffyRootCerts *FluffyRootCerts
	String          *string
}

func (x *MischievousRootCerts) UnmarshalJSON(data []byte) error {
	x.FluffyRootCerts = nil
	var c FluffyRootCerts
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyRootCerts = &c
	}
	return nil
}

func (x *MischievousRootCerts) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyRootCerts != nil, x.FluffyRootCerts, false, nil, false, nil, false)
}

type AmbitiousCo2 struct {
	Double         *float64
	FluffyCo2Array []FluffyCo2
}

func (x *AmbitiousCo2) UnmarshalJSON(data []byte) error {
	x.FluffyCo2Array = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.FluffyCo2Array, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *AmbitiousCo2) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.FluffyCo2Array != nil, x.FluffyCo2Array, false, nil, false, nil, false, nil, false)
}

type AmbitiousPerformance struct {
	Double                 *float64
	FluffyPerformanceArray []FluffyPerformance
}

func (x *AmbitiousPerformance) UnmarshalJSON(data []byte) error {
	x.FluffyPerformanceArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.FluffyPerformanceArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *AmbitiousPerformance) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.FluffyPerformanceArray != nil, x.FluffyPerformanceArray, false, nil, false, nil, false, nil, false)
}

type AmbitiousSize struct {
	Double          *float64
	FluffySizeArray []FluffySize
}

func (x *AmbitiousSize) UnmarshalJSON(data []byte) error {
	x.FluffySizeArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.FluffySizeArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *AmbitiousSize) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.FluffySizeArray != nil, x.FluffySizeArray, false, nil, false, nil, false, nil, false)
}

type Data struct {
	AnythingMap      map[string]interface{}
	AnythingMapArray []map[string]interface{}
}

func (x *Data) UnmarshalJSON(data []byte) error {
	x.AnythingMapArray = nil
	x.AnythingMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.AnythingMapArray, false, nil, true, &x.AnythingMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Data) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.AnythingMapArray != nil, x.AnythingMapArray, false, nil, x.AnythingMap != nil, x.AnythingMap, false, nil, false)
}

type StickyMetadatum struct {
	PurpleMetadatum *PurpleMetadatum
	String          *string
}

func (x *StickyMetadatum) UnmarshalJSON(data []byte) error {
	x.PurpleMetadatum = nil
	var c PurpleMetadatum
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleMetadatum = &c
	}
	return nil
}

func (x *StickyMetadatum) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleMetadatum != nil, x.PurpleMetadatum, false, nil, false, nil, false)
}

type Proto struct {
	String      *string
	StringArray []string
}

func (x *Proto) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Proto) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}

type CunningCA struct {
	FluffyCA *FluffyCA
	String   *string
}

func (x *CunningCA) UnmarshalJSON(data []byte) error {
	x.FluffyCA = nil
	var c FluffyCA
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyCA = &c
	}
	return nil
}

func (x *CunningCA) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyCA != nil, x.FluffyCA, false, nil, false, nil, false)
}

type CunningCERT struct {
	FluffyCERT *FluffyCERT
	String     *string
}

func (x *CunningCERT) UnmarshalJSON(data []byte) error {
	x.FluffyCERT = nil
	var c FluffyCERT
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyCERT = &c
	}
	return nil
}

func (x *CunningCERT) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyCERT != nil, x.FluffyCERT, false, nil, false, nil, false)
}

type CunningKey struct {
	FluffyKey *FluffyKey
	String    *string
}

func (x *CunningKey) UnmarshalJSON(data []byte) error {
	x.FluffyKey = nil
	var c FluffyKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyKey = &c
	}
	return nil
}

func (x *CunningKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyKey != nil, x.FluffyKey, false, nil, false, nil, false)
}

type BraggadociousCERTChain struct {
	String             *string
	TentacledCERTChain *TentacledCERTChain
}

func (x *BraggadociousCERTChain) UnmarshalJSON(data []byte) error {
	x.TentacledCERTChain = nil
	var c TentacledCERTChain
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TentacledCERTChain = &c
	}
	return nil
}

func (x *BraggadociousCERTChain) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TentacledCERTChain != nil, x.TentacledCERTChain, false, nil, false, nil, false)
}

type BraggadociousPrivateKey struct {
	String              *string
	TentacledPrivateKey *TentacledPrivateKey
}

func (x *BraggadociousPrivateKey) UnmarshalJSON(data []byte) error {
	x.TentacledPrivateKey = nil
	var c TentacledPrivateKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TentacledPrivateKey = &c
	}
	return nil
}

func (x *BraggadociousPrivateKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TentacledPrivateKey != nil, x.TentacledPrivateKey, false, nil, false, nil, false)
}

type BraggadociousRootCerts struct {
	String             *string
	TentacledRootCerts *TentacledRootCerts
}

func (x *BraggadociousRootCerts) UnmarshalJSON(data []byte) error {
	x.TentacledRootCerts = nil
	var c TentacledRootCerts
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TentacledRootCerts = &c
	}
	return nil
}

func (x *BraggadociousRootCerts) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TentacledRootCerts != nil, x.TentacledRootCerts, false, nil, false, nil, false)
}

type HTTPStepBody struct {
	FluffyBody *FluffyBody
	String     *string
}

func (x *HTTPStepBody) UnmarshalJSON(data []byte) error {
	x.FluffyBody = nil
	var c FluffyBody
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyBody = &c
	}
	return nil
}

func (x *HTTPStepBody) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyBody != nil, x.FluffyBody, false, nil, false, nil, false)
}

type CunningBody struct {
	String             *string
	TentacledBodyArray []TentacledBody
}

func (x *CunningBody) UnmarshalJSON(data []byte) error {
	x.TentacledBodyArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.TentacledBodyArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *CunningBody) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.TentacledBodyArray != nil, x.TentacledBodyArray, false, nil, false, nil, false, nil, false)
}

type IndigoBodySize struct {
	Double              *float64
	FluffyBodySizeArray []FluffyBodySize
}

func (x *IndigoBodySize) UnmarshalJSON(data []byte) error {
	x.FluffyBodySizeArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.FluffyBodySizeArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndigoBodySize) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.FluffyBodySizeArray != nil, x.FluffyBodySizeArray, false, nil, false, nil, false, nil, false)
}

type CunningCo2 struct {
	Double            *float64
	TentacledCo2Array []TentacledCo2
}

func (x *CunningCo2) UnmarshalJSON(data []byte) error {
	x.TentacledCo2Array = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.TentacledCo2Array, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *CunningCo2) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.TentacledCo2Array != nil, x.TentacledCo2Array, false, nil, false, nil, false, nil, false)
}

type IndigoCooky struct {
	FluffyCookyArray []FluffyCooky
	String           *string
}

func (x *IndigoCooky) UnmarshalJSON(data []byte) error {
	x.FluffyCookyArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.FluffyCookyArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndigoCooky) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.FluffyCookyArray != nil, x.FluffyCookyArray, false, nil, false, nil, false, nil, false)
}

type IndigoHeader struct {
	FluffyHeaderArray []FluffyHeader
	String            *string
}

func (x *IndigoHeader) UnmarshalJSON(data []byte) error {
	x.FluffyHeaderArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.FluffyHeaderArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndigoHeader) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.FluffyHeaderArray != nil, x.FluffyHeaderArray, false, nil, false, nil, false, nil, false)
}

type CunningPerformance struct {
	Double                    *float64
	TentacledPerformanceArray []TentacledPerformance
}

func (x *CunningPerformance) UnmarshalJSON(data []byte) error {
	x.TentacledPerformanceArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.TentacledPerformanceArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *CunningPerformance) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.TentacledPerformanceArray != nil, x.TentacledPerformanceArray, false, nil, false, nil, false, nil, false)
}

type IndigoRequestSize struct {
	Double                 *float64
	FluffyRequestSizeArray []FluffyRequestSize
}

func (x *IndigoRequestSize) UnmarshalJSON(data []byte) error {
	x.FluffyRequestSizeArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.FluffyRequestSizeArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndigoRequestSize) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.FluffyRequestSizeArray != nil, x.FluffyRequestSizeArray, false, nil, false, nil, false, nil, false)
}

type IndigoDaysUntilExpiration struct {
	Double                         *float64
	FluffyDaysUntilExpirationArray []FluffyDaysUntilExpiration
}

func (x *IndigoDaysUntilExpiration) UnmarshalJSON(data []byte) error {
	x.FluffyDaysUntilExpirationArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.FluffyDaysUntilExpirationArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndigoDaysUntilExpiration) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.FluffyDaysUntilExpirationArray != nil, x.FluffyDaysUntilExpirationArray, false, nil, false, nil, false, nil, false)
}

type IndigoSelector struct {
	FluffySelectorArray []FluffySelector
	String              *string
}

func (x *IndigoSelector) UnmarshalJSON(data []byte) error {
	x.FluffySelectorArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.FluffySelectorArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndigoSelector) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.FluffySelectorArray != nil, x.FluffySelectorArray, false, nil, false, nil, false, nil, false)
}

type CunningSize struct {
	Double             *float64
	TentacledSizeArray []TentacledSize
}

func (x *CunningSize) UnmarshalJSON(data []byte) error {
	x.TentacledSizeArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.TentacledSizeArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *CunningSize) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.TentacledSizeArray != nil, x.TentacledSizeArray, false, nil, false, nil, false, nil, false)
}

type IndigoStatus struct {
	Double            *float64
	FluffyStatusArray []FluffyStatus
	String            *string
}

func (x *IndigoStatus) UnmarshalJSON(data []byte) error {
	x.FluffyStatusArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, &x.String, true, &x.FluffyStatusArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndigoStatus) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, x.String, x.FluffyStatusArray != nil, x.FluffyStatusArray, false, nil, false, nil, false, nil, false)
}

type IndigoStatusText struct {
	FluffyStatusTextArray []FluffyStatusText
	String                *string
}

func (x *IndigoStatusText) UnmarshalJSON(data []byte) error {
	x.FluffyStatusTextArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.FluffyStatusTextArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndigoStatusText) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.FluffyStatusTextArray != nil, x.FluffyStatusTextArray, false, nil, false, nil, false, nil, false)
}

type IndigoXpath struct {
	FluffyXpathArray []FluffyXpath
	String           *string
}

func (x *IndigoXpath) UnmarshalJSON(data []byte) error {
	x.FluffyXpathArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.FluffyXpathArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndigoXpath) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.FluffyXpathArray != nil, x.FluffyXpathArray, false, nil, false, nil, false, nil, false)
}

type FormDatumValue struct {
	FormDatumClass *FormDatumClass
	String         *string
}

func (x *FormDatumValue) UnmarshalJSON(data []byte) error {
	x.FormDatumClass = nil
	var c FormDatumClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FormDatumClass = &c
	}
	return nil
}

func (x *FormDatumValue) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FormDatumClass != nil, x.FormDatumClass, false, nil, false, nil, false)
}

type Query struct {
	AnythingMapMap      map[string]map[string]interface{}
	AnythingMapMapArray []map[string]map[string]interface{}
}

func (x *Query) UnmarshalJSON(data []byte) error {
	x.AnythingMapMapArray = nil
	x.AnythingMapMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.AnythingMapMapArray, false, nil, true, &x.AnythingMapMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Query) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.AnythingMapMapArray != nil, x.AnythingMapMapArray, false, nil, x.AnythingMapMap != nil, x.AnythingMapMap, false, nil, false)
}

type MagentaCA struct {
	String      *string
	TentacledCA *TentacledCA
}

func (x *MagentaCA) UnmarshalJSON(data []byte) error {
	x.TentacledCA = nil
	var c TentacledCA
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TentacledCA = &c
	}
	return nil
}

func (x *MagentaCA) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TentacledCA != nil, x.TentacledCA, false, nil, false, nil, false)
}

type MagentaCERT struct {
	String        *string
	TentacledCERT *TentacledCERT
}

func (x *MagentaCERT) UnmarshalJSON(data []byte) error {
	x.TentacledCERT = nil
	var c TentacledCERT
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TentacledCERT = &c
	}
	return nil
}

func (x *MagentaCERT) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TentacledCERT != nil, x.TentacledCERT, false, nil, false, nil, false)
}

type MagentaKey struct {
	String       *string
	TentacledKey *TentacledKey
}

func (x *MagentaKey) UnmarshalJSON(data []byte) error {
	x.TentacledKey = nil
	var c TentacledKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TentacledKey = &c
	}
	return nil
}

func (x *MagentaKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TentacledKey != nil, x.TentacledKey, false, nil, false, nil, false)
}

type CERTChain1 struct {
	StickyCERTChain *StickyCERTChain
	String          *string
}

func (x *CERTChain1) UnmarshalJSON(data []byte) error {
	x.StickyCERTChain = nil
	var c StickyCERTChain
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.StickyCERTChain = &c
	}
	return nil
}

func (x *CERTChain1) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.StickyCERTChain != nil, x.StickyCERTChain, false, nil, false, nil, false)
}

type PrivateKey1 struct {
	StickyPrivateKey *StickyPrivateKey
	String           *string
}

func (x *PrivateKey1) UnmarshalJSON(data []byte) error {
	x.StickyPrivateKey = nil
	var c StickyPrivateKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.StickyPrivateKey = &c
	}
	return nil
}

func (x *PrivateKey1) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.StickyPrivateKey != nil, x.StickyPrivateKey, false, nil, false, nil, false)
}

type RootCerts1 struct {
	StickyRootCerts *StickyRootCerts
	String          *string
}

func (x *RootCerts1) UnmarshalJSON(data []byte) error {
	x.StickyRootCerts = nil
	var c StickyRootCerts
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.StickyRootCerts = &c
	}
	return nil
}

func (x *RootCerts1) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.StickyRootCerts != nil, x.StickyRootCerts, false, nil, false, nil, false)
}

type MagentaBody struct {
	StickyBodyArray []StickyBody
	String          *string
}

func (x *MagentaBody) UnmarshalJSON(data []byte) error {
	x.StickyBodyArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StickyBodyArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MagentaBody) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StickyBodyArray != nil, x.StickyBodyArray, false, nil, false, nil, false, nil, false)
}

type FriskyCA struct {
	StickyCA *StickyCA
	String   *string
}

func (x *FriskyCA) UnmarshalJSON(data []byte) error {
	x.StickyCA = nil
	var c StickyCA
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.StickyCA = &c
	}
	return nil
}

func (x *FriskyCA) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.StickyCA != nil, x.StickyCA, false, nil, false, nil, false)
}

type FriskyCERT struct {
	StickyCERT *StickyCERT
	String     *string
}

func (x *FriskyCERT) UnmarshalJSON(data []byte) error {
	x.StickyCERT = nil
	var c StickyCERT
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.StickyCERT = &c
	}
	return nil
}

func (x *FriskyCERT) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.StickyCERT != nil, x.StickyCERT, false, nil, false, nil, false)
}

type FriskyKey struct {
	StickyKey *StickyKey
	String    *string
}

func (x *FriskyKey) UnmarshalJSON(data []byte) error {
	x.StickyKey = nil
	var c StickyKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.StickyKey = &c
	}
	return nil
}

func (x *FriskyKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.StickyKey != nil, x.StickyKey, false, nil, false, nil, false)
}

type CERTChain2 struct {
	IndigoCERTChain *IndigoCERTChain
	String          *string
}

func (x *CERTChain2) UnmarshalJSON(data []byte) error {
	x.IndigoCERTChain = nil
	var c IndigoCERTChain
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IndigoCERTChain = &c
	}
	return nil
}

func (x *CERTChain2) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.IndigoCERTChain != nil, x.IndigoCERTChain, false, nil, false, nil, false)
}

type PrivateKey2 struct {
	IndigoPrivateKey *IndigoPrivateKey
	String           *string
}

func (x *PrivateKey2) UnmarshalJSON(data []byte) error {
	x.IndigoPrivateKey = nil
	var c IndigoPrivateKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IndigoPrivateKey = &c
	}
	return nil
}

func (x *PrivateKey2) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.IndigoPrivateKey != nil, x.IndigoPrivateKey, false, nil, false, nil, false)
}

type RootCerts2 struct {
	IndigoRootCerts *IndigoRootCerts
	String          *string
}

func (x *RootCerts2) UnmarshalJSON(data []byte) error {
	x.IndigoRootCerts = nil
	var c IndigoRootCerts
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IndigoRootCerts = &c
	}
	return nil
}

func (x *RootCerts2) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.IndigoRootCerts != nil, x.IndigoRootCerts, false, nil, false, nil, false)
}

type FriskyBody struct {
	IndigoBodyArray []IndigoBody
	String          *string
}

func (x *FriskyBody) UnmarshalJSON(data []byte) error {
	x.IndigoBodyArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.IndigoBodyArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *FriskyBody) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.IndigoBodyArray != nil, x.IndigoBodyArray, false, nil, false, nil, false, nil, false)
}

type IndecentBodySize struct {
	Double                 *float64
	TentacledBodySizeArray []TentacledBodySize
}

func (x *IndecentBodySize) UnmarshalJSON(data []byte) error {
	x.TentacledBodySizeArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.TentacledBodySizeArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndecentBodySize) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.TentacledBodySizeArray != nil, x.TentacledBodySizeArray, false, nil, false, nil, false, nil, false)
}

type MagentaCo2 struct {
	Double         *float64
	StickyCo2Array []StickyCo2
}

func (x *MagentaCo2) UnmarshalJSON(data []byte) error {
	x.StickyCo2Array = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.StickyCo2Array, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MagentaCo2) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.StickyCo2Array != nil, x.StickyCo2Array, false, nil, false, nil, false, nil, false)
}

type IndecentCooky struct {
	String              *string
	TentacledCookyArray []TentacledCooky
}

func (x *IndecentCooky) UnmarshalJSON(data []byte) error {
	x.TentacledCookyArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.TentacledCookyArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndecentCooky) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.TentacledCookyArray != nil, x.TentacledCookyArray, false, nil, false, nil, false, nil, false)
}

type IndecentHeader struct {
	String               *string
	TentacledHeaderArray []TentacledHeader
}

func (x *IndecentHeader) UnmarshalJSON(data []byte) error {
	x.TentacledHeaderArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.TentacledHeaderArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndecentHeader) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.TentacledHeaderArray != nil, x.TentacledHeaderArray, false, nil, false, nil, false, nil, false)
}

type MagentaPerformance struct {
	Double                 *float64
	StickyPerformanceArray []StickyPerformance
}

func (x *MagentaPerformance) UnmarshalJSON(data []byte) error {
	x.StickyPerformanceArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.StickyPerformanceArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MagentaPerformance) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.StickyPerformanceArray != nil, x.StickyPerformanceArray, false, nil, false, nil, false, nil, false)
}

type IndecentRequestSize struct {
	Double                    *float64
	TentacledRequestSizeArray []TentacledRequestSize
}

func (x *IndecentRequestSize) UnmarshalJSON(data []byte) error {
	x.TentacledRequestSizeArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.TentacledRequestSizeArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndecentRequestSize) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.TentacledRequestSizeArray != nil, x.TentacledRequestSizeArray, false, nil, false, nil, false, nil, false)
}

type IndecentDaysUntilExpiration struct {
	Double                            *float64
	TentacledDaysUntilExpirationArray []TentacledDaysUntilExpiration
}

func (x *IndecentDaysUntilExpiration) UnmarshalJSON(data []byte) error {
	x.TentacledDaysUntilExpirationArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.TentacledDaysUntilExpirationArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndecentDaysUntilExpiration) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.TentacledDaysUntilExpirationArray != nil, x.TentacledDaysUntilExpirationArray, false, nil, false, nil, false, nil, false)
}

type IndecentSelector struct {
	String                 *string
	TentacledSelectorArray []TentacledSelector
}

func (x *IndecentSelector) UnmarshalJSON(data []byte) error {
	x.TentacledSelectorArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.TentacledSelectorArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndecentSelector) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.TentacledSelectorArray != nil, x.TentacledSelectorArray, false, nil, false, nil, false, nil, false)
}

type MagentaSize struct {
	Double          *float64
	StickySizeArray []StickySize
}

func (x *MagentaSize) UnmarshalJSON(data []byte) error {
	x.StickySizeArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.StickySizeArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MagentaSize) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.StickySizeArray != nil, x.StickySizeArray, false, nil, false, nil, false, nil, false)
}

type IndecentStatus struct {
	Double               *float64
	String               *string
	TentacledStatusArray []TentacledStatus
}

func (x *IndecentStatus) UnmarshalJSON(data []byte) error {
	x.TentacledStatusArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, &x.String, true, &x.TentacledStatusArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndecentStatus) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, x.String, x.TentacledStatusArray != nil, x.TentacledStatusArray, false, nil, false, nil, false, nil, false)
}

type IndecentStatusText struct {
	String                   *string
	TentacledStatusTextArray []TentacledStatusText
}

func (x *IndecentStatusText) UnmarshalJSON(data []byte) error {
	x.TentacledStatusTextArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.TentacledStatusTextArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndecentStatusText) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.TentacledStatusTextArray != nil, x.TentacledStatusTextArray, false, nil, false, nil, false, nil, false)
}

type IndecentXpath struct {
	String              *string
	TentacledXpathArray []TentacledXpath
}

func (x *IndecentXpath) UnmarshalJSON(data []byte) error {
	x.TentacledXpathArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.TentacledXpathArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *IndecentXpath) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.TentacledXpathArray != nil, x.TentacledXpathArray, false, nil, false, nil, false, nil, false)
}

type Headers struct {
	Bool        *bool
	StringArray []string
}

func (x *Headers) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Headers) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}

type CERTChain3 struct {
	IndecentCERTChain *IndecentCERTChain
	String            *string
}

func (x *CERTChain3) UnmarshalJSON(data []byte) error {
	x.IndecentCERTChain = nil
	var c IndecentCERTChain
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IndecentCERTChain = &c
	}
	return nil
}

func (x *CERTChain3) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.IndecentCERTChain != nil, x.IndecentCERTChain, false, nil, false, nil, false)
}

type PrivateKey3 struct {
	IndecentPrivateKey *IndecentPrivateKey
	String             *string
}

func (x *PrivateKey3) UnmarshalJSON(data []byte) error {
	x.IndecentPrivateKey = nil
	var c IndecentPrivateKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IndecentPrivateKey = &c
	}
	return nil
}

func (x *PrivateKey3) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.IndecentPrivateKey != nil, x.IndecentPrivateKey, false, nil, false, nil, false)
}

type RootCerts3 struct {
	IndecentRootCerts *IndecentRootCerts
	String            *string
}

func (x *RootCerts3) UnmarshalJSON(data []byte) error {
	x.IndecentRootCerts = nil
	var c IndecentRootCerts
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IndecentRootCerts = &c
	}
	return nil
}

func (x *RootCerts3) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.IndecentRootCerts != nil, x.IndecentRootCerts, false, nil, false, nil, false)
}

type FriskyCo2 struct {
	Double         *float64
	IndigoCo2Array []IndigoCo2
}

func (x *FriskyCo2) UnmarshalJSON(data []byte) error {
	x.IndigoCo2Array = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.IndigoCo2Array, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *FriskyCo2) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.IndigoCo2Array != nil, x.IndigoCo2Array, false, nil, false, nil, false, nil, false)
}

type FriskyPerformance struct {
	Double                 *float64
	IndigoPerformanceArray []IndigoPerformance
}

func (x *FriskyPerformance) UnmarshalJSON(data []byte) error {
	x.IndigoPerformanceArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.IndigoPerformanceArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *FriskyPerformance) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.IndigoPerformanceArray != nil, x.IndigoPerformanceArray, false, nil, false, nil, false, nil, false)
}

type FriskySize struct {
	Double          *float64
	IndigoSizeArray []IndigoSize
}

func (x *FriskySize) UnmarshalJSON(data []byte) error {
	x.IndigoSizeArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.IndigoSizeArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *FriskySize) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.IndigoSizeArray != nil, x.IndigoSizeArray, false, nil, false, nil, false, nil, false)
}

type IndigoMetadatum struct {
	FluffyMetadatum *FluffyMetadatum
	String          *string
}

func (x *IndigoMetadatum) UnmarshalJSON(data []byte) error {
	x.FluffyMetadatum = nil
	var c FluffyMetadatum
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyMetadatum = &c
	}
	return nil
}

func (x *IndigoMetadatum) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyMetadatum != nil, x.FluffyMetadatum, false, nil, false, nil, false)
}

type MischievousCA struct {
	IndigoCA *IndigoCA
	String   *string
}

func (x *MischievousCA) UnmarshalJSON(data []byte) error {
	x.IndigoCA = nil
	var c IndigoCA
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IndigoCA = &c
	}
	return nil
}

func (x *MischievousCA) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.IndigoCA != nil, x.IndigoCA, false, nil, false, nil, false)
}

type MischievousCERT struct {
	IndigoCERT *IndigoCERT
	String     *string
}

func (x *MischievousCERT) UnmarshalJSON(data []byte) error {
	x.IndigoCERT = nil
	var c IndigoCERT
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IndigoCERT = &c
	}
	return nil
}

func (x *MischievousCERT) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.IndigoCERT != nil, x.IndigoCERT, false, nil, false, nil, false)
}

type MischievousKey struct {
	IndigoKey *IndigoKey
	String    *string
}

func (x *MischievousKey) UnmarshalJSON(data []byte) error {
	x.IndigoKey = nil
	var c IndigoKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IndigoKey = &c
	}
	return nil
}

func (x *MischievousKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.IndigoKey != nil, x.IndigoKey, false, nil, false, nil, false)
}

type CERTChain4 struct {
	HilariousCERTChain *HilariousCERTChain
	String             *string
}

func (x *CERTChain4) UnmarshalJSON(data []byte) error {
	x.HilariousCERTChain = nil
	var c HilariousCERTChain
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.HilariousCERTChain = &c
	}
	return nil
}

func (x *CERTChain4) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.HilariousCERTChain != nil, x.HilariousCERTChain, false, nil, false, nil, false)
}

type PrivateKey4 struct {
	HilariousPrivateKey *HilariousPrivateKey
	String              *string
}

func (x *PrivateKey4) UnmarshalJSON(data []byte) error {
	x.HilariousPrivateKey = nil
	var c HilariousPrivateKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.HilariousPrivateKey = &c
	}
	return nil
}

func (x *PrivateKey4) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.HilariousPrivateKey != nil, x.HilariousPrivateKey, false, nil, false, nil, false)
}

type RootCerts4 struct {
	HilariousRootCerts *HilariousRootCerts
	String             *string
}

func (x *RootCerts4) UnmarshalJSON(data []byte) error {
	x.HilariousRootCerts = nil
	var c HilariousRootCerts
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.HilariousRootCerts = &c
	}
	return nil
}

func (x *RootCerts4) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.HilariousRootCerts != nil, x.HilariousRootCerts, false, nil, false, nil, false)
}

type MischievousBody struct {
	IndecentBodyArray []IndecentBody
	String            *string
}

func (x *MischievousBody) UnmarshalJSON(data []byte) error {
	x.IndecentBodyArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.IndecentBodyArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MischievousBody) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.IndecentBodyArray != nil, x.IndecentBodyArray, false, nil, false, nil, false, nil, false)
}

type BraggadociousCA struct {
	IndecentCA *IndecentCA
	String     *string
}

func (x *BraggadociousCA) UnmarshalJSON(data []byte) error {
	x.IndecentCA = nil
	var c IndecentCA
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IndecentCA = &c
	}
	return nil
}

func (x *BraggadociousCA) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.IndecentCA != nil, x.IndecentCA, false, nil, false, nil, false)
}

type BraggadociousCERT struct {
	IndecentCERT *IndecentCERT
	String       *string
}

func (x *BraggadociousCERT) UnmarshalJSON(data []byte) error {
	x.IndecentCERT = nil
	var c IndecentCERT
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IndecentCERT = &c
	}
	return nil
}

func (x *BraggadociousCERT) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.IndecentCERT != nil, x.IndecentCERT, false, nil, false, nil, false)
}

type BraggadociousKey struct {
	IndecentKey *IndecentKey
	String      *string
}

func (x *BraggadociousKey) UnmarshalJSON(data []byte) error {
	x.IndecentKey = nil
	var c IndecentKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IndecentKey = &c
	}
	return nil
}

func (x *BraggadociousKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.IndecentKey != nil, x.IndecentKey, false, nil, false, nil, false)
}

type CERTChain5 struct {
	AmbitiousCERTChain *AmbitiousCERTChain
	String             *string
}

func (x *CERTChain5) UnmarshalJSON(data []byte) error {
	x.AmbitiousCERTChain = nil
	var c AmbitiousCERTChain
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.AmbitiousCERTChain = &c
	}
	return nil
}

func (x *CERTChain5) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.AmbitiousCERTChain != nil, x.AmbitiousCERTChain, false, nil, false, nil, false)
}

type PrivateKey5 struct {
	AmbitiousPrivateKey *AmbitiousPrivateKey
	String              *string
}

func (x *PrivateKey5) UnmarshalJSON(data []byte) error {
	x.AmbitiousPrivateKey = nil
	var c AmbitiousPrivateKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.AmbitiousPrivateKey = &c
	}
	return nil
}

func (x *PrivateKey5) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.AmbitiousPrivateKey != nil, x.AmbitiousPrivateKey, false, nil, false, nil, false)
}

type RootCerts5 struct {
	AmbitiousRootCerts *AmbitiousRootCerts
	String             *string
}

func (x *RootCerts5) UnmarshalJSON(data []byte) error {
	x.AmbitiousRootCerts = nil
	var c AmbitiousRootCerts
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.AmbitiousRootCerts = &c
	}
	return nil
}

func (x *RootCerts5) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.AmbitiousRootCerts != nil, x.AmbitiousRootCerts, false, nil, false, nil, false)
}

type AvgUnion struct {
	AvgElementArray []AvgElement
	Double          *float64
}

func (x *AvgUnion) UnmarshalJSON(data []byte) error {
	x.AvgElementArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.AvgElementArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *AvgUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.AvgElementArray != nil, x.AvgElementArray, false, nil, false, nil, false, nil, false)
}

type MaxUnion struct {
	Double          *float64
	MaxElementArray []MaxElement
}

func (x *MaxUnion) UnmarshalJSON(data []byte) error {
	x.MaxElementArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.MaxElementArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MaxUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.MaxElementArray != nil, x.MaxElementArray, false, nil, false, nil, false, nil, false)
}

type MedUnion struct {
	Double          *float64
	MedElementArray []MedElement
}

func (x *MedUnion) UnmarshalJSON(data []byte) error {
	x.MedElementArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.MedElementArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MedUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.MedElementArray != nil, x.MedElementArray, false, nil, false, nil, false, nil, false)
}

type MinUnion struct {
	Double          *float64
	MinElementArray []MinElement
}

func (x *MinUnion) UnmarshalJSON(data []byte) error {
	x.MinElementArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.MinElementArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MinUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.MinElementArray != nil, x.MinElementArray, false, nil, false, nil, false, nil, false)
}

type P95Union struct {
	Double          *float64
	P95ElementArray []P95Element
}

func (x *P95Union) UnmarshalJSON(data []byte) error {
	x.P95ElementArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.P95ElementArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *P95Union) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.P95ElementArray != nil, x.P95ElementArray, false, nil, false, nil, false, nil, false)
}

type P99Union struct {
	Double          *float64
	P99ElementArray []P99Element
}

func (x *P99Union) UnmarshalJSON(data []byte) error {
	x.P99ElementArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.P99ElementArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *P99Union) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.P99ElementArray != nil, x.P99ElementArray, false, nil, false, nil, false, nil, false)
}

type CERTChain6 struct {
	CunningCERTChain *CunningCERTChain
	String           *string
}

func (x *CERTChain6) UnmarshalJSON(data []byte) error {
	x.CunningCERTChain = nil
	var c CunningCERTChain
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.CunningCERTChain = &c
	}
	return nil
}

func (x *CERTChain6) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.CunningCERTChain != nil, x.CunningCERTChain, false, nil, false, nil, false)
}

type PrivateKey6 struct {
	CunningPrivateKey *CunningPrivateKey
	String            *string
}

func (x *PrivateKey6) UnmarshalJSON(data []byte) error {
	x.CunningPrivateKey = nil
	var c CunningPrivateKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.CunningPrivateKey = &c
	}
	return nil
}

func (x *PrivateKey6) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.CunningPrivateKey != nil, x.CunningPrivateKey, false, nil, false, nil, false)
}

type RootCerts6 struct {
	CunningRootCerts *CunningRootCerts
	String           *string
}

func (x *RootCerts6) UnmarshalJSON(data []byte) error {
	x.CunningRootCerts = nil
	var c CunningRootCerts
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.CunningRootCerts = &c
	}
	return nil
}

func (x *RootCerts6) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.CunningRootCerts != nil, x.CunningRootCerts, false, nil, false, nil, false)
}

type MischievousCo2 struct {
	Double           *float64
	IndecentCo2Array []IndecentCo2
}

func (x *MischievousCo2) UnmarshalJSON(data []byte) error {
	x.IndecentCo2Array = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.IndecentCo2Array, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MischievousCo2) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.IndecentCo2Array != nil, x.IndecentCo2Array, false, nil, false, nil, false, nil, false)
}

type MischievousPerformance struct {
	Double                   *float64
	IndecentPerformanceArray []IndecentPerformance
}

func (x *MischievousPerformance) UnmarshalJSON(data []byte) error {
	x.IndecentPerformanceArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.IndecentPerformanceArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MischievousPerformance) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.IndecentPerformanceArray != nil, x.IndecentPerformanceArray, false, nil, false, nil, false, nil, false)
}

type MischievousSize struct {
	Double            *float64
	IndecentSizeArray []IndecentSize
}

func (x *MischievousSize) UnmarshalJSON(data []byte) error {
	x.IndecentSizeArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.IndecentSizeArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MischievousSize) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.IndecentSizeArray != nil, x.IndecentSizeArray, false, nil, false, nil, false, nil, false)
}

type IndecentMetadatum struct {
	String             *string
	TentacledMetadatum *TentacledMetadatum
}

func (x *IndecentMetadatum) UnmarshalJSON(data []byte) error {
	x.TentacledMetadatum = nil
	var c TentacledMetadatum
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TentacledMetadatum = &c
	}
	return nil
}

func (x *IndecentMetadatum) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TentacledMetadatum != nil, x.TentacledMetadatum, false, nil, false, nil, false)
}

type CA1 struct {
	HilariousCA *HilariousCA
	String      *string
}

func (x *CA1) UnmarshalJSON(data []byte) error {
	x.HilariousCA = nil
	var c HilariousCA
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.HilariousCA = &c
	}
	return nil
}

func (x *CA1) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.HilariousCA != nil, x.HilariousCA, false, nil, false, nil, false)
}

type CERT1 struct {
	HilariousCERT *HilariousCERT
	String        *string
}

func (x *CERT1) UnmarshalJSON(data []byte) error {
	x.HilariousCERT = nil
	var c HilariousCERT
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.HilariousCERT = &c
	}
	return nil
}

func (x *CERT1) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.HilariousCERT != nil, x.HilariousCERT, false, nil, false, nil, false)
}

type Key1 struct {
	HilariousKey *HilariousKey
	String       *string
}

func (x *Key1) UnmarshalJSON(data []byte) error {
	x.HilariousKey = nil
	var c HilariousKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.HilariousKey = &c
	}
	return nil
}

func (x *Key1) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.HilariousKey != nil, x.HilariousKey, false, nil, false, nil, false)
}

type CERTChain7 struct {
	MagentaCERTChain *MagentaCERTChain
	String           *string
}

func (x *CERTChain7) UnmarshalJSON(data []byte) error {
	x.MagentaCERTChain = nil
	var c MagentaCERTChain
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.MagentaCERTChain = &c
	}
	return nil
}

func (x *CERTChain7) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.MagentaCERTChain != nil, x.MagentaCERTChain, false, nil, false, nil, false)
}

type PrivateKey7 struct {
	MagentaPrivateKey *MagentaPrivateKey
	String            *string
}

func (x *PrivateKey7) UnmarshalJSON(data []byte) error {
	x.MagentaPrivateKey = nil
	var c MagentaPrivateKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.MagentaPrivateKey = &c
	}
	return nil
}

func (x *PrivateKey7) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.MagentaPrivateKey != nil, x.MagentaPrivateKey, false, nil, false, nil, false)
}

type RootCerts7 struct {
	MagentaRootCerts *MagentaRootCerts
	String           *string
}

func (x *RootCerts7) UnmarshalJSON(data []byte) error {
	x.MagentaRootCerts = nil
	var c MagentaRootCerts
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.MagentaRootCerts = &c
	}
	return nil
}

func (x *RootCerts7) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.MagentaRootCerts != nil, x.MagentaRootCerts, false, nil, false, nil, false)
}

type BraggadociousBody struct {
	HilariousBodyArray []HilariousBody
	String             *string
}

func (x *BraggadociousBody) UnmarshalJSON(data []byte) error {
	x.HilariousBodyArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.HilariousBodyArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *BraggadociousBody) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.HilariousBodyArray != nil, x.HilariousBodyArray, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
