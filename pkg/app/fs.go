package app

import (
	"context"
	"io"
	"os"
	"sync"
	"time"

	"github.com/cloudwego/hertz/internal/nocopy"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/protocol"
)

var (
	errDirIndexRequired   = errors.NewPublic("directory index required")
	errNoCreatePermission = errors.NewPublic("no 'create file' permissions")

	rootFSOnce sync.Once
	rootFS     = &FS{
		Root:               "/",
		GenerateIndexPages: true,
		Compress:           true,
		AcceptByteRange:    true,
	}
	rootFSHandler  HandlerFunc
	strInvalidHost = []byte("invalid-host")
)

type PathRewriteFunc func(ctx *RequestContext) []byte

type FS struct {
	noCopy nocopy.NoCopy //lint:ignore U1000 until noCopy is used

	Root string

	IndexNames []string

	GenerateIndexPages bool

	Compress bool

	AcceptByteRange bool

	PathRewrite PathRewriteFunc

	PathNotFound HandlerFunc

	CacheDuration time.Duration

	CompressedFileSuffix string

	once sync.Once
	h    HandlerFunc
}

type byteRangeUpdater interface {
	UpdateByteRange(startPos, endPos int) error
}

type fsSmallFileReader struct {
	ff       *fsFile
	startPos int
	endPos   int
}

func (r *fsSmallFileReader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *fsSmallFileReader) UpdateByteRange(startPos, endPos int) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *fsSmallFileReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *fsSmallFileReader) WriteTo(w io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ServeFile(ctx *RequestContext, path string) { _ = "STUB: not implemented"; return }

func (fs *FS) NewRequestHandler() HandlerFunc { _ = "STUB: not implemented"; return *new(HandlerFunc) }

func (fs *FS) initRequestHandler() { _ = "STUB: not implemented"; return }

type fsHandler struct {
	root                 string
	indexNames           []string
	pathRewrite          PathRewriteFunc
	pathNotFound         HandlerFunc
	generateIndexPages   bool
	compress             bool
	acceptByteRange      bool
	cacheDuration        time.Duration
	compressedFileSuffix string

	cache           map[string]*fsFile
	compressedCache map[string]*fsFile
	cacheLock       sync.Mutex

	smallFileReaderPool sync.Pool
}

type bigFileReader struct {
	f  *os.File
	ff *fsFile
	r  io.Reader
	lr io.LimitedReader
}

func (r *bigFileReader) UpdateByteRange(startPos, endPos int) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *bigFileReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *bigFileReader) WriteTo(w io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *bigFileReader) Close() error { _ = "STUB: not implemented"; return nil }

func (h *fsHandler) cleanCache(pendingFiles []*fsFile) []*fsFile {
	_ = "STUB: not implemented"
	return nil
}

func (h *fsHandler) compressAndOpenFSFile(filePath string) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fsHandler) newCompressedFSFile(filePath string) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fsHandler) compressFileNolock(f *os.File, fileInfo os.FileInfo, filePath, compressedFilePath string) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fsHandler) openFSFile(filePath string, mustCompress bool) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fsHandler) newFSFile(f *os.File, fileInfo os.FileInfo, compressed bool) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fsHandler) createDirIndex(base *protocol.URI, dirPath string, mustCompress bool) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fsHandler) openIndexFile(ctx *RequestContext, dirPath string, mustCompress bool) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ff *fsFile) decReadersCount() { _ = "STUB: not implemented"; return }

func (ff *fsFile) bigFileReader() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (ff *fsFile) NewReader() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (ff *fsFile) smallFileReader() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (h *fsHandler) handleRequest(c context.Context, ctx *RequestContext) {
	_ = "STUB: not implemented"
	return
}

type fsFile struct {
	h             *fsHandler
	f             *os.File
	dirIndex      []byte
	contentType   string
	contentLength int
	compressed    bool

	lastModified    time.Time
	lastModifiedStr []byte

	t            time.Time
	readersCount int

	bigFiles     []*bigFileReader
	bigFilesLock sync.Mutex
}

func (ff *fsFile) Release() { _ = "STUB: not implemented"; return }

func (ff *fsFile) isBig() bool { _ = "STUB: not implemented"; return false }

func cleanCacheNolock(cache map[string]*fsFile, pendingFiles, filesToRelease []*fsFile, cacheDuration time.Duration) ([]*fsFile, []*fsFile) {
	_ = "STUB: not implemented"
	return nil, nil
}

func stripTrailingSlashes(path []byte) []byte { _ = "STUB: not implemented"; return nil }

func isFileCompressible(f *os.File, minCompressRatio float64) bool {
	_ = "STUB: not implemented"
	return false
}

//nolint:errcheck

var (
	filesLockMap     = make(map[string]*sync.Mutex)
	filesLockMapLock sync.Mutex
)

func getFileLock(absPath string) *sync.Mutex { _ = "STUB: not implemented"; return nil }

func fileExtension(path string, compressed bool, compressedFileSuffix string) string {
	_ = "STUB: not implemented"
	return ""
}

func readFileHeader(f *os.File, compressed bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fsModTime(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func ParseByteRange(byteRange []byte, contentLength int) (startPos, endPos int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func NewVHostPathRewriter(slashesCount int) PathRewriteFunc {
	_ = "STUB: not implemented"
	return *new(PathRewriteFunc)
}

func stripLeadingSlashes(path []byte, stripSlashes int) []byte {
	_ = "STUB: not implemented"
	return nil
}

func ServeFileUncompressed(ctx *RequestContext, path string) { _ = "STUB: not implemented"; return }

func NewPathSlashesStripper(slashesCount int) PathRewriteFunc {
	_ = "STUB: not implemented"
	return *new(PathRewriteFunc)
}
