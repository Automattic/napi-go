package entry

/*
#cgo CFLAGS: -DDEBUG
#cgo CFLAGS: -D_DEBUG
#cgo CFLAGS: -DV8_ENABLE_CHECKS
#cgo CFLAGS: -DNAPI_EXPERIMENTAL
#cgo CFLAGS: -I/usr/local/include/node
#cgo CXXFLAGS: -std=c++11

#cgo darwin LDFLAGS: -Wl,-undefined
#cgo darwin LDFLAGS: -Wl,dynamic_lookup
#cgo darwin LDFLAGS: -Wl,-search_paths_first
#cgo darwin LDFLAGS: -arch arm64

#cgo linux LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo LDFLAGS: -L${SRCDIR}
#cgo LDFLAGS: -stdlib=libc++
*/
import "C"
