package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"

	"github.com/Alitindrawan24/go-books-api/internal/constant"
	"github.com/Alitindrawan24/go-books-api/internal/pkg/config"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type key string

const (
	envDirName       = "LOG_DIR"
	envEnableLogging = "LOG_ENABLE"
	envLoggingMode   = "LOG_MODE"

	keyAppName       = "app"
	keyRequestID key = "request_id"
)

const maskedValue = "*****"

var (
	credentialKeys = []string{
		"username",
		"password",
		"password_confirmation",
		"new_password",
		"old_password",
		"secret",
		"token",
		"access_token",

		"AccessToken",
		"Password",
		"PasswordConfirmation",
		"TwoFactorRecoveryPhraseCode",
		"TwoFactorSecret",

		"X-CLIENT-KEY",
		"X-SIGNATURE",
	}
)

var (
	appName = ""
)

// Init initializes logger with app name and configures output destination
func Init(name string) {
	dir := config.Get(envDirName)
	enableLogging := config.Get(envEnableLogging)
	loggingMode := config.GetWIthDefault(envLoggingMode, "file")

	path := dir + name + ".log"
	appName = name // set app name

	logrus.SetLevel(logrus.TraceLevel)

	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if enableLogging == "enable" {

		if loggingMode == "stdout" {
			// Output to stdout instead of the default stderr
			// Can be any io.Writer, see below for File example
			logrus.SetOutput(os.Stdout)
		}

		if loggingMode == "file" {
			// Output to file instead
			logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Fatal(err)
			}
			logrus.SetOutput(logFile)
		}
	}
}

// InitLogCtx initializes context with unique request ID for log tracing
func InitLogCtx(ctx context.Context, requestID string) context.Context {
	if GetRequestID(ctx) == "" {
		if requestID == "" {
			requestID = uuid.NewString()
		}

		return SetRequestID(ctx, requestID)
	}

	return ctx
}

// Info logs informational messages
func Info(ctx context.Context, metadata interface{}, err error, message string) {
	logrus.WithFields(logrus.Fields(toMap(ctx, metadata, err))).Info(message)
}

// Warning logs warning messages
func Warning(ctx context.Context, metadata interface{}, err error, message string) {
	logrus.WithFields(logrus.Fields(toMap(ctx, metadata, err))).Warning(message)
}

// Error logs error messages
func Error(ctx context.Context, metadata interface{}, err error, message string) {
	logrus.WithFields(logrus.Fields(toMap(ctx, metadata, err))).Error(message)
}

// Fatal logs fatal messages and exits the program
func Fatal(ctx context.Context, metadata interface{}, err error, message string) {
	logrus.WithFields(logrus.Fields(toMap(ctx, metadata, err))).Fatal(message)
}

// Trace logs detailed trace messages for debugging
func Trace(ctx context.Context, metadata interface{}, err error, message string) {
	logrus.WithFields(logrus.Fields(toMap(ctx, metadata, err))).Trace(message)
}

// toMap converts metadata to map and enriches with context data
func toMap(ctx context.Context, val interface{}, err error) map[string]interface{} {
	result := make(map[string]interface{})
	result["request_id"] = GetRequestID(ctx)
	result["app"] = appName
	result["ip_address"] = "-"
	ipAddress, ok := ctx.Value(constant.ClientIPAddress).(string)
	if ok {
		result["ip_address"] = ipAddress
	}

	_, dirFile, lineNum, _ := runtime.Caller(2)
	result["file"] = fmt.Sprintf("%s:%d", dirFile, lineNum)

	if err != nil {
		result["error"] = err.Error()
	}

	if val == nil {
		return result
	}

	dataType := reflect.ValueOf(val).Kind()
	switch dataType {
	case reflect.Struct:
		valStruct := struct {
			Metadata interface{} `json:"metadata"`
		}{val}
		data, _ := json.Marshal(valStruct)
		_ = json.Unmarshal(data, &result)
	default:
		result["metadata"] = val
	}

	if metadata, ok := result["metadata"]; ok {
		maskSensitive(metadata)
	}

	return result
}

// maskSensitive recursively masks sensitive credential fields
func maskSensitive(data interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			if isCredentialKey(key) {
				v[key] = maskedValue
				continue
			}
			maskSensitive(value)
		}
	case []interface{}:
		for i := range v {
			maskSensitive(v[i])
		}
	}
}

// isCredentialKey checks if key contains sensitive credential data
func isCredentialKey(key string) bool {
	for _, sensitiveKey := range credentialKeys {
		if key == sensitiveKey {
			return true
		}
	}

	return false
}

// GetRequestID retrieves request ID from context
func GetRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(keyRequestID).(string)
	if !ok {
		return ""
	}

	return requestID
}

// SetRequestID stores request ID in context
func SetRequestID(ctx context.Context, randomKey string) context.Context {
	return context.WithValue(ctx, keyRequestID, randomKey)
}

// hasField checks if struct has specified field name
func hasField(s interface{}, fieldName string) bool {
	structType := reflect.TypeOf(s)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if field.Name == fieldName {
			return true
		}
	}

	return false
}
