package ossvc

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_file_Read(t *testing.T) {
	t.Run("[]byte", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockFileEntry := contract.NewMockIFileEntry(ctrl)
		self := NewFile(mockFileEntry)

		wd, _ := os.Getwd()
		filePath := filepath.Join(wd, "file-read")
		defer os.Remove(filePath)

		mockFileEntry.EXPECT().GetPath().Return(filePath)

		err := os.WriteFile(
			filePath,
			[]byte("[]byte"),
			os.ModePerm,
		)
		assert.Nil(t, err)

		var res []byte
		err = self.Read(&res)
		assert.Nil(t, err)
		assert.EqualValues(
			t,
			res,
			[]byte("[]byte"),
		)
	})

	t.Run("string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockFileEntry := contract.NewMockIFileEntry(ctrl)
		self := NewFile(mockFileEntry)

		wd, _ := os.Getwd()
		filePath := filepath.Join(wd, "file-read")
		defer os.Remove(filePath)

		mockFileEntry.EXPECT().GetPath().Return(filePath)

		err := os.WriteFile(
			filePath,
			[]byte("str"),
			os.ModePerm,
		)
		assert.Nil(t, err)

		var res string
		err = self.Read(&res)
		assert.Nil(t, err)
		assert.EqualValues(t, res, "str")
	})
}
