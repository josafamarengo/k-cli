package cmd

import (
	"os"
    "testing"
	"io/ioutil"

    "github.com/stretchr/testify/assert"
)

func TestParseValue(t *testing.T) {
    server, token := parseValue("server-url token-value")
    assert.Equal(t, "server-url", server)
    assert.Equal(t, "token-value", token)

    server, token = parseValue("server-url")
    assert.Equal(t, "server-url", server)
    assert.Equal(t, "", token)
}

func TestAddContext(t *testing.T) {
    tempFile, err := ioutil.TempFile("", "contexts.txt")
    assert.NoError(t, err)
    defer os.Remove(tempFile.Name())

    contextFile = tempFile.Name() // Override contextFile for test
    addContext("test-context", "https://server-url", "token-value", "openshift")

    contents, err := ioutil.ReadFile(tempFile.Name())
    assert.NoError(t, err)
    assert.Contains(t, string(contents), "os-test-context=https://server-url")
    assert.Contains(t, string(contents), "os-test-context-token=token-value")
}

func TestRemoveContext(t *testing.T) {
	// Cria um arquivo temporário para o teste
	tmpfile, err := ioutil.TempFile("", "contexts.txt")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	// Preenche o arquivo com conteúdo de exemplo
	contextFile = tmpfile.Name()
	ioutil.WriteFile(contextFile, []byte("test-context=https://test-server\ntest-context-token=test-token\n"), 0644)

	removeContext("test-context")

	content, err := ioutil.ReadFile(contextFile)
	if err != nil {
		t.Fatalf("Error reading temp file: %v", err)
	}

	expected := ""
	if string(content) != expected {
		t.Errorf("Expected content %q, but got %q", expected, string(content))
	}
}

func TestSetContextOpenShift(t *testing.T) {
	// Cria um arquivo temporário para o teste
	tmpfile, err := ioutil.TempFile("", "contexts.txt")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	contextFile = tmpfile.Name()
	ioutil.WriteFile(contextFile, []byte("os-test-context=https://test-server test-token\n"), 0644)

	// Testa se o contexto de OpenShift é definido corretamente
	setContext("os-test-context")
	// Aqui você pode adicionar verificações para garantir que o comando oc login foi chamado corretamente
}

func TestSetContextRancher(t *testing.T) {
	// Cria um arquivo temporário para o teste
	tmpfile, err := ioutil.TempFile("", "contexts.txt")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	contextFile = tmpfile.Name()
	ioutil.WriteFile(contextFile, []byte("test-context=https://test-server test-token\n"), 0644)

	// Testa se o contexto de Rancher é definido corretamente
	setContext("test-context")
	// Aqui você pode adicionar verificações para garantir que o comando rancher login foi chamado corretamente
}

func TestListContexts(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "contexts.txt")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	contextFile = tmpfile.Name()
	ioutil.WriteFile(contextFile, []byte("test-context=https://test-server\ntest-context-token=test-token\n"), 0644)

	// Redireciona a saída para verificar o output da função
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	listContexts()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old

	expected := "Here is the list of available contexts to configure:\n   test-context = https://test-server\n"
	if string(out) != expected {
		t.Errorf("Expected %q, but got %q", expected, string(out))
	}
}
