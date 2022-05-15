package repositories

import (
	"errors"
	"github.com/aasumitro/svc-tgbotgo/src/domain"
	"github.com/aasumitro/svc-tgbotgo/src/helpers"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepositoryContract {
	return &userRepository{db: db}
}

func (repo userRepository) Profile(TelegramId string) (res int64, err error) {
	result := repo.db.First(&domain.User{}, "telegram_id = ?", TelegramId)

	return result.RowsAffected, result.Error
}

func (repo userRepository) Connect(Username string, Code string, TelegramId string) (res int64, err error) {
	user := domain.User{}
	findUser := repo.db.Where("username = ?", Username).First(&user)

	if findUser.Error != nil {
		return 0, errors.New(
			"Integrasi Telegram dan OkSetda Absensi **[GAGAL]** \n---USERNAME TIDAK DITEMUKAN---",
		)
	}

	if user.TelegramID == TelegramId {
		return 0, errors.New(
			"Akun Telegram dan Sistem OkSetda Absensi **[TELAH]** terintegrasi",
		)
	}

	if !helpers.SecretCodeVerify(Code, user.IntegrationCode) {
		return 0, errors.New(
			"Integrasi Telegram dan OkSetda Absensi **[GAGAL]** \n---KODE TIDAK VALID---",
		)
	}

	updateUser := repo.db.Model(domain.User{}).Where(
		"id = ?", user.ID,
	).Updates(map[string]interface{}{
		"telegram_id": TelegramId, "integration_code": nil,
	})

	if updateUser.Error != nil {
		return 0, errors.New(
			"Integrasi Telegram dan OkSetda Absensi **[GAGAL]** \n---TIDAK DAPAT MEMPERBAHARUI DATA---",
		)
	}

	return updateUser.RowsAffected, nil
}