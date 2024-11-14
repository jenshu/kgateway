// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/enterprise/options/ai/ai.proto

package ai

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *SingleAuthToken) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SingleAuthToken)
	if !ok {
		that2, ok := that.(SingleAuthToken)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.AuthTokenSource.(type) {

	case *SingleAuthToken_Inline:
		if _, ok := target.AuthTokenSource.(*SingleAuthToken_Inline); !ok {
			return false
		}

		if strings.Compare(m.GetInline(), target.GetInline()) != 0 {
			return false
		}

	case *SingleAuthToken_SecretRef:
		if _, ok := target.AuthTokenSource.(*SingleAuthToken_SecretRef); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSecretRef()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSecretRef()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSecretRef(), target.GetSecretRef()) {
				return false
			}
		}

	case *SingleAuthToken_Passthrough_:
		if _, ok := target.AuthTokenSource.(*SingleAuthToken_Passthrough_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetPassthrough()).(equality.Equalizer); ok {
			if !h.Equal(target.GetPassthrough()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetPassthrough(), target.GetPassthrough()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.AuthTokenSource != target.AuthTokenSource {
			return false
		}
	}

	return true
}

// Equal function
func (m *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Llm.(type) {

	case *UpstreamSpec_Openai:
		if _, ok := target.Llm.(*UpstreamSpec_Openai); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOpenai()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOpenai()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOpenai(), target.GetOpenai()) {
				return false
			}
		}

	case *UpstreamSpec_Mistral_:
		if _, ok := target.Llm.(*UpstreamSpec_Mistral_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetMistral()).(equality.Equalizer); ok {
			if !h.Equal(target.GetMistral()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetMistral(), target.GetMistral()) {
				return false
			}
		}

	case *UpstreamSpec_Anthropic_:
		if _, ok := target.Llm.(*UpstreamSpec_Anthropic_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAnthropic()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAnthropic()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAnthropic(), target.GetAnthropic()) {
				return false
			}
		}

	case *UpstreamSpec_AzureOpenai:
		if _, ok := target.Llm.(*UpstreamSpec_AzureOpenai); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAzureOpenai()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAzureOpenai()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAzureOpenai(), target.GetAzureOpenai()) {
				return false
			}
		}

	case *UpstreamSpec_Multi:
		if _, ok := target.Llm.(*UpstreamSpec_Multi); !ok {
			return false
		}

		if h, ok := interface{}(m.GetMulti()).(equality.Equalizer); ok {
			if !h.Equal(target.GetMulti()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetMulti(), target.GetMulti()) {
				return false
			}
		}

	case *UpstreamSpec_Gemini_:
		if _, ok := target.Llm.(*UpstreamSpec_Gemini_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGemini()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGemini()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGemini(), target.GetGemini()) {
				return false
			}
		}

	case *UpstreamSpec_VertexAi:
		if _, ok := target.Llm.(*UpstreamSpec_VertexAi); !ok {
			return false
		}

		if h, ok := interface{}(m.GetVertexAi()).(equality.Equalizer); ok {
			if !h.Equal(target.GetVertexAi()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetVertexAi(), target.GetVertexAi()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Llm != target.Llm {
			return false
		}
	}

	return true
}

// Equal function
func (m *RouteSettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteSettings)
	if !ok {
		that2, ok := that.(RouteSettings)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetPromptEnrichment()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPromptEnrichment()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPromptEnrichment(), target.GetPromptEnrichment()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPromptGuard()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPromptGuard()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPromptGuard(), target.GetPromptGuard()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRag()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRag()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRag(), target.GetRag()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSemanticCache()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSemanticCache()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSemanticCache(), target.GetSemanticCache()) {
			return false
		}
	}

	if len(m.GetDefaults()) != len(target.GetDefaults()) {
		return false
	}
	for idx, v := range m.GetDefaults() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetDefaults()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetDefaults()[idx]) {
				return false
			}
		}

	}

	if m.GetRouteType() != target.GetRouteType() {
		return false
	}

	return true
}

// Equal function
func (m *FieldDefault) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FieldDefault)
	if !ok {
		that2, ok := that.(FieldDefault)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetField(), target.GetField()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetValue()).(equality.Equalizer); ok {
		if !h.Equal(target.GetValue()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetValue(), target.GetValue()) {
			return false
		}
	}

	if m.GetOverride() != target.GetOverride() {
		return false
	}

	return true
}

