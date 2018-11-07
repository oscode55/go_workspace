/*类访问标识符常量16位*/
package heap

const (
	ACC_PUBLIC       = 0x0001 // class field method              1
	ACC_PRIVATE      = 0x0002 //       field method		        10
	ACC_PROTECTED    = 0x0004 //       field method            100
	ACC_STATIC       = 0x0008 //       field method           1000
	ACC_FINAL        = 0x0010 // class field method          10000
	ACC_SUPER        = 0x0020 // class                      100000
	ACC_SYNCHRONIZED = 0x0020 //             method         100000
	ACC_VOLATILE     = 0x0040 //       field               1000000
	ACC_BRIDGE       = 0x0040 //             method        1000000
	ACC_TRANSIENT    = 0x0080 //       field              10000000
	ACC_VARARGS      = 0x0080 //             method       10000000
	ACC_NATIVE       = 0x0100 //             method      100000000
	ACC_INTERFACE    = 0x0200 // class                   1000000000
	ACC_ABSTRACT     = 0x0400 // class       method     10000000000
	ACC_STRICT       = 0x0800 //             method      1
	ACC_SYNTHETIC    = 0x1000 // class field method      1
	ACC_ANNOTATION   = 0x2000 // class                   1
	ACC_ENUM         = 0x4000 // class field             1
)
