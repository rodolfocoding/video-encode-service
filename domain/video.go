package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// `valid:"uuid"` <- isso é chamado de tag, e é utilizado como referencia de validação para lib govalidator
type Video struct {
	ID         string    `valid:"uuid"`
	ResourceID string    `valid:"notnull"`
	FilePath   string    `valid:"notnull"`
	CreatedAt  time.Time `valid:"-"`
}

// no go a funcão init é como se fosse o contructor da classe, ou seja, a primeira funcão a ser chamada.
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {

	//_ é chamado de blank identifier
	_, err := govalidator.ValidateStruct(video)

	if err != nil {
		return err
	}

	//Normalmente retornamos nil quando queremos dizer que nao houve erro
	return nil
}
