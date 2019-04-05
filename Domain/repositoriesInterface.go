package Domain

type repositoryInterface interface {
	GetPerson(id int) PersonDomainModelInterface
}
