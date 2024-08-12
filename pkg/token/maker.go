package token

type Maker interface {
	CreateTokenPaseto()

	VerifyTokenPaseto()
}