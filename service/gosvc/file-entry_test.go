package gosvc

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_fileEntry_CopyTo(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		wd, _ := os.Getwd()
		srcFilePath := filepath.Join(wd, "copy-to-src.txt")
		defer os.Remove(srcFilePath)

		err := os.WriteFile(
			srcFilePath,
			[]byte("src"),
			fs.ModePerm,
		)
		assert.Nil(t, err)

		dstFilePath := filepath.Join(wd, "copy-to-dst.txt")
		defer os.Remove(dstFilePath)

		mockFileFactory := contract.NewMockIFileFactory(ctrl)
		mockFilePath := contract.NewMockIFilePath(ctrl)
		file := newFileEntry(mockFileFactory, mockFilePath, srcFilePath)

		mockFile := contract.NewMockIFileEntry(ctrl)
		mockFileFactory.EXPECT().BuildFileEntry(dstFilePath).Return(mockFile)

		mockFile.EXPECT().Exists().Return(false)

		mockFile.EXPECT().GetPath().Return(dstFilePath)

		mockFilePath.EXPECT().Join(dstFilePath).Return(dstFilePath)

		err = file.CopyTo(dstFilePath)
		assert.Nil(t, err)

		res, err := os.ReadFile(dstFilePath)
		assert.Nil(t, err)
		assert.Equal(
			t,
			string(res),
			"src",
		)
	})
}

func Test_fileEntry_Exists(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		wd, _ := os.Getwd()
		res := newFileEntry(
			nil,
			nil,
			path.Join(wd, "file.go"),
		).Exists()
		assert.True(t, res)
	})
}
