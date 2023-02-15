package ns

import "io/fs"

type FS struct {
	mounts map[string][]fs.FS
}

func (ns *FS) init() {
	if ns.mounts == nil {
		ns.mounts = make(map[string][]fs.FS)
	}
}

func (ns *FS) bind(mp string, fsys fs.FS, before bool) error {
	if !fs.ValidPath(mp) {
		return &fs.PathError{
			Op:   "bind",
			Path: mp,
			Err:  fs.ErrInvalid,
		}
	}

	if !before {
		ns.mounts[mp] = append(ns.mounts[mp], fsys)
		return nil
	}

	m := ns.mounts[mp]
	m = append(m, m[1:]...)
	m[0] = fsys
	ns.mounts[mp] = m
	return nil
}

func (ns *FS) BindBefore(mp string, fsys fs.FS) error {
	return ns.bind(mp, fsys, true)
}

func (ns *FS) BindAfter(mp string, fsys fs.FS) error {
	return ns.bind(mp, fsys, false)
}

func (ns *FS) Open(name string) (fs.File, error) {
	panic("Not implemented.")
}
