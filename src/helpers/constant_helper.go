package helpers

const (
	DbDriverNotAvailable = "database driver (%s) is not available, check env file for more information!"

	CmdStartMsg = "halo %s, saya %s gunakan kode ini %d untuk melengkapi data pendaftaran! \n\njika anda sudah memiliki akun OkSetda Absensi dan Akun anda belum terintegrasi dengan Telegram Messenger untuk mendapatkan notifikasi, silahkan periksa email atau hubungi admin untuk mendapatkan kode integrasi. \n\nJika sudah memiliki kode integrasi silahkan ketikan: \nCONNECT#[USERNAME]#[INTR_CODE] \n\ncatatan: \ntanpa tanda kurung siku '[]' dan spasi \ncontohnya: \nCONNECT#aasumitro#45@34SAda \n\n/start - untuk memulai kembali \n/profile - untuk melihat status"

	CmdNotFound = `
		%s tidak dapat membalas, perintah (%s) tidak diketahui mohon periksa kembali daftar perintah :
			/start - untuk memulai kembali
			/profile - untuk melihat status
	`
)
