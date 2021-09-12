package recursion

import "os"

func recursive(baseDir string) {
	// get file and directory
	dirs, _ := os.ReadDir(baseDir)

	if len(dirs) > 0 {
		for _, dir := range dirs {
			dirName := baseDir + "/" + dir.Name()

			// hanya menampilkan file-file saja
			if !dir.IsDir() {
				println(dirName)
			}

			// running recursive function
			recursive(dirName)
		}
	}
}

func Run(baseDir string) {
	recursive(baseDir)
}
