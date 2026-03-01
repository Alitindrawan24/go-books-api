package auth

import "context"

func (usecase *UseCase) GetToken(ctx context.Context) (string, error) {
	token, err := usecase.authRepository.GetToken(ctx)
	if err != nil {
		return "", err
	}

	return token, nil
}
