package auth

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
)

type Users struct {
	Fullname string
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
func (s *AuthService) Register(fullname, user, passwd, confirm string){
	if confirm != passwd {
		panic("Password harus sama dengan confirm password")
	}
	
	if _, checkData := s.DB[user]; checkData{
		panic("User sudah terdaftar")
	}

	hashPwd := hashPasswd(passwd) //input passwd kita hash dulu, baru masukkan ke data user
	s.DB[user] = &Users{Fullname: fullname, Username: user, Password: hashPwd}
	fmt.Println("Registrasi sukses!")
}

// method login
func (s *AuthService) Login(user string, passwd string) (*Users, error) {
	username, checkData := s.DB[user]
	if !checkData {
		return nil, errors.New("Username tidak ditemukan")
	}

	if username.Password != hashPasswd(passwd){
		return nil, errors.New("Password tidak sesuai")
	}
	fmt.Printf("Selamat Datang %s!\n", user)
	return username, nil
}