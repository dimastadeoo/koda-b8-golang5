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
	DB []*Users
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

	for _, data := range s.DB{
		if data.Username == user{
			panic("User sudah terdaftar")
		}
	}

	hashPwd := hashPasswd(passwd) //input passwd kita hash dulu, baru masukkan ke data user
	s.DB = append(s.DB, &Users{Fullname: fullname, Username: user, Password: hashPwd})
	fmt.Println("Registrasi sukses! ")
}

// method login
func (s *AuthService) Login(user string, passwd string) (*Users, error) {
	for _, data := range s.DB{
		if data.Username == user && data.Password == hashPasswd(passwd){
			fmt.Printf("Selamat Datang %s!\n", data.Fullname)
			return data, nil
		}
	}
	return nil, errors.New("User atau Password tidak sesuai")
}


// cange password
func (s *AuthService) ChangePasswd(user, oldPass, newPass, confirm string) error{
	if confirm != newPass {
		return  errors.New("Password Baru harus sesuai dengan confirm password")
	}

	for _, data := range s.DB{
		if data.Password == hashPasswd(oldPass){
			data.Password = hashPasswd(newPass)
			fmt.Printf("Password %s Berhasil diubah \n", data.Fullname)
			return nil
		}
	}
	return  errors.New("Password Lama Tidak sesuai")
	
}