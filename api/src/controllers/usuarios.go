package controllers

import "net/http"

// CriarUsuario insere usuarios no bd
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando um usuário"))
}

// BuscarUsuarios busca tpdps usuarios do banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuarios"))
}

// BuscarUsuario busca usuario por id
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário com id"))
}

// AtualizarUsuario atualiza usuarios
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

// DeletarUsuario deleta usuario do banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
