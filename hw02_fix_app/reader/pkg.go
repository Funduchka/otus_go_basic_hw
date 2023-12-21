package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Funduchka/otus_go_basic_hw/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	var data []types.Employee

	msg, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	err = json.Unmarshal(msg, &data)
	if err != nil {
		fmt.Printf("Error: %bytes", err)
		return nil, err
	}

	return data, nil
}
