package services

import (
	"URLSHORTENER/internal/models"
	"URLSHORTENER/internal/store"
	"log"
	"math/rand"
	"time"
)

// random string or url generation
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func GetShort(l models.Long) {

	s := String(5)
	store.Url_map[l.Long] = s
}

func GetMap() map[string]string {
	return store.Url_map
}

func GetLongFromShort(s string) string {
	for key, value := range store.Url_map {
		if value == s {
			return key
		}
	}
	log.Panic("url not saved")
	return ""
}
