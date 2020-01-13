package dto


type UsuarioDTO struct {
	Id          int64  `json:"id,omitempty"`
	Nome        string `json:"nome,omitempty"`
	Sobrenome   string `json:"sobrenome,omitempty"`
	Email       string `json:"email,omitempty"`
	Status      string `json:"status,omitempty"`
	Senha       string `json:"senha,omitempty"`
	DataCriacao string `json:"data_criacao,omitempty"`
}