// Equal function
func (m *Postgres) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Postgres)
	if !ok {
		that2, ok := that.(Postgres)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetConnectionString(), target.GetConnectionString()) != 0 {
		return false
	}

	if strings.Compare(m.GetCollectionName(), target.GetCollectionName()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Embedding) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Embedding)
	if !ok {
		that2, ok := that.(Embedding)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Embedding.(type) {

	case *Embedding_Openai:
		if _, ok := target.Embedding.(*Embedding_Openai); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOpenai()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOpenai()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOpenai(), target.GetOpenai()) {
				return false
			}
		}

	case *Embedding_AzureOpenai:
		if _, ok := target.Embedding.(*Embedding_AzureOpenai); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAzureOpenai()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAzureOpenai()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAzureOpenai(), target.GetAzureOpenai()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Embedding != target.Embedding {
			return false
		}
	}

	return true
}

// Equal function
func (m *SemanticCache) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SemanticCache)
	if !ok {
		that2, ok := that.(SemanticCache)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetDatastore()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDatastore()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDatastore(), target.GetDatastore()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetEmbedding()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEmbedding()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEmbedding(), target.GetEmbedding()) {
			return false
		}
	}

	if m.GetTtl() != target.GetTtl() {
		return false
	}

	if m.GetMode() != target.GetMode() {
		return false
	}

	return true
}

// Equal function
func (m *RAG) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RAG)
	if !ok {
		that2, ok := that.(RAG)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetDatastore()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDatastore()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDatastore(), target.GetDatastore()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetEmbedding()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEmbedding()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEmbedding(), target.GetEmbedding()) {
			return false
		}
	}

	if strings.Compare(m.GetPromptTemplate(), target.GetPromptTemplate()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *AIPromptEnrichment) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AIPromptEnrichment)
	if !ok {
		that2, ok := that.(AIPromptEnrichment)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetPrepend()) != len(target.GetPrepend()) {
		return false
	}
	for idx, v := range m.GetPrepend() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPrepend()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPrepend()[idx]) {
				return false
			}
		}

	}

	if len(m.GetAppend()) != len(target.GetAppend()) {
		return false
	}
	for idx, v := range m.GetAppend() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppend()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppend()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *AIPromptGuard) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AIPromptGuard)
	if !ok {
		that2, ok := that.(AIPromptGuard)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRequest()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequest()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequest(), target.GetRequest()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResponse()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponse()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponse(), target.GetResponse()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *SingleAuthToken_Passthrough) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SingleAuthToken_Passthrough)
	if !ok {
		that2, ok := that.(SingleAuthToken_Passthrough)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *UpstreamSpec_CustomHost) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec_CustomHost)
	if !ok {
		that2, ok := that.(UpstreamSpec_CustomHost)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetHost(), target.GetHost()) != 0 {
		return false
	}

	if m.GetPort() != target.GetPort() {
		return false
	}

	return true
}

// Equal function
func (m *UpstreamSpec_OpenAI) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec_OpenAI)
	if !ok {
		that2, ok := that.(UpstreamSpec_OpenAI)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetAuthToken()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAuthToken()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAuthToken(), target.GetAuthToken()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCustomHost()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCustomHost()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCustomHost(), target.GetCustomHost()) {
			return false
		}
	}

	if strings.Compare(m.GetModel(), target.GetModel()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *UpstreamSpec_AzureOpenAI) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec_AzureOpenAI)
	if !ok {
		that2, ok := that.(UpstreamSpec_AzureOpenAI)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetEndpoint(), target.GetEndpoint()) != 0 {
		return false
	}

	if strings.Compare(m.GetDeploymentName(), target.GetDeploymentName()) != 0 {
		return false
	}

	if strings.Compare(m.GetApiVersion(), target.GetApiVersion()) != 0 {
		return false
	}

	switch m.AuthTokenSource.(type) {

	case *UpstreamSpec_AzureOpenAI_AuthToken:
		if _, ok := target.AuthTokenSource.(*UpstreamSpec_AzureOpenAI_AuthToken); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAuthToken()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAuthToken()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAuthToken(), target.GetAuthToken()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.AuthTokenSource != target.AuthTokenSource {
			return false
		}
	}

	return true
}

