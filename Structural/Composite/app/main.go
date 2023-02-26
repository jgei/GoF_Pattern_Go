// Composite is a structural design pattern that allows composing objects into a tree-like structure and work with
// the it as if it was a singular object.

// Composite became a pretty popular solution for the most problems that require building a tree structure.
// Composite’s great feature is the ability to run methods recursively over the whole tree structure and sum up the results.

// Conceptual Example

// Let’s try to understand the Composite pattern with an example of an operating system’s file system.
// In the file system, there are two types of objects: files and folders. There are cases when files and folders
// should be treated to be the same way. This is where the Composite pattern comes in handy.

// Imagine that you need to run a search for a particular keyword in your file system.
// This search operation applies to both files and folders. For a file, it will just look into the contents
// of the file; for a folder, it will go through all files of that folder to find that keyword.

package main

import (
	"composite_lib"
)

func main() {
	file1 := &composite_lib.File{Name: "File1"}
	file2 := &composite_lib.File{Name: "File2"}
	file3 := &composite_lib.File{Name: "File3"}

	folder1 := &composite_lib.Folder{
		Name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &composite_lib.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
