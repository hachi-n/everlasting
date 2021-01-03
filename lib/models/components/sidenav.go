package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/hachi-n/everlasting/lib/pkgerutil"
	_ "github.com/hachi-n/everlasting/pack"
	"github.com/rivo/tview"
	"path/filepath"
)

type SideNav struct {
	app       *tview.Application
	primitive *tview.TreeView
}

var (
	sideNavComponent *SideNav
)

const (
	readme = "README.md"
)

func NewSideNav(app *tview.Application, resourcesDir string) *SideNav {
	rootDir := resourcesDir
	root := tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	add := func(target *tview.TreeNode, path string) {
		fileinfos, err := pkgerutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, fileinfo := range fileinfos {
			name := fileinfo.Name()
			if name == ".keep" || name == readme {
				continue
			}

			node := tview.NewTreeNode(name).
				SetReference(filepath.Join(path, name)).
				SetSelectable(true)

			if fileinfo.IsDir() {
				node.SetColor(tcell.ColorGreen)
				node.SetSelectedFunc(func() {
					menuComponent.updateViewFromFile(filepath.Join(path, name, readme))
				})
			} else {
				node.SetSelectedFunc(func() {
					modal := NewModal(app, filepath.Join(path, name))
					modal.Display()
				})
			}
			target.AddChild(node)
		}
	}

	add(root, rootDir)

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			menuComponent.updateViewFromString(introduction)
			return
		}

		children := node.GetChildren()
		if len(children) == 0 {
			path := reference.(string)
			add(node, path)
		} else {
			node.SetExpanded(!node.IsExpanded())
		}
	})

	sideNavComponent = new(SideNav)
	sideNavComponent.app = app
	sideNavComponent.primitive = tree

	return sideNavComponent
}

func (p SideNav) Primitive() tview.Primitive {
	p.primitive = sideNavComponent.primitive
	return sideNavComponent.primitive
}
