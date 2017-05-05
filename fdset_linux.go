//-----------------------------------------------------------------------------
/*
FDSET macros/functions missing from the standard golang library.
*/
//-----------------------------------------------------------------------------

package fdset

//-----------------------------------------------------------------------------

/*

#include <sys/select.h>

static void _FD_ZERO(void *set) {
  FD_ZERO((fd_set*)set);
}

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

static int _FD_MAX(void *set) {
  int fd;
  for (fd = FD_SETSIZE-1; fd >= 0; fd --) {
    if (FD_ISSET(fd, (fd_set*)set)) {
      break;
    }
  }
  return fd;
}

*/
import "C"

//-----------------------------------------------------------------------------

import (
	"syscall"
	"unsafe"
)

//-----------------------------------------------------------------------------

// Zero all fd bits.
func Zero(set *syscall.FdSet) {
	C._FD_ZERO(unsafe.Pointer(set))
}

// Clear the fd bit.
func Clr(fd int, set *syscall.FdSet) {
	C._FD_CLR(C.int(fd), unsafe.Pointer(set))
}

// Set the fd bit.
func Set(fd int, set *syscall.FdSet) {
	C._FD_SET(C.int(fd), unsafe.Pointer(set))
}

// Is the fd set?
func IsSet(fd int, set *syscall.FdSet) bool {
	return C._FD_ISSET(C.int(fd), unsafe.Pointer(set)) != 0
}

// Return the highest numbered fd in the set, or -1 if no bits are set.
func Max(set *syscall.FdSet) int {
	return int(C._FD_MAX(unsafe.Pointer(set)))
}

//-----------------------------------------------------------------------------
