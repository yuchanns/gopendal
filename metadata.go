package opendal

import (
	"context"
	"time"
	"unsafe"

	"github.com/jupiterrider/ffi"
)

type Metadata struct {
	contentLength uint64
	isFile        bool
	isDir         bool
	lastModified  time.Time
}

func newMetadata(ctx context.Context, inner *opendalMetadata) *Metadata {
	getLength := getFFI[metaContentLength](ctx, symMetadataContentLength)
	isFile := getFFI[metaIsFile](ctx, symMetadataIsFile)
	isDir := getFFI[metaIsDir](ctx, symMetadataIsDir)
	getLastModified := getFFI[metaLastModified](ctx, symMetadataLastModified)

	var lastModified time.Time
	ms := getLastModified(inner)
	if ms != -1 {
		lastModified = time.UnixMilli(ms)
	}

	free := getFFI[metaFree](ctx, symMetadataFree)
	defer free(inner)

	return &Metadata{
		contentLength: getLength(inner),
		isFile:        isFile(inner),
		isDir:         isDir(inner),
		lastModified:  lastModified,
	}
}

func (m *Metadata) ContentLength() uint64 {
	return m.contentLength
}

func (m *Metadata) IsFile() bool {
	return m.isFile
}

func (m *Metadata) IsDir() bool {
	return m.isDir
}

func (m *Metadata) LastModified() time.Time {
	return m.lastModified
}

type metaContentLength func(m *opendalMetadata) uint64

const symMetadataContentLength = "opendal_metadata_content_length"

var withMetaContentLength = withFFI(ffiOpts{
	sym:    symMetadataContentLength,
	rType:  &ffi.TypeUint64,
	aTypes: []*ffi.Type{&ffi.TypePointer},
}, func(ctx context.Context, ffiCall func(rValue unsafe.Pointer, aValues ...unsafe.Pointer)) metaContentLength {
	return func(m *opendalMetadata) uint64 {
		var length uint64
		ffiCall(
			unsafe.Pointer(&length),
			unsafe.Pointer(&m),
		)
		return length
	}
})

type metaIsFile func(m *opendalMetadata) bool

const symMetadataIsFile = "opendal_metadata_is_file"

var withMetaIsFile = withFFI(ffiOpts{
	sym:    symMetadataIsFile,
	rType:  &ffi.TypeUint8,
	aTypes: []*ffi.Type{&ffi.TypePointer},
}, func(ctx context.Context, ffiCall func(rValue unsafe.Pointer, aValues ...unsafe.Pointer)) metaIsFile {
	return func(m *opendalMetadata) bool {
		var result uint8
		ffiCall(
			unsafe.Pointer(&result),
			unsafe.Pointer(&m),
		)
		return result == 1
	}
})

type metaIsDir func(m *opendalMetadata) bool

const symMetadataIsDir = "opendal_metadata_is_dir"

var withMetaIsDir = withFFI(ffiOpts{
	sym:    symMetadataIsDir,
	rType:  &ffi.TypeUint8,
	aTypes: []*ffi.Type{&ffi.TypePointer},
}, func(ctx context.Context, ffiCall func(rValue unsafe.Pointer, aValues ...unsafe.Pointer)) metaIsDir {
	return func(m *opendalMetadata) bool {
		var result uint8
		ffiCall(
			unsafe.Pointer(&result),
			unsafe.Pointer(&m),
		)
		return result == 1
	}
})

type metaLastModified func(m *opendalMetadata) int64

const symMetadataLastModified = "opendal_metadata_last_modified_ms"

var withMetaLastModified = withFFI(ffiOpts{
	sym:    symMetadataLastModified,
	rType:  &ffi.TypeSint64,
	aTypes: []*ffi.Type{&ffi.TypePointer},
}, func(ctx context.Context, ffiCall func(rValue unsafe.Pointer, aValues ...unsafe.Pointer)) metaLastModified {
	return func(m *opendalMetadata) int64 {
		var result int64
		ffiCall(
			unsafe.Pointer(&result),
			unsafe.Pointer(&m),
		)
		return result
	}
})

type metaFree func(m *opendalMetadata)

const symMetadataFree = "opendal_metadata_free"

var withMetaFree = withFFI(ffiOpts{
	sym:    symMetadataFree,
	rType:  &ffi.TypeVoid,
	aTypes: []*ffi.Type{&ffi.TypePointer},
}, func(ctx context.Context, ffiCall func(rValue unsafe.Pointer, aValues ...unsafe.Pointer)) metaFree {
	return func(m *opendalMetadata) {
		ffiCall(
			nil,
			unsafe.Pointer(&m),
		)
	}
})
