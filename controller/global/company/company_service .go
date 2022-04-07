package company

import (
	"fmt"
	"log"

	"backend/models/global"

	"github.com/mashingan/smapping"
)

//CompanyService is a ....
type CompanyService interface {
	InsertCompany(b global.COMPANY) global.COMPANY
	UpdateCompany(b global.COMPANY) global.COMPANY
	DeleteCompany(b global.COMPANY)
	GetAllCompany() []global.COMPANY
	FindCompanyByID(modulID uint64) global.COMPANY
	IsAllowedToEdit(userID string, modulID uint64) bool
}

type companyService struct {
	companyRepository CompanyRepository
}

//NewCompanyService .....
func NewCompanyService(companyRepo CompanyRepository) CompanyService {
	return &companyService{
		companyRepository: companyRepo,
	}
}

func (service *companyService) InsertCompany(b global.COMPANY) global.COMPANY {
	company := global.COMPANY{}
	err := smapping.FillStruct(&company, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.companyRepository.InsertCompany(company)
	return res
}

func (service *companyService) UpdateCompany(b global.COMPANY) global.COMPANY {
	res := service.companyRepository.UpdateCompany(b)
	return res
}

func (service *companyService) GetAllCompany() []global.COMPANY {
	return service.companyRepository.GetAllCompany()
}

func (service *companyService) FindCompanyByID(modulID uint64) global.COMPANY {
	return service.companyRepository.FindCompanyByID(modulID)
}

func (service *companyService) DeleteCompany(b global.COMPANY) {
	service.companyRepository.DeleteCompany(b)
}

func (service *companyService) IsAllowedToEdit(userID string, companyID uint64) bool {
	b := service.companyRepository.FindCompanyByID(companyID)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return userID == id
}
