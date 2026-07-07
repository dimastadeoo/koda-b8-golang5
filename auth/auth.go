package auth

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
)

type Users struct {
	Username string
	Password string
}

// untuk simpan data user
type AuthService struct {
	DB map[string]*Users
}

// fungsi hash passwd jadi md5
func hashPasswd (password string) string{
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// method register
func (s *AuthService) Register(user string, passwd string){
	if _, checkData := s.DB[user]; checkData{
		panic("User sudah terdaftar")
	}

	hashPwd := hashPasswd(passwd) //input passwd kita hash dulu, baru masukkan ke data user
	s.DB[user] = &Users{Username: user, Password: hashPwd}
	fmt.Println("Registrasi sukses!")
}

// method login
func (s *AuthService) Login(user string, passwd string) error {
	username, checkData := s.DB[user]
	if !checkData {
		return errors.New("Username tidak ditemukan")
	}

	if username.Password != hashPasswd(passwd){
		return errors.New("Password tidak sesuai")
	}
	fmt.Printf("Selamat Datang %s!\n", user)
	return nil
}