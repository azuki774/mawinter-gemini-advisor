package service

import "context"

// MockGeminiClient is a mock implementation of the GeminiClient interface.
type MockGeminiClient struct {
	PostFunc func(ctx context.Context, prompt string) (string, error)
}

// Post calls the PostFunc if it's set, otherwise returns default values.
func (m *MockGeminiClient) Post(ctx context.Context, prompt string) (string, error) {
	if m.PostFunc != nil {
		return m.PostFunc(ctx, prompt)
	}
	// Default mock behavior
	return "mock gemini response", nil
}

// MockMawinterClient is a mock implementation of the MawinterClient interface.
type MockMawinterClient struct {
	GetMonthlyDataFunc func(yyyymm string) (string, error)
}

// GetMonthlyData calls the GetMonthlyDataFunc if it's set, otherwise returns default values.
func (m *MockMawinterClient) GetMonthlyData(yyyymm string) (string, error) {
	if m.GetMonthlyDataFunc != nil {
		return m.GetMonthlyDataFunc(yyyymm)
	}
	// Default mock behavior
	return `{"data": "mock monthly data"}`, nil
}

// MockFileOperator is a mock implementation of the FileOperator interface.
type MockFileOperator struct {
	LoadTxtFileFunc  func(filePath string) (string, error)
	WriteTxtFileFunc func(filePath string, data string) error

	// Optional: fields to record calls for verification
	WriteTxtFileCalledWith map[string]string // map[filePath]data
}

// LoadTxtFile calls the LoadTxtFileFunc if it's set, otherwise returns default values.
func (m *MockFileOperator) LoadTxtFile(filePath string) (string, error) {
	if m.LoadTxtFileFunc != nil {
		return m.LoadTxtFileFunc(filePath)
	}
	// Default mock behavior
	return "mock file content", nil
}

// WriteTxtFile calls the WriteTxtFileFunc if it's set, otherwise records the call.
func (m *MockFileOperator) WriteTxtFile(filePath string, data string) error {
	if m.WriteTxtFileCalledWith == nil {
		m.WriteTxtFileCalledWith = make(map[string]string)
	}
	m.WriteTxtFileCalledWith[filePath] = data

	if m.WriteTxtFileFunc != nil {
		return m.WriteTxtFileFunc(filePath, data)
	}
	// Default mock behavior
	return nil
}
