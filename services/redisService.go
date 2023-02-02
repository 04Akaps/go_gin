package services

// type struct를 짜기는 공수가 들어가니.. 일단 간단하게 string값으로 fix
type RedisService interface {
	SetRedis(string, string) error
	GetRedis(string) (string, error)
	DeleteRedis(string) error
}