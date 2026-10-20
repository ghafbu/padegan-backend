package filepkg

import (
	"fmt"
	"strings"
	"time"
)

const FormatFileName = "%s-%s.%s"

type File struct {
	FileEntity
	file       string
	extensions [][]string
}

func New() *File {
	return &File{}
}

func (f *File) SetExtensions(extensions [][]string) *File {
	f.extensions = extensions

	return f
}

func (f *File) SetFile(file string) *File {
	f.file = file

	return f
}

func (f *File) SetContentType(contentType string) *File {
	f.ContentType = contentType

	return f
}

func (f *File) SetSize(size int64) *File {
	f.Size = size

	return f
}

func (f *File) IsValidExtension(extension string) bool {
	extensionToLower := f.toLower(extension)

	for _, extensionRaw := range f.extensions {
		for _, extensionInRow := range extensionRaw {
			if extensionInRow == extensionToLower {
				return true
			}
		}
	}

	return false
}

func (f *File) GetInfoFile() FileEntity {
	f.setInfoFile()

	return FileEntity{
		FullName:    f.getFullName(),
		ContentType: f.getContentType(),
		Format:      f.getFormat(),
		Name:        f.getName(),
		Hash:        f.getHash(),
		Size:        f.getSize(),
	}
}

func (f *File) setInfoFile() {
	f.setHash()
	f.setFullName()
	f.setFormat()
	f.setName()
}

func (f *File) setFullName() {
	filename, extension := f.sanitizeExtension(f.toLower(f.file))

	f.FullName = fmt.Sprintf(
		"%s.%s",
		f.appendRandomIntegerToFilename(filename), extension,
	)
}

func (f *File) setName() {
	split := f.splitFile()
	f.Name = strings.Join(split[:len(split)-1], ".")
}

func (f *File) setFormat() {
	split := f.splitFile()
	f.Format = split[len(split)-1]
}

func (f *File) setHash() {
	f.Hash = fmt.Sprint(time.Now().Nanosecond())
}

func (f *File) getContentType() string {
	return f.ContentType
}

func (f *File) getFullName() string {
	return f.FullName
}

func (f *File) getName() string {
	return f.Name
}

func (f *File) getFormat() string {
	return f.Format
}

func (f *File) getHash() string {
	return f.Hash
}

func (f *File) getSize() int64 {
	return f.Size
}

func (f *File) sanitizeExtension(s string) (string, string) {
	split := f.splitFile()

	return split[0], split[len(split)-1]
}

func (f *File) toLower(s string) string {
	return strings.ToLower(s)
}

func (f *File) appendRandomIntegerToFilename(s string) string {
	return fmt.Sprintf("%s-%s", s, f.getHash())
}

func (f *File) splitFile() []string {
	return strings.Split(f.file, ".")
}
