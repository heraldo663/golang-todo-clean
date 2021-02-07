package auth

import (
	"heraldo663/todo/shared/utils/password"
)

type IAuthUseCase interface {
	Login(d LoginDTO) (*AuthResponse, error)
	Signup(d SignupDTO) (*AuthResponse, error)
}

type authUseCase struct {
	userRepository IUserRepository
	jwt            password.IJwt
	bcrypt         password.IBcrypt
}

// NewAuthUseCase -> auth useCase
func NewAuthUseCase(
	userRepository IUserRepository,
	jwt password.IJwt,
	bcrypt password.IBcrypt,
) IAuthUseCase {
	return &authUseCase{
		userRepository: userRepository,
		jwt:            jwt,
		bcrypt:         bcrypt,
	}
}

func (s *authUseCase) Login(d LoginDTO) (*AuthResponse, error) {
	user, err := s.userRepository.FindUserByEmail(d.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.Verify(user.Password, user.Password); err != nil {
		return nil, err
	}

	t := s.jwt.Generate(&password.TokenPayload{
		ID: user.ID,
	})

	return &AuthResponse{
		User: &UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
		Auth: &AccessResponse{
			Token: t,
		},
	}, nil
}

func (s *authUseCase) Signup(d SignupDTO) (*AuthResponse, error) {
	user := &User{
		Name:     d.Name,
		Password: s.bcrypt.Generate(d.Password),
		Email:    d.Email,
	}

	user, err := s.userRepository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	// generate access token
	t := jwt.Generate(&password.TokenPayload{
		ID: user.ID,
	})

	return &AuthResponse{
		User: &UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
		Auth: &AccessResponse{
			Token: t,
		},
	}, nil
}
