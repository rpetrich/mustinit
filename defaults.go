package mustinit

type packageDefault struct {
	requirement initRequirement `mustinit:"true"`
	types       map[string]TypeRequirements
}

var defaultRequirements = map[string]packageDefault{
	// go packages
	"archive": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"archive/tar": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"archive/zip": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"bufio": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"builtin": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"bytes": packageDefault{
		requirement: initRequirementSkip,
	},
	"compress/bzip2": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"compress/flate": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"compress/gzip": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"compress/lzw": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"compress/zlib": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"container/heap": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"container/list": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"container/ring": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"context": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/aes": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/cipher": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/des": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/dsa": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/ecdsa": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/ed25519": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/elliptic": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/hmac": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/md5": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/rand": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/rc4": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/rsa": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/sha1": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/sha256": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/sha512": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/subtle": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/tls": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/x509": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"crypto/x509/pkix": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"database": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"database/sql": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"database/sql/driver": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"debug": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"debug/dwarf": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"LineEntry": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"debug/elf": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"Rela64": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Rela32": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"debug/gosym": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"debug/macho": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"debug/pe": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"FileHeader": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"OptionalHeader32": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"OptionalHeader64": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"debug/plan9obj": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"encoding": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"encoding/ascii85": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"encoding/asn1": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"encoding/base32": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"encoding/base64": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"encoding/binary": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"encoding/csv": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"encoding/gob": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"encoding/hex": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"encoding/json": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"encoding/pem": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"encoding/xml": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"errors": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"expvar": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"flag": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"FlagSet": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"fmt": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"go": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"go/ast": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"go/build": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"go/constant": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"go/doc": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"go/format": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"go/importer": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"go/parser": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"go/printer": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"go/scanner": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"go/token": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"go/types": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"Package": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"hash": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"hash/adler32": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"hash/crc32": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"hash/crc64": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"hash/fnv": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"hash/maphash": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"html": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"html/template": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"image": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"image/color": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"image/color/template": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"image/draw": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"image/gif": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"image/jpeg": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"image/png": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"index": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"index/suffixarray": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"io": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"io/ioutil": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"log": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"log/syslog": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"math": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"math/big": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"Float": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Int": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Rat": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"math/bits": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"math/cmplx": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"math/rand": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"mime": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"mime/multipart": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"mime/quotedprintable": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/http": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"connLRU": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"net/http/cgi": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/http/cookiejar": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/http/fcgi": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/http/httptest": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/http/httptrace": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/http/httputil": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/http/pprof": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/mail": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/rpc": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/rpc/jsonrpc": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/smtp": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/textproto": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"net/url": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"os": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"os/exec": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"os/signal": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"os/user": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"path": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"path/filepath": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"plugin": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"reflect": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"Value": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"regexp": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"regexp/syntax": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"runtime": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"MemStats": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"runtime/cgo": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"runtime/debug": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"runtime/msan": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"runtime/pprof": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"runtime/race": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"runtime/trace": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"sort": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"strconv": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"strings": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"Builder": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"sync": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"Cond": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Map": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Mutex": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Once": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Pool": TypeRequirements{
				IsRequired: true,
				RequiredFields: map[string]struct{}{
					"New": struct{}{},
				},
			},
			"RWMutex": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"WaitGroup": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"sync/atomic": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"Value": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"syscall": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"Cmsghdr": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"FdSet": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"PtraceRegs": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Rlimit": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Rusage": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Statfs_t": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Stat_t": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Sysinfo_t": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Timeval": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Tms": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Utsname": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Ustat_t": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Utimbuf": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"WaitStatus": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"syscall/js": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"testing": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"testing/iotest": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"testing/quick": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"text": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"text/scanner": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			"Scanner": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"text/tabwriter": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"text/template": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"text/template/parse": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"time": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
		types: map[string]TypeRequirements{
			// Duration and Time should really not be in this list, but so much
			// of the standard library abuses them
			"Duration": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
			"Time": TypeRequirements{
				IsRequired:     false,
				RequiredFields: nil,
			},
		},
	},
	"unicode": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"unicode/utf16": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"unicode/utf8": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
	"unsafe": packageDefault{
		requirement: initRequirementValues | initRequirementSkip,
	},
}
