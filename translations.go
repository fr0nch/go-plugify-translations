package translations

/*
#include "translations.h"
#cgo noescape LoadTranslation
#cgo noescape ReloadTransalations
#cgo noescape HasTranslation
#cgo noescape HasTranslatedForLanguage
#cgo noescape Translate
#cgo noescape TranslateFormat
*/
import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"unsafe"
	"github.com/untrustedmodders/go-plugify"
)

var _ = errors.New("")
var _ = reflect.TypeOf(0)
var _ = runtime.GOOS
var _ = unsafe.Sizeof(0)
var _ = plugify.ApiVersion

// Generated from translations (group: translations)

var _LoadTranslation = func(paths []string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__paths := plugify.ConstructVectorString(paths)
	plugify.Block {
		Try: func() {
			__native := C.LoadTranslation((*C.Vector)(unsafe.Pointer(&__paths)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyVectorString(&__paths)
		},
	}.Do()
	return __retVal
}

// LoadTranslation 
//  @brief Loads (and merges) one or more translation files.
//
//  @param paths: Paths to the translation files. Relative paths are resolved against the configs directory.
//
//  @return Empty string on success, or an error message on failure.
func LoadTranslation(paths []string) string {
	return _LoadTranslation(paths)
}

var _ReloadTransalations = func() {
	C.ReloadTransalations()
}

// ReloadTransalations 
//  @brief Clears and reloads every translation file previously passed to LoadTranslation.
//
func ReloadTransalations() {
	_ReloadTransalations()
}

var _HasTranslation = func(key string) bool {
	var __retVal bool
	__key := plugify.ConstructString(key)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.HasTranslation((*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// HasTranslation 
//  @brief Checks whether a translation key exists for any language.
//
//  @param key: The translation key to check.
//
//  @return True if the key exists, false otherwise.
func HasTranslation(key string) bool {
	return _HasTranslation(key)
}

var _HasTranslatedForLanguage = func(key string, lang string) bool {
	var __retVal bool
	__key := plugify.ConstructString(key)
	__lang := plugify.ConstructString(lang)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.HasTranslatedForLanguage((*C.String)(unsafe.Pointer(&__key)), (*C.String)(unsafe.Pointer(&__lang))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
			plugify.DestroyString(&__lang)
		},
	}.Do()
	return __retVal
}

// HasTranslatedForLanguage 
//  @brief Checks whether a translation key has a value for a specific language.
//
//  @param key: The translation key to check.
//  @param lang: The language code to check.
//
//  @return True if the key is translated for the language, false otherwise.
func HasTranslatedForLanguage(key string, lang string) bool {
	return _HasTranslatedForLanguage(key, lang)
}

var _Translate = func(lang string, key string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__lang := plugify.ConstructString(lang)
	__key := plugify.ConstructString(key)
	plugify.Block {
		Try: func() {
			__native := C.Translate((*C.String)(unsafe.Pointer(&__lang)), (*C.String)(unsafe.Pointer(&__key)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__lang)
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// Translate 
//  @brief Translates a key into the specified language.
//
//  @param lang: The target language code.
//  @param key: The translation key.
//
//  @return The translated string, or an empty string if not found.
func Translate(lang string, key string) string {
	return _Translate(lang, key)
}

var _TranslateFormat = func(lang string, key string, args []any) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__lang := plugify.ConstructString(lang)
	__key := plugify.ConstructString(key)
	__args := plugify.ConstructVectorVariant(args)
	plugify.Block {
		Try: func() {
			__native := C.TranslateFormat((*C.String)(unsafe.Pointer(&__lang)), (*C.String)(unsafe.Pointer(&__key)), (*C.Vector)(unsafe.Pointer(&__args)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__lang)
			plugify.DestroyString(&__key)
			plugify.DestroyVectorVariant(&__args)
		},
	}.Do()
	return __retVal
}

// TranslateFormat 
//  @brief Translates a key into the specified language and formats it with the given arguments.
//
//  @param lang: The target language code.
//  @param key: The translation key.
//  @param args: Format arguments to substitute into the translated string.
//
//  @return The formatted translated string, or an empty string if not found.
func TranslateFormat(lang string, key string, args []any) string {
	return _TranslateFormat(lang, key, args)
}

