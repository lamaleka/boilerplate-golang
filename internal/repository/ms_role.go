package repository

import (
	"github.com/lamaleka/boilerplate-golang/common/constant"
	"github.com/lamaleka/boilerplate-golang/common/errs"

	"gorm.io/gorm"
)

type MsRoleRepository struct {
	db *DB
}

func NewMsRoleRepository(db *DB) *MsRoleRepository {
	return &MsRoleRepository{db: db}
}

func (r *MsRoleRepository) GetAll(params *MsRoleParams) (*[]MsRole, error) {
	var data *[]MsRole
	query := r.db.Model(&data)
	if params.Keyword != "" {
		query.Where("LOWER(title) LIKE LOWER(?)", "%"+params.Keyword+"%")
	}
	query.Count(&params.TotalData)
	query.Offset(params.Offset).Limit(params.PageSize)
	err := query.Find(&data).Error
	return data, err
}

func (r *MsRoleRepository) FindByUUID(uuid string) (*MsRole, error) {
	var model *MsRole
	err := r.db.First(&model, constant.WhereUUID, uuid).Error
	if err == gorm.ErrRecordNotFound {
		err = errs.ErrRecordNotFound("Master Data Role")
	}
	return model, err
}

func (r *MsRoleRepository) FindByUUIDs(uuids []string) ([]MsRole, error) {
	var data []MsRole
	err := r.db.Where(constant.WhereUUIDIn, uuids).Find(&data).Error
	return data, err
}

func (r *MsRoleRepository) FindByIDs(ids []int) ([]MsRole, error) {
	var data []MsRole
	err := r.db.Where(constant.WhereIDIn, ids).Find(&data).Error
	return data, err
}

func (r *MsRoleRepository) Create(data *MsRole) (*MsRole, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *MsRoleRepository) Update(oldData *MsRole, newData *MsRole) (*MsRole, error) {
	err := r.db.Model(&oldData).Updates(&newData).Error
	return oldData, err
}

func (r *MsRoleRepository) Delete(uuid string) error {
	res := r.db.Delete(&MsRole{}, constant.WhereUUID, uuid)
	if res.RowsAffected == 0 {
		return errs.ErrNothingDeleted
	}
	return res.Error
}
