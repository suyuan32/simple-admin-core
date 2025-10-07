package i18n

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

// Translator handles i18n translation with fallback support
type Translator struct {
	locales map[string]map[string]interface{}
	mu      sync.RWMutex
}

// NewTranslator creates a new Translator instance and loads all locale files
func NewTranslator() (*Translator, error) {
	t := &Translator{
		locales: make(map[string]map[string]interface{}),
	}

	// Load all locale files from embedded FS
	entries, err := LocaleFS.ReadDir("locale")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		if !strings.HasSuffix(filename, ".json") {
			continue
		}

		// Extract locale code from filename (e.g., "zh-TW.json" -> "zh-TW")
		locale := strings.TrimSuffix(filename, ".json")

		data, err := LocaleFS.ReadFile("locale/" + filename)
		if err != nil {
			logx.Errorf("Failed to read locale file %s: %v", filename, err)
			continue
		}

		var messages map[string]interface{}
		if err := json.Unmarshal(data, &messages); err != nil {
			logx.Errorf("Failed to parse locale file %s: %v", filename, err)
			continue
		}

		t.locales[locale] = messages
		logx.Infof("Loaded locale: %s", locale)
	}

	return t, nil
}

// NormalizeLocale normalizes various locale formats to standard codes
// Handles Accept-Language header formats and converts to our internal format
func NormalizeLocale(acceptLang string) string {
	// Convert to lowercase for case-insensitive matching
	lang := strings.ToLower(strings.TrimSpace(acceptLang))

	// Extract primary language tag (before ';' quality value)
	if idx := strings.Index(lang, ";"); idx != -1 {
		lang = lang[:idx]
	}
	lang = strings.TrimSpace(lang)

	// Handle Traditional Chinese variants
	if strings.HasPrefix(lang, "zh-tw") ||
		strings.HasPrefix(lang, "zh-hant-tw") ||
		strings.HasPrefix(lang, "zh-hant") {
		return "zh-TW"
	}

	// Handle Simplified Chinese variants
	if strings.HasPrefix(lang, "zh-cn") ||
		strings.HasPrefix(lang, "zh-hans") ||
		lang == "zh" {
		return "zh"
	}

	// Handle English variants
	if strings.HasPrefix(lang, "en") {
		return "en"
	}

	// Default to Simplified Chinese
	return "zh"
}

// Trans translates a key for the given locale with fallback chain
// Fallback order: requested locale → zh (Simplified Chinese) → en (English)
func (t *Translator) Trans(locale, key string) string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	// Normalize locale
	normalizedLocale := NormalizeLocale(locale)

	// Try requested locale
	if msg := t.getMessageFromLocale(normalizedLocale, key); msg != "" {
		return msg
	}

	// Fallback to Simplified Chinese (if not already trying zh)
	if normalizedLocale != "zh" {
		if msg := t.getMessageFromLocale("zh", key); msg != "" {
			return msg
		}
	}

	// Fallback to English
	if normalizedLocale != "en" {
		if msg := t.getMessageFromLocale("en", key); msg != "" {
			return msg
		}
	}

	// No translation found, return key itself
	logx.Debugf("Translation not found for key: %s (locale: %s)", key, locale)
	return key
}

// getMessageFromLocale retrieves a message from a specific locale
// Supports nested keys using dot notation (e.g., "common.success")
func (t *Translator) getMessageFromLocale(locale, key string) string {
	messages, ok := t.locales[locale]
	if !ok {
		return ""
	}

	// Split key by dots for nested access
	parts := strings.Split(key, ".")
	var current interface{} = messages

	for _, part := range parts {
		if m, ok := current.(map[string]interface{}); ok {
			current, ok = m[part]
			if !ok {
				return ""
			}
		} else {
			return ""
		}
	}

	// Convert final value to string
	if str, ok := current.(string); ok {
		return str
	}

	return ""
}

// GetAvailableLocales returns all loaded locale codes
func (t *Translator) GetAvailableLocales() []string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	locales := make([]string, 0, len(t.locales))
	for locale := range t.locales {
		locales = append(locales, locale)
	}
	return locales
}
