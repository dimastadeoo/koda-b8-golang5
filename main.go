package main

import (
	"authentication-flow/auth"
	"fmt"
	"os"
)

func userMenu(s *auth.AuthService, u *auth.Users){
	LoopUserMenu:
	for {
		fmt.Println("------------------------------------------------------")
		fmt.Println("1. List User \n2. Ganti Password \n3. Logout")
		choice := auth.ReadString("Pilih: ")
		switch choice {
		case "1":
			fmt.Printf("* %-15s | %-20s\n", "USERNAME", "NAMA LENGKAP")
			for _, Alluser := range s.DB{
				fmt.Printf("- %-15s | %-20s\n", Alluser.Username, Alluser.Fullname)
			}
		case "2":
			op := auth.ReadString("Masukkan Password Lama: ")
			np := auth.ReadString("Masukkan Password Baru: ")
			cp := auth.ReadString("Konfirmasi Password Baru: ")
			if err := s.ChangePasswd(u.Username, op, np, cp); err != nil{
				fmt.Println(err)
			}
		case "3":
			auth.WaitForKey("Tekan Enter untuk Logout dan kembali ke menu awal....")
			break LoopUserMenu
		default: 
			fmt.Println("Pilih hanya angka 1 - 3")
			continue
		}	
	}
}


func main(){
	service := &auth.AuthService{DB: []*auth.Users{}}
	defer func (){
		if r := recover(); r != nil {
			fmt.Println("Error: ", r)
		}
	}()

	for {
		fmt.Println("-------Selamat Datang---------")
		fmt.Println("1. Register\n2. Login\n3. Exit")
		choice := auth.ReadString("Pilih: ")

		switch choice {
		case "1":
			fn := auth.ReadString("Nama Lengkap: ")
			un := auth.ReadString("Username: ")
			pw := auth.ReadString("Password: ")
			cp := auth.ReadString("Konfirmasi Password: ")
			service.Register(fn, un, pw, cp)
			auth.WaitForKey("Tekan Enter untuk Kembali ke menu awal....")
		case "2":
			un := auth.ReadString("Username: ")
			pw := auth.ReadString("Password: ")
			user, err := service.Login(un, pw)
			if err != nil{
				fmt.Println(err)
			}else{
				userMenu(service, user)
			}
		case "3":
			os.Exit(0)
		default: 
			fmt.Println("Pilih hanya angka 1 - 3")
			continue
		}
	}
}