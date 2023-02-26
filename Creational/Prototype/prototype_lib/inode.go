package prototype_lib

type Inode interface {
	Print(string)
	Clone() Inode
}
