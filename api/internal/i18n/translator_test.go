package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormalizeLocale(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Traditional Chinese variants
		{
			name:     "zh-TW",
			input:    "zh-TW",
			expected: "zh-TW",
		},
		{
			name:     "zh-tw lowercase",
			input:    "zh-tw",
			expected: "zh-TW",
		},
		{
			name:     "zh-Hant-TW",
			input:    "zh-Hant-TW",
			expected: "zh-TW",
		},
		{
			name:     "zh-hant-tw lowercase",
			input:    "zh-hant-tw",
			expected: "zh-TW",
		},
		{
			name:     "zh-Hant",
			input:    "zh-Hant",
			expected: "zh-TW",
		},
		{
			name:     "zh-hant lowercase",
			input:    "zh-hant",
			expected: "zh-TW",
		},
		// Simplified Chinese variants
		{
			name:     "zh-CN",
			input:    "zh-CN",
			expected: "zh",
		},
		{
			name:     "zh-cn lowercase",
			input:    "zh-cn",
			expected: "zh",
		},
		{
			name:     "zh-Hans",
			input:    "zh-Hans",
			expected: "zh",
		},
		{
			name:     "zh-hans lowercase",
			input:    "zh-hans",
			expected: "zh",
		},
		{
			name:     "zh only",
			input:    "zh",
			expected: "zh",
		},
		// English variants
		{
			name:     "en-US",
			input:    "en-US",
			expected: "en",
		},
		{
			name:     "en-us lowercase",
			input:    "en-us",
			expected: "en",
		},
		{
			name:     "en-GB",
			input:    "en-GB",
			expected: "en",
		},
		{
			name:     "en only",
			input:    "en",
			expected: "en",
		},
		// Accept-Language header format
		{
			name:     "with quality value",
			input:    "zh-TW;q=0.9",
			expected: "zh-TW",
		},
		{
			name:     "with spaces",
			input:    "  zh-TW  ",
			expected: "zh-TW",
		},
		{
			name:     "with quality and spaces",
			input:    "  zh-tw ; q=0.8  ",
			expected: "zh-TW",
		},
		// Unknown locale - defaults to zh
		{
			name:     "unknown locale",
			input:    "fr-FR",
			expected: "zh",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "zh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeLocale(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewTranslator(t *testing.T) {
	trans, err := NewTranslator()
	require.NoError(t, err)
	require.NotNil(t, trans)

	// Check that locales are loaded
	locales := trans.GetAvailableLocales()
	assert.NotEmpty(t, locales, "Should have loaded at least one locale")

	// Check for expected locales
	localeMap := make(map[string]bool)
	for _, locale := range locales {
		localeMap[locale] = true
	}

	assert.True(t, localeMap["zh"], "Should have loaded zh locale")
	assert.True(t, localeMap["en"], "Should have loaded en locale")
	assert.True(t, localeMap["zh-TW"], "Should have loaded zh-TW locale")
}

func TestTrans(t *testing.T) {
	trans, err := NewTranslator()
	require.NoError(t, err)

	tests := []struct {
		name     string
		locale   string
		key      string
		expected string
	}{
		// Test nested key access
		{
			name:     "zh nested key",
			locale:   "zh",
			key:      "common.success",
			expected: "成功",
		},
		{
			name:     "zh-TW nested key",
			locale:   "zh-TW",
			key:      "common.success",
			expected: "成功",
		},
		{
			name:     "en nested key",
			locale:   "en",
			key:      "common.success",
			expected: "Successfully",
		},
		// Test fallback chain: zh-TW → zh → en
		{
			name:     "fallback from zh-TW to zh",
			locale:   "zh-TW",
			key:      "common.success",
			expected: "成功", // Should find in zh-TW or fallback to zh
		},
		{
			name:     "fallback from fr to zh",
			locale:   "fr-FR",
			key:      "common.success",
			expected: "成功", // fr not available, fallback to zh (normalized default)
		},
		// Test locale normalization in Trans
		{
			name:     "zh-hant-tw normalization",
			locale:   "zh-Hant-TW",
			key:      "common.success",
			expected: "成功",
		},
		{
			name:     "zh-hans normalization",
			locale:   "zh-Hans",
			key:      "common.success",
			expected: "成功",
		},
		// Test non-existent keys
		{
			name:     "non-existent key returns key itself",
			locale:   "zh",
			key:      "nonexistent.key.path",
			expected: "nonexistent.key.path",
		},
		// Test different key paths
		{
			name:     "login key",
			locale:   "zh",
			key:      "login.loginSuccessTitle",
			expected: "登录成功",
		},
		{
			name:     "zh-TW login key with Taiwan terminology",
			locale:   "zh-TW",
			key:      "login.loginSuccessTitle",
			expected: "登入成功", // Taiwan uses 登入 not 登录
		},
		{
			name:     "route key",
			locale:   "zh",
			key:      "route.dashboard",
			expected: "控制台",
		},
		{
			name:     "zh-TW route key",
			locale:   "zh-TW",
			key:      "route.dashboard",
			expected: "控制台",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trans.Trans(tt.locale, tt.key)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetMessageFromLocale(t *testing.T) {
	trans, err := NewTranslator()
	require.NoError(t, err)

	tests := []struct {
		name     string
		locale   string
		key      string
		expected string
	}{
		{
			name:     "valid key",
			locale:   "zh",
			key:      "common.success",
			expected: "成功",
		},
		{
			name:     "invalid locale",
			locale:   "invalid",
			key:      "common.success",
			expected: "",
		},
		{
			name:     "invalid key",
			locale:   "zh",
			key:      "invalid.key",
			expected: "",
		},
		{
			name:     "deep nested key",
			locale:   "zh",
			key:      "mcms.email.subject",
			expected: "【Simple Admin】 的验证码",
		},
		{
			name:     "zh-TW deep nested key",
			locale:   "zh-TW",
			key:      "mcms.email.subject",
			expected: "【Simple Admin】 的驗證碼",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trans.getMessageFromLocale(tt.locale, tt.key)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetAvailableLocales(t *testing.T) {
	trans, err := NewTranslator()
	require.NoError(t, err)

	locales := trans.GetAvailableLocales()
	assert.NotEmpty(t, locales)

	// Check minimum expected locales
	localeMap := make(map[string]bool)
	for _, locale := range locales {
		localeMap[locale] = true
	}

	assert.True(t, localeMap["zh"], "Should include zh locale")
	assert.True(t, localeMap["en"], "Should include en locale")
	assert.True(t, localeMap["zh-TW"], "Should include zh-TW locale")
}

func TestConcurrentAccess(t *testing.T) {
	trans, err := NewTranslator()
	require.NoError(t, err)

	// Test concurrent reads
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				_ = trans.Trans("zh-TW", "common.success")
				_ = trans.GetAvailableLocales()
			}
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
