//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------

package fdset

//-----------------------------------------------------------------------------

/*
#include <sys/select.h>

static void _FD_CLR(int fd, void *set) {
  if ((unsigned int)fd < (unsigned int)FD_SETSIZE) {
        FD_CLR(fd, (fd_set*)set);
  }
}

static void _FD_SET(int fd, void *set) {
  if ((unsigned int)fd < (unsigned int)FD_SETSIZE) {
        FD_SET(fd, (fd_set*)set);
  }
}

static int _FD_ISSET (int fd, void *set) {
  if ((unsigned int)fd < (unsigned int)FD_SETSIZE) {
        return FD_ISSET(fd, (fd_set*)set);
  }
  return 0;
}

*/
import "C"

//-----------------------------------------------------------------------------

import (
	"syscall"
	"unsafe"
)

//-----------------------------------------------------------------------------

func Clr(fd int, set *syscall.FdSet) {
	C._FD_CLR(C.int(fd), unsafe.Pointer(set))
}

func Set(fd int, set *syscall.FdSet) {
	C._FD_SET(C.int(fd), unsafe.Pointer(set))
}

func IsSet(fd int, set *syscall.FdSet) bool {
	return C._FD_ISSET(C.int(fd), unsafe.Pointer(set)) != 0
}

//-----------------------------------------------------------------------------