// Equal function
func (m *UpstreamSpec_Gemini) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec_Gemini)
	if !ok {
		that2, ok := that.(UpstreamSpec_Gemini)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetModel(), target.GetModel()) != 0 {
		return false
	}

	if strings.Compare(m.GetApiVersion(), target.GetApiVersion()) != 0 {
		return false
	}

	switch m.AuthTokenSource.(type) {

	case *UpstreamSpec_Gemini_AuthToken:
		if _, ok := target.AuthTokenSource.(*UpstreamSpec_Gemini_AuthToken); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAuthToken()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAuthToken()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAuthToken(), target.GetAuthToken()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.AuthTokenSource != target.AuthTokenSource {
			return false
		}
	}

	return true
}

// Equal function
func (m *UpstreamSpec_VertexAI) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec_VertexAI)
	if !ok {
		that2, ok := that.(UpstreamSpec_VertexAI)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetModel(), target.GetModel()) != 0 {
		return false
	}

	if strings.Compare(m.GetApiVersion(), target.GetApiVersion()) != 0 {
		return false
	}

	if strings.Compare(m.GetProjectId(), target.GetProjectId()) != 0 {
		return false
	}

	if strings.Compare(m.GetLocation(), target.GetLocation()) != 0 {
		return false
	}

	if strings.Compare(m.GetModelPath(), target.GetModelPath()) != 0 {
		return false
	}

	if m.GetPublisher() != target.GetPublisher() {
		return false
	}

	switch m.AuthTokenSource.(type) {

	case *UpstreamSpec_VertexAI_AuthToken:
		if _, ok := target.AuthTokenSource.(*UpstreamSpec_VertexAI_AuthToken); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAuthToken()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAuthToken()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAuthToken(), target.GetAuthToken()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.AuthTokenSource != target.AuthTokenSource {
			return false
		}
	}

	return true
}

// Equal function
func (m *UpstreamSpec_Mistral) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec_Mistral)
	if !ok {
		that2, ok := that.(UpstreamSpec_Mistral)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetAuthToken()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAuthToken()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAuthToken(), target.GetAuthToken()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCustomHost()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCustomHost()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCustomHost(), target.GetCustomHost()) {
			return false
		}
	}

	if strings.Compare(m.GetModel(), target.GetModel()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *UpstreamSpec_Anthropic) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec_Anthropic)
	if !ok {
		that2, ok := that.(UpstreamSpec_Anthropic)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetAuthToken()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAuthToken()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAuthToken(), target.GetAuthToken()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCustomHost()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCustomHost()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCustomHost(), target.GetCustomHost()) {
			return false
		}
	}

	if strings.Compare(m.GetVersion(), target.GetVersion()) != 0 {
		return false
	}

	if strings.Compare(m.GetModel(), target.GetModel()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *UpstreamSpec_MultiPool) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec_MultiPool)
	if !ok {
		that2, ok := that.(UpstreamSpec_MultiPool)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetPriorities()) != len(target.GetPriorities()) {
		return false
	}
	for idx, v := range m.GetPriorities() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPriorities()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPriorities()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *UpstreamSpec_MultiPool_Backend) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec_MultiPool_Backend)
	if !ok {
		that2, ok := that.(UpstreamSpec_MultiPool_Backend)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Llm.(type) {

	case *UpstreamSpec_MultiPool_Backend_Openai:
		if _, ok := target.Llm.(*UpstreamSpec_MultiPool_Backend_Openai); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOpenai()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOpenai()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOpenai(), target.GetOpenai()) {
				return false
			}
		}

	case *UpstreamSpec_MultiPool_Backend_Mistral:
		if _, ok := target.Llm.(*UpstreamSpec_MultiPool_Backend_Mistral); !ok {
			return false
		}

		if h, ok := interface{}(m.GetMistral()).(equality.Equalizer); ok {
			if !h.Equal(target.GetMistral()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetMistral(), target.GetMistral()) {
				return false
			}
		}

	case *UpstreamSpec_MultiPool_Backend_Anthropic:
		if _, ok := target.Llm.(*UpstreamSpec_MultiPool_Backend_Anthropic); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAnthropic()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAnthropic()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAnthropic(), target.GetAnthropic()) {
				return false
			}
		}

	case *UpstreamSpec_MultiPool_Backend_AzureOpenai:
		if _, ok := target.Llm.(*UpstreamSpec_MultiPool_Backend_AzureOpenai); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAzureOpenai()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAzureOpenai()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAzureOpenai(), target.GetAzureOpenai()) {
				return false
			}
		}

	case *UpstreamSpec_MultiPool_Backend_Gemini:
		if _, ok := target.Llm.(*UpstreamSpec_MultiPool_Backend_Gemini); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGemini()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGemini()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGemini(), target.GetGemini()) {
				return false
			}
		}

	case *UpstreamSpec_MultiPool_Backend_VertexAi:
		if _, ok := target.Llm.(*UpstreamSpec_MultiPool_Backend_VertexAi); !ok {
			return false
		}

		if h, ok := interface{}(m.GetVertexAi()).(equality.Equalizer); ok {
			if !h.Equal(target.GetVertexAi()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetVertexAi(), target.GetVertexAi()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Llm != target.Llm {
			return false
		}
	}

	return true
}

