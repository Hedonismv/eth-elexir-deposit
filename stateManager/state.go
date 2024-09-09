package stateManager

import (
	"encoding/json"
	"os"
)

type State map[string]bool

func LoadState(filePath string) (State, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(State), nil //! return empty state if file does not exist
		}
		return nil, err
	}
	defer file.Close()

	var state State
	if err := json.NewDecoder(file).Decode(&state); err != nil {
		return nil, err
	}

	return state, nil
}

func SaveState(filePath string, state State) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(state)
}
