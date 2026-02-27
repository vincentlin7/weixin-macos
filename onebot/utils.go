package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func SaveBase64Image(base64Data string) (string, string, error) {
	rawContents := base64Data
	if strings.HasPrefix(base64Data, "base64://") {
		rawContents = strings.TrimPrefix(base64Data, "base64://")
	} else if idx := strings.Index(base64Data, ","); idx != -1 {
		rawContents = base64Data[idx+1:]
	}
	
	data, err := base64.StdEncoding.DecodeString(rawContents)
	if err != nil {
		return "", "", fmt.Errorf("base64 decode failed: %v", err)
	}
	salt := []byte(fmt.Sprintf("\n#md5_salt_%d_%d#", time.Now().UnixNano(), rand.Intn(10000)))
	data = append(data, salt...)
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(1000) // 生成 0-999 的随机数
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d_%d.%s", randomNumber, timestamp, DetectImageFormat(data))
	targetPath := config.ImagePath + fileName
	dir := filepath.Dir(targetPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", "", fmt.Errorf("create directory failed: %v", err)
	}
	
	err = os.WriteFile(targetPath, data, 0644)
	if err != nil {
		return "", "", fmt.Errorf("write file failed: %v", err)
	}
	
	md5Str, err := GetFileMD5(targetPath)
	if err != nil {
		return "", "", fmt.Errorf("get file md5 failed: %v", err)
	}
	
	return targetPath, md5Str, nil
}

func GetFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func DetectImageFormat(data []byte) string {
	if len(data) < 12 {
		return "unknown"
	}
	
	switch {
	case bytes.HasPrefix(data, []byte{0xFF, 0xD8, 0xFF}):
		return "jpg"
	case bytes.HasPrefix(data, []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}):
		return "png"
	case bytes.HasPrefix(data, []byte("GIF87a")) || bytes.HasPrefix(data, []byte("GIF89a")):
		return "gif"
	case bytes.HasPrefix(data, []byte{0x42, 0x4D}):
		return "bmp"
	case bytes.HasPrefix(data, []byte("RIFF")) && bytes.HasPrefix(data[8:], []byte("WEBP")):
		return "webp"
	default:
		return "unknown"
	}
}
