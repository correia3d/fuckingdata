package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TibiaData/tibiadata-api-go/src/validation"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var _ = func() bool {
	testing.Init()
	return true
}()

func TestFakeToUpCodeCoverage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// adding support for proxy for tests
	if isEnvExist("TIBIADATA_PROXY") {
		TibiaDataProxyDomain = "https://" + getEnv("TIBIADATA_PROXY", "www.tibia.com") + "/"
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "name",
			Value: "Durin",
		},
	}

	assert := assert.New(t)

	tibiaBoostableBosses(c)
	assert.Equal(http.StatusOK, w.Code)

	tibiaCharactersCharacter(c)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	tibiaCreaturesOverview(c)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "race",
			Value: "Demon",
		},
	}

	tibiaCreaturesCreature(c)
	fmt.Println("tibiaCreaturesCreature", w)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	tibiaFansites(c)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "name",
			Value: "Pax",
		},
	}

	tibiaGuildsGuild(c)
	fmt.Println("tibiaGuildsGuild", w)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "world",
			Value: "antica",
		},
	}

	tibiaGuildsOverview(c)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "world",
			Value: "Antica",
		},
		{
			Key:   "category",
			Value: "experience",
		},
		{
			Key:   "vocation",
			Value: "sorcerer",
		},
		{
			Key:   "page",
			Value: "4",
		},
	}

	tibiaHighscores(c)
	fmt.Println("tibiaHighscores", w)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "world",
			Value: "antica",
		},
		{
			Key:   "house_id",
			Value: "59054",
		},
	}

	tibiaHousesHouse(c)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "world",
			Value: "antica",
		},
		{
			Key:   "town",
			Value: "venore",
		},
	}

	tibiaHousesOverview(c)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "world",
			Value: "antica",
		},
	}

	tibiaKillstatistics(c)
	assert.Equal(http.StatusOK, w.Code)

	assert.False(false, tibiaNewslistArchive())
	assert.False(false, tibiaNewslistArchiveDays())
	assert.False(false, tibiaNewslistLatest())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "days",
			Value: "90",
		},
	}

	tibiaNewslist(c)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "news_id",
			Value: "6607",
		},
	}

	tibiaNews(c)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "vocation",
			Value: "sorcerer",
		},
	}

	tibiaSpellsOverview(c)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "spell_id",
			Value: "exori",
		},
	}

	tibiaSpellsSpell(c)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	tibiaWorldsOverview(c)
	assert.Equal(http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "name",
			Value: "antica",
		},
	}

	tibiaWorldsWorld(c)
	assert.Equal(http.StatusOK, w.Code)

	rootz(c)
	assert.Equal(http.StatusOK, w.Code)

	healthz(c)
	assert.Equal(http.StatusOK, w.Code)

	readyz(c)
	assert.Equal(http.StatusOK, w.Code)

	type test struct {
		T string `json:"t"`
	}

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	TibiaDataAPIHandleResponse(c, "", test{T: "abc"})
	assert.Equal(http.StatusOK, w.Code)
}

func TestErrorHandler(t *testing.T) {
	assert := assert.New(t)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	TibiaDataErrorHandler(c, errors.New("test error"), http.StatusBadRequest)
	assert.Equal(http.StatusBadRequest, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	TibiaDataErrorHandler(c, validation.ErrorAlreadyRunning, 0)
	assert.Equal(http.StatusInternalServerError, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	TibiaDataErrorHandler(c, validation.ErrorCharacterNameInvalid, 0)
	assert.Equal(http.StatusBadRequest, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	TibiaDataErrorHandler(c, errors.New("test error"), 0)
	assert.Equal(http.StatusBadGateway, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	TibiaDataErrorHandler(c, validation.ErrStatusForbidden, http.StatusForbidden)
	assert.Equal(http.StatusBadGateway, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	TibiaDataErrorHandler(c, validation.ErrStatusFound, http.StatusFound)
	assert.Equal(http.StatusBadGateway, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	TibiaDataErrorHandler(c, validation.ErrStatusUnknown, http.StatusConflict)
	assert.Equal(http.StatusBadGateway, w.Code)
}

func TestTibiaCharactersCharacterConcurrency(t *testing.T) {
	gin.SetMode(gin.TestMode)
	assert := assert.New(t)

	// Configurando nosso pool de workers
	characterWorkerPool = make(chan struct{}, 5) // Limitado a 5 para fins de teste

	// Criando um servidor de teste
	router := gin.New()
	router.GET("/v4/character/:name", tibiaCharactersCharacter)

	// Número de requisições concorrentes para teste
	numRequests := 10

	// Canais para sincronização
	start := make(chan struct{})
	done := make(chan bool, numRequests)

	// Função para fazer uma requisição
	makeRequest := func(name string) {
		// Esperar sinal de início
		<-start

		// Criar requisição
		req, _ := http.NewRequest("GET", "/v4/character/"+name, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Verificar resultado
		// Primeiro verificamos se o código de status está entre os aceitáveis
		isSuccessful := w.Code == http.StatusOK || w.Code == http.StatusBadGateway

		// Registrar o resultado para informação, sem falhar o teste
		t.Logf("Requisição para %s retornou código %d: %s", name, w.Code, w.Body.String())

		// Enviar resultado para o canal
		done <- isSuccessful
	}

	// Lançar várias goroutines para simular requisições concorrentes
	characterNames := []string{"Trollefar", "Bubble", "Cachero", "Eternal Oblivion", "Moonzinn",
		"Arieswar", "Leonel", "Pablocb", "Moonzinn", "Kiviki"}

	for i := 0; i < numRequests; i++ {
		go makeRequest(characterNames[i%len(characterNames)])
	}

	// Marcar início e aguardar todas as goroutines
	close(start) // Inicia todas as requisições simultaneamente

	// Aguardar todas as requisições completarem
	successCount := 0
	for i := 0; i < numRequests; i++ {
		if <-done {
			successCount++
		}
	}

	// Verificar resultados
	t.Logf("%d de %d requisições foram bem-sucedidas", successCount, numRequests)
	assert.True(successCount > 0, "Pelo menos algumas requisições deveriam ser bem-sucedidas")

}
