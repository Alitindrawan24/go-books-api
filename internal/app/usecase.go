package app

type UseCases struct {
}

func NewUseCase(repository *Repositories) *UseCases {
	useCases := &UseCases{}
	return useCases
}
