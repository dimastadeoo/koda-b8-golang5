package auth

import (
	"fmt"
	"errors"
	"crypto/md5"
	"encoding/hex"
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
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func (s *AuthService) Register(user string, passwd string){
	if _, checkData := s.DB[user]; checkData{
		panic("User sudah terdaftar")
	}

	hashPwd := hashPasswd(passwd) //input passwd kita hash dulu, baru masukkan ke data user
	s.DB[user] = &Users{Username: user, Password: hashPwd}
	fmt.Println("Registrasi sukses!")

}

