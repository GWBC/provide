package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func Get(ctx context.Context, url string, retry int) string {
	for range retry + 1 {
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return ""
		}

		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36 Edg/139.0.0.0")

		transport := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		cli := &http.Client{
			Transport: transport,
		}

		resp, err := cli.Do(req)
		if err != nil {
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		return string(body)
	}

	return ""
}

func removeAllWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

func removeComments(datas string) string {
	var result bytes.Buffer
	scanner := bufio.NewScanner(strings.NewReader(datas))

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		if len(trimmedLine) == 0 {
			continue
		}

		if commentIndex := strings.Index(trimmedLine, "//"); commentIndex != -1 {
			if commentIndex == 0 {
				continue
			}

			if commentIndex := strings.LastIndex(trimmedLine, "//"); commentIndex != -1 {
				byteLine := []byte(trimmedLine)
				index := max(commentIndex-1, 0)

				if byteLine[index] == ' ' || byteLine[index] == ',' {
					line = trimmedLine[:commentIndex]
				}

				if len(line) == 0 {
					continue
				}
			}
		}

		result.WriteString(line)
		result.WriteString("\n")
	}

	return strings.TrimSpace(result.String())
}

func Pwd() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}

	return filepath.Dir(exePath)
}

func removeGarbage(s string) string {
	var result []rune
	for _, r := range s {
		// 排除整个变体选择符区块
		if r >= '\uFE00' && r <= '\uFE0F' {
			continue
		}

		if unicode.IsPrint(r) || unicode.IsSpace(r) {
			result = append(result, r)
		}
	}

	return string(result)
}
