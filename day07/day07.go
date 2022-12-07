package day07

import (
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day07/input.txt", func(line string) {
		input = append(input, line)
	})
	folders := Parse(input)
	fmt.Printf("Part 1: %d\n", Part1(folders))
	fmt.Printf("Part 2: %d\n", Part2(folders))
}

type File struct {
	Size int
	Path string
}
type Folder struct {
	Files      []File
	SubFolders []*Folder
	Path       string
}

func appenPath(currentPath, subPath string) string {
	result := fmt.Sprintf("%s/%s", currentPath, subPath)
	if currentPath == "/" {
		result = fmt.Sprintf("%s%s", currentPath, subPath)
	}
	return result
}

func Parse(outputs []string) map[string]*Folder {
	result := make(map[string]*Folder)
	currentPath := ""
	for _, line := range outputs {
		if line[0] == '$' { //Command
			if string(line[2:4]) == "cd" {
				switch string(line[5:]) {
				case "/":
					currentPath = "/"
					break
				case "..":
					li := strings.LastIndex(currentPath, "/")
					if li != -1 {
						currentPath = currentPath[:li]
						if currentPath == "" {
							currentPath = "/"
						}
					}
					break
				default:
					currentPath = appenPath(currentPath, string(line[5:]))
					break
				}
				if _, ok := result[currentPath]; !ok {
					result[currentPath] = &Folder{Path: currentPath}
				}
			}
		} else {
			if string(line[0:3]) == "dir" {
				subFolder := appenPath(currentPath, string(line[4:]))
				if _, ok := result[subFolder]; !ok {
					result[subFolder] = &Folder{
						Path: subFolder,
					}
				}
				result[currentPath].SubFolders = append(result[currentPath].SubFolders, result[subFolder])
			} else {
				f := strings.Split(line, " ")
				result[currentPath].Files = append(result[currentPath].Files, File{
					Size: utils.Atoi(f[0]),
					Path: f[1],
				})
			}
		}
	}
	return result
}

func folderSize(folder *Folder) int {
	size := 0
	for _, sub := range folder.SubFolders {
		size += folderSize(sub)
	}
	for _, file := range folder.Files {
		size += file.Size
	}
	return size
}

func Part1(folders map[string]*Folder) int {
	result := 0
	for _, folder := range folders {
		size := folderSize(folder)
		if size <= 100000 {
			result += size
		}
	}
	return result
}

func Part2(folders map[string]*Folder) int {
	sizes := make([]int, 0)
	totalSize := 0
	for _, folder := range folders {
		size := folderSize(folder)
		sizes = append(sizes, size)
		if folder.Path == "/" {
			totalSize = size
		}
	}
	leastSpace := 30000000 - (70000000 - totalSize)
	smallest := 30000000
	for _, s := range sizes {
		if s >= leastSpace {
			smallest = utils.Min(smallest, s)
		}
	}

	return smallest
}