// Equal function
func (m *UpstreamSpec_MultiPool_Priority) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec_MultiPool_Priority)
	if !ok {
		that2, ok := that.(UpstreamSpec_MultiPool_Priority)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetPool()) != len(target.GetPool()) {
		return false
	}
	for idx, v := range m.GetPool() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPool()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPool()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *Embedding_OpenAI) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Embedding_OpenAI)
	if !ok {
		that2, ok := that.(Embedding_OpenAI)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.AuthTokenSource.(type) {

	case *Embedding_OpenAI_AuthToken:
		if _, ok := target.AuthTokenSource.(*Embedding_OpenAI_AuthToken); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAuthToken()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAuthToken()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAuthToken(), target.GetAuthToken()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.AuthTokenSource != target.AuthTokenSource {
			return false
		}
	}

	return true
}

// Equal function
func (m *Embedding_AzureOpenAI) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Embedding_AzureOpenAI)
	if !ok {
		that2, ok := that.(Embedding_AzureOpenAI)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetApiVersion(), target.GetApiVersion()) != 0 {
		return false
	}

	if strings.Compare(m.GetEndpoint(), target.GetEndpoint()) != 0 {
		return false
	}

	if strings.Compare(m.GetDeploymentName(), target.GetDeploymentName()) != 0 {
		return false
	}

	switch m.AuthTokenSource.(type) {

	case *Embedding_AzureOpenAI_AuthToken:
		if _, ok := target.AuthTokenSource.(*Embedding_AzureOpenAI_AuthToken); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAuthToken()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAuthToken()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAuthToken(), target.GetAuthToken()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.AuthTokenSource != target.AuthTokenSource {
			return false
		}
	}

	return true
}

// Equal function
func (m *SemanticCache_Redis) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SemanticCache_Redis)
	if !ok {
		that2, ok := that.(SemanticCache_Redis)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetConnectionString(), target.GetConnectionString()) != 0 {
		return false
	}

	if m.GetScoreThreshold() != target.GetScoreThreshold() {
		return false
	}

	return true
}

// Equal function
func (m *SemanticCache_Weaviate) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SemanticCache_Weaviate)
	if !ok {
		that2, ok := that.(SemanticCache_Weaviate)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetHost(), target.GetHost()) != 0 {
		return false
	}

	if m.GetHttpPort() != target.GetHttpPort() {
		return false
	}

	if m.GetGrpcPort() != target.GetGrpcPort() {
		return false
	}

	if m.GetInsecure() != target.GetInsecure() {
		return false
	}

	return true
}

// Equal function
func (m *SemanticCache_DataStore) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SemanticCache_DataStore)
	if !ok {
		that2, ok := that.(SemanticCache_DataStore)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Datastore.(type) {

	case *SemanticCache_DataStore_Redis:
		if _, ok := target.Datastore.(*SemanticCache_DataStore_Redis); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRedis()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRedis()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRedis(), target.GetRedis()) {
				return false
			}
		}

	case *SemanticCache_DataStore_Weaviate:
		if _, ok := target.Datastore.(*SemanticCache_DataStore_Weaviate); !ok {
			return false
		}

		if h, ok := interface{}(m.GetWeaviate()).(equality.Equalizer); ok {
			if !h.Equal(target.GetWeaviate()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetWeaviate(), target.GetWeaviate()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Datastore != target.Datastore {
			return false
		}
	}

	return true
}

