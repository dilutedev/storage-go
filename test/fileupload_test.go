package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	storage_go "github.com/supabase-community/storage-go"
)

var (
	rawUrl = "https://abc.supabase.co/storage/v1"
	token  = ""
)

func TestUpload(t *testing.T) {
	file, err := os.Open("dummy.txt")
	if err != nil {
		panic(err)
	}
	c := storage_go.NewClient(rawUrl, token, map[string]string{})
	resp, err := c.UploadFile("test1", "test.txt", file, "")
	assert.NotNil(t, err)
	fmt.Println(resp)

	// resp = c.UploadFile("test1", "hola.txt", []byte("hello world"))
	// fmt.Println(resp)
}

func TestUpdate(t *testing.T) {
	file, err := os.Open("dummy.txt")
	if err != nil {
		panic(err)
	}
	c := storage_go.NewClient(rawUrl, token, map[string]string{})
	resp, err := c.UpdateFile("test1", "test.txt", file, "")
	assert.NotNil(t, err)
	fmt.Println(resp)
}

func TestMoveFile(t *testing.T) {
	c := storage_go.NewClient(rawUrl, token, map[string]string{})
	resp, err := c.MoveFile("test1", "test.txt", "random/test.txt")
	assert.NotNil(t, err)
	fmt.Println(resp)
}

func TestSignedUrl(t *testing.T) {
	c := storage_go.NewClient(rawUrl, token, map[string]string{})
	resp, err := c.CreateSignedUrl("test1", "file_example_MP4_480_1_5MG.mp4", 120)
	assert.NotNil(t, err)
	fmt.Println(resp)
}

func TestPublicUrl(t *testing.T) {
	c := storage_go.NewClient(rawUrl, token, map[string]string{})
	resp := c.GetPublicUrl("shield", "book.pdf")

	fmt.Println(resp)
}

func TestDeleteFile(t *testing.T) {
	c := storage_go.NewClient(rawUrl, token, map[string]string{})
	resp, err := c.RemoveFile("shield", []string{"book.pdf"})
	assert.NotNil(t, err)
	fmt.Println(resp)
}

func TestListFile(t *testing.T) {
	c := storage_go.NewClient(rawUrl, token, map[string]string{})
	resp, err := c.ListFiles("test1", "", storage_go.FileSearchOptions{
		Limit:  10,
		Offset: 0,
		SortByOptions: storage_go.SortBy{
			Column: "",
			Order:  "",
		},
	})
	assert.NotNil(t, err)
	fmt.Println(resp)
}

func TestCreateUploadSignedUrl(t *testing.T) {
	c := storage_go.NewClient(rawUrl, token, map[string]string{"apiKey": token})
	resp, err := c.CreateSignedUploadUrl("your-bucket-id", "book.pdf")
	assert.NotNil(t, err)
	fmt.Println(resp, err)
}

func TestUploadToSignedUrl(t *testing.T) {
	c := storage_go.NewClient(rawUrl, token, map[string]string{"apiKey": token})
	file, err := os.Open("dummy.txt")
	if err != nil {
		panic(err)
	}
	resp, err := c.UploadToSignedUrl("signed-url-response", file)

	fmt.Println(resp, err)
}

func TestDownloadFile(t *testing.T) {
	c := storage_go.NewClient(rawUrl, token, map[string]string{})
	resp, err := c.DownloadFile("your-bucket-id", "book.pdf")
	if err != nil {
		t.Fatalf("DownloadFile failed: %v", err)
	}

	err = os.WriteFile("book.pdf", resp, 0644)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
}
