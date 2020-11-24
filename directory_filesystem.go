package phpfuncs

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// Basename - Returns trailing name component of path.
// Original : https://www.php.net/manual/en/function.basename.php
// Given a string containing the path to a file or directory, this function will return the trailing name component.
func Basename(path string) string {
	return filepath.Base(path)
}

// Chgrp - Changes file group.
// Original : https://www.php.net/manual/en/function.chgrp.php
// Attempts to change the group of the file filename to group.
// Only the superuser may change the group of a file arbitrarily; other users may change the group of a file to any group of which that user is a member.
func Chgrp(name string, uid, gid int) error {
	return os.Chown(name, uid, gid)
}

// Chmod - Changes file mode
// Original : https://www.php.net/manual/en/function.chmod.php
//Attempts to change the mode of the specified file to that given in mode.
func Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

// Chown - Changes file owner.
// Original : https://www.php.net/manual/en/function.chown.php
// Attempts to change the owner of the file filename to user user. Only the superuser may change the owner of a file.
func Chown(name string, uid int, gid int) error {
	return os.Chown(name, uid, gid)
}

// Copy - Copies file
// Original : https://www.php.net/manual/en/function.copy.php
// Makes a copy of the file source to dest.
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// Delete - Deletes a file.
// Original : https://www.php.net/manual/en/function.delete.php
// Deletes filename. Similar to the Unix C unlink() function. An E_WARNING level error will be generated on failure.
func Delete(name string) error {
	return os.Remove(name)
}

// DirName - Returns a parent directory's path
// Original : https://www.php.net/manual/en/function.dirname.php
// Given a string containing the path of a file or directory, this function will return the parent directory's path that is levels up from the current directory.
func DirName(path string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(path)
}

//DiskStatus struct
type DiskStatus struct {
	Free string `json:"Free"`
}

// DiskFreeSpace - Returns available space on filesystem or disk partition
// Original : https://www.php.net/manual/en/function.disk-free-space.php
// Given a string containing a directory, this function will return the number of bytes available on the corresponding filesystem or disk partition.
// DEVELOPER NOTE : PROBABLY WORKING ON ONLY LINUX AND MAC. TO-DO : WINDOWS
func DiskFreeSpace(path string) (disk DiskStatus) {
	stat := syscall.Statfs_t{}
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return
	}
	disk.Free = ByteCountIEC(stat.Bfree * uint64(stat.Bsize))
	return
}

// FClose - Closes an open file pointer
// Original : https://www.php.net/manual/en/function.fclose.php
// The file pointed to by handle is closed.
func FClose(file *os.File) error {
	return file.Close()
}

// FOpen - Opens file
// Original : https://www.php.net/manual/en/function.fopen.php
// fopen() binds a named resource, specified by filename, to a stream.
// NOT COMPLETED
// func FOpen(file string, mode int) (os.file, error) {
// 	f, err := os.OpenFile(file, mode, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := f.Close(); err != nil {
// 		log.Fatal(err)
// 	}

// 	defer file.Close()
// }

// FileExists - Checks whether a file or directory exists.
// Original : https://www.php.net/manual/en/function.file-exists.php
// Checks whether a file or directory exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// FileMime - Gets file modification time
// Original : https://www.php.net/manual/en/function.filemtime.php
// This function returns the time when the data blocks of a file were being written to, that is, the time when the content of the file was changed.
func FileMime(file string) time.Time {
	fi, err := os.Stat(file)
	if err != nil {
		return time.Time{}
	}
	return fi.ModTime()
}

// FilePerms - Gets file permissions.
// Original : https://www.php.net/manual/en/function.fileperms.php
// Gets permissions for the given file.
func FilePerms(path string) os.FileMode{
    p, _ := os.Open(path)
    m, _ := p.Stat()
		p.Close()
    return m.Mode().Perm()
}

// FileSize - Gets file permissions.
// Original : https://www.php.net/manual/en/function.filesize.php
// Gets permissions for the given file.
func FileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
    if err != nil {
        return 0, err
    }
    return fi.Size(), nil
}

// FileType - Gets file type.
// Original : https://www.php.net/manual/en/function.filetype.php
// Returns the type of the given file.
func FileType(fs string) (string, error) {
		f, err := os.Open(fs)
		if err != nil {
			return "", err
		}
		defer f.Close()
		buffer := make([]byte, 512)
		fff, err := f.Read(buffer)
		if err != nil {
			fmt.Println(fff)
			return "", err
		}
		contentType := http.DetectContentType(buffer)
		return contentType, nil
}


// IsDir - Tells whether the filename is a directory.
// Original : https://www.php.net/manual/en/function.is-dir.php
// Tells whether the given filename is a directory.
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.IsDir()
}

// IsFile - Tells whether the filename is a regular file.
// Original : https://www.php.net/manual/en/function.is-file.php
// Tells whether the given file is a regular file.
func IsFile(name string) bool{
	file, err := os.Stat(name)
	return err == nil && file.Mode().IsRegular()
}

// IsLink - Tells whether the filename is a symbolic link.
// Original : https://www.php.net/manual/en/function.is-link.php
// Tells whether the given file is a symbolic link.
func IsLink(path string) bool{
	_, err := os.Readlink(path)
	return err == nil
}

// IsReadable - Tells whether a file exists and is readable.
// Original : https://www.php.net/manual/en/function.is-readable.php
// Tells whether a file exists and is readable.
func IsReadable(path string) bool{
      file, err := os.OpenFile(path, os.O_WRONLY, 0666)
      file.Close()
			return err == nil
}

// IsWritable - Tells whether the filename is writable.
// Original : https://www.php.net/manual/en/function.is-writable.php
// Returns TRUE if the filename exists and is writable. The filename argument may be a directory name allowing you to check if a directory is writable.
func IsWritable(path string) bool {
	file, err := os.OpenFile(path, os.O_WRONLY, 0)
	file.Close()
	return err == nil
}

// IsWriteable - Tells whether the filename is writable.
// Original : https://www.php.net/manual/en/function.is-writeable.php
// Returns TRUE if the filename exists and is writable. The filename argument may be a directory name allowing you to check if a directory is writable.
func IsWriteable(path string) bool {
	return IsWritable(path)
}

// MkDir - Makes directory.
// Original : https://www.php.net/manual/en/function.mkdir.php
// Attempts to create the directory specified by pathname.
func MkDir(path string, mode os.FileMode) error {
	return os.Mkdir(path, mode)
}

// ByteCountIEC - Bytecount & Humanize Bytes
func ByteCountIEC(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
