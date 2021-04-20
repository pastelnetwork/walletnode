package badger

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	v3Badger "github.com/dgraph-io/badger/v3"
	"github.com/pastelnetwork/go-commons/errors"
	"github.com/pastelnetwork/walletnode/storage"
)

var (
	chatDB storage.KeyValue
)

func tearDown(err error) {
	fmt.Println(errors.Errorf("error caused %v", err))
	os.Exit(0)
}

func TestMain(m *testing.M) {
	tmpDir, err := ioutil.TempDir("", "chat")
	if err != nil {
		tearDown(err)
	}
	fmt.Println("Created temporary directory", tmpDir)
	cfg := storage.NewConfig()
	cfg.ChatDBDir = tmpDir
	chatDB = NewChatDB(cfg)
	if chatDB == nil {
		tearDown(fmt.Errorf("can not start chatdb"))
	}
	code := m.Run()
	fmt.Println("Deleting", tmpDir)
	os.RemoveAll(tmpDir)
	os.Exit(code)
}

func TestChatDBSet(t *testing.T) {
	type args struct {
		key   []byte
		value []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Storing a new value into DB [hello-world]: OK",
			args: args{
				key:   []byte("hello"),
				value: []byte("world"),
			},
			wantErr: false,
		},
		{
			name: "Storing a new value into DB [abc-def]: OK",
			args: args{
				key:   []byte("abc"),
				value: []byte("def"),
			},
			wantErr: false,
		},
		{
			name: "Storing a new value into DB [___-&^%]: OK",
			args: args{
				key:   []byte("___"),
				value: []byte("&^%"),
			},
			wantErr: false,
		},
		{
			name: "Storing a new value into DB [-]: Error",
			args: args{
				key:   []byte(""),
				value: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := chatDB.Set(string(tt.args.key), tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatDBGet(t *testing.T) {
	var empty []byte
	type args struct {
		key []byte
	}
	tests := []struct {
		name       string
		args       args
		wantResult []byte
		wantErr    bool
	}{
		{
			name:       "Getting world by hello: OK",
			args:       args{key: []byte("hello")},
			wantErr:    false,
			wantResult: []byte("world"),
		},
		{
			name:       "Getting def by abc: OK",
			args:       args{key: []byte("abc")},
			wantErr:    false,
			wantResult: []byte("def"),
		},
		{
			name:       "Getting  &^% by ___: OK",
			args:       args{key: []byte("___")},
			wantErr:    false,
			wantResult: []byte("&^%"),
		},
		{
			name:       "Error: Key doesnt exists",
			args:       args{key: []byte("testKey")},
			wantErr:    true,
			wantResult: empty,
		},
		{
			name:       "Error: Empty key",
			args:       args{key: []byte("")},
			wantErr:    true,
			wantResult: empty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := chatDB.Get(string(tt.args.key))
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Get() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestChatDBDelete(t *testing.T) {

	type args struct {
		key []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Deleting abc: OK",
			args:    args{key: []byte("abc")},
			wantErr: false,
		},
		{
			name:    "Deleting abcd: OK",
			args:    args{key: []byte("abcd")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := chatDB.Delete(string(tt.args.key)); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := chatDB.Get(string(tt.args.key)); err != v3Badger.ErrKeyNotFound {
				t.Errorf("Delete() function didnt delete data by key %v", tt.args.key)
			}
		})
	}
}
