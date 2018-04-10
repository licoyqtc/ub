package api


import (
	"github.com/cybergarage/go-net-upnp/net/upnp/http"
	"fmt"
	"syscall"
	"encoding/json"
	gofstab "github.com/deniswernert/go-fstab"

)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bavail * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

type diskRsp struct {
	Err_no  int		`json:"err_no"`
	Err_msg string	`json:"err_msg"`
}


var GB float64 = 1024 * 1024 * 1024


func Getdiskinfo(req *http.Request,rsp http.ResponseWriter) {

	ret := lgRsp{}


	mounts ,_ :=gofstab.ParseSystem()
	for _,val := range mounts{
		//fmt.Printf("%v\n",val.File)
		if val.File == "swap"||val.File == "/dev/shm"||val.File == "/dev/pts"||val.File == "/proc"||val.File =="/sys"{
			continue
		}
		disk := DiskUsage(val.File)
		//fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
		//fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
		//fmt.Printf("Free: %.2f GB\n", float64(disk.Free)/float64(GB))
		diskall:=float64(disk.All)/float64(GB)
		diskfree:= float64(disk.Free)/float64(GB)
		dfpercent:=float64(diskfree/diskall)
		fmt.Printf("%s %.2f%%\n",val.File, dfpercent*100)
	}

	ret.Err_msg = "success"
	r , _ := json.Marshal(ret)
	rsp.Write(r)
}