package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Define o diretório base
	mDIREC := "<PATH>"

	// Busca pelos diretórios que contêm o arquivo external-secrets.yaml
	directories := findDirectories(mDIREC, "external-secrets.yaml")

	// Loop para percorrer cada diretório encontrado e remover os arquivos desejados
	for _, dir := range directories {
		// Verifica se o arquivo secret.yaml existe e o remove
		secretFile := filepath.Join(dir, "secret.yaml")
		if fileExists(secretFile) {
			err := os.Remove(secretFile)
			if err != nil {
				fmt.Printf("Erro ao remover %s: %v\n", secretFile, err)
			} else {
				fmt.Printf("Removed %s\n", secretFile)
			}
		}

		// Verifica se o arquivo configmap.yaml existe e o remove
		configMapFile := filepath.Join(dir, "configmap.yaml")
		if fileExists(configMapFile) {
			err := os.Remove(configMapFile)
			if err != nil {
				fmt.Printf("Erro ao remover %s: %v\n", configMapFile, err)
			} else {
				fmt.Printf("Removed %s\n", configMapFile)
			}
		}
	}
}

// findDirectories encontra todos os diretórios que contêm o arquivo especificado
func findDirectories(root string, filename string) []string {
	var directories []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == filename {
			directories = append(directories, filepath.Dir(path))
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Erro ao buscar diretórios: %v\n", err)
	}
	return directories
}

// fileExists verifica se um arquivo existe
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
