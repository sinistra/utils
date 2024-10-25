package utils

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCreateDirIfNotExist(t *testing.T) {
	tests := []struct {
		name      string
		dir       string
		wantErr   bool
		removeDir bool
	}{
		{
			name:      "success",
			dir:       "testdir",
			wantErr:   false,
			removeDir: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateDirIfNotExist(tt.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDirIfNotExist() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.removeDir {
				err = os.RemoveAll(tt.dir)
				require.NoError(t, err)
			}
		})
	}
}
