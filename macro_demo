#define _IOC_TYPECHECK(t) (sizeof(t))
#define _IOC_READ	2U
#define _IOC_WRITE	1U
#define _IOC_SIZEBITS	14
#define _IOC_NRBITS	8
#define _IOC_TYPEBITS	8

#define _IOC_NRMASK	((1 << _IOC_NRBITS)-1)
#define _IOC_TYPEMASK	((1 << _IOC_TYPEBITS)-1)
#define _IOC_SIZEMASK	((1 << _IOC_SIZEBITS)-1)
#define _IOC_DIRMASK	((1 << _IOC_DIRBITS)-1)

#define _IOC_NRSHIFT	0
#define _IOC_TYPESHIFT	(_IOC_NRSHIFT+_IOC_NRBITS)
#define _IOC_SIZESHIFT	(_IOC_TYPESHIFT+_IOC_TYPEBITS)
#define _IOC_DIRSHIFT	(_IOC_SIZESHIFT+_IOC_SIZEBITS)

#define _IOC(dir,type,nr,size) \
	(((dir)  << _IOC_DIRSHIFT) | \
	 ((type) << _IOC_TYPESHIFT) | \
	 ((nr)   << _IOC_NRSHIFT) | \
	 ((size) << _IOC_SIZESHIFT))

typedef __signed__ int __s32;

struct binder_version {
	/* driver protocol version -- increment with incompatible change */
	__s32       protocol_version;
};

#define _IOWR(type,nr,size)	_IOC(_IOC_READ|_IOC_WRITE,(type),(nr),(_IOC_TYPECHECK(size)))

//http://androidxref.com/kernel_3.18/xref/include/uapi/asm-generic/ioctl.h

//http://androidxref.com/kernel_3.18/xref/drivers/staging/android/uapi/binder.h
//http://androidxref.com/kernel_3.18/xref/drivers/staging/android/binder.c
//http://aospxref.com/android-9.0.0_r61/xref/frameworks/native/cmds/servicemanager/binder.c#100
//http://aospxref.com/android-9.0.0_r61/xref/frameworks/native/cmds/servicemanager/binder.c

#include <stdio.h>

int main() {
    unsigned int binder_ver = _IOWR('b', 9, struct binder_version);
    printf("%u", binder_ver);
    return 0;
}