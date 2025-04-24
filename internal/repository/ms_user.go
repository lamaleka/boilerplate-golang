package repository

import (
	"strings"

	"github.com/lamaleka/boilerplate-golang/common/constant"
	"github.com/lamaleka/boilerplate-golang/common/errs"

	"gorm.io/gorm"
)

type MsUserRepository struct {
	db                   *DB
	MsEmployeeRepository *MsEmployeeRepository
}

func NewMsUserRepository(db *DB, msEmployeeRepository *MsEmployeeRepository) *MsUserRepository {
	return &MsUserRepository{
		db:                   db,
		MsEmployeeRepository: msEmployeeRepository,
	}
}

func (r *MsUserRepository) GetAll(params *MsUserParams) (*[]MsUser, error) {
	var data *[]MsUser
	query := r.db.Model(&data)
	query.Preload("Roles.Permissions").Preload("Roles").Preload("Employee")
	params.MapSearch(query)
	query.Count(&params.TotalData)
	query.Offset(params.Offset).Limit(params.PageSize)
	err := query.Find(&data).Error

	return data, err
}

func (r *MsUserRepository) FindByUsername(username string) (*MsUser, error) {
	var model *MsUser
	err := r.db.Preload("Roles.Permissions").
		Preload("Roles").
		Preload("Employee").
		Where(constant.WhereUsername, username).
		Where(constant.WhereIsActive, true).
		First(&model).Error
	if err == gorm.ErrRecordNotFound {
		err = errs.ErrRecordNotFound("Pengguna")
	}
	return model, err
}
func (r *MsUserRepository) FindByUUID(uuid string) (*MsUser, error) {
	var model *MsUser
	err := r.db.Preload("Roles.Permissions").Preload("Roles").Preload("Employee").First(&model, constant.WhereUUID, uuid).Error
	if err == gorm.ErrRecordNotFound {
		err = errs.ErrRecordNotFound("Master Data User")
	}
	return model, err
}

func (r *MsUserRepository) FindByUUIDs(uuids []string) ([]MsUser, error) {
	var data []MsUser
	err := r.db.Where(constant.WhereUUIDIn, uuids).Find(&data).Error
	return data, err
}

func (r *MsUserRepository) ResetPassword(data *MsUser, newPassword string) (*MsUser, error) {
	err := r.db.Model(&data).Update("password", newPassword).Error
	return data, err
}

func (r *MsUserRepository) Create(data *MsUser) (*MsUser, error) {
	trx := r.db.Begin()
	if err := trx.Create(&data).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			var oldData MsUser
			errFindRemoved := trx.Unscoped().
				Where(constant.WhereUsername, &data.Username).
				First(&oldData).Error
			if errFindRemoved != nil {
				defer trx.Rollback()
				return nil, errs.ErrUserExists
			}
			data.DeletedAt = nil
			errUpdateRemoved := trx.Model(&data).Unscoped().Updates(&data).Error
			if errUpdateRemoved != nil {
				defer trx.Rollback()
				return nil, errUpdateRemoved
			}
		} else {
			defer trx.Rollback()
			return nil, err
		}
	}
	updateData := map[string]interface{}{"user_id": data.ID}

	if err := trx.Model(&MsEmployee{}).Where(constant.WhereID, data.EmployeeID).Updates(updateData).Error; err != nil {
		defer trx.Rollback()
		return nil, err
	}

	if err := trx.Commit().Error; err != nil {
		defer trx.Rollback()
		return nil, err
	}
	return data, nil
}

func (r *MsUserRepository) Update(oldData *MsUser, newData *MsUser) (*MsUser, error) {
	err := r.db.Model(&oldData).Updates(&newData).Error
	return oldData, err
}

func (r *MsUserRepository) Delete(uuid string) error {
	var data MsUser
	trx := r.db.Begin()
	res := trx.Where(constant.WhereUUID, uuid).First(&data)
	if res.Error != nil {
		return res.Error
	}
	res = trx.Delete(&data)
	if res.RowsAffected == 0 {
		return errs.ErrNothingDeleted
	}
	if err := res.Error; err != nil {
		return err
	}
	updateData := map[string]interface{}{"user_id": nil}

	if err := trx.Model(&MsEmployee{}).Where(constant.WhereID, data.EmployeeID).Updates(updateData).Error; err != nil {
		defer trx.Rollback()
		return err
	}

	if err := trx.Commit().Error; err != nil {
		defer trx.Rollback()
		return err
	}
	return nil
}
