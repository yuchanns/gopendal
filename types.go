package opendal

import (
	"unsafe"

	"github.com/jupiterrider/ffi"
)

var (
	typeResultOperatorNew = ffi.Type{
		Type: ffi.Struct,
		Elements: &[]*ffi.Type{
			&ffi.TypePointer,
			&ffi.TypePointer,
			nil,
		}[0],
	}

	typeResultRead = ffi.Type{
		Type: ffi.Struct,
		Elements: &[]*ffi.Type{
			&ffi.TypePointer,
			&ffi.TypePointer,
			nil,
		}[0],
	}

	typeBytes = ffi.Type{
		Type: ffi.Struct,
		Elements: &[]*ffi.Type{
			&ffi.TypePointer,
			&ffi.TypePointer,
			nil,
		}[0],
	}

	typeResultStat = ffi.Type{
		Type: ffi.Struct,
		Elements: &[]*ffi.Type{
			&ffi.TypePointer,
			&ffi.TypePointer,
			nil,
		}[0],
	}

	typeResultList = ffi.Type{
		Type: ffi.Struct,
		Elements: &[]*ffi.Type{
			&ffi.TypePointer,
			&ffi.TypePointer,
			nil,
		}[0],
	}

	typeResultListerNext = ffi.Type{
		Type: ffi.Struct,
		Elements: &[]*ffi.Type{
			&ffi.TypePointer,
			&ffi.TypePointer,
			nil,
		}[0],
	}

	typeResultOperatorReader = ffi.Type{
		Type: ffi.Struct,
		Elements: &[]*ffi.Type{
			&ffi.TypePointer,
			&ffi.TypePointer,
			nil,
		}[0],
	}

	typeResultReaderRead = ffi.Type{
		Type: ffi.Struct,
		Elements: &[]*ffi.Type{
			&ffi.TypePointer,
			&ffi.TypePointer,
			nil,
		}[0],
	}

	typeResultIsExist = ffi.Type{
		Type: ffi.Struct,
		Elements: &[]*ffi.Type{
			&ffi.TypeUint8,
			&ffi.TypePointer,
			nil,
		}[0],
	}

	typeCapability = ffi.Type{
		Type: ffi.Struct,
		Elements: &[]*ffi.Type{
			&ffi.TypeUint8,   // stat
			&ffi.TypeUint8,   // stat_with_if_match
			&ffi.TypeUint8,   // stat_with_if_none_match
			&ffi.TypeUint8,   // read
			&ffi.TypeUint8,   // read_with_if_match
			&ffi.TypeUint8,   // read_with_if_match_none
			&ffi.TypeUint8,   // read_with_override_cache_control
			&ffi.TypeUint8,   // read_with_override_content_disposition
			&ffi.TypeUint8,   // read_with_override_content_type
			&ffi.TypeUint8,   // write
			&ffi.TypeUint8,   // write_can_multi
			&ffi.TypeUint8,   // write_can_empty
			&ffi.TypeUint8,   // write_can_append
			&ffi.TypeUint8,   // write_with_content_type
			&ffi.TypeUint8,   // write_with_content_disposition
			&ffi.TypeUint8,   // write_with_cache_control
			&ffi.TypePointer, // write_multi_max_size
			&ffi.TypePointer, // write_multi_min_size
			&ffi.TypePointer, // write_multi_align_size
			&ffi.TypePointer, // write_total_max_size
			&ffi.TypeUint8,   // create_dir
			&ffi.TypeUint8,   // delete
			&ffi.TypeUint8,   // copy
			&ffi.TypeUint8,   // rename
			&ffi.TypeUint8,   // list
			&ffi.TypeUint8,   // list_with_limit
			&ffi.TypeUint8,   // list_with_start_after
			&ffi.TypeUint8,   // list_with_recursive
			&ffi.TypeUint8,   // presign
			&ffi.TypeUint8,   // presign_read
			&ffi.TypeUint8,   // presign_stat
			&ffi.TypeUint8,   // presign_write
			&ffi.TypeUint8,   // batch
			&ffi.TypeUint8,   // batch_delete
			&ffi.TypePointer, // batch_max_operations
			&ffi.TypeUint8,   // blocking
			nil,
		}[0],
	}
)

type opendalCapability struct {
	stat                               uint8
	statWithIfmatch                    uint8
	statWithIfNoneMatch                uint8
	read                               uint8
	readWithIfmatch                    uint8
	readWithIfMatchNone                uint8
	readWithOverrideCacheControl       uint8
	readWithOverrideContentDisposition uint8
	readWithOverrideContentType        uint8
	write                              uint8
	writeCanMulti                      uint8
	writeCanEmpty                      uint8
	writeCanAppend                     uint8
	writeWithContentType               uint8
	writeWithContentDisposition        uint8
	writeWithCacheControl              uint8
	writeMultiMaxSize                  uint
	writeMultiMinSize                  uint
	writeMultiAlignSize                uint
	writeTotalMaxSize                  uint
	createDir                          uint8
	delete                             uint8
	copy                               uint8
	rename                             uint8
	list                               uint8
	listWithLimit                      uint8
	listWithStartAfter                 uint8
	listWithRecursive                  uint8
	presign                            uint8
	presignRead                        uint8
	presignStat                        uint8
	presignWrite                       uint8
	batch                              uint8
	batchDelete                        uint8
	batchMaxOperations                 uint
	blocking                           uint8
}

type resultOperatorNew struct {
	op    *opendalOperator
	error *opendalError
}

type opendalOperator struct {
	ptr uintptr
}

type resultRead struct {
	data  *opendalBytes
	error *opendalError
}

type opendalReader struct {
	inner uintptr
}

type resultOperatorReader struct {
	reader *opendalReader
	error  *opendalError
}

type resultReaderRead struct {
	size  uint
	error *opendalError
}

type resultIsExist struct {
	is_exist uint8
	error    *opendalError
}

type resultStat struct {
	meta  *opendalMetadata
	error *opendalError
}

type opendalMetadata struct {
	inner uintptr
}

type opendalBytes struct {
	data *byte
	len  uintptr
}

type opendalError struct {
	code    int32
	message opendalBytes
}

type opendalOperatorInfo struct {
	inner uintptr
}

type opendalResultList struct {
	lister *opendalLister
	err    *opendalError
}

type opendalLister struct {
	inner uintptr
}

type opendalResultListerNext struct {
	entry *opendalEntry
	err   *opendalError
}

type opendalEntry struct {
	inner uintptr
}

func toOpendalBytes(data []byte) opendalBytes {
	var ptr *byte
	l := len(data)
	if l > 0 {
		ptr = &data[0]
	}
	return opendalBytes{
		data: ptr,
		len:  uintptr(l),
	}
}

func parseBytes(b *opendalBytes) (data []byte) {
	if b == nil || b.len == 0 {
		return nil
	}
	data = make([]byte, b.len)
	copy(data, unsafe.Slice(b.data, b.len))
	return
}
