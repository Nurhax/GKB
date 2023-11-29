package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var MenuUtama string
	var botNum, userNum int
	var GKBUser, GKBBot string
	var hasilPermainan int

	//Front Page atau Halaman Utama dari game
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Selamat Datang di Permainan Gunting Kertas Batu ᕕ(⌐■_■)ᕗ ♪♬")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("        .--'''''''''--.")
	fmt.Println("     .'      .---.      '.")
	fmt.Println("    /    .-----------.   \\")
	fmt.Println("   /        .-----.       \\")
	fmt.Println("   |       .-.   .-.       |")
	fmt.Println("   |      /   \\/   \\     |")
	fmt.Println("   \\    | .-. | .-. |    /")
	fmt.Println("     '-._| | | | | | |_.-'")
	fmt.Println("         | '-' | '-' |")
	fmt.Println("         \\___/ \\___/")
	fmt.Println("       _.-'  /   \\  `-._")
	fmt.Println("     .' _.--|     |--._ '.")
	fmt.Println("     ' _...-|     |-..._ '")
	fmt.Println("            |     |")
	fmt.Println("            '.___.'")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Mainkan game?")
	fmt.Println("Pilihan (Iya/Tidak) :")
	fmt.Scan(&MenuUtama)
	//Menu utama, kalau dia gak pilih iya atau tidak. User akan diminta untuk menginput ulang pilihan

	if MenuUtama == "Iya" || MenuUtama == "iya" {
		for MenuUtama == "Iya" || MenuUtama == "iya" {
			//User diminta untuk memasukan pilihannya diantara G K atau B
			fmt.Println("MASUKAN PILIHAN MU!")
			fmt.Println("GUNTING/KERTAS/BATU")
			fmt.Println("-----------------------------------------------------------")
			fmt.Println("    _______")
			fmt.Println("---'   ____)____")
			fmt.Println("          ______)")
			fmt.Println("       __________)")
			fmt.Println("      (____)")
			fmt.Println("---.__(___)")
			fmt.Println("-----------------------------------------------------------")
			fmt.Println("     _______")
			fmt.Println("---'    ____)____")
			fmt.Println("           ______)")
			fmt.Println("          _______)")
			fmt.Println("         _______)")
			fmt.Println("---.__________)")
			fmt.Println("-----------------------------------------------------------")
			fmt.Println("    _______")
			fmt.Println("---'   ____)")
			fmt.Println("      (_____)")
			fmt.Println("      (_____)")
			fmt.Println("      (____)")
			fmt.Println("---.__(___)")
			fmt.Println("-----------------------------------------------------------")
			fmt.Println("Pilihan:")
			fmt.Scan(&GKBUser)
			userNum = objekToAngka(GKBUser)
			fmt.Println("Pilihanmu adalah:", GKBUser)
			gambarGKB(GKBUser)

			//Generate Bot Number
			botNum = randomNumGen(botNum)
			GKBBot = angkaToObjek(botNum)

			//Hitung mundur
			fmt.Println("Lawanmu memilih...")
			time.Sleep(1 * time.Second)
			fmt.Println("3")
			time.Sleep(1 * time.Second)
			fmt.Println("2")
			time.Sleep(1 * time.Second)
			fmt.Println("1")

			//Hasil pilih bot
			fmt.Println("Lawanmu Telah Memilih:", GKBBot)
			gambarGKB(GKBBot)

			time.Sleep(2 * time.Second)
			//Hasil permainan
			hasilPermainan = hasilMain(botNum, userNum)

			gambarHasil(&hasilPermainan)

			time.Sleep(3 * time.Second)
			fmt.Println("Apakah kamu mau main lagi?")
			fmt.Println("Pilihan(Iya/Tidak):")
			fmt.Scan(&MenuUtama)
		}
	} else if MenuUtama == "Tidak" || MenuUtama == "tidak" {
	} else {
		fmt.Scan(&MenuUtama)
	}

	//Kasih WM NURHAX :V
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("TERIMAKASIH TELAH MEMAINKAN GAME INI >W< ")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("███╗░░██╗██╗░░░██╗██████╗░██╗░░██╗░█████╗░██╗░░██╗")
	fmt.Println("████╗░██║██║░░░██║██╔══██╗██║░░██║██╔══██╗╚██╗██╔╝")
	fmt.Println("██╔██╗██║██║░░░██║██████╔╝███████║███████║░╚███╔╝░")
	fmt.Println("██║╚████║██║░░░██║██╔══██╗██╔══██║██╔══██║░██╔██╗░")
	fmt.Println("██║░╚███║╚██████╔╝██║░░██║██║░░██║██║░░██║██╔╝╚██╗")
	fmt.Println("╚═╝░░╚══╝░╚═════╝░╚═╝░░╚═╝╚═╝░░╚═╝╚═╝░░╚═╝╚═╝░░╚═╝")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("SAMPAI JUMPA DILAIN HARI 	⊂(◉‿◉)つ")

	time.Sleep(2 * time.Second)
}

func randomNumGen(bilangan int) int {
	bilangan = rand.Intn(3)
	return bilangan
}

func objekToAngka(Pilihan string) int {
	var hasil int
	if Pilihan == "GUNTING" || Pilihan == "gunting" {
		hasil = 0
	} else if Pilihan == "KERTAS" || Pilihan == "kertas" {
		hasil = 1
	} else if Pilihan == "BATU" || Pilihan == "batu" {
		hasil = 2
	}

	return hasil
}

func angkaToObjek(Angka int) string {
	var hasil string
	if Angka == 0 {
		hasil = "GUNTING"
	} else if Angka == 1 {
		hasil = "KERTAS"
	} else if Angka == 2 {
		hasil = "BATU"
	}

	return hasil
}

func hasilMain(AngkaBot int, AngkaUser int) int {
	var hasil int
	if AngkaBot == (AngkaUser+1)%3 {
		hasil = 1
	} else if AngkaBot == AngkaUser {
		hasil = 2
	} else {
		hasil = 3
	}

	return hasil
}

func gambarHasil(angka *int) {
	if *angka == 1 {
		fmt.Println("SELAMAT KAMU MENANG!!! 	⊂(｡◕‿‿◕｡)つ ")
		fmt.Println("         .* *.               `o`o`")
		fmt.Println("         *. .*              o`o`o`o      ^,^,^ ")
		fmt.Println("           * \\              `o`o`     ^,^,^,^,^")
		fmt.Println("              \\     ***        |       ^,^,^,^,^")
		fmt.Println("               \\   *****       |        /^,^,^")
		fmt.Println("                \\   ***        |       /")
		fmt.Println("    ~@~*~@~      \\   \\         |      /")
		fmt.Println("  ~*~@~*~@~*~     \\   \\        |     /")
		fmt.Println("  ~*~@smd@~*~      \\   \\       |    /     #$#$#        .`'.;.")
		fmt.Println("  ~*~@~*~@~*~       \\   \\      |   /     #$#$#$#   00  .`,.',")
		fmt.Println("    ~@~*~@~ \\       \\   \\     |  /      /#$#$#   /|||  `.,'")
		fmt.Println("_____________\\________\\___\\____|_/______/_________|\\/\\___||______")
	} else if *angka == 2 {
		fmt.Println("SERI BANG...	┻━┻ ︵ヽ(`Д´)ﾉ︵ ┻━┻ ")
	} else if *angka == 3 {
		fmt.Println("KAMU KALAH BANG (҂◡_◡) ᕤ")
	} else {
		fmt.Println("NGACO KAMU BANG")
		fmt.Println("⣿⢋⡼⢣⡙⢦⡙⢆⠓⠈⠄⠀⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⢀⠀⠀⠀⠀⣀⠙⣆")
		fmt.Println("⢥⢫⡜⢣⡝⢦⡙⢦⣢⣤⣴⢬⣤⣥⣄⣀⣀⣢⣵⣶⣬⣤⣀⡀⠠⠀⢀⠈⠣⡈")
		fmt.Println("⢎⡳⡜⣣⢞⡣⣽⠟⢩⣿⣿⣶⡄⠀⣹⡿⣿⠁⣴⣿⣿⣧⡉⠻⣶⠀⠀⠂⠄⠱")
		fmt.Println("⡝⣲⠹⣔⢣⠳⣼⣿⡘⠿⣿⡿⠃⣰⡿⢀⠻⣧⠙⡿⣿⠟⠁⣰⡟⠀⠄⢈⠐⡈")
		fmt.Println("⡜⢥⡛⣬⢣⡛⣤⢛⠿⢷⡶⣴⣾⡿⢶⡥⣖⣹⣟⠳⠶⠶⠿⠋⠀⡀⠂⡀⠂⠐")
		fmt.Println("⡜⢣⡓⢦⢣⡝⢢⢏⠞⣰⣿⣿⣿⣿⣿⢾⣵⣾⣻⢿⣆⡀⠐⠀⠂⠀⠀⠄⠁⠂")
		fmt.Println("⡜⣣⢝⡪⢵⡘⣣⠎⣼⣿⣿⣿⣿⣿⣿⣻⣯⣿⣿⣿⣿⡿⠀⠀⠀⠀⠀⡘⠠⢀")
		fmt.Println("⡜⢥⢎⡵⢣⢣⡕⢪⠄⡉⠘⠉⠘⠻⠿⡟⠟⠁⠋⠋⠁⠀⠀⠀⠀⠀⠁⠀⢂⠄")
		fmt.Println("⡜⣊⠞⡴⣋⠶⣘⢣⠚⡄⡁⠂⠡⠀⠄⡀⠄⠂⠁⠀⠀⠀⠀⠀⠀⠀⠄⢁⠂⠄")
		fmt.Println("⡜⡥⢛⡴⡩⢞⡱⢪⠕⡢⠄⢁⠐⠈⠀⠄⠐⠠⠈⡀⠐⠀⠠⠀⠂⠈⠄⠠⠌⡐")
		fmt.Println("⡳⣜⠣⡜⡱⢣⡚⡕⢪⡑⠌⡀⠂⡁⠈⠄⠂⡀⢁⠀⠄⠂⡀⢀⠠⠁⡌⢢⠑⡠")
		fmt.Println("⣗⢮⢳⡙⣖⢣⡕⢪⢡⡘⣠⢁⡐⠈⡐⠈⠄⠐⠀⠈⠀⠂⠐⠀⠢⢅⡐⢢⡑⠤")
		fmt.Println("⣟⣮⢷⣻⣼⣳⡞⢯⠓⢏⠓⢋⡙⠛⠝⠛⠞⠳⢞⠳⠞⠴⠢⣌⠠⢀⡹⣰⠘⢦")
		fmt.Println("⣿⢾⡿⣽⣾⢷⣿⣯⣿⣞⡾⣖⡶⣹⡜⢶⡚⡵⢢⣆⢦⡴⣤⣤⢖⡮⣵⢣⡟⣮")
		fmt.Println("⣿⡿⣿⢿⣽⣿⣳⣿⢾⣻⣽⡿⣽⣷⣻⣷⣻⣽⣷⣻⣞⣷⣳⣞⡿⣼⣳⢻⣜⣳")

	}
}

func gambarGKB(Pilihan string) {
	if Pilihan == "GUNTING" || Pilihan == "gunting" {
		fmt.Println("    _______")
		fmt.Println("---'   ____)____")
		fmt.Println("          ______)")
		fmt.Println("       __________)")
		fmt.Println("      (____)")
		fmt.Println("---.__(___)")
	} else if Pilihan == "KERTAS" || Pilihan == "kertas" {
		fmt.Println("     _______")
		fmt.Println("---'    ____)____")
		fmt.Println("           ______)")
		fmt.Println("          _______)")
		fmt.Println("         _______)")
		fmt.Println("---.__________)")
	} else if Pilihan == "BATU" || Pilihan == "batu" {
		fmt.Println("---'   ____)")
		fmt.Println("      (_____)")
		fmt.Println("      (_____)")
		fmt.Println("      (____)")
		fmt.Println("---.__(___)")
	}
}
