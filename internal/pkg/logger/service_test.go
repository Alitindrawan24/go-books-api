package logger

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLogCtx(t *testing.T) {
	tests := []struct {
		name      string
		ctx       context.Context
		requestID string
		wantEmpty bool
	}{
		{
			name:      "new context with provided request ID",
			ctx:       context.Background(),
			requestID: "test-request-id",
			wantEmpty: false,
		},
		{
			name:      "new context with empty request ID generates UUID",
			ctx:       context.Background(),
			requestID: "",
			wantEmpty: false,
		},
		{
			name:      "existing context with request ID returns same",
			ctx:       SetRequestID(context.Background(), "existing-id"),
			requestID: "new-id",
			wantEmpty: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InitLogCtx(tt.ctx, tt.requestID)
			requestID := GetRequestID(result)
			
			if tt.wantEmpty {
				assert.Empty(t, requestID)
			} else {
				assert.NotEmpty(t, requestID)
				if tt.requestID != "" && GetRequestID(tt.ctx) == "" {
					assert.Equal(t, tt.requestID, requestID)
				}
			}
		})
	}
}

func TestGetRequestID(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		want string
	}{
		{
			name: "context with request ID",
			ctx:  SetRequestID(context.Background(), "test-id"),
			want: "test-id",
		},
		{
			name: "context without request ID",
			ctx:  context.Background(),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRequestID(tt.ctx)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSetRequestID(t *testing.T) {
	ctx := context.Background()
	requestID := "test-request-id"
	
	result := SetRequestID(ctx, requestID)
	got := GetRequestID(result)
	
	assert.Equal(t, requestID, got)
}

func TestToMap(t *testing.T) {
	ctx := SetRequestID(context.Background(), "test-id")
	
	tests := []struct {
		name     string
		ctx      context.Context
		val      interface{}
		err      error
		expected map[string]interface{}
	}{
		{
			name: "nil value",
			ctx:  ctx,
			val:  nil,
			err:  nil,
			expected: map[string]interface{}{
				"request_id": "test-id",
				"app":        "",
				"ip_address": "-",
			},
		},
		{
			name: "with error",
			ctx:  ctx,
			val:  nil,
			err:  errors.New("test error"),
			expected: map[string]interface{}{
				"request_id": "test-id",
				"app":        "",
				"ip_address": "-",
				"error":      "test error",
			},
		},
		{
			name: "string value",
			ctx:  ctx,
			val:  "test string",
			err:  nil,
			expected: map[string]interface{}{
				"request_id": "test-id",
				"app":        "",
				"ip_address": "-",
				"metadata":   "test string",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toMap(tt.ctx, tt.val, tt.err)
			
			assert.Equal(t, tt.expected["request_id"], result["request_id"])
			assert.Equal(t, tt.expected["app"], result["app"])
			assert.Equal(t, tt.expected["ip_address"], result["ip_address"])
			
			if tt.err != nil {
				assert.Equal(t, tt.expected["error"], result["error"])
			}
			
			if tt.val != nil && tt.val != "" {
				assert.Equal(t, tt.expected["metadata"], result["metadata"])
			}
			
			// File field should exist
			assert.Contains(t, result, "file")
		})
	}
}

func TestHasField(t *testing.T) {
	type TestStruct struct {
		Name  string
		Value int
	}

	tests := []struct {
		name      string
		s         interface{}
		fieldName string
		want      bool
	}{
		{
			name:      "existing field",
			s:         TestStruct{},
			fieldName: "Name",
			want:      true,
		},
		{
			name:      "non-existing field",
			s:         TestStruct{},
			fieldName: "NonExistent",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasField(tt.s, tt.fieldName)
			assert.Equal(t, tt.want, got)
		})
	}
}