// Equal function
func (m *RAG_DataStore) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RAG_DataStore)
	if !ok {
		that2, ok := that.(RAG_DataStore)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Datastore.(type) {

	case *RAG_DataStore_Postgres:
		if _, ok := target.Datastore.(*RAG_DataStore_Postgres); !ok {
			return false
		}

		if h, ok := interface{}(m.GetPostgres()).(equality.Equalizer); ok {
			if !h.Equal(target.GetPostgres()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetPostgres(), target.GetPostgres()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Datastore != target.Datastore {
			return false
		}
	}

	return true
}

// Equal function
func (m *AIPromptEnrichment_Message) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AIPromptEnrichment_Message)
	if !ok {
		that2, ok := that.(AIPromptEnrichment_Message)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetRole(), target.GetRole()) != 0 {
		return false
	}

	if strings.Compare(m.GetContent(), target.GetContent()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *AIPromptGuard_Regex) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AIPromptGuard_Regex)
	if !ok {
		that2, ok := that.(AIPromptGuard_Regex)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetMatches()) != len(target.GetMatches()) {
		return false
	}
	for idx, v := range m.GetMatches() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMatches()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMatches()[idx]) {
				return false
			}
		}

	}

	if len(m.GetBuiltins()) != len(target.GetBuiltins()) {
		return false
	}
	for idx, v := range m.GetBuiltins() {

		if v != target.GetBuiltins()[idx] {
			return false
		}

	}

	if m.GetAction() != target.GetAction() {
		return false
	}

	return true
}

// Equal function
func (m *AIPromptGuard_Webhook) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AIPromptGuard_Webhook)
	if !ok {
		that2, ok := that.(AIPromptGuard_Webhook)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetHost(), target.GetHost()) != 0 {
		return false
	}

	if m.GetPort() != target.GetPort() {
		return false
	}

	if len(m.GetForwardHeaders()) != len(target.GetForwardHeaders()) {
		return false
	}
	for idx, v := range m.GetForwardHeaders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetForwardHeaders()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetForwardHeaders()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *AIPromptGuard_Moderation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AIPromptGuard_Moderation)
	if !ok {
		that2, ok := that.(AIPromptGuard_Moderation)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Moderation.(type) {

	case *AIPromptGuard_Moderation_Openai:
		if _, ok := target.Moderation.(*AIPromptGuard_Moderation_Openai); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOpenai()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOpenai()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOpenai(), target.GetOpenai()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Moderation != target.Moderation {
			return false
		}
	}

	return true
}

// Equal function
func (m *AIPromptGuard_Request) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AIPromptGuard_Request)
	if !ok {
		that2, ok := that.(AIPromptGuard_Request)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetCustomResponse()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCustomResponse()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCustomResponse(), target.GetCustomResponse()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRegex()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRegex()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRegex(), target.GetRegex()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetWebhook()).(equality.Equalizer); ok {
		if !h.Equal(target.GetWebhook()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetWebhook(), target.GetWebhook()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetModeration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetModeration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetModeration(), target.GetModeration()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *AIPromptGuard_Response) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AIPromptGuard_Response)
	if !ok {
		that2, ok := that.(AIPromptGuard_Response)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRegex()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRegex()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRegex(), target.GetRegex()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetWebhook()).(equality.Equalizer); ok {
		if !h.Equal(target.GetWebhook()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetWebhook(), target.GetWebhook()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *AIPromptGuard_Regex_RegexMatch) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AIPromptGuard_Regex_RegexMatch)
	if !ok {
		that2, ok := that.(AIPromptGuard_Regex_RegexMatch)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetPattern(), target.GetPattern()) != 0 {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *AIPromptGuard_Webhook_HeaderMatch) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AIPromptGuard_Webhook_HeaderMatch)
	if !ok {
		that2, ok := that.(AIPromptGuard_Webhook_HeaderMatch)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	if m.GetMatchType() != target.GetMatchType() {
		return false
	}

	return true
}

// Equal function
func (m *AIPromptGuard_Moderation_OpenAI) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AIPromptGuard_Moderation_OpenAI)
	if !ok {
		that2, ok := that.(AIPromptGuard_Moderation_OpenAI)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetModel(), target.GetModel()) != 0 {
		return false
	}

	switch m.AuthTokenSource.(type) {

	case *AIPromptGuard_Moderation_OpenAI_AuthToken:
		if _, ok := target.AuthTokenSource.(*AIPromptGuard_Moderation_OpenAI_AuthToken); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAuthToken()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAuthToken()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAuthToken(), target.GetAuthToken()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.AuthTokenSource != target.AuthTokenSource {
			return false
		}
	}

	return true
}

// Equal function
func (m *AIPromptGuard_Request_CustomResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AIPromptGuard_Request_CustomResponse)
	if !ok {
		that2, ok := that.(AIPromptGuard_Request_CustomResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetMessage(), target.GetMessage()) != 0 {
		return false
	}

	if m.GetStatusCode() != target.GetStatusCode() {
		return false
	}

	return true
}
