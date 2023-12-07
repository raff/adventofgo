package main

import (
	"fmt"
	"sort"

	"github.com/raff/adventofgo/advlib"
)

type item struct {
	name string
	size int
	dir  bool

	children []*item
	parent   *item
}

type sortable []*item

func (s sortable) Len() int           { return len(s) }
func (s sortable) Less(i, j int) bool { return s[i].size < s[j].size }
func (s sortable) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func walk(i *item, prefix string) {
	fmt.Println(prefix, i.name, i.size, i.dir)
	if i.dir {
		for _, c := range i.children {
			walk(c, prefix+" ")
			i.size += c.size
		}
		fmt.Println(prefix, " >", i.size)
	}
}

func listdirs(i *item, d *[]*item) {
	if i.dir {
		*d = append(*d, i)

		for _, c := range i.children {
			listdirs(c, d)
		}
	}
}

func main() {
	r := advlib.NewReader()

	var root, cwd *item

	for {
		line, err := r.Readline()
		if err != nil {
			fmt.Println(err)
			break
		}

		parts := advlib.Split(line)
		fmt.Println(parts)

		if parts[0] == "$" { // command
			if parts[1] == "cd" {
				switch parts[2] {
				case "/":
					root = &item{name: "/", dir: true}
					cwd = root

				case "..":
					if cwd.parent == nil {
						panic("cannot go up")
					}

					cwd = cwd.parent

				default:
					cur := &item{name: parts[2], dir: true, parent: cwd}
					cwd.children = append(cwd.children, cur)
					cwd = cur
				}
			} else if parts[1] == "ls" {
				// list current directory
			}
		} else if parts[0] == "dir" {
			// directory, we'll get there
		} else {
			size := advlib.ParseInt(parts[0])
			name := parts[1]
			cwd.children = append(cwd.children, &item{name: name, size: size, parent: cwd})
			//cwd.size += size
		}
	}

	walk(root, "")

	var dirs []*item

	listdirs(root, &dirs)

	// part 1
	fmt.Println()

	total := 0
	for _, d := range dirs {
		if d.size < 100000 {
			total += d.size
		}
	}

	fmt.Println("part 1:", total)

	// part 2
	fsize := 70000000
	required := 30000000
	available := fsize - root.size
	needed := required - available

	fmt.Println()
	fmt.Println("part 2:")
	fmt.Println("available:", available)
	fmt.Println("needed:", needed)

	sort.Sort(sortable(dirs))

	for _, d := range dirs {
		if d.size > needed {
			fmt.Println(d.name, d.size)
			break
		}
	}
}
