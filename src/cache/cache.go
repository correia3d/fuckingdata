package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	// Cliente Redis global
	Client *redis.Client
	// Contexto para operações do Redis
	Ctx = context.Background()
)

// Setup inicializa o cliente Redis
func Setup(addr, password string, db int) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

// Get retorna um valor do cache
func Get(key string) (string, error) {
	return Client.Get(Ctx, key).Result()
}

// Set define um valor no cache com TTL
func Set(key string, value interface{}, ttl time.Duration) error {
	return Client.Set(Ctx, key, value, ttl).Err()
}

// Delete remove uma chave do cache
func Delete(key string) error {
	return Client.Del(Ctx, key).Err()
}

// FlushAll limpa todo o cache
func FlushAll() error {
	return Client.FlushAll(Ctx).Err()
}

// GetTTL retorna o TTL recomendado para diferentes tipos de dados
// GetTTL retorna o TTL recomendado para diferentes tipos de dados
func GetTTL(dataType string) time.Duration {
    switch dataType {
    case "character":
        return 15 * time.Second
    case "world":
        return 1 * time.Second
    case "guild":
        return 15 * time.Second
    case "highscores":
        return 5 * time.Minute
    default:
        return 5 * time.Second
    }
}