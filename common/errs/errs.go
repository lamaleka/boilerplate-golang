package errs

import (
	"errors"
	"fmt"
)

var (
	ErrNothingDeleted            = errors.New("Gagal menghapus data!")
	ErrNothingUpdated            = errors.New("Gagal mengupdate data!")
	ErrMaxUploadLimit            = errors.New("Ukuran maksimal unggah: 5MB!")
	ErrInvalidUsernameOrPassword = errors.New("Nama pengguna atau kata sandi salah!")
	ErrPasswordMismatch          = errors.New("Kata sandi tidak cocok!")
	ErrEncryptPassword           = errors.New("Gagal membuat kata sandi!")
	ErrUserNotExist              = errors.New("Pengguna tidak terdaftar!")
	ErrGenerateToken             = errors.New("Gagal generate token!")
	ErrUserNotVerified           = errors.New("User belum diverifikasi!")
	ErrInvalidToken              = errors.New("Token tidak valid atau kadaluarsa!")
	ErrUserAlreadyVerified       = errors.New("User sudah diverifikasi!")
	ErrFailedToProcessFile       = errors.New("Gagal memproses file!")
	ErrInvalidFile               = errors.New("Gagal mendapatkan file!")
	ErrInvalidDateTime           = errors.New("Format tanggal tidak valid. Contoh: 2006-01-02 15:04:05!")
	ErrUnableToRetrieveEmployees = errors.New("Gagal mengambil data Karyawan!")
	ErrEmployeeExists            = errors.New("Karyawan sudah terdaftar!")
	ErrUserExists                = errors.New("User sudah terdaftar!")
	ErrUnauthorized              = errors.New("Unauthorized!")
	ErrUnableToRetrieveData      = errors.New("Gagal mengambil data!")
	ErrInvalidDocumentStatus     = errors.New("Status dokumen tidak valid!")
	ErrInvalidDocumentType       = errors.New("Jenis dokumen tidak valid!")
)

func ErrRecordNotFound(instance string) error {
	return errors.New(fmt.Sprintf("%s tidak ditemukan", instance))
}
