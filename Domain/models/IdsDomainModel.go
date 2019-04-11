package models

type IdsDomainModel struct {
	ClientId           int
	CoborrowersIdSlice []int
}

func (i *IdsDomainModel) GetClientId() int {
	return i.ClientId
}

func (i *IdsDomainModel) GetCoborrowersIdSlice() []int {
	return i.CoborrowersIdSlice
}
