package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/yosisa/rfm-server/rfm"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var listen = flag.String("listen", "127.0.0.1:50051", "Listen address")

type server struct{}

func (s *server) ReadDir(ctx context.Context, request *rfm.Request) (*rfm.DirInfo, error) {
	path := expandPath(filepath.Join(request.BaseDir, request.Target))
	du, err := newDiskUsage(path)
	if err != nil {
		return nil, err
	}
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	di := &rfm.DirInfo{
		Path: path,
		DiskUsage: &rfm.DiskUsage{
			Size: du.Size(),
			Free: du.Free(),
		},
		Items: make([]*rfm.FileInfo, len(fis)),
	}
	for i, fi := range fis {
		di.Items[i] = convertFileInfo(fi)
	}
	sort.Sort(byTypeThenName(di.Items))
	return di, nil
}

func (s *server) Find(ctx context.Context, req *rfm.FindRequest) (*rfm.DirInfo, error) {
	root := expandPath(req.BaseDir)
	plen := len(root) + 1
	var m matcher
	if req.Name != "" {
		m = newGlob(req.Name)
	}

	di := &rfm.DirInfo{
		Path: root,
	}
	dirs := make(map[string]*rfm.FileInfo)
	addWithParentDirs := func(r *rfm.FileInfo) {
		stack := []*rfm.FileInfo{r}
		var ok bool
		for {
			dir, _ := filepath.Split(r.Name)
			if dir = strings.TrimSuffix(dir, string(os.PathSeparator)); dir != "" {
				if r, ok = dirs[dir]; ok {
					stack = append(stack, r)
					delete(dirs, dir)
					continue
				}
			}
			break
		}
		for i := len(stack) - 1; i >= 0; i-- {
			di.Items = append(di.Items, stack[i])
		}
	}

	filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				return nil
			}
			return err
		}
		if len(path) < plen {
			return nil
		}

		matched := m == nil || m.Match(fi.Name())
		if !matched && !fi.IsDir() {
			return nil
		}
		rfi := convertFileInfo(fi)
		rfi.Name = path[plen:]
		if matched {
			addWithParentDirs(rfi)
		} else {
			dirs[rfi.Name] = rfi
		}
		return nil
	})

	return di, nil
}

func convertFileInfo(fi os.FileInfo) *rfm.FileInfo {
	owner, _ := getOwner(fi)
	return &rfm.FileInfo{
		Name:    fi.Name(),
		Size:    fi.Size(),
		Mode:    uint32(fi.Mode()),
		ModTime: float64(fi.ModTime().UnixNano()) * 1e-9,
		IsDir:   fi.IsDir(),
		Owner:   owner,
	}
}

var uidCache = struct {
	sync.Mutex
	m map[uint32]string
}{
	m: make(map[uint32]string),
}

func getOwner(fi os.FileInfo) (string, error) {
	uidCache.Lock()
	defer uidCache.Unlock()
	uid := fi.Sys().(*syscall.Stat_t).Uid
	if s, ok := uidCache.m[uid]; ok {
		return s, nil
	}
	u, err := user.LookupId(strconv.Itoa(int(uid)))
	if err != nil {
		return "", err
	}
	s := u.Username
	uidCache.m[uid] = s
	return s, nil
}

type diskUsage struct {
	stat *syscall.Statfs_t
}

func newDiskUsage(path string) (*diskUsage, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return nil, err
	}
	return &diskUsage{&stat}, nil
}

func (d *diskUsage) Size() uint64 {
	return d.stat.Blocks * uint64(d.stat.Bsize)
}

func (d *diskUsage) Free() uint64 {
	return d.stat.Bfree * uint64(d.stat.Bsize)
}

func expandPath(s string) string {
	if s[0] == '~' {
		home := os.Getenv("HOME")
		s = home + s[1:]
	}
	return filepath.Clean(s)
}

type byTypeThenName []*rfm.FileInfo

func (z byTypeThenName) Len() int      { return len(z) }
func (z byTypeThenName) Swap(i, j int) { z[i], z[j] = z[j], z[i] }
func (z byTypeThenName) Less(i, j int) bool {
	switch {
	case z[i].IsDir && !z[j].IsDir:
		return true
	case !z[i].IsDir && z[j].IsDir:
		return false
	default:
		return z[i].Name < z[j].Name
	}
}

func main() {
	flag.Parse()
	ln, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	rfm.RegisterFSServer(s, &server{})
	if err := s.Serve(ln); err != nil {
		log.Fatal(err)
	}
}
