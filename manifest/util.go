package manifest

import "os"

func createFile(path string) {
	_, err := os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return
		}
		defer file.Close()
	}

}
