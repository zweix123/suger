/*
json 并不符合 Go 的设计哲学, 应该将入参设置为符合io.Writer或者具有Encode方法的接口, 但是这很Pythonic[dog]
*/

package json

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save[T any](filepath string, structSlice []T) error {
	// 创建或打开一个文件用于写入
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Error creating file: %w", err)
	}
	defer file.Close()

	// 使用 json.NewEncoder 和 Encode 方法将切片编码为 JSON 格式
	encoder := json.NewEncoder(file)
	err = encoder.Encode(structSlice)
	if err != nil {
		return fmt.Errorf("Error encoding structs to JSON: %w", err)
	}
	return nil
}

func Load[T any](filepath string) ([]T, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var structSlice []T
	err = decoder.Decode(&structSlice)
	if err != nil {
		return nil, fmt.Errorf("Error decoding structs from JSON: %w", err)
	}
	return structSlice, nil
}
