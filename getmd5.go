package codesnippets

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// GetMD5 calculates the md5 hash (in bytes) of a file
func GetMD5(configFile string) ([]byte, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	configMD5 := hash.Sum(nil)
	return configMD5, nil
}

// WriteMD5File takes the md5sum and writes a file.md5 for a file
func WriteMD5File(filePath, fileName string, configMD5Sum []byte) error {
	configMD5SumString := fmt.Sprintf("%x %s\n", configMD5Sum, fileName)
	file, err := os.Create(filepath.Join(filePath, fileName+".md5"))
	if err != nil {
		return fmt.Errorf("cannot create file: %s", err)
	}
	defer file.Close()

	fmt.Fprintf(file, configMD5SumString)
	return nil
}

func TestGetMD5Config(t *testing.T) {
	filename := "text.txt"
	file, err := os.Create(filepath.Join(filename))
	if err != nil {
		t.Errorf("getMD5() Cannot create test file: %v", err)
		return
	}
	fmt.Fprintf(file, "some words")

	hash, err := GetMD5(file.Name())
	if err != nil {
		t.Errorf("getMD5() error = %v", err)
		return
	}
	if fmt.Sprintf("%x", hash) != "45274721987e0624326a67b74e146d61" {
		t.Error("expected 45274721987e0624326a67b74e146d61, got ", fmt.Sprintf("%x", hash))
	}
	defer os.Remove(file.Name())
}

func TestGetMD5(t *testing.T) {

	type args struct {
		configFile string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"md5test", args{"./testdata/leaf9-ncg1-hit3.cfg"}, leaf9md5, false},
		{"md5test", args{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMD5(tt.args.configFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMD5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteMD5File(t *testing.T) {
	type args struct {
		filePath     string
		fileName     string
		configMD5Sum []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"leaf9", args{"./testdata", "leaf9-ncg1-hit3.cfg", leaf9md5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteMD5File(tt.args.filePath, tt.args.fileName, tt.args.configMD5Sum); (err != nil) != tt.wantErr {
				t.Errorf("WriteMD5File() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	md5Files := []string{"leaf9-ncg1-hit3.cfg.md5"}

	for _, file := range md5Files {
		md5File := filepath.Join("testdata", file)
		actual, _ := ioutil.ReadFile(md5File)

		golden := filepath.Join("testdata/golden", file+".golden")
		if *update {
			ioutil.WriteFile(golden, actual, 0644)
		}
		expected, _ := ioutil.ReadFile(golden)

		if !bytes.Equal(actual, expected) {
			t.Errorf("WriteMD5File() write:\n%s\ndid not match golden file:\n%s", actual, expected)
		}
	}
}
