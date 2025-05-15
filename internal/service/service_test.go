package service

import (
	"context"
	"errors"
	"testing"
)

func TestService_Start(t *testing.T) {
	type fields struct {
		GeminiClient   GeminiClient
		MawinterClient MawinterClient
		FileOperator   FileOperator
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful start",
			fields: fields{
				GeminiClient: &MockGeminiClient{
					PostFunc: func(ctx context.Context, prompt string) (string, error) {
						return "gemini response", nil
					},
				},
				MawinterClient: &MockMawinterClient{
					GetMonthlyDataFunc: func(yyyymm string) (string, error) {
						return `{"data": "mawinter data"}`, nil
					},
				},
				FileOperator: &MockFileOperator{
					LoadTxtFileFunc: func(filePath string) (string, error) {
						return "pre prompt", nil
					},
					WriteTxtFileFunc: func(filePath string, data string) error {
						return nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
		{
			name: "failed to load pre_prompt",
			fields: fields{
				GeminiClient:   &MockGeminiClient{},
				MawinterClient: &MockMawinterClient{},
				FileOperator: &MockFileOperator{
					LoadTxtFileFunc: func(filePath string) (string, error) {
						return "", errors.New("ErrLoadFile")
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "failed to get monthly data",
			fields: fields{
				GeminiClient: &MockGeminiClient{},
				MawinterClient: &MockMawinterClient{
					GetMonthlyDataFunc: func(yyyymm string) (string, error) {
						return "", errors.New("ErrGetMonthlyData")
					},
				},
				FileOperator: &MockFileOperator{
					LoadTxtFileFunc: func(filePath string) (string, error) {
						return "pre prompt", nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "failed to get gemini response",
			fields: fields{
				GeminiClient: &MockGeminiClient{
					PostFunc: func(ctx context.Context, prompt string) (string, error) {
						return "", errors.New("ErrGeminiPost")
					},
				},
				MawinterClient: &MockMawinterClient{
					GetMonthlyDataFunc: func(yyyymm string) (string, error) {
						return `{"data": "mawinter data"}`, nil
					},
				},
				FileOperator: &MockFileOperator{
					LoadTxtFileFunc: func(filePath string) (string, error) {
						return "pre prompt", nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "failed to write gemini response",
			fields: fields{
				GeminiClient: &MockGeminiClient{
					PostFunc: func(ctx context.Context, prompt string) (string, error) {
						return "gemini response", nil
					},
				},
				MawinterClient: &MockMawinterClient{
					GetMonthlyDataFunc: func(yyyymm string) (string, error) {
						return `{"data": "mawinter data"}`, nil
					},
				},
				FileOperator: &MockFileOperator{
					LoadTxtFileFunc: func(filePath string) (string, error) {
						return "pre prompt", nil
					},
					WriteTxtFileFunc: func(filePath string, data string) error {
						return errors.New("ErrWriteFile")
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				GeminiClient:   tt.fields.GeminiClient,
				MawinterClient: tt.fields.MawinterClient,
				FileOperator:   tt.fields.FileOperator,
			}
			if err := s.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Service.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
