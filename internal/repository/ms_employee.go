package repository

import (
	"strings"

	"github.com/lamaleka/boilerplate-golang/common/constant"
	"github.com/lamaleka/boilerplate-golang/common/errs"

	"gorm.io/gorm"
)

type MsEmployeeRepository struct {
	db *DB
}

func NewMsEmployeeRepository(db *DB) *MsEmployeeRepository {
	return &MsEmployeeRepository{db: db}
}

func (r *MsEmployeeRepository) GetAll(params *MsEmployeeParams) (*[]MsEmployee, error) {
	var data *[]MsEmployee
	query := r.db.Model(&data)
	params.MapSearch(query)
	query.Offset(params.Offset).Limit(params.PageSize)
	err := query.Find(&data).Error
	return data, err
}
func (r *MsEmployeeRepository) GetAllUnregistered(params *MsEmployeeParams) (*[]MsEmployee, error) {
	var data *[]MsEmployee
	query := r.db.Model(&data).Where("user_id IS NULL")
	if params.Keyword != "" {
		query.Where("LOWER(nama) LIKE LOWER(?)", "%"+params.Keyword+"%").Or("LOWER(badge) LIKE LOWER(?)", "%"+params.Keyword+"%")
	}
	query.Count(&params.TotalData)
	query.Offset(params.Offset).Limit(params.PageSize)
	err := query.Find(&data).Error
	return data, err
}

func (r *MsEmployeeRepository) FindByBadge(badge string) (*MsEmployee, error) {
	var model *MsEmployee
	err := r.db.First(&model, constant.WhereBadge, badge).Error
	if err == gorm.ErrRecordNotFound {
		err = errs.ErrRecordNotFound("Karyawan")
	}
	return model, err
}

func (r *MsEmployeeRepository) FirstOrCreate(data *MsEmployee) (*MsEmployee, error) {
	err := r.db.Where(constant.WhereBadge, data.Badge).FirstOrCreate(&data).Error
	return data, err
}

func (r *MsEmployeeRepository) FindByUUID(uuid string) (*MsEmployee, error) {
	var model *MsEmployee
	err := r.db.First(&model, constant.WhereUUID, uuid).Error
	if err == gorm.ErrRecordNotFound {
		err = errs.ErrRecordNotFound("Master Data Karyawan")
	}
	return model, err
}

func (r *MsEmployeeRepository) FindByUUIDs(uuids []string) ([]MsEmployee, error) {
	var data []MsEmployee
	err := r.db.Where(constant.WhereUUIDIn, uuids).Find(&data).Error
	return data, err
}

func (r *MsEmployeeRepository) Create(data *MsEmployee) (*MsEmployee, error) {
	if err := r.db.Create(&data).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			var oldData MsEmployee
			errFindRemoved := r.db.Unscoped().
				Where(constant.WhereBadge, &data.Badge).
				First(&oldData).Error
			if errFindRemoved != nil {
				return nil, errs.ErrEmployeeExists
			}
			data.DeletedAt = nil
			errUpdateRemoved := r.db.Model(&data).Unscoped().Updates(&data).Error
			if errUpdateRemoved != nil {
				return nil, errUpdateRemoved
			}
		} else {
			return nil, err
		}
	}
	return data, nil
}

func (r *MsEmployeeRepository) Update(oldData *MsEmployee, newData *MsEmployee) (*MsEmployee, error) {
	err := r.db.Model(&oldData).Updates(&newData).Error
	return oldData, err
}
func (r *MsEmployeeRepository) UpdateStatus(uuid string) error {
	var employee MsEmployee
	trx := r.db.Begin()
	if err := trx.First(&employee, constant.WhereUUID, uuid).Error; err != nil {
		defer trx.Rollback()
		return err
	}
	if err := trx.Model(&employee).Update("is_active", gorm.Expr("~is_active")).Error; err != nil {
		defer trx.Rollback()
		return err
	}
	if err := trx.Model(&MsUser{}).Where("employee_id = ?", employee.ID).Update("is_active", gorm.Expr("~is_active")).Error; err != nil {
		defer trx.Rollback()
		return err
	}
	return trx.Commit().Error
}

func (r *MsEmployeeRepository) Delete(uuid string) error {
	res := r.db.Delete(&MsEmployee{}, constant.WhereUUID, uuid)
	if res.RowsAffected == 0 {
		return errs.ErrNothingDeleted
	}
	return res.Error
}
